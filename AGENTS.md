# To AGENTS

## Purpose

You are an implementation agent. Produce production-grade code that is:

- Correct first, then clean, then fast.
- Easy to change, easy to test, easy to review.
- Designed for maintainers, not for the author.

When requirements are ambiguous, prefer the simplest implementation that preserves extensibility and explicit boundaries. Avoid speculative abstractions.

## Universal Priority Order (non-negotiable)

When principles conflict, resolve in this order:

1. **Correctness & Safety**

    - Data integrity, security, and deterministic behavior come first.

2. **Clarity / Maintainability**

    - “Write code for the maintainer”, Principle of Least Surprise.

3. **Testability**

    - Architecture and seams that enable fast, reliable tests.

4. **Cohesion & Coupling**

    - High cohesion, loose coupling; stable boundaries.

5. **Modularity & Evolvability**

    - Acyclic dependencies, composable units, replaceable components.

6. **Performance & Efficiency**

    - Only after the above; measure before optimizing.

7. **Convenience & Brevity**

    - Never at the expense of the above.

## Core Engineering Principles (apply by default)

### Design & Architecture

- **High Cohesion & Loose Coupling**

  - Keep related behavior together; minimize cross-module knowledge.

- **Separation of Concerns**

  - UI concerns, domain concerns, and infrastructure concerns do not mix.

- **Separation of Mechanism and Policy**

  - Mechanisms (how) are reusable; policies (what/why) are domain-specific.

- **SOLID**

  - Especially SRP and DIP. Apply pragmatically; avoid over-abstracting.

- **IoC / Dependency Inversion**

  - Depend on abstractions; infrastructure depends on domain, not vice versa.

- **Acyclic Dependencies Principle**

  - No import cycles; enforce a directed dependency graph.

- **Code to Interface, not Implementation**

  - Define ports/contracts; implement adapters behind them.

- **Composition over Inheritance**

  - Prefer small composable functions/structs; avoid deep class hierarchies.

- **Convention over Configuration**

  - Prefer predictable defaults and standard layouts.

- **Principle of Least Surprise**

  - Use idiomatic patterns for Vue/TypeScript and Go.

- **Fail Fast**

  - Validate inputs early; crash/return errors with clear context.

- **Immutable State Principle**

  - Prefer immutable data transformations; avoid shared mutable state.

- **Avoid Side Effects**

  - Make side effects explicit at boundaries (I/O, network, persistence).

- **Idempotency Principle**

  - For APIs, retries must be safe wherever feasible.

- **Graceful Degradation**

  - Partial failures should degrade functionality without corrupting data.

- **SPoT** (Single Point of Truth)

  - Avoid duplicated logic/config; centralize invariants.

- **DRY**

  - Do not repeat knowledge; however, prefer duplication over premature abstraction.

- **No Magic Numbers**

  - Use named constants; document units and rationale.

### Data & Persistence

- **ACID** where transactional storage is involved.
- Treat migrations, constraints, and invariants as first-class.

### Code Style & Readability

- Prefer explicit names, small units, early returns.
- Prefer domain language in names (ubiquitous language).
- Avoid cleverness. Avoid “one-liners” that hide intent.

## Domain-Driven Design (DDD) Rules

### Strategic Design

- Identify **Bounded Contexts**.

  - Do not leak models across contexts; integrate via contracts.

- Use explicit **Context Mapping** patterns where needed:

  - Anti-corruption layer (ACL) when consuming external/legacy models.

### Tactical Design

- **Entities**

  - Identity-based, lifecycle, invariants enforced in methods.

- **Value Objects**

  - Immutable, equality by value; validate at creation.

- **Aggregates**

  - One aggregate = one consistency boundary.
  - Only aggregate roots are referenced externally.
  - Enforce invariants inside the aggregate; keep it cohesive and small.

- **Domain Services**

  - Only when behavior doesn’t naturally belong to an entity/value object.

- **Repositories**

  - Collection-like abstraction for aggregates.
  - Interface in domain/application; implementation in infrastructure.

- **Domain Events**

  - Emit when something meaningful happened in the domain.
  - Events are facts (past tense), immutable, versioned.

- **Factories**

  - Use when creation is complex and needs invariants enforced.

- **Ubiquitous Language**

  - Keep naming consistent across code, docs, and API contracts.

## Clean Architecture (rules & checks)

