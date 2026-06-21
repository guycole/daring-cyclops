# Quickstart: Ping

## Purpose

Validate the first end-to-end gRPC slice for Daring Cyclops: a CLI `ping`
command that requests server version and current time from the server.

## Prerequisites

- Go 1.26+
- `protoc` installed and available on `PATH`
- Go protobuf plugins installed:
  - `protoc-gen-go`
  - `protoc-gen-go-grpc`
- Working tree checked out on the `001-ping` feature branch

## Contract Reference

- gRPC contract: `specs/001-ping/contracts/ping.proto`
- Message definitions and validation rules: `specs/001-ping/data-model.md`

## Validation Scenario 1: Generate protobuf code

1. Ensure `protoc` and the Go plugins are installed.
2. Run the project's protobuf generation command once it is added during
   implementation.
3. Confirm generated Go stubs are written under `gen/proto/ping/v1/`.

Expected outcome:

- Generated protobuf and gRPC Go code exists and builds cleanly.

## Validation Scenario 2: Server ping unit behavior

1. Run the server unit tests for the ping service.
2. Confirm the service returns a non-empty version string and valid timestamp.

Expected outcome:

- Unit tests pass and verify the service returns both required response fields.

## Validation Scenario 3: In-memory gRPC integration

1. Run the integration test suite using an in-memory gRPC listener.
2. Confirm the client stub receives a successful `PingResponse`.

Expected outcome:

- Integration tests pass without requiring an external listening port.

## Validation Scenario 4: CLI ping success path

1. Start the server locally.
2. Run the CLI client `ping` command against the server address.
3. Inspect the command output.

Expected outcome:

- Output includes a labeled server version value.
- Output includes a labeled server time value.
- Command exits successfully.

## Validation Scenario 5: CLI ping failure path

1. Stop the server or target an unreachable server address.
2. Run the CLI client `ping` command.

Expected outcome:

- The command reports failure clearly.
- The command does not print fabricated version or time values.
- The command exits unsuccessfully.