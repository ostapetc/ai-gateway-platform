# AI Gateway Platform

Go microservices monorepo. AI routing layer that proxies requests to OpenAI, Ollama, and Anthropic with async processing via NATS JetStream.

## Architecture

```
Client → API Gateway → Router Service → [OpenAI | Ollama | Anthropic]
                                  ↓
                         NATS JetStream → Async Workers
```

Shared infrastructure: PostgreSQL, Redis, NATS JetStream, Prometheus, Grafana, OpenTelemetry.

## Monorepo Layout

```
services/          # one directory per microservice
  <service-name>/
libs/              # shared libraries (no business logic)
  config/
  logger/
  tracing/
  nats/
  protobuf/
deploy/
  docker/
  k8s/
  helm/
  local-dev/
tools/             # code generation scripts, makefile helpers
docs/
  architecture/
  ADR/
```

## Framework: go-zero

All services use [go-zero](https://go-zero.dev/). Follow these conventions:

- Generate RPC services with `goctl rpc protoc`. Never hand-write gRPC boilerplate.
- Generate REST APIs with `goctl api go`. Use `.api` descriptor files.
- Service config lives in `etc/<service>.yaml`, loaded via `zrpc.MustNewServer` / `rest.MustNewServer`.
- Use `go-zero` built-in middleware for logging, tracing, and metrics — do not reimplement them.
- Dependency injection happens in `main.go` via `svccontext.go` (ServiceContext pattern).

## Domain-Driven Design

Each service owns its bounded context. Never import domain types across service boundaries — communicate via protobuf contracts.

**Domain layer must contain:**
- Entities with identity (e.g., `Request`, `Provider`, `Usage`)
- Value objects (immutable, comparable by value)
- Repository interfaces (not implementations)
- Domain errors (`domain.ErrProviderUnavailable`, etc.)
- No framework imports, no I/O

**Application layer must contain:**
- Services with one method, examples: UserUpdater.Update(), ChatProcessor.Process()
- Command/Query objects (plain structs)
- No HTTP/gRPC types leaking in

## SOLID

- **S**: one use case per struct. If a use case file grows past ~150 lines, split it.
- **O**: extend behavior via new use cases or middleware, not by modifying existing ones.
- **L**: repository implementations must be substitutable with any other implementation of the same interface.
- **I**: keep repository interfaces narrow — one interface per use case if usage differs.
- **D**: domain and application layers depend on interfaces. Infrastructure provides implementations. Wire in `main.go`.

## gRPC Conventions

- All inter-service communication is gRPC. No direct HTTP calls between services.
- Proto files live in `services/<name>/api/<name>.proto`.
- Generated code goes in `services/<name>/api/<name>`. Commit generated code.
- Use `context.Context` for deadlines and cancellation — always propagate it.
- Return domain errors mapped to gRPC status codes in the infrastructure layer, not in the domain.
- Streaming RPCs for long-running AI responses (server-side streaming).

## Async / NATS JetStream

- Use `libs/nats` wrapper for pub/sub. Do not use raw `nats.go` client outside the lib.
- Event names: `<service>.<entity>.<past-tense-verb>` — e.g., `router.request.completed`.
- Consumers are idempotent — messages may be redelivered.
- Dead-letter queue for failed messages: `<subject>.dlq`.

## Shared Libraries (`libs/`)

- Libraries contain zero business logic — only cross-cutting infrastructure concerns.
- Each lib exposes a single public API; internals are unexported.
- No lib may import from `services/`.

## Code Style
- `gofmt` + `goimports` on every file. Run `golangci-lint` before committing.
- Error wrapping: `fmt.Errorf("operationName: %w", err)`. No bare `errors.New` inside functions that call other functions.
- No `init()` functions.
- Constructors return `(T, error)` — never panic on invalid config.
- Context is always the first parameter, named `ctx`.
- Unexport everything that does not need to be part of a public API.
- Use table-driven tests with `t.Run`.

## Testing
- Domain and application layers: pure unit tests, no external dependencies.
- Infrastructure layer: integration tests using testcontainers.
- Use `go test ./...` from service root.
- Test files alongside source: `foo_test.go` in the same package for white-box, `foo_test` package for black-box.
