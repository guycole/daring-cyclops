# Feature Specification: Game Session Lifecycle

**Feature Branch**: `002-game-session`

**Created**: 2026-06-20

**Status**: Draft

**Input**: User description: "game session lifecycle — foundational feature enabling multiple players to connect to the server and participate in a shared game instance, including session creation with game options, side selection, pre-game lobby, ship activation, and kill-queue respawn"

## User Scenarios & Testing *(mandatory)*

<!--
  IMPORTANT: User stories should be PRIORITIZED as user journeys ordered by importance.
  Each user story/journey must be INDEPENDENTLY TESTABLE - meaning if you implement just ONE of them,
  you should still have a viable MVP (Minimum Viable Product) that delivers value.

  Assign priorities (P1, P2, P3, etc.) to each story, where P1 is the most critical.
  Think of each story as a standalone slice of functionality that can be:
  - Developed independently
  - Tested independently
  - Deployed independently
  - Demonstrated to users independently
-->

### User Story 1 - First Player Creates and Configures a Game Session (Priority: P1)

As the first player to connect, I can create a new game session and configure
its options so that the game runs with the rules I choose and other players can
join under those rules.

**Why this priority**: A game session must exist before any other lifecycle
activity can occur. Establishing session creation and configuration is the
smallest meaningful unit of work and is the prerequisite for every other story.

**Independent Test**: Connect a single client to a server with no active
sessions, confirm the server creates a new session in the waiting state, prompt
the first player for game options, accept the options, and confirm the session
is configured and ready for additional players to join.

**Acceptance Scenarios**:

1. **Given** no game session exists on the server, **When** the first player
   connects, **Then** the server creates a new session in the waiting state and
   prompts the player to set game options before proceeding.
2. **Given** the first player is prompted for game options, **When** the player
   selects Regular mode with Romulans enabled and black holes disabled, **Then**
   the session is configured with those options and the player is placed in the
   pre-game lobby.
3. **Given** the first player is prompted for game options, **When** the player
   selects Tournament mode and supplies a named seed, **Then** the session is
   configured for reproducible play using that seed.
4. **Given** the first player has set game options, **When** additional players
   connect, **Then** those players are not re-prompted for game options; they
   inherit the session configuration set by the first player.

---

### User Story 2 - Additional Players Join a Session and Select a Side (Priority: P2)

As a player joining an existing session, I can choose the Federation or Empire
side so that I play as part of my preferred faction.

**Why this priority**: Side selection is the gateway into the game for every
player after the first. It must work correctly before the lobby or gameplay can
be meaningful.

**Independent Test**: Start a session as the first player. Connect a second
client. Confirm the second player is presented with side options, can select a
side, and is placed in the pre-game lobby for that side.

**Acceptance Scenarios**:

1. **Given** an active session with at least one open slot on each side, **When**
   a new player connects, **Then** the player is offered a choice between
   Federation and Empire.
2. **Given** a new player has chosen a side with available capacity, **When**
   the choice is confirmed, **Then** the server assigns the player to that side
   and places them in the pre-game lobby.
3. **Given** a new player's preferred side is full (9 ships active), **When**
   the player attempts to select that side, **Then** the server informs the
   player the side is full and offers the option to join the other side or wait.
4. **Given** the server detects a significant imbalance in team sizes, **When**
   a new player attempts to join the larger side, **Then** the server may
   redirect the player to the smaller side to restore balance, with a clear
   explanation.
5. **Given** a player has previously been on a side and was destroyed, **When**
   that player rejoins the session, **Then** the server automatically reassigns
   them to their previous side without asking for a new side selection.

---

### User Story 3 - Player Uses the Pre-Game Lobby (Priority: P3)

As a player in the pre-game lobby, I can run lobby commands to orient myself
before entering the game so that I can make an informed decision about when and
how to activate my ship.

**Why this priority**: The lobby is the buffer state between connection and
active gameplay. Players need access to informational commands before
committing to a ship.

