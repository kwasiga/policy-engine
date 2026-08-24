# API

## POST /evaluate

```json
{
  "principal": {"type": "User", "id": "alice"},
  "action": {"type": "Action", "id": "viewDocument"},
  "resource": {"type": "Document", "id": "doc-123"},
  "context": {"ip": "10.0.0.1"},
  "entities": [
    {"uid": {"type": "Document", "id": "doc-123"}, "attrs": {"owner": "alice"}, "parents": []}
  ]
}
```

Response:

```json
{
  "decision": "ALLOW",
  "determining_policy_ids": ["..."],
  "errors": [],
  "evaluated_at": "2026-08-23T20:00:00Z",
  "latency_micros": 1200
}
```

## /policies

- `GET /policies` — list, optional `?status=` filter
- `POST /policies` — create (starts at version 1, status `draft`)
- `GET /policies/{id}` — fetch current version
- `PUT /policies/{id}` — update (creates a new version, advances `current_version`)
- `DELETE /policies/{id}` — soft-delete (status → `archived`)
- `POST /policies/{id}/rollback` — `{"to_version": n}`, points `current_version` at an older version

See `proto/policy_engine.proto` for the gRPC equivalent.
