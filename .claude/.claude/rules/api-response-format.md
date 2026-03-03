---
name: API Response Format
description: Enforce a unified API response structure across all Go and TypeScript endpoints
paths:
  - "**/*.go"
  - "**/*.ts"
  - "**/*.tsx"
---

# API Response Format

## Applicability

- Applies to: `backend-developer`, `frontend-developer`, `code-reviewer`

## Rule Content

### Unified Response Structure

All API responses must use this format:

```json
{
  "code": 0,
  "message": "success",
  "data": <T | null>
}
```

- `code`: `0` for success, non-zero for errors
- `message`: A human-readable description of the result
- `data`: The response payload, or `null` when no data is returned

### Pagination Response Structure

All paginated endpoints must return data in this format:

```json
{
  "code": 0,
  "message": "success",
  "data": [<T>],
  "pagination": {
    "page": 1,
    "pageSize": 20,
    "total": 150,
    "totalPages": 8
  }
}
```

You must include all four pagination fields (`page`, `pageSize`, `total`, `totalPages`).

### HTTP Status Code Mapping

You must use HTTP status codes that match their semantic meaning:

| Status Code | Usage |
|---|---|
| 200 OK | Successful GET, PUT, PATCH, DELETE |
| 201 Created | Successful POST that creates a resource |
| 400 Bad Request | Malformed request, validation failure |
| 401 Unauthorized | Missing or invalid authentication |
| 403 Forbidden | Authenticated but insufficient permissions |
| 404 Not Found | Resource does not exist |
| 500 Internal Server Error | Unhandled server-side error |

### Error Code Documentation

You must document every non-zero error code in the endpoint's JSDoc (TypeScript) or Swagger annotation (Go). Each error code entry must include the numeric code, the message, and the condition that triggers it.

### Language-Specific Implementation

**Go**: You must use the `cwsutil` package's response helpers to construct all API responses. You must not construct response JSON manually.

**TypeScript / Next.js**: You must use `NextResponse.json()` with the unified format. You must not return plain objects or use `res.send()` / `res.json()` outside the unified structure.

## Violation Determination

- API endpoint returns a response not matching the unified format (`code`, `message`, `data`) → Violation
- Paginated endpoint omits any of the four pagination fields → Violation
- HTTP status code does not match its semantic meaning (e.g., returning 200 for validation errors) → Violation
- Non-zero error code exists without JSDoc or Swagger documentation → Violation
- Go endpoint constructs response JSON manually instead of using `cwsutil` helpers → Violation
- TypeScript endpoint uses `res.send()` or `res.json()` instead of `NextResponse.json()` with the unified format → Violation

## Exceptions

- Health check endpoints (`/health`, `/ready`) may return a plain `{ "status": "ok" }` response without the unified wrapper.
