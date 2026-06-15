# Cyclops Project Constitution

## Purpose
Daring Cyclops is a multiplayer space strategy game inspired by [MegaWars](https://en.wikipedia.org/wiki/MegaWars). The project consists of a Go command-line client and a Go backend server. The server may host multiple game instances concurrently, and a client connects to a specific game instance. This constitution defines the governing technical and operational principles for the project and establishes the rules by which future specifications and implementation decisions will be made.

---

## 1. Governance
- This constitution is the highest-level governing document for Daring Cyclops unless replaced by a later constitution approved by the core maintainers.
- Amendments may be proposed through tracked issues and require approval by a majority of active admins.
- Disputes regarding technical direction, architecture, or contributor conduct are resolved by admin review and vote. Ties defer to the principal author or authors.

---

## 2. Core Roles
- User: A participant who interacts with the game through the client.
- Admin: A maintainer with authority to manage releases, infrastructure, governance, moderation, and project policy.
- Maintainer: A contributor trusted to review changes and help evolve project architecture and specifications.

---

## 3. Product and Architecture Invariants
The following principles are mandatory unless explicitly amended:

- The project consists of a Go command-line client and a Go backend server.
- Client and server communication must use gRPC.
- Network contracts must be defined in Protocol Buffers as the source of truth for client/server APIs.
- A single server process or deployment may host multiple independent game instances concurrently.
- Each client session must target a specific game instance.
- The server is authoritative for multiplayer game state, simulation state, and rule enforcement within each game instance.
- Clients are responsible for user interaction and presentation, but must not be trusted as authoritative for game state.
- Game state, simulation time, scheduled events, and player actions must be isolated per game instance.
- The server must run on Linux and support both local or LAN-hosted play and cloud-hosted deployment, including AWS.
- Cloud infrastructure, especially AWS infrastructure, must be managed declaratively with Terraform.
- Non-cloud deployments must also use reproducible, version-controlled deployment artifacts and configuration.

---

## 4. Simulation and Time Model
Daring Cyclops is governed by virtual time rather than wall-clock time for gameplay progression.

- Core gameplay logic must operate on simulation time.
- Each game instance maintains its own simulation timeline.
- Scheduled gameplay events must be expressible in terms of future simulation cycles or ticks within a given game instance.
- The server is the authority for simulation timing and event execution order for each game instance.
- Simulation behavior should be designed for deterministic testing wherever practical.
- Changes to time progression, event ordering, or scheduling semantics require explicit review because they directly affect game rules and fairness.

---

## 5. Interfaces and Compatibility
- All protocol changes must be tracked through issues and reviewed before implementation.
- Breaking changes to protobuf or gRPC interfaces require explicit approval and a migration plan.
- Client/server compatibility expectations must be documented as the protocol evolves.
- Public interfaces should be versioned when needed to support orderly evolution.

---

## 6. Security and Identity
- Authentication and authorization are not fixed at this stage of the project.
- Early phases may operate with open client access.
- The architecture must keep authentication, authorization, and session control pluggable so stronger identity and access controls can be introduced later without major redesign.
- No implementation choice in this area should create unnecessary coupling to a single provider or vendor before requirements are settled.

---

## 7. Persistence and Data Boundaries
- Database and persistence technology are not yet fixed.
- The architecture must isolate persistence concerns behind clear interfaces so storage strategy can evolve without rewriting core game logic.
- Any long-term decision about persistence, recovery, or historical game data must be captured in a tracked specification before becoming a hard dependency.
- If hosted deployments require persistence, backup and restore expectations must be documented before production use.

---

## 8. Operational Principles
- Infrastructure changes must be version-controlled and reviewed.
- Deployments should be reproducible across supported environments.
- Observability, logging, and operational diagnostics must be considered part of the server design, though specific tooling may evolve over time.
- Production-oriented hosting decisions must document operational assumptions, including deployment, recovery, and upgrade behavior.

---

## 9. Contribution and Change Control
- Features, bugs, architectural changes, protocol changes, and infrastructure changes must be tracked through GitHub Issues.
- Code and documentation changes require review before merging.
- Changes affecting protocol contracts, simulation semantics, deployment architecture, or governance must receive explicit maintainer approval.
- Detailed implementation direction should live in follow-on specifications governed by this constitution.

---

## 10. SpecKit Governance
- This application is managed through SpecKit.
- The constitution defines stable project rules and architectural invariants.
- Detailed behavior, implementation plans, infrastructure designs, simulation rules, and gameplay systems should be defined in separate specifications derived from this constitution.
- Future specifications must remain consistent with the principles established here unless this constitution is amended first.

---

## 11. Adoption
By contributing to, deploying, or administering Daring Cyclops, maintainers and contributors agree to follow the governance and architectural principles defined in this constitution.

**Version**: 0.1.0 | **Ratified**: 2026-06-15 | **Last Amended**: 2026-06-15
