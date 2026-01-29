# Test Templates

## Table of Contents
- [Pytest API Tests](#pytest-api-tests)
- [Playwright E2E Tests](#playwright-e2e-tests)
- [Test Case Template](#test-case-template)
- [Test Report Template](#test-report-template)

---

## Pytest API Tests

### Complete API Test Suite

```python
import pytest
import requests
from typing import Generator
import concurrent.futures

BASE_URL = "http://localhost:8080"

@pytest.fixture
def auth_headers() -> dict:
    """Get authenticated headers"""
    resp = requests.post(f"{BASE_URL}/api/login", json={
        "email": "test@example.com",
        "password": "password123"
    })
    token = resp.json().get("token")
    return {"Authorization": f"Bearer {token}"}


class TestAuthentication:
    """Auth middleware tests"""

    def test_no_token_returns_401(self):
        resp = requests.get(f"{BASE_URL}/api/users")
        assert resp.status_code == 401
        assert resp.json()["code"] == "UNAUTHORIZED"

    def test_invalid_token_returns_401(self):
        resp = requests.get(
            f"{BASE_URL}/api/users",
            headers={"Authorization": "Bearer invalid_token"}
        )
        assert resp.status_code == 401

    def test_valid_token_returns_200(self, auth_headers):
        resp = requests.get(f"{BASE_URL}/api/users", headers=auth_headers)
        assert resp.status_code == 200


class TestRequestID:
    """Traceability tests"""

    def test_request_id_in_response(self, auth_headers):
        resp = requests.get(f"{BASE_URL}/api/users", headers=auth_headers)
        assert "X-Request-ID" in resp.headers
        assert len(resp.headers["X-Request-ID"]) == 36  # UUID format

    def test_request_id_passthrough(self, auth_headers):
        custom_id = "test-request-123"
        resp = requests.get(
            f"{BASE_URL}/api/users",
            headers={**auth_headers, "X-Request-ID": custom_id}
        )
        assert resp.headers["X-Request-ID"] == custom_id


class TestDocsEndpoint:
    """Self-describing system tests"""

    def test_docs_accessible(self):
        resp = requests.get(f"{BASE_URL}/api/docs")
        assert resp.status_code == 200

    def test_docs_not_empty(self):
        resp = requests.get(f"{BASE_URL}/api/docs")
        content = resp.json().get("content", "")
        assert len(content) > 100, "Documentation is too short"

    def test_docs_is_markdown(self):
        resp = requests.get(f"{BASE_URL}/api/docs")
        content = resp.json().get("content", "")
        assert "#" in content, "Documentation should contain markdown headers"


class TestInputValidation:
    """Fuzzing and edge case tests"""

    def test_oversized_input(self, auth_headers):
        """Server should not panic on large input"""
        payload = {"name": "x" * 100000}
        resp = requests.post(
            f"{BASE_URL}/api/users",
            json=payload,
            headers=auth_headers
        )
        assert resp.status_code != 500, "Server panicked on large input"

    def test_null_values(self, auth_headers):
        """Server should handle null gracefully"""
        payload = {"name": None, "email": None}
        resp = requests.post(
            f"{BASE_URL}/api/users",
            json=payload,
            headers=auth_headers
        )
        assert resp.status_code in [400, 422], "Should reject null values"

    def test_empty_body(self, auth_headers):
        """Server should handle empty body"""
        resp = requests.post(
            f"{BASE_URL}/api/users",
            json={},
            headers=auth_headers
        )
        assert resp.status_code in [400, 422]

    def test_sql_injection_attempt(self, auth_headers):
        """Server should sanitize SQL injection"""
        payload = {"name": "'; DROP TABLE users; --"}
        resp = requests.post(
            f"{BASE_URL}/api/users",
            json=payload,
            headers=auth_headers
        )
        assert resp.status_code != 500


class TestConcurrency:
    """Concurrent request tests"""

    def test_concurrent_requests(self, auth_headers):
        """Server should handle concurrent requests"""
        def make_request():
            return requests.get(
                f"{BASE_URL}/api/users",
                headers=auth_headers
            )

        with concurrent.futures.ThreadPoolExecutor(max_workers=10) as executor:
            futures = [executor.submit(make_request) for _ in range(50)]
            results = [f.result() for f in futures]

        success_count = sum(1 for r in results if r.status_code == 200)
        assert success_count >= 45, f"Too many failures: {50 - success_count}/50"
```

---

## Playwright E2E Tests

### Complete E2E Test Suite

```typescript
// tests/e2e/auth.spec.ts
import { test, expect } from '@playwright/test';

test.describe('Authentication Flow', () => {
  test('should redirect to login when not authenticated', async ({ page }) => {
    await page.goto('/dashboard');
    await expect(page).toHaveURL(/.*login/);
  });

  test('should show error on invalid credentials', async ({ page }) => {
    await page.goto('/login');
    await page.fill('[data-testid="email"]', 'wrong@email.com');
    await page.fill('[data-testid="password"]', 'wrongpassword');
    await page.click('[data-testid="submit"]');

    await expect(page.locator('[data-testid="error"]')).toBeVisible();
  });

  test('should login successfully with valid credentials', async ({ page }) => {
    await page.goto('/login');
    await page.fill('[data-testid="email"]', 'test@example.com');
    await page.fill('[data-testid="password"]', 'password123');
    await page.click('[data-testid="submit"]');

    await expect(page).toHaveURL('/dashboard');
  });

  test('should persist session across page reload', async ({ page }) => {
    // Login first
    await page.goto('/login');
    await page.fill('[data-testid="email"]', 'test@example.com');
    await page.fill('[data-testid="password"]', 'password123');
    await page.click('[data-testid="submit"]');
    await expect(page).toHaveURL('/dashboard');

    // Reload and check still authenticated
    await page.reload();
    await expect(page).toHaveURL('/dashboard');
  });
});

test.describe('Error Handling', () => {
  test('should show 404 page for non-existent route', async ({ page }) => {
    await page.goto('/non-existent-page');
    await expect(page.locator('text=404')).toBeVisible();
  });

  test('should handle network error gracefully', async ({ page, context }) => {
    await context.route('**/api/**', route => route.abort());
    await page.goto('/dashboard');

    await expect(page.locator('[data-testid="error-boundary"]')).toBeVisible();
  });
});

test.describe('Form Validation', () => {
  test('should show validation errors for empty form', async ({ page }) => {
    await page.goto('/register');
    await page.click('[data-testid="submit"]');

    await expect(page.locator('[data-testid="email-error"]')).toBeVisible();
    await expect(page.locator('[data-testid="password-error"]')).toBeVisible();
  });

  test('should validate email format', async ({ page }) => {
    await page.goto('/register');
    await page.fill('[data-testid="email"]', 'invalid-email');
    await page.click('[data-testid="submit"]');

    await expect(page.locator('[data-testid="email-error"]')).toContainText('valid email');
  });
});
```

---

## Test Case Template

```markdown
# Test Case: [TC-XXX] [Feature Name]

## Metadata
- **Author**: QA Engineer
- **Date**: YYYY-MM-DD
- **Priority**: Critical | High | Medium | Low
- **Type**: Functional | Integration | E2E | Security

## Preconditions
- [ ] [Precondition 1]
- [ ] [Precondition 2]

## Test Data
| Field | Value |
|-------|-------|
| Email | test@example.com |
| Password | password123 |

## Test Steps

| Step | Action | Expected Result |
|------|--------|-----------------|
| 1 | Navigate to /login | Login page displayed |
| 2 | Enter email | Email field populated |
| 3 | Enter password | Password field masked |
| 4 | Click Submit | Redirect to dashboard |

## Edge Cases
- [ ] Empty email
- [ ] Invalid email format
- [ ] Wrong password
- [ ] Account locked

## Pass Criteria
- [ ] All steps complete without error
- [ ] Response time < 2 seconds
- [ ] No console errors
```

---

## Test Report Template

```markdown
# Test Report

**Date**: YYYY-MM-DD
**Environment**: Staging | Production
**Build**: v1.2.3
**Tester**: QA Engineer

---

## Summary

| Metric | Value |
|--------|-------|
| Total Tests | 50 |
| Passed | 45 |
| Failed | 3 |
| Skipped | 2 |
| Pass Rate | 90% |

---

## Failed Tests

### 1. [TC-001] Login with invalid email

**Status**: FAILED
**Severity**: High

**Expected**: Error message "Invalid email format"
**Actual**: Server returned 500 Internal Error

**Bug ID**: BUG-123
**Assigned To**: Backend Team

---

### 2. [TC-015] Concurrent user registration

**Status**: FAILED
**Severity**: Critical

**Expected**: All registrations succeed
**Actual**: Race condition causes duplicate key error

**Bug ID**: BUG-124
**Assigned To**: Backend Team

---

## Recommendations

1. Fix critical bugs before release
2. Add rate limiting to prevent abuse
3. Improve error messages for user clarity

---

## Sign-off

- [ ] QA Lead Reviewed
- [ ] Dev Lead Acknowledged
- [ ] PM Notified

**Final Status**: REJECTED / APPROVED
```
