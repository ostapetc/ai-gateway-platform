COMPOSE    := docker compose -f deploy/local-dev/docker-compose.yml
KUBECONFIG := $(HOME)/.kube/timeweb_config.yaml
KUBECTL    := kubectl --kubeconfig $(KUBECONFIG)
K8S_DIR    := deploy/k8s
NAMESPACE  := ai-gateway
REGISTRY   := ghcr.io/ostapetc/ai-gateway-platform
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

# ── Build & test ──────────────────────────────────────────────────────────────

.PHONY: build
build: ## Build all services (production images)
	$(COMPOSE) build users posts comments

.PHONY: test
test: ## Run tests for all services
	cd services/users    && go test ./...
	cd services/posts    && go test ./...
	cd services/comments && go test ./...

.PHONY: lint
lint: ## Run golangci-lint on all services
	cd services/users    && golangci-lint run ./...
	cd services/posts    && golangci-lint run ./...
	cd services/comments && golangci-lint run ./...

.PHONY: tidy
tidy: ## Run go mod tidy on all services
	cd services/users    && go mod tidy
	cd services/posts    && go mod tidy
	cd services/comments && go mod tidy

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
k8s-apply: ## Apply all Kubernetes manifests (namespace → infra → apps → ingress)
	$(KUBECTL) apply -f $(K8S_DIR)/namespace.yaml
	$(KUBECTL) apply -f $(K8S_DIR)/infra/
	$(KUBECTL) apply -f $(K8S_DIR)/apps/
	$(KUBECTL) apply -f $(K8S_DIR)/ingress.yaml

.PHONY: k8s-delete
k8s-delete: ## Delete all Kubernetes resources in the namespace
	$(KUBECTL) delete namespace $(NAMESPACE) --ignore-not-found

.PHONY: k8s-status
k8s-status: ## Show pod and service status
	$(KUBECTL) get pods,svc,ingress -n $(NAMESPACE)

.PHONY: k8s-logs
k8s-logs: ## Tail logs for a service: make k8s-logs SVC=users
	$(KUBECTL) logs -n $(NAMESPACE) -l app=$(SVC) -f

.PHONY: docker-push
docker-push: ## Build and push production images to registry
	docker build -t $(REGISTRY)/users:$(TAG)    services/users
	docker build -t $(REGISTRY)/posts:$(TAG)    services/posts
	docker build -t $(REGISTRY)/comments:$(TAG) services/comments
	docker push $(REGISTRY)/users:$(TAG)
	docker push $(REGISTRY)/posts:$(TAG)
	docker push $(REGISTRY)/comments:$(TAG)

# ── Help ──────────────────────────────────────────────────────────────────────

.PHONY: help
help: ## Show this help
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-18s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help
