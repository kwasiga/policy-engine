#!/usr/bin/env bash
# Applies migrations/*.sql in order against DATABASE_URL. Requires `psql`.
set -euo pipefail

DATABASE_URL="${DATABASE_URL:-postgres://policy_engine:policy_engine@localhost:5432/policy_engine?sslmode=disable}"

for f in "$(dirname "$0")"/../migrations/*.sql; do
  echo "applying $f"
  psql "$DATABASE_URL" -f "$f"
done
