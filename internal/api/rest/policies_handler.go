package rest

import "net/http"

// CRUD + versioning/rollback for /policies. Every mutating handler must
// call s.cache.Invalidate(ctx) after a successful write so the cedar-agent
// sidecar picks up the change immediately.

func (s *Server) handleListPolicies(w http.ResponseWriter, r *http.Request) {
	// TODO: s.repo.List(ctx, statusFilter); write JSON array.
	http.Error(w, "not implemented", http.StatusNotImplemented)
}

func (s *Server) handleCreatePolicy(w http.ResponseWriter, r *http.Request) {
	// TODO: decode body, s.repo.Create(...), s.cache.Invalidate(ctx).
	http.Error(w, "not implemented", http.StatusNotImplemented)
}

func (s *Server) handleGetPolicy(w http.ResponseWriter, r *http.Request) {
	// TODO: parse {id}, s.repo.Get(ctx, id).
	http.Error(w, "not implemented", http.StatusNotImplemented)
}

func (s *Server) handleUpdatePolicy(w http.ResponseWriter, r *http.Request) {
	// TODO: decode body, s.repo.Update(...), s.cache.Invalidate(ctx).
	http.Error(w, "not implemented", http.StatusNotImplemented)
}

func (s *Server) handleDeletePolicy(w http.ResponseWriter, r *http.Request) {
	// TODO: s.repo.Delete (soft-delete to 'archived'), s.cache.Invalidate(ctx).
	http.Error(w, "not implemented", http.StatusNotImplemented)
}

func (s *Server) handleRollbackPolicy(w http.ResponseWriter, r *http.Request) {
	// TODO: decode {"to_version": n}, s.repo.Rollback(...), s.cache.Invalidate(ctx).
	http.Error(w, "not implemented", http.StatusNotImplemented)
}