**Independent Test**: Connect a player, complete side selection, and confirm
the player is in the lobby state. Run each lobby command (HELP, NEWS, USERS,
POINTS, SUMMARY, TIME, GRIPE, QUIT) and verify each produces a response or
action appropriate to the command. Confirm no gameplay action or ship
activation occurs from these commands.

**Acceptance Scenarios**:

1. **Given** a player has completed side selection, **When** they are placed in
   the lobby, **Then** the server presents a lobby prompt and the player can
   enter lobby commands.
2. **Given** a player is in the lobby, **When** they run HELP, **Then** the
   server returns a list of available lobby commands.
3. **Given** a player is in the lobby, **When** they run USERS, **Then** the
   server returns a list of currently connected players across all sides.
4. **Given** a player is in the lobby, **When** they run TIME, **Then** the
   server returns the current simulation time (stardate) for the session.
5. **Given** a player is in the lobby, **When** they run QUIT, **Then** the
   server disconnects the player cleanly and removes them from the session.
6. **Given** a player is in the lobby, **When** they attempt a command not in
   the allowed lobby set, **Then** the server rejects the command with a clear
   message indicating it is not available in the lobby.

---

### User Story 4 - Player Activates a Ship to Enter Gameplay (Priority: P4)

As a player in the pre-game lobby, I can run the ACTIVATE command to select
and enter a ship so that I can begin participating in the game.

**Why this priority**: ACTIVATE is the transition that moves a player from
observer to active participant. Without it, the lobby has no exit path into
gameplay.

**Independent Test**: Place a player in the lobby, run the ACTIVATE command,
and confirm the server moves the player out of the lobby state, assigns them a
ship, and transitions them to active play. Confirm the ship now appears as
active on that player's side.

**Acceptance Scenarios**:

1. **Given** a player is in the lobby, **When** they run ACTIVATE, **Then** the
   server transitions the player out of the lobby and begins ship selection.
2. **Given** a player has completed ship selection via ACTIVATE, **When** the
   ship is assigned, **Then** the player is in the active state and the ship
   counts toward the side's active ship total.
3. **Given** a side already has 9 active ships, **When** a lobby player on
   that side attempts to ACTIVATE, **Then** the server informs the player the
   side is at capacity and the player remains in the lobby.

---

### User Story 5 - Killed Player Waits in the Respawn Queue (Priority: P5)

As a player whose ship has been destroyed, I can see my respawn countdown and
re-enter the game when my wait expires so that I can continue participating
after death.

**Why this priority**: The kill queue is the recovery path for destroyed ships.
It is important for the full lifecycle but depends on active gameplay, making
it lower priority than session creation, joining, lobby, and activation.

**Independent Test**: Simulate ship destruction for a player. Confirm the
player is placed in the kill queue, is shown a countdown, cannot reactivate
until the countdown reaches zero, and is eligible to re-enter the lobby (and
then ACTIVATE again) once the timer expires.

**Acceptance Scenarios**:

1. **Given** a player's ship is destroyed, **When** the destruction is
   processed, **Then** the player is placed in the kill queue and shown the
   remaining wait time.
2. **Given** a player is in the kill queue, **When** they check their status,
   **Then** the server shows the current countdown to respawn eligibility.
3. **Given** a player's kill queue countdown has expired, **When** the player
   attempts to re-enter, **Then** the server allows them to return to the
   pre-game lobby and eventually activate a new ship.
4. **Given** the kill queue already contains 10 entries, **When** another ship
   is destroyed, **Then** the server handles the overflow condition and informs
   the player of the delay without crashing or losing state.
5. **Given** a player is killed and re-enters via the lobby, **When** they
   select ACTIVATE, **Then** the server assigns them to their previous side,
   not a new side.

---

### User Story 6 - Game Session Ends When a Side Wins (Priority: P6)

