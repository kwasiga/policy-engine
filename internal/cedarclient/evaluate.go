package cedarclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kwasi/policy-engine/internal/domain"
)

type authorizeRequest struct {
	Principal string          `json:"principal"`
	Action    string          `json:"action"`
	Resource  string          `json:"resource"`
	Context   map[string]any  `json:"context"`
	Entities  []domain.Entity `json:"entities"`
}

type authorizeResponse struct {
	Decision string   `json:"decision"` // "Allow" | "Deny"
	Reasons  []string `json:"reasons"`  // matched policy IDs
	Errors   []string `json:"errors"`
}

// Authorize calls the sidecar's POST /v1/is_authorized endpoint. Determining
// policy IDs come back as strings and are parsed to uuid.UUID by the caller
// (internal/api) since this package stays free of the domain/uuid coupling.
func (c *Client) Authorize(ctx context.Context, req domain.DecisionRequest) (*authorizeResponse, error) {
	body := authorizeRequest{
		Principal: req.Principal.String(),
		Action:    req.Action.String(),
		Resource:  req.Resource.String(),
		Context:   req.Context,
		Entities:  req.Entities,
	}
	payload, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("marshal authorize request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+"/v1/is_authorized", bytes.NewReader(payload))
	if err != nil {
		return nil, fmt.Errorf("build authorize request: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+c.authToken)

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("call cedar-agent: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("cedar-agent returned status %d", resp.StatusCode)
	}

	var out authorizeResponse
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, fmt.Errorf("decode authorize response: %w", err)
	}
	return &out, nil
}
