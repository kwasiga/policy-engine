package cache

import (
	"context"
	"time"

	"github.com/kwasi/policy-engine/internal/domain"
)

// Invalidate reloads the active policy set from Postgres and pushes it to
// the cedar-agent sidecar. Call this after every policy create/update/
// rollback/status-change so the sidecar's compiled set never drifts from
// the source of truth in Postgres.
func (c *PolicyCache) Invalidate(ctx context.Context) ([]domain.PolicyRecord, error) {
	records, err := c.repo.ListActive(ctx)
	if err != nil {
		return nil, err
	}

	if err := c.cedar.SyncPolicies(ctx, records); err != nil {
		// Deliberately not swallowed: if the sidecar sync fails we keep the
		// previous cached set rather than reporting success on a policy
		// change that never actually reached the evaluator.
		return nil, err
	}

	c.mu.Lock()
	c.active = records
	c.lastSynced = time.Now()
	c.mu.Unlock()

	return records, nil
}
