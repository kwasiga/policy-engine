package rest

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/kwasi/policy-engine/internal/domain"
)

// handleEvaluate is POST /evaluate: decodes a DecisionRequest, calls the
// cedar-agent sidecar via s.cedar, returns allow/deny with reasoning, and
// fires the result at s.auditor — allowed and denied decisions alike.
func (s *Server) handleEvaluate(w http.ResponseWriter, r *http.Request) {
	started := time.Now()

	var req domain.DecisionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	resp, err := s.cedar.Authorize(r.Context(), req)
	if err != nil {
		// Fail closed: a sidecar error must never be reported as an ALLOW.
		http.Error(w, "evaluation unavailable", http.StatusServiceUnavailable)
		return
	}

	result := domain.DecisionResult{
		Decision:      domain.Decision(resp.Decision),
		Errors:        resp.Errors,
		EvaluatedAt:   time.Now(),
		LatencyMicros: time.Since(started).Microseconds(),
		// TODO: parse resp.Reasons (policy ID strings) into
		// DeterminingPolicyIDs ([]uuid.UUID).
	}

	s.auditor.Log(req, result)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
