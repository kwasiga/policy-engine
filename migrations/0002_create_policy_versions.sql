-- Append-only version history. Every create/update inserts a new row here;
-- policies.current_version is updated to point at it. Rollback = set
-- current_version back to an older row's version (no data is deleted).
CREATE TABLE policy_versions (
    policy_id    UUID NOT NULL REFERENCES policies (id) ON DELETE CASCADE,
    version      INTEGER NOT NULL,
    policy_text  TEXT NOT NULL,
    status       policy_status NOT NULL,
    created_by   TEXT NOT NULL,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT now(),
    PRIMARY KEY (policy_id, version)
);

CREATE INDEX idx_policy_versions_policy_id ON policy_versions (policy_id);
