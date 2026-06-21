# Quickstart: Game Session Lifecycle

## Prerequisites

- Go 1.26.3 installed
- `protoc`, `protoc-gen-go`, `protoc-gen-go-grpc` on `$PATH` (same as 001-ping)
- Server binary built: `go build ./cmd/cyclopsd/`
- Client binary built: `go build ./cmd/cyclops/`

---

## 1. Generate the session protobuf stubs

After copying the contract to `proto/session/v1/session.proto`, regenerate:

```bash
make proto
```

The generated Go files appear in `gen/proto/session/v1/`.

---

## 2. Start the server

```bash
./cyclopsd
# Server listens on 127.0.0.1:50051 by default
```

---

## 3. Validate: First player creates a session (User Story 1)

```bash
# Create a regular game session with Romulans enabled
./cyclops session create --player alice --mode regular --romulans

# Expected output:
# Session created: <uuid>
# Mode: Regular
# Romulans: enabled
# Black holes: disabled
```

To validate tournament mode:

```bash
./cyclops session create --player alice --mode tournament --name "GrandPrix2026"

# Expected output:
# Session created: <uuid>
# Mode: Tournament (GrandPrix2026)
# Romulans: enabled
# Black holes: disabled
```

---

## 4. Validate: Second player joins and selects a side (User Story 2)

```bash
# List available sessions (find the session_id from step 3)
./cyclops session list

# Join, requesting the Federation side
./cyclops session join --session <uuid> --player bob --side federation

# Expected output:
# Joined session <uuid>
# Side: Federation
# There are Romulans in this game.
```

To test balance enforcement, fill one side to capacity (9 ships) and attempt to
join it with a 10th player:

```bash
./cyclops session join --session <uuid> --player tenth --side federation

# Expected output includes something like:
# Federation is at capacity. You have been assigned to the Empire.
```

---

## 5. Validate: Pre-game lobby commands (User Story 3)

```bash
# Run USERS in the lobby
./cyclops session lobby --session <uuid> --player bob --command users

# Expected: list of currently connected players, their sides, and states

# Run TIME in the lobby
./cyclops session lobby --session <uuid> --player bob --command time

# Expected: current simulation stardate

# Run HELP in the lobby
./cyclops session lobby --session <uuid> --player bob --command help

# Expected: list of valid lobby commands

# Run an invalid command (should be rejected)
./cyclops session lobby --session <uuid> --player bob --command warp

# Expected: error message "This command unavailable in Pre-game"
```

---

## 6. Validate: Ship activation (User Story 4)

```bash
./cyclops session activate --session <uuid> --player bob

# Expected output:
# Ship assigned: Lexington (Federation)
# Session state: Active
```

Attempt activation when the side is full (9 ships active):

```bash
# Expected output:
# All Federation ships are currently in use. Please wait or try again later.
```

---

## 7. Validate: Kill queue countdown (User Story 5)

Use the integration test harness to destroy a player's ship, then check status:

```bash
go test ./tests/integration/... -run TestKillQueueRespawn -v

# The test confirms:
# - Player enters kill queue after destruction
# - GetKillQueueStatus returns ticks_remaining > 0
# - Player cannot activate a ship while in the kill queue
# - After respawn_wait_ticks elapse, player transitions back to lobby
```

For manual verification with respawn_wait_ticks = 0 (instant respawn):

```bash
./cyclops session kill-queue --session <uuid> --player bob
# Expected: in_kill_queue: false (already eligible)
```

---

## 8. Validate: Session end (User Story 6)

```bash
go test ./tests/integration/... -run TestSessionEndVictory -v

# The test confirms:
# - Session transitions to ENDED when one side is eliminated
# - GetSessionStatus returns SESSION_STATE_ENDED
# - All subsequent gameplay commands return an error indicating session is over
```

---

## 9. Validate: Concurrent sessions (SC-007)

```bash
go test ./tests/integration/... -run TestConcurrentSessions -v

# The test creates two sessions, runs actions in both, and confirms:
# - State changes in session A do not appear in session B
# - Both sessions respond to GetSessionStatus independently
```

---

## Running all session tests

```bash
go test ./... -run TestSession
```

For the full test suite:

```bash
go test ./...
```
