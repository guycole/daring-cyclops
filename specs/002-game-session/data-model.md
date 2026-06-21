# Data Model: Game Session Lifecycle

## Entities

### GameSession

The top-level entity representing one independent game instance on the server.

| Field              | Type           | Notes |
|--------------------|----------------|-------|
| `id`               | string (UUID)  | Server-generated at creation; opaque to clients |
| `state`            | SessionState   | WAITING → ACTIVE → ENDED; transitions are one-way |
| `config`           | SessionConfig  | Set once by the first player; immutable after creation |
| `created_at`       | Timestamp      | Wall-clock time of session creation |
| `stardate`         | float64        | Current simulation time; advances per tick |
| `federation_slots` | []ShipSlot     | Fixed 9 slots for Federation ships |
| `empire_slots`     | []ShipSlot     | Fixed 9 slots for Empire ships |
| `kill_queue`       | []KillQueueEntry | Max 10 entries; FIFO |

**Lifecycle constraints**:
- A session is created in `WAITING` state when the first player calls `CreateSession`.
- It transitions to `ACTIVE` when the first `ActivateShip` call succeeds.
- It transitions to `ENDED` when victory conditions are met or manually closed.
- No new `JoinSession` or `ActivateShip` calls are accepted in `ENDED` state.

**State machine**:
```
[WAITING] --first ActivateShip--> [ACTIVE] --victory condition--> [ENDED]
```

---

### SessionConfig

Immutable options set by the first player during session creation. Stored inside
the parent `GameSession`.

| Field                  | Type     | Notes |
|------------------------|----------|-------|
| `game_mode`            | GameMode | REGULAR or TOURNAMENT |
| `tournament_name`      | string   | Required when game_mode = TOURNAMENT; used to derive seed |
| `tournament_seed`      | int64    | Derived from tournament_name via FNV-64a hash; 0 for REGULAR |
| `romulans_enabled`     | bool     | Default: true |
| `black_holes_enabled`  | bool     | Default: false |
| `respawn_wait_ticks`   | int32    | Simulation ticks before a killed player may re-enter; default: 120 |
| `balance_threshold`    | int32    | Max allowed side-size difference before forced balance; default: 3 |

---

### ShipSlot

One of 18 fixed berths in a session (9 Federation, 9 Empire). A ship slot
holds the per-player state for that berth.

| Field          | Type      | Notes |
|----------------|-----------|-------|
| `slot_index`   | int32     | 0–8 within its side |
| `side`         | Side      | FEDERATION or EMPIRE |
| `ship_name`    | string    | Canonical ship name (e.g., "Lexington"); fixed per slot |
| `player_name`  | string    | Display name of the occupying player; empty if unoccupied |
| `player_state` | PlayerState | UNOCCUPIED, LOBBY, ACTIVE, KILL_QUEUE |
| `side_history` | []Side    | Not persisted; tracked in kill queue for respawn routing |

**Availability rule**: A slot is available for a new player when `player_state`
is `UNOCCUPIED`. A returning player from the kill queue is routed to any open
slot on their previous side.

---

### Player (session-scoped)

A transient entity representing a connected player within a session. Not
persisted beyond the session lifetime.

| Field          | Type        | Notes |
|----------------|-------------|-------|
| `player_name`  | string      | Provided by the player at join time |
| `side`         | Side        | UNKNOWN until side selection; FEDERATION or EMPIRE after |
| `player_state` | PlayerState | LOBBY, ACTIVE, or KILL_QUEUE |
| `ship_name`    | string      | Assigned at ActivateShip; empty while in lobby |
| `prior_side`   | Side        | Remembered across deaths; used to auto-reassign on respawn |

---

### KillQueueEntry

Represents a destroyed player waiting to respawn. The kill queue is a server-
managed list on the parent `GameSession`; max 10 entries.

| Field           | Type      | Notes |
|-----------------|-----------|-------|
| `player_name`   | string    | Display name |
| `side`          | Side      | The side the player will return to |
| `ship_name`     | string    | The ship that was destroyed |
| `killed_at`     | float64   | Simulation stardate at time of death |
| `eligible_at`   | float64   | `killed_at + config.respawn_wait_ticks`; player may re-enter after this stardate |

**Overflow rule**: If the kill queue already has 10 entries when a new player
dies, the oldest entry is evicted (the longest-waiting player becomes eligible
immediately) to make room.

---

## Enumerations

### SessionState

| Value     | Meaning |
|-----------|---------|
| `WAITING` | No ship has been activated; game options have been set |
| `ACTIVE`  | At least one ship is active in the session |
| `ENDED`   | Victory condition met; no further gameplay accepted |

### GameMode

| Value        | Meaning |
|--------------|---------|
| `REGULAR`    | Standard randomised game |
| `TOURNAMENT` | Seeded game; same seed produces same initial galaxy state |

### Side

| Value         | Meaning |
|---------------|---------|
| `UNKNOWN`     | Player has not yet completed side selection |
| `FEDERATION`  | Human / Federation side |
| `EMPIRE`      | Klingon / Empire side |

### PlayerState

| Value        | Meaning |
|--------------|---------|
| `UNOCCUPIED` | Slot has no player |
| `LOBBY`      | Player connected and side-selected; has not activated a ship |
| `ACTIVE`     | Player's ship is live in the game |
| `KILL_QUEUE` | Player's ship was destroyed; waiting for respawn countdown |

---

## Key Relationships

```
GameSession 1 --has--> 1 SessionConfig
GameSession 1 --has--> 9 ShipSlot (Federation)
GameSession 1 --has--> 9 ShipSlot (Empire)
GameSession 1 --has--> 0..10 KillQueueEntry
ShipSlot 1 --may hold--> 0..1 Player
KillQueueEntry 0..1 --references--> 1 Player (by player_name + side)
```

---

## Validation Rules

- `session.state` transitions are one-way: WAITING → ACTIVE → ENDED.
- `config.game_mode == TOURNAMENT` requires `config.tournament_name` to be
  non-empty.
- `config.tournament_seed` MUST be derived from `tournament_name` server-side;
  it is not accepted directly from the client.
- Side counts per session MUST NOT exceed 9 active ships (PlayerState == ACTIVE)
  per side.
- Kill queue length MUST NOT exceed 10; overflow evicts the oldest entry.
- `player.player_state` may only advance LOBBY → ACTIVE or ACTIVE → KILL_QUEUE
  or KILL_QUEUE → LOBBY; it cannot go directly from LOBBY to KILL_QUEUE.
- A session in ENDED state MUST reject all `ActivateShip`, `LobbyCommand`, and
  `JoinSession` requests.
