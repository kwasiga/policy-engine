package cedarclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kwasi/policy-engine/internal/domain"
)

type policyUpload struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

// SyncPolicies replaces the sidecar's entire active policy set via PUT
// /v1/policies. Called by internal/cache whenever the Postgres-backed
// policy set changes (create/update/rollback/disable) so the sidecar's
// in-memory compiled set never drifts from the source of truth.
func (c *Client) SyncPolicies(ctx context.Context, active []domain.PolicyRecord) error {
	uploads := make([]policyUpload, 0, len(active))
	for _, p := range active {
		uploads = append(uploads, policyUpload{ID: p.ID.String(), Content: p.PolicyText})
	}

	payload, err := json.Marshal(uploads)
	if err != nil {
		return fmt.Errorf("marshal policy set: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, c.baseURL+"/v1/policies", bytes.NewReader(payload))
	if err != nil {
		return fmt.Errorf("build sync request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.authToken)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("call cedar-agent: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("cedar-agent policy sync returned status %d", resp.StatusCode)
	}
	return nil
}
