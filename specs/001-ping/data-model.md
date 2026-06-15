# Data Model: Ping

## Overview

This feature introduces a minimal request/response contract rather than durable
domain storage. The primary entities are transport messages exchanged between the
CLI client and gRPC server.

## Entities

### PingRequest

- **Purpose**: Invokes the server ping operation.
- **Fields**: None for the initial version.
- **Validation rules**:
  - Request must be syntactically valid for the protobuf contract.
  - Request carries no user-supplied payload in v1.
- **State transitions**:
  - Created by client
  - Sent to server
  - Consumed by ping service

### PingResponse

- **Purpose**: Returns minimal server identity and time information.
- **Fields**:
  - `server_version`: string
  - `server_time`: timestamp
- **Validation rules**:
  - `server_version` must be non-empty in successful responses.
  - `server_time` must be present in successful responses.
  - `server_time` represents the server's current time.
- **State transitions**:
  - Constructed by server ping service
  - Returned over gRPC to client
  - Rendered as labeled CLI output

## Relationships

- A `PingRequest` produces exactly one `PingResponse` on successful unary RPC
  execution.
- A failed request produces an RPC error instead of a successful `PingResponse`.

## Derived View Data

### Ping Output

- **Purpose**: Human-readable CLI rendering of the ping response.
- **Derived from**: `PingResponse`
- **Display requirements**:
  - Show server version label and value
  - Show server time label and value
  - Omit raw protobuf formatting from normal CLI output