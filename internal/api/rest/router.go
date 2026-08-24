// Package rest exposes the REST surface: POST /evaluate for real-time
// decisions and /policies for CRUD + versioning/rollback management.
package rest

import (
	"net/http"

	"github.com/kwasi/policy-engine/internal/audit"
	"github.com/kwasi/policy-engine/internal/cache"
	"github.com/kwasi/policy-engine/internal/cedarclient"
	"github.com/kwasi/policy-engine/internal/storage"
)

type Server struct {
	cedar   *cedarclient.Client
	repo    *storage.PolicyRepo
	cache   *cache.PolicyCache
	auditor *audit.Logger
}

func NewServer(cedar *cedarclient.Client, repo *storage.PolicyRepo, c *cache.PolicyCache, auditor *audit.Logger) *Server {
	return &Server{cedar: cedar, repo: repo, cache: c, auditor: auditor}
}

func (s *Server) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", s.handleHealth)
	mux.HandleFunc("POST /evaluate", s.handleEvaluate)
	mux.HandleFunc("GET /policies", s.handleListPolicies)
	mux.HandleFunc("POST /policies", s.handleCreatePolicy)
	mux.HandleFunc("GET /policies/{id}", s.handleGetPolicy)
	mux.HandleFunc("PUT /policies/{id}", s.handleUpdatePolicy)
	mux.HandleFunc("DELETE /policies/{id}", s.handleDeletePolicy)
	mux.HandleFunc("POST /policies/{id}/rollback", s.handleRollbackPolicy)
	return mux
}
