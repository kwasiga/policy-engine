# Builds only the Go API. The cedar-agent sidecar is a separate published
# image (see docker-compose.yml and sidecar/cedar-agent/README.md) — it is
# not built from this Dockerfile.
FROM golang:1.22-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /out/server ./cmd/server

FROM gcr.io/distroless/static-debian12
COPY --from=build /out/server /server
EXPOSE 8080 9090
ENTRYPOINT ["/server"]
