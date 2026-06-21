# Research: Game Session Lifecycle

## Decision: Model the session as a server-side state machine with three states

### Rationale

The original DECWAR source (SETUP.FOR, DECWAR.FOR) demonstrates a clear lifecycle:
a game starts in a setup/waiting phase (first player sets options), transitions to
active once a ship is activated, and ends when one side is eliminated. Encoding
this explicitly in the server as a state machine (WAITING → ACTIVE → ENDED) makes
transitions auditable, testable, and resistant to race conditions in a concurrent
Go server.

### Alternatives considered

- Boolean `active` flag only: rejected because it cannot represent the ENDED
  state distinctly from WAITING, which leads to ambiguous handling of
  post-session commands.
- Event-sourced state: rejected as premature for this feature; the constitution
  defers persistence architecture decisions, and event sourcing adds complexity
  that is not justified for the first lifecycle slice.

---

## Decision: Unary gRPC RPCs for all session lifecycle operations in this feature

### Rationale

Session lifecycle operations (create, join, lobby commands, activate, status
queries) are discrete request/response interactions. Bidirectional streaming is
appropriate for real-time gameplay (movement, combat) but adds complexity that
is not needed for the session setup and lobby phase. Unary RPCs compose cleanly
with the existing ping service and allow straightforward client implementation.
Streaming can be layered on in a later gameplay spec without breaking the session
lifecycle contract.

### Alternatives considered

- Bidirectional streaming for the entire session: rejected because it entangles
  session setup complexity with streaming lifecycle management before the
  game-loop spec exists.
- Server streaming for lobby commands: rejected because lobby responses are
  discrete answers, not ongoing push streams.

---

## Decision: Session identity uses a server-generated UUID string

### Rationale

UUIDs are the standard stateless opaque identifier in Go services. Clients pass
the session ID in subsequent calls, which allows the server to route requests to
the correct isolated game instance. UUID v4 generation is available in the
standard library (`crypto/rand`) without additional dependencies.

### Alternatives considered

- Integer sequence: rejected because sequences require shared state and are less
  safe in a concurrent context without coordination.
- Human-readable names: rejected because they create collision risk and suggest
  implicit state that the server must manage.

---

## Decision: Session discovery via ListSessions RPC; session creation via CreateSession

### Rationale

In the original DECWAR, all players connected to the same shared time-sharing
process, so discovery was implicit. In a network service hosting multiple
concurrent sessions, clients need a way to enumerate available sessions before
joining. A `ListSessions` RPC returning session summaries (ID, state, player
counts per side) is the minimal interface that enables this without requiring
out-of-band coordination. `CreateSession` is a separate call that returns the
new session ID and the first-player configuration prompt state.

### Alternatives considered

- Implicit session auto-creation on JoinSession: rejected because it makes it
  impossible to distinguish a first player from a joiner, and first players need
  to supply game options before the session is usable.
- Discovery via environment variable or static config: rejected because it
  prevents the server from hosting multiple concurrent sessions correctly.

---

## Decision: Kill queue countdown is a configurable session parameter in simulation ticks

### Rationale

The original DECWAR had a `KWAIT` constant (set to 0 ms in the utexas build,
but ~2 minutes in production). Making it a session-scoped configuration value
(defaulting to 120 simulation ticks) respects the constitution's simulation-time
model and allows tests to set it to 0 for fast iteration.

### Alternatives considered

- Wall-clock timer: rejected because the constitution mandates simulation time
  governs gameplay progression, and using wall-clock time for respawn creates
  inconsistency between the timer and game speed.
- Hard-coded constant: rejected because tests would need to wait real time.

---

## Decision: Pre-game lobby is a distinct named player state, not a boolean flag

### Rationale

The original DECWAR distinguishes the `PG>` pre-game prompt from the in-game
`Command:` prompt by checking `who .eq. 0` (player has not yet activated a ship).
In a gRPC server with concurrent players, encoding lobby vs active as a named
enum value (`LOBBY`, `ACTIVE`, `KILL_QUEUE`) makes routing of command dispatch
explicit, testable, and extensible to future states (e.g., spectator mode).

### Alternatives considered

- Boolean `inLobby` flag: rejected because adding future states requires schema
  changes, and the three-state model is already known from the spec.

---

## Decision: Side balance enforcement uses a configurable imbalance threshold

### Rationale

From DECNWS.RNO v2.0: "If the teams are too skewed you're forced to join the
smaller team." The exact threshold is not stated in the source. A configurable
threshold (default: force join when difference ≥ 3 ships) lets the server enforce
balance without hard-coding a constant that may need tuning.

### Alternatives considered

- Force join whenever any imbalance exists: rejected because a difference of 1
  ship is acceptable early in a game and forcing defection at that point is
  hostile UX.
- No enforcement: rejected because the original game explicitly enforced balance
  and the spec requires it.

---

## Decision: Tournament seed is an int64, derived from the supplied name string via hashing

### Rationale

The original DECWAR uses `setran(iabs(tknlst(i)))` — it converts the token value
(the tournament name) directly to an integer seed for the random generator. In
Go, deriving a deterministic int64 seed from a string via `fnv.New64a().Sum64()`
is idiomatic and avoids forcing callers to supply raw integers.

### Alternatives considered

- Client supplies a raw int64: rejected because it is less usable from a CLI
  (users name tournaments, they do not generate integers).
- UUID-based seed: rejected because UUIDs are not reproducible by the player.
