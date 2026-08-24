#!/usr/bin/env bash
# Generates internal/api/grpc/pb from proto/policy_engine.proto.
# Requires protoc, protoc-gen-go, protoc-gen-go-grpc on PATH.
set -euo pipefail

mkdir -p internal/api/grpc/pb
protoc \
  --go_out=internal/api/grpc/pb --go_opt=paths=source_relative \
  --go-grpc_out=internal/api/grpc/pb --go-grpc_opt=paths=source_relative \
  proto/policy_engine.proto
