# Spec Templates

## Table of Contents
- [Feature Spec Template](#feature-spec-template)
- [Bug Fix Spec Template](#bug-fix-spec-template)
- [Product Manual Template](#product-manual-template)
- [User Manual Template](#user-manual-template)

---

## Feature Spec Template

```markdown
# Spec: [Feature Name]

**Date**: YYYY-MM-DD
**Spec ID**: YYYYMMDD-XX
**Author**: PM
**Status**: Draft | Review | Approved | Implemented

---

## 1. Background

### Problem Statement
[What problem does this feature solve?]

### Business Value
[Why is this worth building?]

### Success Metrics
- [ ] [Metric 1]
- [ ] [Metric 2]

---

## 2. User Stories

### Primary User Story
As a [user type],
I want [capability],
So that [benefit].

### Secondary Stories
- As a [user], I want [action] so that [benefit]
- As a [user], I want [action] so that [benefit]

---

## 3. Acceptance Criteria

### Happy Path
- [ ] Given [context], when [action], then [expected result]
- [ ] Given [context], when [action], then [expected result]

### Edge Cases
- [ ] Given [edge case], when [action], then [handling]

### Error Handling
- [ ] Given [error condition], when [action], then [error response]

---

## 4. Technical Requirements

### Data Model Changes
```
Table: [table_name]
- field_name: type (constraints)
- field_name: type (constraints)
```

### API Changes
```
POST /api/v1/[endpoint]
Request:
{
  "field": "type"
}
Response:
{
  "success": true,
  "data": {}
}
```

### UI Changes
- [ ] [Component/Page]: [Change description]
- [ ] [Component/Page]: [Change description]

---

## 5. Out of Scope

- [Feature/capability NOT included]
- [Feature/capability NOT included]

---

## 6. Dependencies

- [ ] [Dependency 1]
- [ ] [Dependency 2]

---

## 7. Risks & Mitigations

| Risk | Impact | Mitigation |
|------|--------|------------|
| [Risk] | High/Med/Low | [Mitigation strategy] |

---

## 8. Sign-off

- [ ] PM Approved
- [ ] Architect Approved
- [ ] Tech Lead Assigned
```

---

## Bug Fix Spec Template

```markdown
# Bug Fix: [Brief Description]

**Date**: YYYY-MM-DD
**Spec ID**: YYYYMMDD-XX-fix
**Severity**: Critical | High | Medium | Low
**Reported By**: [Name/Source]

---

## 1. Bug Description

### Current Behavior
[What happens now - the bug]

### Expected Behavior
[What should happen]

### Steps to Reproduce
1. [Step 1]
2. [Step 2]
3. [Step 3]

### Environment
- Browser/Device: [Details]
- OS: [Details]
- Version: [App version]

---

## 2. Root Cause Analysis

[Technical explanation of why this bug occurs]

---

## 3. Proposed Fix

### Code Changes
- [ ] [File]: [Change description]

### Testing Checklist
- [ ] Unit test added
- [ ] Regression test passed
- [ ] Manual QA verified

---

## 4. Impact Assessment

- **Users Affected**: [Number/Percentage]
- **Data Impact**: [None/Recoverable/Permanent]
- **Rollback Plan**: [How to revert if needed]
```

---

## Product Manual Template

```markdown
# Product Manual

**Last Updated**: YYYY-MM-DD
**Version**: X.X.X

---

## Overview

[1-2 sentence product description]

---

## Features

### 1. [Feature Name]

**Description**: [What this feature does]

**User Flow**:
1. [Step 1]
2. [Step 2]
3. [Step 3]

**Technical Notes**:
- API: `[endpoint]`
- Data: `[relevant tables/fields]`

---

### 2. [Feature Name]

[Same structure as above]

---

## Data Model

### Core Entities

| Entity | Description | Key Fields |
|--------|-------------|------------|
| User | System users | id, email, role |
| [Entity] | [Description] | [Fields] |

---

## API Reference

### Endpoints

| Method | Path | Description |
|--------|------|-------------|
| GET | /api/v1/users | List users |
| POST | /api/v1/users | Create user |

---

## Configuration

| Setting | Default | Description |
|---------|---------|-------------|
| [Setting] | [Value] | [What it does] |

---

## Changelog

### [Version] - YYYY-MM-DD
- Added: [Feature]
- Fixed: [Bug]
- Changed: [Modification]
```

---

## User Manual Template

```markdown
# [Product Name] User Guide

Welcome to [Product Name]! This guide will help you get started.

---

## Getting Started

### Quick Start

1. **[First Step]**: [Simple instruction]
2. **[Second Step]**: [Simple instruction]
3. **[Third Step]**: [Simple instruction]

You're ready to go!

---

## Features

### [Feature Name]

[Friendly description of what this does and why it's useful]

**How to use it:**

1. [Clear step]
2. [Clear step]
3. [Clear step]

**Tips:**
- [Helpful tip]
- [Helpful tip]

---

### [Feature Name]

[Same friendly structure]

---

## FAQ

### [Common Question]?

[Clear, helpful answer]

### [Common Question]?

[Clear, helpful answer]

---

## Need Help?

- **Email**: support@example.com
- **Docs**: https://docs.example.com

---

*Last updated: YYYY-MM-DD*
```
