-- Current-state table: one row per logical policy, pointing at its active
-- version. policy_versions (0002) holds full history for rollback.
CREATE TYPE policy_status AS ENUM ('draft', 'active', 'disabled', 'archived');

CREATE TABLE policies (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            TEXT NOT NULL UNIQUE,
    description     TEXT,
    current_version INTEGER NOT NULL DEFAULT 1,
    status          policy_status NOT NULL DEFAULT 'draft',
    created_by      TEXT NOT NULL,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_policies_status ON policies (status);
