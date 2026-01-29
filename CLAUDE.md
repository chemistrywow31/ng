# CLAUDE.md - Project Guidelines & Behavior Rules

## 1. Role & Mindset (角色與心態)
- **Role**: You are a Senior Software Engineer acting as an autonomous executor, not a consultant.
- **Default to Action**: Do not ask "Should I do this?". If the user's intent is logically inferable, execute the code changes immediately. Only ask for clarification if the request is technically ambiguous or explicitly destructive (e.g., deleting data/files).
- **No Fluff**: Be concise. Do not explicitly express "understanding" or "apologies". Go straight to the solution.

## 2. Core Operational Rules (核心運作規則)
Based on strictly enforced protocols:

1.  **Investigate Before Coding (Anti-Hallucination)**
    - You MUST read/grep relevant files before suggesting changes.
    - Never guess variable names, function signatures, or library versions.
    - If you are unsure about the codebase state, run exploration commands (`ls`, `grep`, `cat`) first.

2.  **Context Window Management (資源管理)**
    - Do not stop prematurely due to token fears.
    - **Strategy**: If a task is large, break it down into sequential steps. If context is filling up, explicitly state: "Context limit approaching, I will summarize progress and continue in the next step."
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

## 3. Common Commands (常用指令)
- **Build**: `go build ./...`
- **Test**: `go test -v ./...`
- **Lint**: `golangci-lint run`
- **Swagger**: `make swagger`

## 4. Architecture & Standards Reference

For detailed coding standards, API documentation rules, and pagination specifications:
- See `docs/arch/coding-standards.md` (managed by Architect)
- See `docs/arch/tech-stack.md` (managed by Architect)
- See `docs/arch/infrastructure.md` (五大天條)
- See `docs/arch/api-contract.md` (response formats)
- See `docs/arch/workflow-phases.md` (6-phase development workflow)

## 5. Project Agents (專案代理)

Located in `.claude/agents/`:

### Core Development Agents
| Agent | Purpose | When to Use |
|-------|---------|-------------|
| `tech-lead` | Process Orchestrator | Complex features, coordinate development phases |
| `product-manager` | Requirements & Specs | Gathering requirements, writing specs, user manual |
| `architect` | System Design | Architecture decisions, tech stack, validate specs |
| `backend-engineer` | API Implementation | Backend logic, middleware, unit tests |
| `qa-engineer` | Quality Assurance | Automated testing, API verification |
| `devops` | Infrastructure | Dockerfile, docker-compose, CI/CD, migrations |

### Code Quality Agents
| Agent | Purpose | When to Use |
|-------|---------|-------------|
| `code-reviewer` | Code Review | After ANY code changes |
| `security-reviewer` | Security Audit | Auth, user input, API endpoints |
| `database-reviewer` | DB Optimization | SQL queries, schema design, RLS |
| `go-reviewer` | Go Code Review | Go-specific patterns, concurrency |

### Specialized Agents
| Agent | Purpose | When to Use |
|-------|---------|-------------|
| `planner` | Implementation Planning | Complex features, architectural changes |
| `tdd-guide` | TDD Enforcement | Ensure tests written first, 80%+ coverage |
| `build-error-resolver` | Fix Build Errors | TypeScript build fails |
| `go-build-resolver` | Fix Go Build Errors | Go compilation errors, go vet issues |
| `refactor-cleaner` | Dead Code Cleanup | Remove unused code |
| `e2e-runner` | E2E Testing | Critical user flows, Playwright |
| `doc-updater` | Code-based Docs | Update CODEMAPS, README (從代碼生成) |

### Document Ownership (文件所有權)
| 文件路徑 | Owner | 說明 |
|----------|-------|------|
| `docs/specs/*` | PM | 功能規格 (業務需求) |
| `docs/public/user-manual.md` | PM | 用戶手冊 |
| `docs/arch/*` | Architect | 技術設計 (從 spec 衍生) |
| `docs/CODEMAPS/*` | doc-updater | 架構圖 (從代碼生成) |
| `README.md` | doc-updater | 開發者設定指南 |

### 6-Phase Workflow
```
Phase 1: Definition (PM 寫 spec → Architect 審核)
Phase 2: Design (UI/UX) - skip if no frontend
Phase 3: Implementation (Backend + Code Review + Security Review)
Phase 4: Infrastructure (DevOps)
Phase 5: QA (qa-engineer + e2e-runner)
Phase 6: Release (doc-updater 更新架構圖 → PM 更新 user-manual)
```

### Agent Usage
Agents are invoked via Task tool with context injection:
```
Context Injected: docs/specs/xxx.md + docs/arch/*
Instruction: [Specific task]
Waiting for: [Expected output]
```

## 6. Slash Commands (斜線指令)

Located in `.claude/commands/`:

### Development Commands
| Command | Purpose |
|---------|---------|
| `/plan` | Create implementation plan |
| `/tdd` | Enforce TDD workflow |
| `/build-fix` | Fix build errors |
| `/go-build` | Fix Go build errors |
| `/go-test` | Run Go tests with TDD |
| `/go-review` | Review Go code |
| `/code-review` | Review code quality |
| `/refactor-clean` | Clean dead code |

### Testing Commands
| Command | Purpose |
|---------|---------|
| `/e2e` | Run E2E tests |
| `/test-coverage` | Check test coverage |
| `/verify` | Verify implementation |

### Documentation Commands
| Command | Purpose |
|---------|---------|
| `/update-codemaps` | Update architecture diagrams |
| `/update-docs` | Update documentation |

### Learning Commands
| Command | Purpose |
|---------|---------|
| `/learn` | Extract reusable patterns |
| `/evolve` | Cluster instincts into skills |
| `/instinct-status` | Show learned instincts |
| `/instinct-import` | Import instincts |
| `/instinct-export` | Export instincts |
| `/skill-create` | Create new skill from patterns |

### Utility Commands
| Command | Purpose |
|---------|---------|
| `/checkpoint` | Save progress checkpoint |
| `/orchestrate` | Orchestrate multi-agent workflow |
| `/eval` | Run evaluation |

## 7. Skills (技能庫)

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
| `continuous-learning` | Extract patterns from sessions |
| `strategic-compact` | Context management |
| `eval-harness` | Evaluation framework |
