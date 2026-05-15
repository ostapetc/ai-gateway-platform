COMPOSE    := docker compose -f deploy/local-dev/docker-compose.yml
KUBECONFIG := $(HOME)/.kube/timeweb_config.yaml
KUBECTL    := kubectl --kubeconfig $(KUBECONFIG)
K8S_DIR    := deploy/k8s
NAMESPACE  := ai-gateway
REGISTRY   ?= freelikeatruth
TAG        ?= latest

# ── Dev environment ───────────────────────────────────────────────────────────

.PHONY: up
up: ## Start all services with hot-reload (air)
	$(COMPOSE) up -d

.PHONY: down
down: ## Stop and remove all containers
	$(COMPOSE) down

.PHONY: restart
restart: ## Restart all services
	$(COMPOSE) restart

.PHONY: rebuild
rebuild: ## Rebuild and restart app services (pick up Dockerfile changes)
	$(COMPOSE) up -d --build users posts comments

.PHONY: logs
logs: ## Tail logs from all services (Ctrl+C to stop)
	$(COMPOSE) logs -f

.PHONY: logs-users
logs-users: ## Tail users service logs
	$(COMPOSE) logs -f users

.PHONY: logs-posts
logs-posts: ## Tail posts service logs
	$(COMPOSE) logs -f posts

.PHONY: logs-comments
logs-comments: ## Tail comments service logs
	$(COMPOSE) logs -f comments

.PHONY: ps
ps: ## Show running containers and ports
	$(COMPOSE) ps

.PHONY: endpoints
endpoints: ## Print all HTTP endpoints (docker-compose)
	@printf "\n  Docker Compose HTTP Endpoints\n\n"
	@printf "  \033[36m%-14s\033[0m %s\n" "users"       "http://localhost:8881"
	@printf "  \033[36m%-14s\033[0m %s\n" "posts"       "http://localhost:8882"
	@printf "  \033[36m%-14s\033[0m %s\n" "comments"    "http://localhost:8884"
	@printf "  \033[36m%-14s\033[0m %s\n" "grafana"     "http://localhost:3000"
	@printf "  \033[36m%-14s\033[0m %s\n" "prometheus"  "http://localhost:9090"
	@printf "  \033[36m%-14s\033[0m %s\n" "nats-mon"    "http://localhost:8222"
	@printf "  \033[36m%-14s\033[0m %s\n" "otel-zpages" "http://localhost:55679"
	@printf "\n"

# ── Build & test ──────────────────────────────────────────────────────────────

.PHONY: build
build: ## Build all services (production images)
	$(COMPOSE) build users posts comments

.PHONY: test
test: ## Run tests for all services
	go test ./services/...

.PHONY: fmt
fmt: ## Run gofmt on all services
	gofmt -w services/

.PHONY: vet
vet: ## Run go vet on all services
	go vet ./services/...

.PHONY: lint
lint: ## Run golangci-lint on all services
	golangci-lint run ./services/...

.PHONY: tidy
tidy: ## Run go mod tidy
	go mod tidy

# ── Infrastructure ────────────────────────────────────────────────────────────

.PHONY: infra-up
infra-up: ## Start only infrastructure services (postgres, redis, nats, etc.)
	$(COMPOSE) up -d postgres redis nats otelcol prometheus grafana

.PHONY: infra-down
infra-down: ## Stop infrastructure services
	$(COMPOSE) stop postgres redis nats otelcol prometheus grafana

# ── Utilities ─────────────────────────────────────────────────────────────────

.PHONY: psql
psql: ## Open a psql shell in the postgres container
	$(COMPOSE) exec postgres psql -U platform -d platform

.PHONY: redis-cli
redis-cli: ## Open a redis-cli shell
	$(COMPOSE) exec redis redis-cli

.PHONY: nats-sub
nats-sub: ## Subscribe to all NATS subjects (requires nats CLI)
	$(COMPOSE) exec nats nats sub ">"

.PHONY: clean
clean: ## Remove containers, volumes, and dev build artifacts
	$(COMPOSE) down -v
	find services -type d -name tmp | xargs rm -rf

# ── Kubernetes ────────────────────────────────────────────────────────────────

.PHONY: k8s-apply
k8s-apply: ## Apply all Kubernetes manifests (namespace → infra → apps → cronjobs → ingress)
	$(KUBECTL) apply -f $(K8S_DIR)/namespace.yaml
	$(KUBECTL) apply -R -f $(K8S_DIR)/infra/
	$(KUBECTL) apply -R -f $(K8S_DIR)/apps/
	$(KUBECTL) apply -R -f $(K8S_DIR)/cronjobs/
	$(KUBECTL) apply -R -f $(K8S_DIR)/ingress/

