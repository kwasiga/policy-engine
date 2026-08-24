// Package audit decouples audit-log writes from the request path. /evaluate
// must return the decision immediately; recording it to Postgres happens
// on a buffered worker so a slow or momentarily-down database never adds
// latency to (or fails) an authorization decision.
package audit

import (
	"context"
	"log/slog"

	"github.com/kwasi/policy-engine/internal/domain"
	"github.com/kwasi/policy-engine/internal/storage"
)

type record struct {
	req    domain.DecisionRequest
	result domain.DecisionResult
}

type Logger struct {
	repo  *storage.AuditRepo
	queue chan record
}

// New starts a background worker draining a bounded queue into Postgres.
// bufferSize controls how many decisions can be in flight before Log()
// starts blocking callers — size it to your expected burst, not steady
// state, since audit writes should never be the bottleneck.
func New(repo *storage.AuditRepo, bufferSize int) *Logger {
	l := &Logger{repo: repo, queue: make(chan record, bufferSize)}
	go l.run()
	return l
}

// Log enqueues a decision for durable audit storage. Never blocks on
// Postgres itself; only blocks if the buffer is full, which should only
// happen under sustained overload.
func (l *Logger) Log(req domain.DecisionRequest, result domain.DecisionResult) {
	l.queue <- record{req: req, result: result}
}

func (l *Logger) run() {
	for rec := range l.queue {
		if err := l.repo.Record(context.Background(), rec.req, rec.result); err != nil {
			// TODO: route to a dead-letter table or metrics counter instead
			// of just logging — a silently dropped audit row is a
			// compliance gap, not just an error.
			slog.Error("audit log write failed", "error", err)
		}
	}
}
