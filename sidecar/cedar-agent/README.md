# cedar-agent sidecar

The Go API never links Cedar directly. Instead, `internal/cedarclient`
talks over localhost HTTP to [`cedar-agent`](https://github.com/cedar-policy/cedar-agent),
AWS's small Rust HTTP service that wraps the `cedar-policy` crate.

## Why a sidecar

- Keeps the Go binary a pure Go build (no cgo, no cross-compiling Rust).
- `cedar-agent` owns policy parsing/compilation and holds the compiled
  policy set in memory — that satisfies "in-memory cache layer for compiled
  policies" at the Cedar layer itself; `internal/cache` on the Go side
  mirrors the active set so Postgres, the cache, and the sidecar never
  drift apart (see `internal/cache/invalidation.go`).
- Tradeoff: every `/evaluate` call costs one extra localhost HTTP hop.
  `internal/cedarclient.Client` uses a 50ms timeout and callers must treat
  a sidecar failure as fail-closed DENY, never a silent ALLOW.

## Endpoints used

- `PUT /v1/policies` — replace the active policy set (called by
  `internal/cache.PolicyCache.Invalidate`)
- `POST /v1/is_authorized` — evaluate one request (called by
  `internal/cedarclient.Client.Authorize`)
- `GET /health` — liveness, used by the Go service's own `/healthz` and by
  `docker-compose.yml`'s healthcheck

## Running locally

`docker-compose.yml` runs `cedar-agent` as its own container on
`localhost:8180`. To run it standalone instead:

```sh
docker run -p 8180:8180 ghcr.io/cedar-policy/cedar-agent:latest
```

Set `CEDAR_AGENT_AUTH_TOKEN` in `.env` to match whatever the sidecar is
configured to require.
