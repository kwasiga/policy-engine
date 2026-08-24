package rest

import "net/http"

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	// TODO: also ping Postgres and the cedar-agent sidecar's /health.
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}`))
}
