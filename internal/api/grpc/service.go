package grpc

// TODO: implement pb.PolicyEngineServer methods (Evaluate, CreatePolicy,
// GetPolicy, ListPolicies, UpdatePolicy, DeletePolicy, RollbackPolicy) once
// the generated pb package exists. Each should delegate to the same
// s.cedar / s.repo / s.cache / s.auditor calls as internal/api/rest, so
// REST and gRPC stay behaviorally identical — consider extracting a shared
// internal/service layer if that duplication grows.
