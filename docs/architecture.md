# Architecture

```
            ┌─────────────┐        ┌──────────────────┐
 REST/gRPC  │             │ HTTP   │                   │
 ─────────▶ │  Go API     │──────▶ │  cedar-agent      │
            │ (this repo) │  :8180 │  sidecar (Rust)   │
            │             │◀────── │  compiled policies│
            └──────┬──────┘        └───────────────────┘
                   │
        ┌──────────┴──────────┐
        ▼                     ▼
  ┌───────────┐        ┌─────────────┐
  │ Postgres  │        │ in-memory   │
  │ policies, │        │ PolicyCache │
  │ versions, │        │ (this repo) │
  │ audit_log │        └─────────────┘
  └───────────┘
```

- **Source of truth**: Postgres (`policies` + `policy_versions`) — full
  history, instant rollback via `current_version`.
- **Evaluation**: delegated to the `cedar-agent` sidecar over localhost
  HTTP; the Go process never links Cedar directly (see
  `sidecar/cedar-agent/README.md` for why).
- **Cache**: `internal/cache.PolicyCache` mirrors the active policy set and
  is the single choke point that keeps Postgres, itself, and the sidecar in
  sync — every mutating `/policies` write must call `Invalidate`.
- **Audit**: every `/evaluate` call (allow and deny) is queued to
  `internal/audit.Logger`, which writes to `audit_log` off the request path.
- **Fail-closed**: sidecar timeout or error → `503`, never a silent allow.