As a player in an active game, I am notified when the game ends and my side's
outcome is reported so that the session can be closed and results recorded.

**Why this priority**: Session termination completes the lifecycle. It is
necessary for correctness but depends on all prior lifecycle stages being
functional.

**Independent Test**: Reach a state where one side has no active ships and no
eligible players in the kill queue. Confirm the server transitions the session
to the ended state, notifies all connected players, and prevents new gameplay
actions from being processed.

**Acceptance Scenarios**:

1. **Given** all ships on one side are destroyed and no players remain eligible
   to respawn, **When** the session evaluates victory conditions, **Then** the
   server transitions the session to the ended state and declares the
   surviving side the winner.
2. **Given** the session has ended, **When** any player attempts to issue a
   gameplay command, **Then** the server rejects the command and informs the
   player the session is over.
3. **Given** the session has ended, **When** a player disconnects, **Then**
   the server handles the disconnection cleanly without error.

---

### Edge Cases

- A player disconnects mid-side-selection; the server must release any
  reserved slot and not leave the session in an inconsistent state.
- A player in the lobby disconnects; the server must remove them from the
  session cleanly.
- Both sides fill simultaneously, reaching the 18-ship maximum; the server
  must correctly track all 18 ships and enforce the cap.
- A player attempts to connect to a session that has already ended; the server
  must reject the connection with a clear status message.
- The server is hosting multiple concurrent sessions; a player's actions in
  one session must not affect any other session.
- The first player disconnects before setting game options; the server must
  either assign defaults or offer the next connecting player the opportunity
  to set options.
- Tournament mode is requested without supplying a seed; the server must
  require the seed before accepting the configuration.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The server MUST support multiple concurrent independent game
  sessions, each with isolated state, simulation time, and player rosters.
- **FR-002**: A game session MUST exist in exactly one of three lifecycle
  states: **waiting** (no ship has been activated), **active** (at least one
  ship is active), or **ended** (session is closed).
- **FR-003**: The first player to connect to a new session MUST be prompted to
  configure game options before any lobby or ship activation is available.
- **FR-004**: Game options MUST include:
  - Mode: Regular or Tournament
  - Romulan NPC: enabled or disabled (default: enabled)
  - Black holes: enabled or disabled (default: disabled)
- **FR-005**: Tournament mode MUST require a named seed that governs randomized
  game elements for reproducibility.
- **FR-006**: A game session MUST support exactly two sides: Federation and
  Empire, with a maximum of 9 simultaneously active ships per side (18 total).
- **FR-007**: New players joining a session MUST be offered a choice of side,
  subject to availability and server-enforced balance rules.
- **FR-008**: The server MUST enforce a maximum of 9 active ships per side; a
  player attempting to join a full side MUST be offered the option to join the
  other side or wait.
- **FR-009**: The server MUST enforce team balance and MAY force a player to
  join the smaller side when an imbalance exceeds a configurable threshold.
- **FR-010**: A player who reconnects after having been on a side previously
  (including after death) MUST be automatically reassigned to their prior side.
- **FR-011**: Every player MUST pass through a pre-game lobby state upon
  connection and side assignment before activating a ship.
- **FR-012**: The lobby MUST accept the following commands and no others:
  HELP, NEWS, USERS, POINTS, SUMMARY, TIME, GRIPE, QUIT, ACTIVATE.
- **FR-013**: The lobby MUST present a distinct server-side prompt that
  differs from the in-game prompt so players can tell which state they are in.
- **FR-014**: The ACTIVATE command MUST be the sole mechanism to transition a
  player from the lobby state to the active ship state.
- **FR-015**: When a player's ship is destroyed, the player MUST be placed in
  a kill queue; the kill queue MUST hold a maximum of 10 entries.
- **FR-016**: A player in the kill queue MUST be shown a countdown timer
  indicating how long until they are eligible to re-enter the lobby.
- **FR-017**: A player MUST NOT be able to re-enter the game (return to lobby
  or activate a ship) until their kill queue countdown expires.
