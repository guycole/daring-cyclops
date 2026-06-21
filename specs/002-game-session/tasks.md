# Tasks: Game Session Lifecycle

**Input**: Design documents from `/specs/002-game-session/`

**Prerequisites**: plan.md (required), spec.md (required for user stories), research.md, data-model.md, contracts/session.proto, quickstart.md

**Tests**: No dedicated TDD test tasks. The Polish phase includes unit and integration tests as implementation verification. The spec did not explicitly request a TDD workflow.

**Organization**: Tasks are grouped by user story to enable independent implementation and testing of each story.

## Format: `[ID] [P?] [Story] Description`

- **[P]**: Can run in parallel (different files, no dependencies on incomplete tasks)
- **[Story]**: Which user story this task belongs to (US1–US6)
- Exact file paths are included in all descriptions

---

## Phase 1: Setup (Shared Infrastructure)

**Purpose**: Copy the protobuf contract into the live source tree and regenerate stubs.

- [ ] T001 Copy specs/002-game-session/contracts/session.proto to proto/session/v1/session.proto
- [ ] T002 [P] Update the proto Makefile target to include proto/session/v1/session.proto alongside the existing ping target in Makefile
- [ ] T003 Regenerate all protobuf stubs by running `make proto`; confirm gen/proto/session/v1/ contains session.pb.go and session_grpc.pb.go

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: Core server-side data structures and the service wiring that all user stories depend on.

**⚠️ CRITICAL**: No user story work can begin until this phase is complete.

- [ ] T004 Implement GameSession struct with SessionState (WAITING/ACTIVE/ENDED), one-way transition guards (WAITING→ACTIVE→ENDED only), federation_slots and empire_slots as fixed [9]ShipSlot arrays, kill_queue as []KillQueueEntry, and current_stardate float64 in internal/server/session/session.go
- [ ] T005 Add SessionConfig struct with all fields and apply-defaults function (romulans_enabled=true, black_holes_enabled=false, respawn_wait_ticks=120, balance_threshold=3); add DeriveTournamentSeed(name string) int64 using hash/fnv FNV-64a in internal/server/session/session.go
- [ ] T006 Add ShipSlot struct (slot_index, side, ship_name, player_name, PlayerState) with 9 canonical Federation ship names and 9 canonical Empire ship names as package-level constants; add PlayerState type (UNOCCUPIED/LOBBY/ACTIVE/KILL_QUEUE) in internal/server/session/session.go
- [ ] T007 Add KillQueueEntry struct (player_name, side, ship_name, killed_at_stardate, eligible_at_stardate) and KillQueue methods: Enqueue (enforce max-10 by evicting the oldest entry on overflow), Dequeue (remove by player_name), Lookup (find by player_name), EligiblePlayers (return entries where eligible_at ≤ given stardate) in internal/server/session/session.go
- [ ] T008 Implement Registry type: sync.RWMutex-guarded map[string]*GameSession with Create (generates UUID v4 session ID via crypto/rand), Get, List (optionally include ENDED), and Delete methods in internal/server/session/service.go
- [ ] T009 Implement SessionService struct (embeds Registry, implements sessionv1.SessionServiceServer) and register it with the gRPC server in cmd/cyclopsd/main.go
- [ ] T010 [P] Add `session` top-level sub-command routing in cmd/cyclops/main.go: parse args[0]=="session" and dispatch to sub-commands (create, list, join, lobby, activate, kill-queue); print usage if no sub-command given

**Checkpoint**: Foundation ready — all user story phases may now begin in priority order.

---

## Phase 3: User Story 1 — First Player Creates and Configures a Game Session (Priority: P1) 🎯 MVP

**Goal**: A single player can create a session, set game options, and receive a session ID that other players can use to join.

**Independent Test**: Start the server. Run `./cyclops session create --player alice --mode regular --romulans`. Confirm the CLI prints a UUID session ID and the confirmed config. Run `./cyclops session list` and confirm the new session appears in state WAITING.

### Implementation for User Story 1

- [ ] T011 [US1] Implement CreateSession RPC handler: validate config (require non-empty tournament_name when game_mode=TOURNAMENT; reject raw tournament_seed from client), call ApplyDefaults, call DeriveTournamentSeed, create session via Registry.Create, return session_id and confirmed config in internal/server/session/service.go
- [ ] T012 [US1] Implement ListSessions RPC handler: call Registry.List(req.IncludeEnded), build SessionSummary for each session (count ACTIVE and LOBBY slots per side), return ListSessionsResponse in internal/server/session/service.go
- [ ] T013 [US1] Implement GetSessionStatus RPC handler: look up session by ID (return gRPC NotFound if absent), marshal all federation_slots and empire_slots, kill_queue, and current_stardate into SessionStatus in internal/server/session/service.go
- [ ] T014 [P] [US1] Implement `session create` and `session list` CLI sub-commands in internal/client/session/run.go: parse flags (--player, --mode, --tournament-name, --romulans, --black-holes), dial gRPC, call CreateSession or ListSessions, pass response to output.go formatters; return non-zero exit code on RPC error
- [ ] T015 [P] [US1] Implement CreateSessionResponse and SessionSummary output formatters in internal/client/session/output.go: print session ID, mode, romulans/black-holes flags for create; print tabular session list (ID, state, Fed active/lobby, Emp active/lobby) for list

