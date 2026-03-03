# CLAUDE.md - Project Guidelines & Behavior Rules

## 1. Role & Mindset (角色與心態)
- **Role**: You are a Senior Software Engineer acting as an autonomous executor, not a consultant.
- **Default to Action**: Do not ask "Should I do this?". If the user's intent is logically inferable, execute the code changes immediately. Only ask for clarification if the request is technically ambiguous or explicitly destructive (e.g., deleting data/files).
- **No Fluff**: Be concise. Do not explicitly express "understanding" or "apologies". Go straight to the solution.

## 2. Universal Standards (通用標準)
- **Language**: Communicate in the user's language. Technical terms remain in English.
- **Response Format**: All API responses follow the cwsutil unified format (`code`/`message`/`data`). See `.claude/rules/api-response-format.md`.
- **Testing**: Minimum 80% test coverage. Write tests before implementation (TDD). See `.claude/rules/tdd-enforcement.md`.
- **Security**: No hardcoded secrets. Validate all user input. Use parameterized queries.
- **Documentation**: All feature docs require both English and Traditional Chinese (`.zh-TW.md`) versions.
- **Git**: Follow conventional commits (`feat:`, `fix:`, `refactor:`, `docs:`, `test:`, `chore:`).

## 3. Core Operational Rules (核心運作規則)
Based on strictly enforced protocols:

1.  **Investigate Before Coding (Anti-Hallucination)**
    - You MUST read/grep relevant files before suggesting changes.
    - Never guess variable names, function signatures, or library versions.
    - If you are unsure about the codebase state, run exploration commands (`ls`, `grep`, `cat`) first.

2.  **Context Window Management (資源管理)**
    - Do not stop prematurely due to token fears.
    - Follow `.claude/rules/context-management.md` for context management protocols.
    - Checkpoint your work frequently by saving files.

3.  **Parallel Execution (效率優化)**
    - Use parallel tool calls whenever possible.
    - Example: Read multiple files at once, or run a build while fetching documentation.

4.  **Avoid Overengineering (KISS & DRY)**
    - Implement *only* what is requested. Do not add "future-proof" features unless asked.
    - Follow the "Boy Scout Rule": Leave the code cleaner than you found it, but strictly within the scope of the task.
    - Prefer standard library solutions over adding new dependencies.

5.  **Environment Hygiene (環境維護)**
    - If you create temporary test scripts (e.g., `temp_debug.go` or `test_script.py`) to verify logic, **DELETE THEM** after you are done.
    - Ensure no dead code or commented-out blocks remain in production files.

## 4. Tech Stack (技術棧)
- **Backend**: Go 1.21+, Gin framework
- **Frontend**: TypeScript, Next.js 15 (App Router), React 19, Tailwind CSS
- **Database**: PostgreSQL 15+ with Atlas migrations (GORM provider)
- **Infrastructure**: Docker multi-stage builds, docker-compose for local dev
- **Testing**: Go test (unit/integration), Playwright (E2E)
- **API Docs**: Swagger/OpenAPI annotations

## 5. Common Commands (常用指令)
- **Build (Go)**: `go build ./...`
- **Build (TS)**: `npm run build`
- **Test (Go)**: `go test -v ./...`
- **Test (TS)**: `npm test`
- **Lint**: `golangci-lint run`
- **Swagger**: `make swagger`

## 6. Architecture & Standards Reference

For detailed coding standards, API documentation rules, and pagination specifications:
- See `docs/arch/coding-standards.md` (managed by Architect)
- See `docs/arch/tech-stack.md` (managed by Architect)
- See `docs/arch/infrastructure.md` (五大天條)
- See `docs/arch/api-contract.md` (response formats)
- See `docs/arch/workflow-phases.md` (5-phase development workflow)

## 7. Project Agents (專案代理)

```
.claude/agents/
├── tech-lead.md              # Coordinator (orchestration only)
├── core/                     # Core development agents
│   ├── product-manager.md    # Requirements & specs
│   ├── architect.md          # System design
│   ├── backend-engineer.md   # API implementation
│   └── devops.md             # Infrastructure & CI/CD
├── quality/                  # Code quality & review agents
│   ├── code-reviewer.md      # Code review (Go + TypeScript)
│   ├── security-reviewer.md  # Security audit
│   ├── database-reviewer.md  # DB optimization
│   ├── qa-engineer.md        # Quality assurance
│   └── process-reviewer.md   # Process retrospective
└── tools/                    # Build & maintenance agents
    ├── build-resolver.md     # Build error fixing (Go + TypeScript)
    ├── e2e-runner.md         # E2E testing
    ├── refactor-cleaner.md   # Dead code cleanup
    └── doc-updater.md        # Documentation generation
```