- **FR-018**: All session state transitions, side assignments, lobby
  interactions, ship activations, and kill queue management MUST be executed
  and enforced exclusively by the server.
- **FR-019**: Gameplay time progression within a session MUST be governed by
  simulation time (stardates), not wall-clock time.
- **FR-020**: A session MUST transition to the ended state when one side
  achieves victory or all ships and eligible respawns for one side are
  exhausted.
- **FR-021**: Once a session has ended, no further gameplay commands MUST be
  accepted by that session.

### Key Entities

- **Game Session**: An independent game instance with a lifecycle state
  (waiting, active, ended), a configuration (game options), a simulation
  clock, and a roster of players and their ships. Sessions are isolated from
  each other.
- **Game Options**: Configuration established by the first player at session
  creation. Includes mode (Regular/Tournament), tournament seed (Tournament
  mode only), Romulan NPC presence, and black hole presence.
- **Side**: One of two factions in a session (Federation or Empire). Each side
  has a roster of up to 9 active ships, its own identity, and participates in
  the session's victory conditions.
- **Player**: A connected user associated with a session and a side. Has a
  lifecycle state: in-lobby, active (ship is live), or in-kill-queue (waiting
  to respawn). A player's side assignment persists across deaths.
- **Ship**: A player's active in-game vessel. Exactly one ship per active
  player. A destroyed ship places the player in the kill queue.
- **Pre-Game Lobby**: A server-side player state representing a connected
  player who has completed side selection but has not yet activated a ship.
  The lobby exposes a restricted command set and a distinct prompt.
- **Kill Queue**: A server-managed list of players waiting to respawn after
  ship destruction. Holds a maximum of 10 entries. Each entry carries a
  countdown timer. When the timer expires the player may return to the lobby.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The first player can connect to an empty server, create a
  session, set game options, and reach the lobby prompt in a single
  uninterrupted flow.
- **SC-002**: Up to 18 players across both sides can be simultaneously active
  in a single game session with no observable cross-player interference.
- **SC-003**: A new player can join an active session, complete side selection,
  use at least one lobby command, and activate a ship without requiring any
  out-of-band knowledge of session state.
- **SC-004**: All lobby commands respond within the same interaction cycle as
  any other server response; there is no observable difference in
  responsiveness between lobby commands and a successful ping.
- **SC-005**: A killed player's respawn countdown is visible and accurate; the
  player transitions back to the lobby automatically when the countdown
  expires.
- **SC-006**: When a session ends, all connected players receive a clear
  notification of the outcome before being disconnected or returned to an idle
  state.
- **SC-007**: Two concurrently running sessions on the same server produce no
  cross-session effects; a player in session A cannot observe or influence
  session B.
- **SC-008**: Tournament mode sessions with the same seed produce the same
  initial game state across independent server runs.

## Assumptions

- Player authentication and persistent identity are out of scope for this
  feature; the server distinguishes players by their connection for the
  duration of a session only.
- Detailed ship capabilities, combat mechanics, Romulan NPC behavior, and
  black hole effects are gameplay details addressed in future specifications;
  this feature defines the lifecycle framework those systems will operate within.
- The exact kill queue countdown duration is a configurable session parameter;
  the precise default value is deferred to the simulation time specification.
- The NEWS and GRIPE lobby commands may require persistent storage; this
  feature defines the command interface but defers storage architecture to a
  separate specification.
- The exact stardate tick rate and simulation time advancement rules are
  defined in the simulation time specification; this feature only requires
  that the TIME lobby command can report the current simulation time.
- The balance threshold that triggers forced side assignment is a configurable
  server parameter; the exact default is not fixed in this specification.
- A single client connects to a single game session at a time; connecting to
  multiple sessions simultaneously is out of scope.
- The POINTS and SUMMARY lobby commands report session-scoped statistics;
  cross-session leaderboards and historical records are out of scope.