**Checkpoint**: User Story 1 fully functional and independently demonstrable.

---

## Phase 4: User Story 2 — Additional Players Join a Session and Select a Side (Priority: P2)

**Goal**: A second player can look up the session ID, choose Federation or Empire, and be placed in the lobby state — with balance enforcement and capacity checks.

**Independent Test**: Create a session as alice. Run `./cyclops session join --session <id> --player bob --side federation`. Confirm the CLI shows "Joined session, Side: Federation". Fill Federation to 9 players and attempt a 10th — confirm redirect or capacity error.

### Implementation for User Story 2

- [ ] T016 [US2] Implement JoinSession RPC handler: look up session (NotFound if absent, FailedPrecondition if ENDED), detect returning player via kill queue prior_side, enforce 9-slot cap per side (Unavailable if full and other side also full), apply balance_threshold redirect (force player to smaller side if |fed_count - emp_count| ≥ balance_threshold), place player in LOBBY state (set ShipSlot.player_name + LOBBY for a newly-allocated tracking slot), return assigned_side and redirect metadata in internal/server/session/service.go
- [ ] T017 [US2] Add Player tracking to GameSession: map[playerName]*PlayerInfo (side, ship_slot_index, prior_side), CountActive(side) and CountLobby(side) helper methods, and LookupPriorSide(playerName) that checks kill queue and player map in internal/server/session/session.go
- [ ] T018 [P] [US2] Implement `session join` CLI sub-command in internal/client/session/run.go: parse flags (--session, --player, --side), dial gRPC, call JoinSession, pass response to output.go; print redirect notice if side_redirected is true
- [ ] T019 [P] [US2] Implement JoinSessionResponse output formatter in internal/client/session/output.go: print assigned side, redirect reason if redirected, romulans/black-holes presence announcement

**Checkpoint**: User Story 2 fully functional and independently demonstrable.

---

## Phase 5: User Story 3 — Player Uses the Pre-Game Lobby (Priority: P3)

**Goal**: A player in the lobby can run HELP, NEWS, USERS, POINTS, SUMMARY, TIME, GRIPE, and QUIT — and the server rejects any other command.

**Independent Test**: Join a session as bob (lobby state). Run each of the 8 lobby commands via `./cyclops session lobby`. Confirm each returns appropriate output. Run an invalid command and confirm rejection with "This command unavailable in Pre-game".

### Implementation for User Story 3

- [ ] T020 [US3] Implement LobbyCommand RPC handler: look up session and player, return FailedPrecondition if player is not in PLAYER_STATE_LOBBY, dispatch to per-command handler via switch (HELP, NEWS, USERS, POINTS, SUMMARY, TIME, GRIPE, QUIT), return output string and session_ended=true for QUIT; return InvalidArgument for LOBBY_COMMAND_UNSPECIFIED in internal/server/session/service.go
- [ ] T021 [US3] Implement lobby command handler methods on GameSession in internal/server/session/session.go: LobbyHelp() returns command list, LobbyUsers() returns player name/side/state table, LobbyTime() returns current_stardate as a formatted string, LobbyPoints() returns stub "no points yet", LobbyNews() returns stub "no news", LobbySummary() returns ship roster summary, LobbyGripe(text) records text to a gripe slice on the session (deferred storage), LobbyQuit(playerName) removes player from tracking map and clears their ship slot
- [ ] T022 [P] [US3] Implement `session lobby` CLI sub-command in internal/client/session/run.go: parse flags (--session, --player, --command, --arg), dial gRPC, call LobbyCommand, print output; if session_ended exit cleanly with message "You have left the game"
- [ ] T023 [P] [US3] Implement LobbyCommandResponse output formatter in internal/client/session/output.go: print command output verbatim; print "You have left the game." suffix when session_ended is true

**Checkpoint**: User Story 3 fully functional and independently demonstrable.

---

## Phase 6: User Story 4 — Player Activates a Ship to Enter Gameplay (Priority: P4)

**Goal**: A lobby player can run ACTIVATE (via ActivateShip RPC), be assigned an available ship on their side, and transition to active state — transitioning the session to ACTIVE on first activation.