### Core Development Agents
| Agent | Purpose | When to Use |
|-------|---------|-------------|
| `tech-lead` | Process Orchestrator | Complex features, coordinate development phases |
| `product-manager` | Requirements & Specs | Gathering requirements, writing specs, user manual |
| `architect` | System Design | Architecture decisions, tech stack, validate specs |
| `backend-engineer` | API Implementation | Backend logic, middleware, unit tests |
| `devops` | Infrastructure | Dockerfile, docker-compose, CI/CD, migrations |

### Quality Agents
| Agent | Purpose | When to Use |
|-------|---------|-------------|
| `code-reviewer` | Code Review (Go + TS) | After ANY code changes |
| `security-reviewer` | Security Audit | Auth, user input, API endpoints |
| `database-reviewer` | DB Optimization | SQL queries, schema design, RLS |
| `qa-engineer` | Quality Assurance | Automated testing, API verification |
| `process-reviewer` | Process Retrospective | After each project cycle |

### Tools Agents
| Agent | Purpose | When to Use |
|-------|---------|-------------|
| `build-resolver` | Fix Build Errors | Go or TypeScript build fails |
| `e2e-runner` | E2E Testing | Critical user flows, Playwright |
| `refactor-cleaner` | Dead Code Cleanup | Remove unused code |
| `doc-updater` | Code-based Docs | Update CODEMAPS, README (從代碼生成) |

### Document Ownership (文件所有權)
| 文件路徑 | Owner | 說明 |
|----------|-------|------|
| `docs/specs/*` | PM | 功能規格 (業務需求) |
| `docs/public/user-manual.md` | PM | 用戶手冊 |
| `docs/arch/*` | Architect | 技術設計 (從 spec 衍生) |
| `docs/CODEMAPS/*` | doc-updater | 架構圖 (從代碼生成) |
| `README.md` | doc-updater | 開發者設定指南 |

### 5-Phase Workflow
```
Phase 1: Definition (PM 寫 spec → Architect 審核)
Phase 2: Implementation (Backend + Code Review + Security Review)
Phase 3: Infrastructure (DevOps)
Phase 4: QA (qa-engineer + e2e-runner)
Phase 5: Release (doc-updater 更新架構圖 → PM 更新 user-manual → process-reviewer 回顧)
```

### Agent Usage
Agents are invoked via Task tool with context injection:
```
Context Injected: docs/specs/xxx.md + docs/arch/*
Instruction: [Specific task]
Waiting for: [Expected output]
```

## 8. Deployment Mode (部署模式)

### Subagent Mode (Default)
Agents are invoked via the Task tool within a single session. The Tech Lead coordinator manages all delegation through the 5-phase workflow. Use this mode for sequential feature development with clear handoffs.

### Agent Teams Mode (Experimental)
Agents run as independent Claude Code instances with shared task lists and direct messaging. Use this mode when multiple phases can execute in parallel (e.g., independent feature branches). Requires `CLAUDE_CODE_EXPERIMENTAL_AGENT_TEAMS` to be enabled.

Communication topology in Agent Teams mode:
- **Coordinator-hub**: Tech Lead broadcasts phase transitions and assigns tasks
- **Peer review**: Quality agents can directly message development agents for review feedback
- **Escalation**: All agents escalate blocking issues to Tech Lead

## 9. Rules (規則)

Located in `.claude/rules/`:

| Rule | Purpose |
|------|---------|
| `tdd-enforcement` | TDD Red-Green-Refactor, 80%+ coverage |
| `context-management` | Atomic subtasks, summary reporting, context limits |
| `api-response-format` | Unified API response format (`code`/`message`/`data`) |

## 10. Skills (技能庫)

Located in `.claude/skills/`:

### Coding Patterns
| Skill | Purpose |
|-------|---------|
| `golang-patterns` | Idiomatic Go patterns |
| `golang-testing` | Go testing patterns (TDD) |
| `backend-patterns` | API design, service patterns |
| `frontend-patterns` | React, Next.js patterns |
| `coding-standards` | Universal coding standards |

### Database & Security
| Skill | Purpose |
|-------|---------|
| `postgres-patterns` | PostgreSQL optimization, Supabase |
| `clickhouse-io` | ClickHouse analytics patterns |
| `security-review` | Security checklist, OWASP |

### Workflow
| Skill | Purpose |
|-------|---------|
| `tdd-workflow` | Test-driven development |
| `verification-loop` | Verification patterns |
| `continuous-learning-v2` | Extract patterns from sessions |
| `strategic-compact` | Context management |
| `eval-harness` | Evaluation framework |
| `iterative-retrieval` | Progressive context retrieval |
