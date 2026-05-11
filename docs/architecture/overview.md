         Client
            |
            v
       API Gateway
            |
  +--------------------+
  | Router Service     |
  +--------------------+
      |       |      |
      v       v      v
   OpenAI   Ollama   Anthropic

        Async Layer
             |
             v
       NATS JetStream
             |
             v
        Async Workers

Shared Infrastructure:
- PostgreSQL
- Redis
- Prometheus
- Grafana
- OpenTelemetry