**Independent Test**: Join as bob (lobby), call `./cyclops session activate --session <id> --player bob`. Confirm the CLI shows ship name, side, and session state. Call `./cyclops session status --session <id>` and confirm bob's slot is now ACTIVE.

### Implementation for User Story 4

- [ ] T024 [US4] Implement ActivateShip RPC handler: look up session and player, return FailedPrecondition if player is not LOBBY or session is ENDED, call AssignShip to find and claim a slot, transition session to ACTIVE if this is the first ACTIVE ship (session.state == WAITING), enforce 9-ship active cap per side (return Unavailable with message if full), return ship_name, side, session_state in internal/server/session/service.go
- [ ] T025 [US4] Implement AssignShip(playerName string, side Side, preferred string) method on GameSession: scan the side's 9 ShipSlots for UNOCCUPIED state (prefer the slot matching preferred ship name if non-empty), set slot.player_name = playerName and slot.player_state = ACTIVE, update the player tracking map, return the assigned ship name or error if all 9 slots are already ACTIVE in internal/server/session/session.go
- [ ] T026 [P] [US4] Implement `session activate` CLI sub-command in internal/client/session/run.go: parse flags (--session, --player, --ship), dial gRPC, call ActivateShip, pass response to output formatter
- [ ] T027 [P] [US4] Implement ActivateShipResponse output formatter in internal/client/session/output.go: print "Ship assigned: <name> (<side>)" and "Session state: <state>"

**Checkpoint**: User Story 4 fully functional and independently demonstrable.

---

## Phase 7: User Story 5 — Killed Player Waits in the Respawn Queue (Priority: P5)

**Goal**: When a ship is destroyed, the player is queued with a countdown. GetKillQueueStatus shows remaining ticks. Once eligible, the player can return to the lobby.

**Independent Test**: Use the integration test harness to call KillPlayer for an active ship. Call `./cyclops session kill-queue --session <id> --player bob` and confirm ticks_remaining > 0. Advance the simulation clock past eligible_at and confirm the player may re-enter the lobby.

### Implementation for User Story 5

- [ ] T028 [US5] Implement KillPlayer(playerName string) method on GameSession in internal/server/session/session.go: find active slot by playerName, set slot.player_state = KILL_QUEUE, create KillQueueEntry with killed_at_stardate = current_stardate and eligible_at_stardate = killed_at + float64(config.respawn_wait_ticks), call KillQueue.Enqueue (oldest-eviction overflow), retain prior_side in PlayerInfo map
- [ ] T029 [US5] Implement GetKillQueueStatus RPC handler in internal/server/session/service.go: look up player in session's kill queue, compute ticks_remaining = int32(max(0, entry.eligible_at_stardate - session.current_stardate)), return in_kill_queue=true with current_stardate, eligible_at, and ticks_remaining; return in_kill_queue=false if player not found in queue
- [ ] T030 [US5] Implement RespawnToLobby(playerName string) method on GameSession in internal/server/session/session.go: validate that entry.eligible_at_stardate ≤ current_stardate (return error if not yet eligible), dequeue from kill queue, look up prior_side, find an UNOCCUPIED slot on prior_side (redirect to other side if prior_side is at 9 LOBBY+ACTIVE), set slot to LOBBY state and update player tracking map
- [ ] T031 [P] [US5] Implement `session kill-queue` CLI sub-command in internal/client/session/run.go: parse flags (--session, --player), call GetKillQueueStatus, pass response to output formatter
- [ ] T032 [P] [US5] Implement GetKillQueueStatusResponse output formatter in internal/client/session/output.go: if in_kill_queue print "Waiting to respawn: <ticks_remaining> ticks remaining (eligible at stardate <eligible_at>)"; if not in queue print "Not in kill queue"

**Checkpoint**: User Story 5 fully functional and independently demonstrable.

---

## Phase 8: User Story 6 — Game Session Ends When a Side Wins (Priority: P6)

**Goal**: When one side has no active ships and no queued respawns, the session transitions to ENDED. Subsequent commands are rejected. All GetSessionStatus responses reflect the winner.

**Independent Test**: Use the integration test harness to activate and kill all ships on one side with no kill queue entries remaining. Call GetSessionStatus and confirm state is SESSION_STATE_ENDED. Attempt JoinSession and confirm FailedPrecondition.

### Implementation for User Story 6

