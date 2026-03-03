---
name: Context Management
description: Control context size and information flow to prevent token bloat and maintain agent focus
---

# Context Management

## Applicability

- Applies to: All agents (coordinator and all worker agents)

## Rule Content

### Atomic Subtask Decomposition

The Tech Lead (coordinator) must break every task into focused, atomic subtasks. Each subtask must target a single concern — one file group, one feature slice, or one logical unit. You must not assign a subtask that spans multiple unrelated modules or features.

### Summary-Based Reporting

Agents must report results as concise summaries containing:

- What was done (one to three sentences)
- Files created or modified (list of paths)
- Decisions made and their rationale (bullet points)
- Open issues or blockers (if any)

You must not dump full file contents, complete logs, or raw command output into reports. Summarize before passing information to the next agent.

### Context Injection Scoping

You must include only the files directly relevant to the current subtask when injecting context. You must not inject entire directories. Specify exact file paths or use targeted glob patterns scoped to the current concern.

### Token Usage Monitoring

You must use the `/strategic-compact` skill to check token usage at these logical checkpoints:

- After completing each subtask
- Before starting a subtask estimated to consume significant context
- When switching between distinct areas of the codebase

### Context Limit Handling

When the context limit approaches 70% utilization, you must:

1. Checkpoint current progress by writing a summary to the task output
2. Complete the current atomic unit of work
3. Continue remaining work in a new step with a fresh context window

You must not attempt to continue complex work when context is near capacity.

### Large Code Review Splitting

You must split large code reviews by module or functional area. Each review pass must focus on one module. You must not review an entire codebase or multiple unrelated modules in a single context window.

## Violation Determination

- Coordinator assigns a subtask spanning multiple unrelated modules without decomposition → Violation
- Agent reports contain full file dumps or raw command output exceeding 50 lines instead of summaries → Violation
- Context injection includes an entire directory instead of specific files → Violation
- Agent continues working past 70% context utilization without checkpointing → Violation
- Code review covers multiple unrelated modules in a single pass → Violation

## Exceptions

- When debugging a cross-module issue that requires simultaneous visibility of multiple files, the coordinator may authorize a broader context injection, documented in the subtask description with a rationale.
