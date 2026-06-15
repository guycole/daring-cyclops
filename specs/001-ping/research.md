# Research: Ping

## Decision: Use a minimal versioned protobuf contract

### Rationale

The first client/server slice should keep the contract small and explicit. A
versioned package such as `cyclops.ping.v1` allows future protocol evolution
without redefining the first RPC later.

### Alternatives considered

- Unversioned protobuf package: rejected because it makes future evolution and
  compatibility discipline weaker.
- Ad hoc JSON over TCP or HTTP: rejected because the constitution requires gRPC
  and protobuf as the client/server contract.

## Decision: Represent server time with `google.protobuf.Timestamp`

### Rationale

`google.protobuf.Timestamp` is the standard protobuf type for time values and
maps cleanly to Go. It avoids custom parsing rules and preserves precise server
time semantics.

### Alternatives considered

- String timestamp: rejected because it creates custom formatting/parsing rules
  in the contract.
- Integer epoch field: rejected because it is less expressive and less standard
  than the protobuf timestamp type.

## Decision: Source server version from build metadata with a default `dev`

### Rationale

The ping feature only needs a stable, visible server version string. A package
variable with a default and optional build-time override is simple and avoids
premature release-management complexity.

### Alternatives considered

- Read version from a file at runtime: rejected because it adds file management
  complexity for the first slice.
- Hardcode a constant in the ping service: rejected because version ownership is
  cleaner in a dedicated build-info package.

## Decision: Use separate CLI and server binaries in a single Go module

### Rationale

The constitution requires a Go CLI client and a Go backend server. Separate
`cmd/` entrypoints with shared internal packages and generated protobuf code are
the simplest layout that respects that boundary.

### Alternatives considered

- Single binary for both roles: rejected because it blurs the required client vs
  server separation.
- Multi-module repository: rejected because it adds unnecessary packaging
  overhead for an initial ping feature.

## Decision: Plan unit tests plus `bufconn` and CLI integration tests

### Rationale

The first useful confidence layer is end-to-end client/server behavior. Unit
tests cover formatting and service logic, while `bufconn` covers gRPC behavior
without requiring real sockets.

### Alternatives considered

- Socket-only integration tests: rejected because they are slower and more
  brittle for the first RPC.
- Unit tests only: rejected because they would not prove the first end-to-end
  gRPC slice works.

## Decision: Treat protobuf generation tooling as a required local build dependency

### Rationale

The feature depends on generated protobuf code, so the implementation and
validation workflow must explicitly require `protoc` and the Go protobuf
plugins.

### Alternatives considered

- Delay code generation decisions: rejected because the ping feature directly
  depends on the protobuf contract.
- Check in hand-written stubs: rejected because protobuf generation should
  remain the source of truth.