# Architecture Document Templates

## Table of Contents
- [tech-stack.md Template](#tech-stackmd-template)
- [coding-standards.md Template](#coding-standardsmd-template)
- [infrastructure.md Template](#infrastructuremd-template)
- [api-contract.md Template](#api-contractmd-template)

---

## tech-stack.md Template

```markdown
# Tech Stack

## Core Language
- **Language**: [Go 1.21 / Node.js 20 / Python 3.12]
- **Version Lock**: STRICT. No upgrades without Architecture approval.

## Framework
- **Web Framework**: [Gin / Express / FastAPI]
- **Reason**: [Performance / Community / Type Safety]

## Database
- **Primary DB**: [PostgreSQL 15 / MySQL 8]
- **ORM**: [GORM / Prisma / SQLAlchemy]
- **Migration Tool**: [golang-migrate / Prisma Migrate / Alembic]

## Caching
- **Cache Layer**: [Redis 7 / Memcached]
- **Use Cases**: Session, Rate Limiting, Hot Data

## Message Queue (if applicable)
- **Queue System**: [RabbitMQ / Kafka / NATS]
- **Use Cases**: [Async Processing / Event Sourcing]

## Key Libraries
| Purpose | Library | Version | Locked |
|---------|---------|---------|--------|
| Logging | zap | 1.26.x | YES |
| Config | viper | 1.18.x | YES |
| Validation | validator | 10.x | YES |
| Testing | testify | 1.8.x | YES |

## Forbidden Technologies
- [ ] console.log / print statements in production
- [ ] [Specific framework/library to avoid]
- [ ] Direct SQL queries (use ORM)
```

---

## coding-standards.md Template

```markdown
# Coding Standards

## Directory Structure

### Backend (Go)
```
/cmd
  /api           # Main entry point
/internal
  /config        # Configuration loading
  /middleware    # HTTP middleware (Auth, Logger, RequestID)
  /handler       # HTTP handlers (Controllers)
  /service       # Business logic
  /repository    # Data access layer
  /model         # Domain models
  /dto           # Data transfer objects
/pkg             # Shared utilities (can be imported by external)
/docs
  /arch          # Architecture documents
  /specs         # PM specifications
  /public        # User-facing docs
/migrations      # Database migrations
```

### Frontend (React/Vue)
```
/src
  /components    # Reusable UI components
  /pages         # Route-level components
  /hooks         # Custom React hooks
  /services      # API client layer
  /store         # State management
  /utils         # Helper functions
  /types         # TypeScript definitions
```

## Naming Conventions

### Files
- Go: `snake_case.go`
- TypeScript/React: `PascalCase.tsx` for components, `camelCase.ts` for utils
- Test files: `*_test.go` or `*.test.ts`

### Variables & Functions
- Go: `camelCase` for private, `PascalCase` for public
- TypeScript: `camelCase` for variables/functions, `PascalCase` for types/interfaces

### Constants
- `SCREAMING_SNAKE_CASE` for true constants
- `PascalCase` for enum-like constants in Go

## Error Handling

### Go Pattern
```go
result, err := doSomething()
if err != nil {
    return fmt.Errorf("doSomething failed: %w", err)
}
```

### TypeScript Pattern
```typescript
try {
  const result = await doSomething();
} catch (error) {
  logger.error('doSomething failed', { error, context });
  throw new AppError('OPERATION_FAILED', error);
}
```

## Constraints
- NO magic numbers. Use named constants.
- NO nested callbacks beyond 2 levels.
- NO function longer than 50 lines. Split if exceeded.
- NO commented-out code in production.
```

---

## infrastructure.md Template

```markdown
# Infrastructure Requirements

## Mandatory Middleware Stack

### 1. Identity Middleware (Auth)
- **Type**: JWT Bearer Token / Session Cookie
- **Algorithm**: HS256 / RS256
- **Token Expiry**: 72 hours (access), 30 days (refresh)
- **Whitelist Routes**: `/login`, `/register`, `/health`, `/public/*`, `/swagger/*`

Implementation Requirements:
- Extract token from `Authorization: Bearer <token>` header
- Validate signature and expiry
- Inject user context into request
- Return 401 on invalid/expired token

### 2. Request ID Middleware (Traceability)
- Generate UUID v4 for each request
- Check for existing `X-Request-ID` header (pass-through from upstream)
- Inject into:
  - All log entries
  - Database query context
  - Response headers
  - Downstream service calls

### 3. Structured Logging Middleware
- **Library**: zap (Go) / winston (Node) / structlog (Python)
- **Format**: JSON
- **Required Fields**:
  ```json
  {
    "timestamp": "ISO8601",
    "level": "info|warn|error",
    "request_id": "uuid",
    "method": "GET|POST|...",
    "path": "/api/...",
    "status": 200,
    "latency_ms": 45,
    "user_id": "optional"
  }
  ```
- **FORBIDDEN**: `print()`, `console.log()`, `fmt.Println()`

### 4. Recovery Middleware (Panic Handler)
- Catch all panics/uncaught exceptions
- Log full stack trace with request context
- Return 500 with generic error message
- Never expose internal errors to client

### 5. Meta-Docs API
- **Endpoint**: `GET /api/docs`
- **Function**: Return contents of `docs/public/user-manual.md`
- **Response**: `{ "content": "<markdown string>" }`
- **Purpose**: Self-describing system

### 6. Swagger API Documentation
- **Endpoint**: `GET /swagger/*`
- **Function**: Serve interactive Swagger UI
- **Purpose**: Technical API contract for developers

Requirements:
- ALL API endpoints MUST have OpenAPI/Swagger annotations
- Include request/response schemas
- Document all error codes
- Include authentication requirements
- Group endpoints by resource/tag

Tools:
- Go: `swaggo/swag` + `swaggo/gin-swagger`
- Node: `swagger-jsdoc` + `swagger-ui-express`
- Python: `FastAPI` (built-in) or `flasgger`

## Health Check
- **Endpoint**: `GET /health`
- **Response**: `{ "status": "ok", "version": "x.x.x", "timestamp": "..." }`
- **Checks**: DB connectivity, Redis connectivity, critical dependencies

## Rate Limiting
- **Strategy**: Token bucket / Sliding window
- **Default**: 100 requests/minute per IP
- **Authenticated**: 1000 requests/minute per user
- **Headers**: `X-RateLimit-Remaining`, `X-RateLimit-Reset`
```

---

## api-contract.md Template

```markdown
# API Contract

## Base URL
- Development: `http://localhost:8080/api/v1`
- Staging: `https://staging.example.com/api/v1`
- Production: `https://api.example.com/v1`

## Authentication
- **Method**: Bearer Token (JWT)
- **Header**: `Authorization: Bearer <token>`
- **Token Endpoint**: `POST /auth/login`

## Request Headers
| Header | Required | Description |
|--------|----------|-------------|
| `Authorization` | Yes* | Bearer token (*except public routes) |
| `Content-Type` | Yes | `application/json` |
| `X-Request-ID` | No | Client-generated trace ID |

## Response Format

### Success Response
```json
{
  "success": true,
  "data": { ... },
  "meta": {
    "request_id": "uuid",
    "timestamp": "ISO8601"
  }
}
```

### Error Response
```json
{
  "success": false,
  "error": {
    "code": "ERROR_CODE",
    "message": "Human readable message",
    "details": { ... }
  },
  "meta": {
    "request_id": "uuid",
    "timestamp": "ISO8601"
  }
}
```

### Pagination
```json
{
  "success": true,
  "data": [ ... ],
  "pagination": {
    "page": 1,
    "per_page": 20,
    "total": 100,
    "total_pages": 5
  }
}
```

## HTTP Status Codes
| Code | Usage |
|------|-------|
| 200 | Success |
| 201 | Created |
| 204 | No Content (successful delete) |
| 400 | Bad Request (validation error) |
| 401 | Unauthorized (missing/invalid token) |
| 403 | Forbidden (insufficient permissions) |
| 404 | Not Found |
| 422 | Unprocessable Entity (business logic error) |
| 429 | Too Many Requests (rate limited) |
| 500 | Internal Server Error |

## Error Codes
| Code | HTTP | Description |
|------|------|-------------|
| `VALIDATION_ERROR` | 400 | Request payload validation failed |
| `UNAUTHORIZED` | 401 | Authentication required |
| `FORBIDDEN` | 403 | Insufficient permissions |
| `NOT_FOUND` | 404 | Resource not found |
| `RATE_LIMITED` | 429 | Too many requests |
| `INTERNAL_ERROR` | 500 | Unexpected server error |
```
