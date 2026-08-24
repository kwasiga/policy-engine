package domain

import (
	"time"

	"github.com/google/uuid"
)

// PolicyStatus is the lifecycle state of a stored policy. Only Active
// policies are synced to the cedar-agent sidecar; Draft/Disabled/Archived
// are excluded from evaluation but retained for history and rollback.
type PolicyStatus string

const (
	StatusDraft    PolicyStatus = "draft"
	StatusActive   PolicyStatus = "active"
	StatusDisabled PolicyStatus = "disabled"
	StatusArchived PolicyStatus = "archived"
)

// PolicyRecord is a stored policy: Cedar source text plus versioning
// metadata. Every update inserts a new row in policy_versions; Version
// tracks which one is current, enabling instant rollback (see
// internal/storage/policy_repo.go).
type PolicyRecord struct {
	ID          uuid.UUID    `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description,omitempty"`
	Version     int32        `json:"version"`
	PolicyText  string       `json:"policy_text"`
	Status      PolicyStatus `json:"status"`
	CreatedBy   string       `json:"created_by"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}