.PHONY: k8s-delete
k8s-delete: ## Delete all Kubernetes resources in the namespace
	$(KUBECTL) delete namespace $(NAMESPACE) --ignore-not-found

.PHONY: k8s-status
k8s-status: ## Show pod and service status
	$(KUBECTL) get pods,svc,ingress -n $(NAMESPACE)

.PHONY: k8s-endpoints
k8s-endpoints: ## Print all HTTP endpoints (k8s ingress)
	@HOST=$$($(KUBECTL) get ingress ai-gateway -n $(NAMESPACE) \
		-o jsonpath='{.status.loadBalancer.ingress[0].ip}{.status.loadBalancer.ingress[0].hostname}' 2>/dev/null); \
	if [ -z "$$HOST" ]; then \
		printf "\n  Ingress has no external IP/hostname yet — is the ingress controller running?\n\n"; \
	else \
		printf "\n  Kubernetes HTTP Endpoints  (host: $$HOST)\n\n"; \
		printf "  \033[36m%-14s\033[0m %s\n" "users"      "http://$$HOST/users"; \
		printf "  \033[36m%-14s\033[0m %s\n" "posts"      "http://$$HOST/posts"; \
		printf "  \033[36m%-14s\033[0m %s\n" "comments"   "http://$$HOST/comments"; \
		printf "  \033[36m%-14s\033[0m %s\n" "grafana"    "http://$$HOST/grafana"; \
		printf "  \033[36m%-14s\033[0m %s\n" "prometheus" "http://$$HOST/prometheus"; \
		printf "\n"; \
	fi

.PHONY: k8s-logs
k8s-logs: ## Tail logs for a service: make k8s-logs SVC=users
	$(KUBECTL) logs -n $(NAMESPACE) -l app=$(SVC) -f

.PHONY: docker-login
docker-login: ## Log in to Docker Hub using .env credentials
	@set -a && . ./.env && set +a && \
		echo "$$DOCKER_PASSWORD" | docker login -u "$$DOCKER_USER" --password-stdin

.PHONY: docker-push
docker-push: ## Build and push production images: make docker-push [REGISTRY=ostapetc] [TAG=latest]
	docker build -f services/users/Dockerfile          -t $(REGISTRY)/users:$(TAG)            .
	docker build -f services/posts/Dockerfile          -t $(REGISTRY)/posts:$(TAG)            .
	docker build -f services/posts-bot-cronjob/Dockerfile -t $(REGISTRY)/posts-bot-cronjob:$(TAG) .
	docker build -f services/comments/Dockerfile       -t $(REGISTRY)/comments:$(TAG)         .
	docker push $(REGISTRY)/users:$(TAG)
	docker push $(REGISTRY)/posts:$(TAG)
	docker push $(REGISTRY)/posts-bot-cronjob:$(TAG)
	docker push $(REGISTRY)/comments:$(TAG)

.PHONY: k8s-set-images
k8s-set-images: ## Update k8s deployments to use current REGISTRY/TAG
	$(KUBECTL) set image deployment/users    users=$(REGISTRY)/users:$(TAG)       -n $(NAMESPACE)
	$(KUBECTL) set image deployment/posts    posts=$(REGISTRY)/posts:$(TAG)       -n $(NAMESPACE)
	$(KUBECTL) set image cronjob/posts-bot-cronjob posts-bot-cronjob=$(REGISTRY)/posts-bot-cronjob:$(TAG) -n $(NAMESPACE)
	$(KUBECTL) set image deployment/comments comments=$(REGISTRY)/comments:$(TAG) -n $(NAMESPACE)

.PHONY: k8s-rollout
k8s-rollout: ## Wait for all app deployments to finish rolling out
	$(KUBECTL) rollout status deployment/users    -n $(NAMESPACE)
	$(KUBECTL) rollout status deployment/posts    -n $(NAMESPACE)
	$(KUBECTL) rollout status deployment/comments -n $(NAMESPACE)

.PHONY: k8s-restart
k8s-restart: ## Rolling restart of app pods (no image change needed)
	$(KUBECTL) rollout restart deployment/users    -n $(NAMESPACE)
	$(KUBECTL) rollout restart deployment/posts    -n $(NAMESPACE)
	$(KUBECTL) rollout restart deployment/comments -n $(NAMESPACE)

.PHONY: k8s-deploy
k8s-deploy: docker-push k8s-apply k8s-restart ## Build & push images, apply manifests, restart pods

.PHONY: deploy
deploy: docker-push k8s-set-images k8s-rollout ## Full deploy: build → push → update k8s → wait for rollout

# ── Help ──────────────────────────────────────────────────────────────────────

.PHONY: help
help: ## Show this help
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-18s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help
