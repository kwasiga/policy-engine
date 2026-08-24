package storage

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/kwasi/policy-engine/internal/domain"
)

// PolicyRepo owns the policies/policy_versions tables: CRUD, versioning,
// and rollback. It does not know about the cedar-agent sidecar or the
// cache — callers (internal/cache) are responsible for re-syncing after a
// mutation.
type PolicyRepo struct {
	pool *pgxpool.Pool
}

func NewPolicyRepo(pool *pgxpool.Pool) *PolicyRepo {
	return &PolicyRepo{pool: pool}
}

// Create inserts a new policy at version 1 and writes the matching
// policy_versions row in the same transaction.
func (r *PolicyRepo) Create(ctx context.Context, name, description, policyText, createdBy string) (*domain.PolicyRecord, error) {
	// TODO: wrap in a transaction; insert into policies then policy_versions.
	return nil, fmt.Errorf("not implemented")
}

// Update inserts a new policy_versions row (version = current + 1) and
// advances policies.current_version to it. The old version's text is left
// untouched in history for Rollback.
func (r *PolicyRepo) Update(ctx context.Context, id uuid.UUID, policyText, updatedBy string) (*domain.PolicyRecord, error) {
	// TODO: fetch current_version, insert version+1, update policies row.
	return nil, fmt.Errorf("not implemented")
}

// Rollback points policies.current_version at an earlier existing version
// without deleting any history.
func (r *PolicyRepo) Rollback(ctx context.Context, id uuid.UUID, toVersion int32) (*domain.PolicyRecord, error) {
	// TODO: verify policy_versions row exists for (id, toVersion), then update.
	return nil, fmt.Errorf("not implemented")
}

func (r *PolicyRepo) Get(ctx context.Context, id uuid.UUID) (*domain.PolicyRecord, error) {
	// TODO: join policies + policy_versions on current_version.
	return nil, fmt.Errorf("not implemented")
}

func (r *PolicyRepo) Delete(ctx context.Context, id uuid.UUID) error {
	// TODO: soft-delete via status = 'archived' rather than a hard DELETE,
	// so audit_log.determining_policy_ids stays resolvable.
	return fmt.Errorf("not implemented")
}

// ListActive returns every policy currently in the "active" status — the
// set that should be synced to the cedar-agent sidecar and used for
// evaluation.
func (r *PolicyRepo) ListActive(ctx context.Context) ([]domain.PolicyRecord, error) {
	// TODO: SELECT ... WHERE status = 'active'.
	return nil, fmt.Errorf("not implemented")
}

// List returns policies with optional status filtering, for the
// GET /policies management endpoint.
func (r *PolicyRepo) List(ctx context.Context, status *domain.PolicyStatus) ([]domain.PolicyRecord, error) {
	// TODO: SELECT ... [WHERE status = $1].
	return nil, fmt.Errorf("not implemented")
}
