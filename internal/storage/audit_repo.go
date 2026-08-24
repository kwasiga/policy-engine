package storage

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/kwasi/policy-engine/internal/domain"
)

// AuditRepo writes every evaluation decision — allow and deny alike — to
// the append-only audit_log table for forensic analysis and compliance
// reporting. Writes should never block or fail the /evaluate response path;
// see internal/audit for the async wrapper that enforces that.
type AuditRepo struct {
	pool *pgxpool.Pool
}

func NewAuditRepo(pool *pgxpool.Pool) *AuditRepo {
	return &AuditRepo{pool: pool}
}

// Record captures one decision plus the request shape that produced it —
// principal/action/resource, context, entity attributes, matched policy
// IDs, and errors — as a single audit_log row.
func (r *AuditRepo) Record(ctx context.Context, req domain.DecisionRequest, result domain.DecisionResult) error {
	// TODO: INSERT INTO audit_log (...).
	return fmt.Errorf("not implemented")
}

// Query supports compliance/forensic lookups by time range, principal, or
// decision outcome, backing a future GET /audit-log endpoint.
func (r *AuditRepo) Query(ctx context.Context, filter AuditFilter) ([]domain.DecisionResult, error) {
	// TODO: SELECT ... WHERE evaluated_at BETWEEN $1 AND $2 [AND principal_id = $3] [AND decision = $4].
	return nil, fmt.Errorf("not implemented")
}

type AuditFilter struct {
	PrincipalID *string
	Decision    *domain.Decision
	// From/To bound evaluated_at; nil means unbounded on that side.
	From, To *string
}
