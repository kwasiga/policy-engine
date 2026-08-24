// Package cache holds the in-memory, read-optimized view of the active
// policy set. Note the split of responsibilities in this sidecar
// architecture: the *compiled* Cedar policy set actually lives inside the
// cedar-agent process (it owns parsing/compilation); this cache exists so
// (a) GET /policies doesn't hit Postgres on every request, and (b) writes
// have a single choke point that keeps Postgres, this cache, and the
// sidecar's compiled set all consistent. See invalidation.go.
package cache

import (
	"context"
	"sync"
	"time"

	"github.com/kwasi/policy-engine/internal/cedarclient"
	"github.com/kwasi/policy-engine/internal/domain"
	"github.com/kwasi/policy-engine/internal/storage"
)

type PolicyCache struct {
	mu          sync.RWMutex
	active      []domain.PolicyRecord
	lastSynced  time.Time

	repo   *storage.PolicyRepo
	cedar  *cedarclient.Client
	ttl    time.Duration
}

func New(repo *storage.PolicyRepo, cedar *cedarclient.Client, ttl time.Duration) *PolicyCache {
	return &PolicyCache{repo: repo, cedar: cedar, ttl: ttl}
}

// Active returns the cached active policy set, transparently refreshing
// from Postgres (and re-syncing the sidecar) if the TTL has expired — a
// belt-and-suspenders guard against a missed Invalidate() call, e.g. from a
// direct DB write outside this process.
func (c *PolicyCache) Active(ctx context.Context) ([]domain.PolicyRecord, error) {
	c.mu.RLock()
	stale := time.Since(c.lastSynced) > c.ttl
	snapshot := c.active
	c.mu.RUnlock()

	if !stale && snapshot != nil {
		return snapshot, nil
	}
	return c.Invalidate(ctx)
}
