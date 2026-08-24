-- Every evaluation decision (allow AND deny) is written here for forensic
-- analysis and compliance reporting. Kept append-only and wide (denormalized
-- entity attrs as JSONB) so audits don't require joining back to mutable
-- entity state that may have since changed.
CREATE TABLE audit_log (
    id                     UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    principal_type         TEXT NOT NULL,
    principal_id           TEXT NOT NULL,
    action_type            TEXT NOT NULL,
    action_id              TEXT NOT NULL,
    resource_type          TEXT NOT NULL,
    resource_id            TEXT NOT NULL,
    context                JSONB NOT NULL DEFAULT '{}',
    entity_attributes      JSONB NOT NULL DEFAULT '{}',
    decision               TEXT NOT NULL CHECK (decision IN ('ALLOW', 'DENY')),
    determining_policy_ids UUID[] NOT NULL DEFAULT '{}',
    evaluation_errors      TEXT[] NOT NULL DEFAULT '{}',
    latency_micros         BIGINT NOT NULL,
    evaluated_at           TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_audit_log_evaluated_at ON audit_log (evaluated_at);
CREATE INDEX idx_audit_log_principal ON audit_log (principal_type, principal_id);
CREATE INDEX idx_audit_log_decision ON audit_log (decision);
CREATE INDEX idx_audit_log_determining_policies ON audit_log USING GIN (determining_policy_ids);