- **Layers** (inward dependency only)

  - Domain (enterprise business rules)
  - Application (use cases, orchestration)
  - Interface Adapters (controllers, presenters, DTO mapping)
  - Infrastructure (DB, HTTP clients, frameworks)

- **Dependency rule**: Dependencies point inward only.

  - Domain has zero dependencies on frameworks.
  - Application depends on domain.
  - Adapters depend on application/domain.
  - Infrastructure depends on adapters/application/domain, never the reverse.

- **Ports & Adapters**

  - Define ports (interfaces) at the boundary of domain/application.
  - Implement adapters in outer layers (HTTP, DB, message bus, external APIs).

- **DTOs and Mapping**

  - Domain objects do not contain JSON/DB annotations.
  - Use explicit mapping at boundaries; avoid “anemic domain” unless requirements are trivial.

## Delivery Expectations

### Non-Goals

- Do not add speculative frameworks or heavy tooling.
- Do not introduce unnecessary generic abstractions.

### Code Must Include

- Clear folder structure, consistent naming.
- Input validation, typed contracts, and explicit error handling.
- Tests at the right level (unit for domain, integration for adapters, minimal E2E where valuable).
- Minimal but sufficient documentation (README updates if needed).

## Stack-Specific Guidance

### Frontend: TypeScript + Vue + Vite

#### Architecture

- Prefer feature-based organization over type-based dumping grounds.

- Separate:

  - `ui/` (components)
  - `features/` (feature modules)
  - `domain/` (frontend domain models if needed)
  - `services/` (API clients)
  - `stores/` (state management)

- Keep Vue components thin:

  - UI concerns in components.
  - Business rules in composables/services.

- Prefer composables:

  - `useXyz()` that expose explicit inputs/outputs.

#### State Management

- Keep state minimal; derive when possible.
- Avoid shared mutable state across unrelated features.
- Prefer immutable update patterns.

#### Error Handling & UX

- Fail fast on programmer errors; handle runtime failures gracefully.
- Centralize API error normalization; present user-friendly messages.

#### TypeScript Rules

- No `any` unless justified and localized.
- Use discriminated unions for state machines (loading/success/error).
- Prefer `unknown` + narrowing over unsafe casts.

### Backend: Golang Microservices

#### Service Boundaries

- Each service owns its data.
- Inter-service calls happen via explicit API contracts; no shared DB tables across services.

#### Packages & Dependencies

- Enforce acyclic imports.
- Keep domain/application packages framework-agnostic.

#### Handlers & Use Cases

- HTTP handlers/controllers do:

  - parse/validate input
  - call application use case
  - map result to response

- Business rules live in domain/application.

#### Error Handling

- Prefer explicit errors with context; wrap errors.
- Return stable error codes/messages at boundaries.
- Do not leak internal errors directly to clients.

#### Concurrency

- Keep concurrency explicit and tested.
- Avoid shared global mutable state.
- Use contexts properly (timeouts, cancellation propagation).

#### Idempotency

- For write endpoints, support idempotency keys where meaningful.
- Ensure retry-safety for network boundaries and message processing.

## Testing Strategy (required)

### Domain (fast unit tests)

- Test invariants and edge cases.
- No I/O.

### Application (use-case tests)

- Use mocks/fakes for ports (repositories, external clients).

### Infrastructure/Adapters (integration tests)

- Verify DB queries, HTTP wiring, serialization, migrations.
- Prefer containerized ephemeral dependencies if needed.

### Test Quality Bar

- Tests must be deterministic.
- Avoid brittle snapshot tests unless essential.
- Prefer behavior-focused assertions.

## Security & Robustness Baselines

- Validate all external input.
- Avoid logging secrets/PII.
- Use least privilege in configs.
- Use safe defaults: timeouts, retries with backoff, circuit-break patterns where relevant.
- Protect against replay/double-submit where applicable.

## Documentation & Communication

- Keep README accurate for running, testing, and basic architecture.
- If you add a new module or boundary, document the contract and rationale briefly.

## PR/Commit Hygiene (if applicable)

- Small, coherent changes.
- Commit messages describe intent.
- No drive-by refactors unless required for the task.

## Output Requirements for the Agent

When you deliver code:

- Provide a short summary of what changed and why in Turkish.
- List key decisions and which principle/constraint drove them in Turkish.
- Identify risks/assumptions and how you mitigated them in Turkish.
- Provide commands to run:

  - lint/format
  - tests
  - local dev
