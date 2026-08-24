// Package grpc hosts the gRPC surface (see proto/policy_engine.proto),
// generated code lands in internal/api/grpc/pb via:
//   protoc --go_out=. --go-grpc_out=. proto/policy_engine.proto
// Not checked in yet — run scripts/gen_proto.sh once the proto stabilizes.
package grpc

import (
	"google.golang.org/grpc"

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

// Register wires this Server into a *grpc.Server once the generated
// pb.RegisterPolicyEngineServer stub exists.
func (s *Server) Register(grpcServer *grpc.Server) {
	// TODO: pb.RegisterPolicyEngineServer(grpcServer, s)
}
