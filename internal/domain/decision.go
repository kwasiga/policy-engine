package domain

import (
	"time"

	"github.com/google/uuid"
)

// DecisionRequest is the inbound payload for POST /evaluate.
type DecisionRequest struct {
	Principal EntityUid      `json:"principal"`
	Action    EntityUid      `json:"action"`
	Resource  EntityUid      `json:"resource"`
	Context   map[string]any `json:"context,omitempty"`
	// Entities supplies attributes/hierarchy needed to resolve the request
	// (e.g. the resource's owning group, the principal's roles).
	Entities []Entity `json:"entities,omitempty"`
}

type Decision string

const (
	Allow Decision = "ALLOW"
	Deny  Decision = "DENY"
)

// DecisionResult is returned from POST /evaluate and is what gets written
// to the audit log verbatim.
type DecisionResult struct {
	Decision              Decision    `json:"decision"`
	DeterminingPolicyIDs  []uuid.UUID `json:"determining_policy_ids"`
	Errors                []string    `json:"errors,omitempty"`
	EvaluatedAt           time.Time   `json:"evaluated_at"`
	LatencyMicros         int64       `json:"latency_micros"`
}
