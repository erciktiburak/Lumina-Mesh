# Lumina-Mesh
[![Version](https://img.shields.io/badge/version-1.0.0--beta-blue.svg)](https://github.com/erciktiburak/Lumina-Mesh)

The Open-Source Serverless AI Mesh for Edge-to-Cloud Orchestration.

**Lumina-Mesh** is a low-latency, event-driven communication network connecting Docker containers, serverless functions, and AI models. Running entirely on Virtual Containers and Microservices, it orchestrates complex AI workflows (e.g., video processing -> sentiment analysis -> Slack notification) using a single YAML file.

## Architecture
Lumina-Mesh uses a decentralized approach where each node is part of a high-speed messaging mesh.

```mermaid
graph TD
    A[Lumina-Node] -->|Pub/Sub| B(NATS JetStream)
    A -->|Discovery| C[Peer Nodes]
    A -->|Tasks| D[Worker Pool]
    D -->|Failures| E[Dead Letter Queue]
    D -->|Runtime| F[WASM Executor]
    G[Lumina-CLI] -->|Monitor| B
```

## Features
- **Core Engine:** Written in Go for optimal concurrency and memory usage.
- **Messaging Layer:** Uses NATS / JetStream for a modern, high-speed pub/sub system.
- **Workflow Definition:** YAML / DSL-based workflows.
- **Runtime Environment:** WebAssembly (WASM) for isolated and cross-platform function execution.
- **AI Processing:** Hugging Face / ONNX integration for local AI inference.
- **Observability:** OpenTelemetry for distributed tracing and Prometheus for metrics.

## Observability
Lumina-Mesh integrates with OpenTelemetry and Prometheus out of the box:
- **Tracing:** Distributed tracing for every workflow step.
- **Metrics:** Real-time node load and message throughput.
- **Visualizer:** A built-in D3.js mesh visualizer on port 8080.

## Getting Started
To start a node, ensure NATS is running:
```bash
docker run -d --name nats -p 4222:4222 nats
go run cmd/lumina/main.go
```

To monitor the mesh:
```bash
go run cmd/lumina-cli/main.go monitor
```

To deploy a serverless workflow:
```bash
go run cmd/lumina-cli/main.go deploy examples/video-workflow.yaml
```

To see the visualizer:
```bash
# The node starts a web server on :8080
open http://localhost:8080
```

## Workflow Example
Workflows are defined in YAML and executed as high-performance WASM sandboxes:
```yaml
name: sentiment-analysis
steps:
  - name: analyze
    action: wasm://ai-sentiment.wasm
    inputs:
      text: "Lumina-Mesh is amazing!"
```

---
*Lumina-Mesh: The Ghost in the Machine (v1.0.0-beta)*
