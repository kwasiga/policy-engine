# policy-engine

Production-grade ABAC access-control engine using [Cedar](https://www.cedarpolicy.com/)
(AWS's policy language) for fine-grained, sub-5ms access decisions at scale.

## Architecture

Go API (REST + gRPC, Postgres-backed policy storage with versioning/rollback,
in-memory cache, async audit logging) delegating actual policy compilation
and evaluation to a `cedar-agent` sidecar over localhost. See
[docs/architecture.md](docs/architecture.md) and
[sidecar/cedar-agent/README.md](sidecar/cedar-agent/README.md) for why.

## Layout

```
cmd/server/           entrypoint — wires config, db, cache, cedar client, servers
internal/domain/       shared types: entities, decision request/result, policy record
internal/cedarclient/  HTTP client to the cedar-agent sidecar
internal/storage/      Postgres: policy CRUD/versioning, audit log
internal/cache/        in-memory active-policy cache + sidecar invalidation
internal/audit/        async audit-log writer (off the /evaluate request path)
internal/api/rest/     REST handlers (/evaluate, /policies)
internal/api/grpc/     gRPC service (proto/policy_engine.proto)
migrations/            Postgres schema
sidecar/cedar-agent/   notes on the Cedar evaluator sidecar
docs/                  architecture + API reference
```

## Running locally

```sh
docker compose up --build
```

Brings up Postgres, the `cedar-agent` sidecar, and the Go API on
`:8080` (REST) / `:9090` (gRPC).

## Status

Scaffold stage — types, wiring, and interfaces are in place; repo/handler
bodies are `TODO` pending implementation (see inline comments).
