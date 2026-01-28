# API User Manual

## Overview

This is a Go backend API service with JWT authentication, structured logging, and Swagger documentation.

## Base URL

- Development: `http://localhost:8080`

## Authentication

All protected endpoints require a JWT token in the Authorization header:

```
Authorization: Bearer <your-token>
```

### Getting a Token

```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username": "admin", "password": "admin123"}'
```

## Endpoints

### System

| Endpoint | Method | Auth | Description |
|----------|--------|------|-------------|
| `/health` | GET | No | Health check |
| `/api/docs` | GET | No | This documentation |
| `/swagger/*` | GET | No | Interactive API docs |

### Auth

| Endpoint | Method | Auth | Description |
|----------|--------|------|-------------|
| `/api/v1/auth/login` | POST | No | Get JWT token |
| `/api/v1/auth/me` | GET | Yes | Get current user |

### Test

| Endpoint | Method | Auth | Description |
|----------|--------|------|-------------|
| `/api/v1/ping` | GET | Yes | Ping test |
| `/api/v1/echo` | POST | Yes | Echo message |

## Response Format

### Success
```json
{
  "code": "SUCCESS",
  "message": "Operation description",
  "data": { ... }
}
```

### Error
```json
{
  "code": "ERROR_CODE",
  "message": "Error description",
  "data": null
}
```

## Error Codes

| Code | HTTP Status | Description |
|------|-------------|-------------|
| `VALIDATION_ERROR` | 400 | Invalid request |
| `UNAUTHORIZED` | 401 | Missing/invalid token |
| `NOT_FOUND` | 404 | Resource not found |
| `INTERNAL_ERROR` | 500 | Server error |
