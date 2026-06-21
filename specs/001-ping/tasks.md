# Tasks: Ping

**Input**: Design documents from `/specs/001-ping/`

**Prerequisites**: plan.md (required), spec.md (required for user stories), research.md, data-model.md, contracts/

**Tests**: No dedicated test-writing tasks are included because the feature spec did not explicitly request a TDD workflow. Validation is captured through quickstart execution and implementation tasks should still include appropriate code-level verification.

**Organization**: Tasks are grouped by user story to enable independent implementation and testing of each story.

## Format: `[ID] [P?] [Story] Description`

- **[P]**: Can run in parallel (different files, no dependencies)
- **[Story]**: Which user story this task belongs to (e.g., US1, US2)
- Include exact file paths in descriptions

## Phase 1: Setup (Shared Infrastructure)

**Purpose**: Create the base Go module, source layout, and protobuf generation path.

- [X] T001 Initialize the Go module and add base gRPC/protobuf dependencies in go.mod
- [X] T002 Create the planned source tree and binary entrypoints in cmd/cyclops/main.go and cmd/cyclopsd/main.go
- [X] T003 [P] Copy the ping contract into the runtime protobuf source path at proto/ping/v1/ping.proto
- [X] T004 [P] Add a protobuf generation command for the repo in Makefile

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: Establish shared pieces required by both the successful ping path and the failure-reporting path.

**⚠️ CRITICAL**: No user story work can begin until this phase is complete

- [X] T005 Create shared build metadata version handling in internal/buildinfo/version.go
- [X] T006 [P] Add shared client connection configuration and dial logic in internal/client/config.go
- [X] T007 [P] Generate Go protobuf and gRPC stubs into gen/proto/pingv1/ from proto/ping/v1/ping.proto
- [X] T008 Wire basic server startup, listen address handling, and gRPC server bootstrap in cmd/cyclopsd/main.go

**Checkpoint**: Foundation ready - user story implementation can now begin in priority order

---

## Phase 3: User Story 1 - Ping the Server (Priority: P1) 🎯 MVP

**Goal**: Allow the CLI client to send a ping request to the server and display the returned server version and current server time.

**Independent Test**: Start the server, run the CLI `ping` command against it, and confirm the CLI prints labeled version and server time values.

### Implementation for User Story 1

- [X] T009 [US1] Implement the ping RPC service to return server version and current server time in internal/server/ping/service.go
- [X] T010 [US1] Register the ping service with the gRPC server in cmd/cyclopsd/main.go
- [X] T011 [US1] Implement the client-side ping command execution and response mapping in internal/client/ping/run.go
- [X] T012 [US1] Add CLI command parsing and `ping` dispatch in cmd/cyclops/main.go
- [X] T013 [US1] Add human-readable ping output formatting for version and server time in internal/client/ping/output.go

**Checkpoint**: User Story 1 should now be fully functional and independently demonstrable

---

## Phase 4: User Story 2 - Report Ping Failures Clearly (Priority: P2)

**Goal**: Ensure the CLI reports ping failures clearly when the server is unreachable or the response is invalid.

**Independent Test**: Run the CLI `ping` command against an unavailable or invalid server target and confirm the CLI reports failure without printing fabricated values.

### Implementation for User Story 2

- [X] T014 [US2] Add client-side ping failure handling for unreachable servers and RPC errors in internal/client/ping/run.go
- [X] T015 [US2] Add response validation for missing version or timestamp fields in internal/client/ping/run.go
- [X] T016 [US2] Update CLI exit behavior and error rendering for failed ping commands in cmd/cyclops/main.go
- [X] T017 [US2] Add server-side defensive response construction for ping replies in internal/server/ping/service.go

**Checkpoint**: User Stories 1 and 2 should both work independently, including successful and failed ping flows

---

## Phase 5: Polish & Cross-Cutting Concerns

**Purpose**: Finalize feature usability, documentation, and end-to-end validation.

- [X] T018 [P] Document ping build and usage steps in README.md
- [X] T019 Validate the implementation against specs/001-ping/quickstart.md
- [X] T020 [P] Clean up generated or helper command documentation in .github/copilot-instructions.md if the plan path changed during implementation

---

## Dependencies & Execution Order

### Phase Dependencies

- **Setup (Phase 1)**: No dependencies - can start immediately
- **Foundational (Phase 2)**: Depends on Setup completion - blocks all user stories
- **User Story 1 (Phase 3)**: Depends on Foundational completion
- **User Story 2 (Phase 4)**: Depends on User Story 1 because it extends the ping command and RPC handling already introduced there
- **Polish (Phase 5)**: Depends on both user stories being complete

### User Story Dependencies

- **User Story 1 (P1)**: Starts after Phase 2 and establishes the MVP gRPC slice
- **User Story 2 (P2)**: Starts after User Story 1 because it refines the same client and server paths with failure handling

### Parallel Opportunities

- `T003` and `T004` can run in parallel after `T001` and `T002`
- `T006` and `T007` can run in parallel after setup is complete
- `T018` and `T020` can run in parallel during the polish phase

---

## Parallel Example: User Story 1 Foundation

```bash
# After setup is complete, shared prerequisites can proceed together:
Task: "Add shared client connection configuration and dial logic in internal/client/config.go"
Task: "Generate Go protobuf and gRPC stubs into gen/proto/pingv1/ from proto/ping/v1/ping.proto"
```

---

## Implementation Strategy

### MVP First (User Story 1 Only)

1. Complete Phase 1: Setup
2. Complete Phase 2: Foundational
3. Complete Phase 3: User Story 1
4. Validate the CLI `ping` success path against the quickstart

### Incremental Delivery

1. Establish Go module, source layout, and protobuf generation support
2. Deliver the first successful ping round-trip as the MVP
3. Add explicit failure handling for unreachable servers and invalid responses
4. Finish with usage documentation and quickstart validation

### Parallel Team Strategy

With multiple developers:

1. One developer completes Go module and binary scaffolding while another prepares the runtime protobuf source and generation command
2. After foundation work, one developer can focus on server ping service wiring while another prepares client connection plumbing
3. Failure-path refinements can follow once the success path is working end to end

---

## Notes

- Tasks are intentionally limited to the first gRPC slice defined in the `001-ping` feature spec.
- Generated protobuf code is treated as part of the implementation output under gen/proto/pingv1/.
- No authentication, persistence, or simulation-time work is included in this feature.