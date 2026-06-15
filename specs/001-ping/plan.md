# Implementation Plan: Ping

**Branch**: `001-ping` | **Date**: 2026-06-15 | **Spec**: `/specs/001-ping/spec.md`

**Input**: Feature specification from `/specs/001-ping/spec.md`

## Summary

Implement the first end-to-end gRPC slice for Daring Cyclops by adding a unary
`Ping` RPC to the Go server and a `ping` command to the Go CLI client. The
server returns a version string and current server time, and the client renders
 those values in human-readable CLI output while reporting connection failures
 clearly.

## Technical Context

**Language/Version**: Go 1.26.3

**Primary Dependencies**: `google.golang.org/grpc`, `google.golang.org/protobuf`,
`google.golang.org/grpc/test/bufconn` for in-memory integration tests

**Storage**: N/A for this feature

**Testing**: `go test` for unit and integration tests

**Target Platform**: Linux server runtime; macOS/Linux developer and CLI runtime

**Project Type**: Separate CLI client and gRPC server binaries in one Go module

**Performance Goals**: Ping completes in under 1 second on a local or LAN-hosted
server under normal conditions

**Constraints**: gRPC with Protocol Buffers as source of truth; no authentication
for this feature; no persistence; server is authoritative for returned time and
version data; protobuf code generation requires `protoc`, `protoc-gen-go`, and
`protoc-gen-go-grpc`

**Scale/Scope**: One unary RPC, one CLI command, one shared protobuf contract,
basic failure handling, and first-slice client/server integration only

## Constitution Check

*GATE: Must pass before Phase 0 research. Re-check after Phase 1 design.*

### Pre-Design Gate Review

- **Go client/server shape**: PASS. The plan preserves the required Go CLI
  client and Go backend server split.
- **gRPC and protobuf contract**: PASS. The feature is centered on a protobuf
  contract and gRPC transport.
- **Server authority**: PASS. The server remains authoritative for returned
  version and server time values.
- **Linux hosting compatibility**: PASS. Nothing in the design conflicts with a
  Linux-hosted server.
- **Replaceable auth and persistence**: PASS. The feature is read-only and adds
  no fixed auth or storage dependency.

### Post-Design Gate Review

- **Interface contract versioning**: PASS. Contract is defined under a `v1`
  namespace to support future evolution.
- **Simulation and game-instance isolation**: PASS. The ping feature does not
  bypass or alter simulation or per-instance authority rules.
- **Operational reproducibility**: PASS. Tooling requirements and validation
  commands are documented in `quickstart.md`.

## Project Structure

### Documentation (this feature)

```text
specs/001-ping/
├── plan.md
├── research.md
├── data-model.md
├── quickstart.md
├── contracts/
│   └── ping.proto
└── tasks.md
```

### Source Code (repository root)

```text
cmd/
├── cyclops/
│   └── main.go
└── cyclopsd/
    └── main.go

internal/
├── buildinfo/
│   └── version.go
├── client/
│   └── ping/
│       └── run.go
└── server/
    └── ping/
        └── service.go

proto/
└── ping/v1/
    └── ping.proto

gen/
└── proto/
    └── ping/v1/

tests/
├── integration/
│   ├── grpc_ping_test.go
│   └── cli_ping_test.go
└── unit/
    ├── ping_service_test.go
    └── ping_output_test.go
```

**Structure Decision**: Use a single Go module with separate `cmd/` binaries for
the CLI client and server, shared generated protobuf code under `gen/`, and
feature-specific client/server logic under `internal/`.

## Phase 0: Research Outcomes

- Use `google.protobuf.Timestamp` for server time instead of a custom string or
  integer encoding.
- Source the server version from a small build-info package with a default value
  such as `dev`, overridable at build time.
- Favor an empty `PingRequest` and a minimal `PingResponse` containing only the
  required fields.
- Prefer in-memory `bufconn` integration testing for the first gRPC slice.

## Complexity Tracking

No constitution violations or complexity exemptions are required for this
feature.
