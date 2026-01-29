# Phase Details

## Table of Contents
- [Phase 1: Definition](#phase-1-definition)
- [Phase 2: Design](#phase-2-design)
- [Phase 3: Implementation](#phase-3-implementation)
- [Phase 4: Infrastructure](#phase-4-infrastructure)
- [Phase 5: Quality Assurance](#phase-5-quality-assurance)
- [Phase 6: Release](#phase-6-release)
- [Rollback Procedures](#rollback-procedures)

---

## Phase 1: Definition

### Actors
- **PM Agent**: Requirements gathering, spec writing
- **Architect Agent**: Technical review, approval

### PM Agent Call

```
Action: Calling PM Agent
Phase: 1 - Definition
Context Injected:
  - docs/current/product-manual.md
Instruction: |
  User request: "[User's original request]"

  1. Interrogate user if requirements unclear
  2. Generate spec: docs/specs/YYYYMMDD-XX-[feature].md
  3. Include: User Stories, Acceptance Criteria, Data Model changes
Waiting for: Spec file path
```

### Architect Agent Call

```
Action: Calling Architect Agent
Phase: 1 - Definition
Context Injected:
  - docs/specs/YYYYMMDD-XX-[feature].md (just created)
  - docs/arch/tech-stack.md
  - docs/arch/coding-standards.md
  - docs/arch/infrastructure.md
  - docs/arch/api-contract.md
Instruction: |
  Review this spec against architecture rules.
  Check:
  - Tech stack compliance
  - Infrastructure requirements (Auth, Logging, RequestID)
  - API contract adherence
  - Database schema changes

  If violations found: REJECT with reason
  If approved: Update docs/arch/* if needed
Waiting for: "Architecture Approved" or "REJECTED: [reason]"
```

### Exit Criteria
- [ ] Spec file exists in docs/specs/
- [ ] Architect has approved
- [ ] Any arch doc updates committed

---

## Phase 2: Design

### Actors
- **UI/UX Agent**: Visual specifications

### Skip Conditions
- No frontend changes required
- API-only feature
- Backend-only feature

### UI/UX Agent Call

```
Action: Calling UI/UX Agent
Phase: 2 - Design
Context Injected:
  - docs/specs/YYYYMMDD-XX-[feature].md
  - docs/arch/tech-stack.md (for CSS framework)
Instruction: |
  Based on spec, generate Design-Spec.md including:
  - Component Tree
  - Design Tokens (colors, spacing, typography)
  - All component states (default, hover, active, disabled, loading, error)
  - Tailwind classes (or CSS as per tech-stack)
  - Responsive breakpoints
Waiting for: Design-Spec.md path
```

### Exit Criteria
- [ ] Design-Spec.md created
- [ ] All states documented
- [ ] Tech-stack alignment verified

---

## Phase 3: Implementation

### Actors
- **Backend Agent**: API implementation
- **Frontend Agent**: UI implementation

### Backend Agent Call

```
Action: Calling Backend Agent
Phase: 3 - Implementation
Context Injected:
  - docs/specs/YYYYMMDD-XX-[feature].md
  - docs/arch/tech-stack.md
  - docs/arch/coding-standards.md
  - docs/arch/infrastructure.md
  - docs/arch/api-contract.md
Instruction: |
  Implement feature per spec. MANDATORY:
  1. Auth Middleware on all protected endpoints
  2. X-Request-ID in all logs and responses
  3. Structured JSON logging (no print/console.log)
  4. Unit tests for ALL exported functions
  5. GET /api/docs endpoint serving user manual

  Produce:
  - Implementation code
  - *_test.go or *.spec.ts files
Waiting for: Code paths + confirmation of test files
```

### Frontend Agent Call

```
Action: Calling Frontend Agent
Phase: 3 - Implementation
Context Injected:
  - docs/specs/YYYYMMDD-XX-[feature].md
  - docs/arch/tech-stack.md
  - docs/arch/coding-standards.md
  - Design-Spec.md (from Phase 2)
Instruction: |
  Implement UI per Design-Spec. MANDATORY:
  1. Pixel-perfect to design spec
  2. TypeScript (no 'any')
  3. Defensive coding (?.  and error boundaries)
  4. Mobile-first responsive
  5. Use Mock Data if API not ready

  Produce:
  - Component files
  - Types/interfaces
Waiting for: Component paths
```

### Checkpoint
Before proceeding to Phase 4:
- [ ] Backend code exists
- [ ] Backend test files exist (HARD REQUIREMENT)
- [ ] Frontend code exists (if applicable)

**If no test files**: REJECT backend, do not proceed.

---

## Phase 4: Infrastructure

### Actors
- **SRE Agent**: Containerization, orchestration

### SRE Agent Call

```
Action: Calling SRE Agent
Phase: 4 - Infrastructure
Context Injected:
  - Codebase (implementation from Phase 3)
  - docs/arch/tech-stack.md
  - package.json / go.mod (dependencies)
Instruction: |
  Code is ready. Generate/update:
  1. Dockerfile (multi-stage, Alpine base)
  2. docker-compose.yml (with healthchecks)
  3. Makefile (make up, make test, make logs)
  4. Any new dependencies in containers

  Ensure one-click startup: `make up`
Waiting for: Infrastructure file paths
```

### Exit Criteria
- [ ] Dockerfile updated/created
- [ ] docker-compose.yml updated/created
- [ ] Makefile updated/created
- [ ] `make up` works

---

## Phase 5: Quality Assurance

### Actors
- **QA Agent**: Testing, verification

### QA Agent Call

```
Action: Calling QA Agent
Phase: 5 - Quality Assurance
Context Injected:
  - docs/specs/YYYYMMDD-XX-[feature].md
  - docs/public/user-manual.md (draft)
  - SRE environment info (Makefile commands)
Instruction: |
  Environment: Run `make up` to start

  Execute:
  1. API-level tests (auth, X-Request-ID, fuzzing)
  2. E2E tests (if UI changes)
  3. Documentation verification:
     - GET /api/docs returns content
     - Content matches user-manual.md

  Pass criteria:
  - All happy paths pass
  - 3+ edge cases tested
  - X-Request-ID present
  - /api/docs working

  Generate test report.
Waiting for: "QA Approved" or "REJECTED: [Bug Report]"
```

### On Rejection
```
If QA REJECTED:
  - Parse Bug Report
  - If code bug → Return to Phase 3 (Backend/Frontend)
  - If doc mismatch → Return to Phase 1 (PM)
  - After fix → Re-run Phase 4 (SRE) → Phase 5 (QA)
```

---

## Phase 6: Release

### Actors
- **PM Agent**: Documentation finalization

### PM Agent Call

```
Action: Calling PM Agent
Phase: 6 - Release
Context Injected:
  - docs/specs/YYYYMMDD-XX-[feature].md
  - docs/current/product-manual.md
  - docs/public/user-manual.md (draft)
Instruction: |
  QA Approved. Execute release:
  1. Merge spec content into docs/current/product-manual.md
  2. Finalize docs/public/user-manual.md
  3. Ensure /api/docs will serve correct content
  4. Close this spec (mark as IMPLEMENTED)
Waiting for: Confirmation of doc updates
```

### Exit Criteria
- [ ] product-manual.md updated
- [ ] user-manual.md published
- [ ] Spec marked as implemented

---

## Rollback Procedures

### From Phase 5 (QA Reject)

```
QA Bug Type: Code Bug
Action:
  1. Route Bug Report to Backend or Frontend Agent
  2. Fix code
  3. Re-run Phase 4 (SRE may need container updates)
  4. Re-run Phase 5 (QA)

QA Bug Type: Documentation Mismatch
Action:
  1. Route to PM Agent
  2. Update user-manual.md OR update spec
  3. If spec changed → Re-run from Phase 1
  4. If only manual changed → Re-run Phase 5
```

### From Phase 1 (Architect Reject)

```
Architect Reject Reason: Tech Stack Violation
Action:
  1. Return spec to PM
  2. PM revises with correct technology
  3. Re-submit to Architect

Architect Reject Reason: Infrastructure Concern
Action:
  1. Architect updates docs/arch/infrastructure.md
  2. PM incorporates new constraints
  3. Re-submit for approval
```