- [ ] T033 [US6] Implement CheckVictory() method on GameSession in internal/server/session/session.go: count ACTIVE slots per side; count KILL_QUEUE entries per side; if one side has 0 ACTIVE + 0 KILL_QUEUE and session.state == ACTIVE, set session.state = ENDED and record winning_side; return (ended bool, winner Side)
- [ ] T034 [US6] Integrate CheckVictory into all mutating RPC handlers in internal/server/session/service.go: call session.CheckVictory() after ActivateShip completes and after any KillPlayer call (future); propagate SessionState in all response messages; add winning_side string field to GetSessionStatusResponse and populate it when state is ENDED
- [ ] T035 [P] [US6] Add ENDED-state guard to JoinSession, LobbyCommand, and ActivateShip RPC handlers in internal/server/session/service.go: check session.state == ENDED before any processing; return gRPC FailedPrecondition with message "session <id> has ended" for all three
- [ ] T036 [P] [US6] Implement session-end output formatting in internal/client/session/output.go: print "Session ended. Winner: <side>" when GetSessionStatusResponse shows ENDED state; handle FailedPrecondition errors from all session sub-commands with message "This session has ended."

**Checkpoint**: User Story 6 fully functional. Full lifecycle (create → join → lobby → activate → die → respawn → victory) is demonstrable end-to-end.

---

## Phase 9: Polish & Cross-Cutting Concerns

**Purpose**: Unit and integration tests, concurrent-session isolation verification, and Makefile hygiene.

- [ ] T037 Add unit tests for GameSession state machine: WAITING→ACTIVE transition on first ActivateShip, ACTIVE→ENDED on CheckVictory, duplicate transition guards, kill queue max-10 eviction, tournament seed determinism (same name → same seed) in internal/server/session/session_test.go
- [ ] T038 [P] Add unit tests for all 7 SessionService RPC handlers using an in-process Registry and bufconn: CreateSession (regular and tournament mode), ListSessions, JoinSession (capacity enforcement, balance redirect, prior-side assignment), LobbyCommand (all 8 commands + rejection of unknown), ActivateShip (first activation transitions session, cap enforcement), GetKillQueueStatus, GetSessionStatus in internal/server/session/service_test.go
- [ ] T039 [P] Add gRPC integration test for full session lifecycle via bufconn: create → join two players (one per side) → LOBBY TIME command → activate both → call KillPlayer on one → GetKillQueueStatus → RespawnToLobby → activate again → kill all of one side → confirm SESSION_STATE_ENDED in tests/integration/grpc_session_test.go
- [ ] T040 [P] Add CLI integration test for session sub-commands against a live test server: `session create`, `session list`, `session join`, `session lobby --command users`, `session activate` in tests/integration/cli_session_test.go
- [ ] T041 [P] Add concurrent-session isolation test: create two sessions A and B, activate a player in A, verify B's GetSessionStatus is unaffected; kill the player in A, verify B remains in WAITING state in tests/integration/grpc_session_test.go

---

## Dependencies

```
Phase 1 (Setup) → Phase 2 (Foundational) → Phase 3 (US1) → Phase 4 (US2) → Phase 5 (US3) → Phase 6 (US4) → Phase 7 (US5) → Phase 8 (US6) → Phase 9 (Polish)
```

Within each Phase 3–8, server-side tasks (service.go / session.go) must complete before the CLI tasks (run.go / output.go) can be functionally tested end-to-end. However, [P]-marked CLI tasks may be written concurrently against the known gRPC interface.

### Parallel Execution Examples

**Phase 3 (US1)**:
- T011 + T012 + T013 can be executed sequentially in service.go
- T014 (run.go) and T015 (output.go) can run in parallel once T011–T013 are complete

**Phase 4 (US2)**:
- T016 (service.go) and T017 (session.go) — T017 must finish before T016 since service calls session methods
- T018 (run.go) and T019 (output.go) are parallel once T016 is done

**Phase 9 (Polish)**:
- T037 (session_test.go) and T038 (service_test.go) fully independent
- T039–T041 (tests/integration/) all independent of each other

---

## Implementation Strategy

**MVP** (deliver first): Phases 1 + 2 + 3 = T001–T015 (15 tasks)
- Gives a working CreateSession / ListSessions / GetSessionStatus slice
- Validates the protobuf contract, the in-memory registry, and the CLI create/list flow
- Directly satisfies SC-001 and SC-008 from the spec

**Increment 2**: + Phase 4 (US2, T016–T019) — adds JoinSession with balance enforcement
**Increment 3**: + Phase 5 (US3, T020–T023) — adds the pre-game lobby and its 8 commands
**Increment 4**: + Phase 6 (US4, T024–T027) — adds ship activation and session ACTIVE transition
**Increment 5**: + Phases 7 + 8 (US5 + US6, T028–T036) — completes the lifecycle with kill queue and session end
**Increment 6**: + Phase 9 (T037–T041) — full test coverage and concurrent-session validation
