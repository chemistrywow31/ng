# Infrastructure Templates

## Table of Contents
- [Dockerfile Templates](#dockerfile-templates)
- [docker-compose Templates](#docker-compose-templates)
- [GitHub Actions Templates](#github-actions-templates)
- [Makefile Templates](#makefile-templates)

---

## Dockerfile Templates

### Go Backend
```dockerfile
# Build Stage
FROM golang:1.21-alpine AS builder
WORKDIR /app

# Install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/server ./cmd/api

# Run Stage
FROM alpine:3.19
RUN apk --no-cache add ca-certificates tzdata
WORKDIR /app

COPY --from=builder /app/server .
COPY --from=builder /app/docs ./docs

EXPOSE 8080
USER nobody:nobody
CMD ["./server"]
```

### Node.js Backend
```dockerfile
# Build Stage
FROM node:20-alpine AS builder
WORKDIR /app

COPY package*.json ./
RUN npm ci --only=production

COPY . .
RUN npm run build

# Run Stage
FROM node:20-alpine
WORKDIR /app

COPY --from=builder /app/dist ./dist
COPY --from=builder /app/node_modules ./node_modules
COPY --from=builder /app/package.json ./

EXPOSE 3000
USER node
CMD ["node", "dist/index.js"]
```

### Next.js Frontend
```dockerfile
# Dependencies
FROM node:20-alpine AS deps
WORKDIR /app
COPY package*.json ./
RUN npm ci

# Builder
FROM node:20-alpine AS builder
WORKDIR /app
COPY --from=deps /app/node_modules ./node_modules
COPY . .
RUN npm run build

# Runner
FROM node:20-alpine AS runner
WORKDIR /app
ENV NODE_ENV=production

RUN addgroup --system --gid 1001 nodejs
RUN adduser --system --uid 1001 nextjs

COPY --from=builder /app/public ./public
COPY --from=builder --chown=nextjs:nodejs /app/.next/standalone ./
COPY --from=builder --chown=nextjs:nodejs /app/.next/static ./.next/static

USER nextjs
EXPOSE 3000
CMD ["node", "server.js"]
```

### Python FastAPI
```dockerfile
# Build Stage
FROM python:3.12-slim AS builder
WORKDIR /app

RUN pip install --no-cache-dir poetry
COPY pyproject.toml poetry.lock ./
RUN poetry export -f requirements.txt --output requirements.txt

# Run Stage
FROM python:3.12-slim
WORKDIR /app

COPY --from=builder /app/requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

COPY . .

EXPOSE 8000
CMD ["uvicorn", "app.main:app", "--host", "0.0.0.0", "--port", "8000"]
```

---

## docker-compose Templates

### Full Stack (Go + Next.js + Postgres + Redis)
```yaml
version: '3.8'

services:
  backend:
    build:
      context: ./backend
      dockerfile: Dockerfile
    environment:
      - DB_HOST=db
      - DB_PORT=5432
      - DB_USER=postgres
      - DB_PASSWORD=postgres
      - DB_NAME=app
      - REDIS_URL=redis://redis:6379
      - JWT_SECRET=dev-secret-change-in-prod
    ports:
      - "8080:8080"
    depends_on:
      db:
        condition: service_healthy
      redis:
        condition: service_healthy
    volumes:
      - ./backend:/app:ro
    restart: unless-stopped

  frontend:
    build:
      context: ./frontend
      dockerfile: Dockerfile
    environment:
      - NEXT_PUBLIC_API_URL=http://localhost:8080
    ports:
      - "3000:3000"
    depends_on:
      - backend
    restart: unless-stopped

  db:
    image: postgres:15-alpine
    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: postgres
      POSTGRES_DB: app
    volumes:
      - ./infra/db/init.sql:/docker-entrypoint-initdb.d/init.sql
      - postgres_data:/var/lib/postgresql/data
    ports:
      - "5432:5432"
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U postgres"]
      interval: 5s
      timeout: 5s
      retries: 5
    restart: unless-stopped

  redis:
    image: redis:7-alpine
    command: redis-server --appendonly yes
    volumes:
      - redis_data:/data
    ports:
      - "6379:6379"
    healthcheck:
      test: ["CMD", "redis-cli", "ping"]
      interval: 5s
      timeout: 5s
      retries: 5
    restart: unless-stopped

volumes:
  postgres_data:
  redis_data:

networks:
  default:
    name: app-network
```

### With Logging (ELK Stack)
```yaml
version: '3.8'

services:
  app:
    build: .
    logging:
      driver: "json-file"
      options:
        max-size: "10m"
        max-file: "3"
    labels:
      - "logging=true"

  filebeat:
    image: elastic/filebeat:8.11.0
    volumes:
      - ./infra/filebeat.yml:/usr/share/filebeat/filebeat.yml:ro
      - /var/lib/docker/containers:/var/lib/docker/containers:ro
      - /var/run/docker.sock:/var/run/docker.sock:ro
    depends_on:
      - elasticsearch

  elasticsearch:
    image: elasticsearch:8.11.0
    environment:
      - discovery.type=single-node
      - xpack.security.enabled=false
    volumes:
      - es_data:/usr/share/elasticsearch/data
    ports:
      - "9200:9200"

  kibana:
    image: kibana:8.11.0
    environment:
      - ELASTICSEARCH_HOSTS=http://elasticsearch:9200
    ports:
      - "5601:5601"
    depends_on:
      - elasticsearch

volumes:
  es_data:
```

---

## GitHub Actions Templates

### Complete CI/CD Pipeline
```yaml
name: CI/CD

on:
  push:
    branches: [main, develop]
  pull_request:
    branches: [main]

env:
  REGISTRY: ghcr.io
  IMAGE_NAME: ${{ github.repository }}

jobs:
  lint:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with:
          go-version: '1.21'
      - name: golangci-lint
        uses: golangci/golangci-lint-action@v3

  test:
    runs-on: ubuntu-latest
    services:
      postgres:
        image: postgres:15-alpine
        env:
          POSTGRES_USER: test
          POSTGRES_PASSWORD: test
          POSTGRES_DB: test
        options: >-
          --health-cmd pg_isready
          --health-interval 10s
          --health-timeout 5s
          --health-retries 5
        ports:
          - 5432:5432
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with:
          go-version: '1.21'
      - run: go test -v -race -coverprofile=coverage.out ./...
      - uses: codecov/codecov-action@v3
        with:
          files: ./coverage.out

  build:
    needs: [lint, test]
    runs-on: ubuntu-latest
    permissions:
      contents: read
      packages: write
    steps:
      - uses: actions/checkout@v4
      - uses: docker/login-action@v3
        with:
          registry: ${{ env.REGISTRY }}
          username: ${{ github.actor }}
          password: ${{ secrets.GITHUB_TOKEN }}
      - uses: docker/build-push-action@v5
        with:
          context: .
          push: ${{ github.event_name != 'pull_request' }}
          tags: ${{ env.REGISTRY }}/${{ env.IMAGE_NAME }}:${{ github.sha }}

  deploy:
    needs: build
    if: github.ref == 'refs/heads/main'
    runs-on: ubuntu-latest
    environment: production
    steps:
      - name: Deploy to production
        run: |
          echo "Deploy ${{ github.sha }} to production"
          # Add deployment commands here
```

### Node.js CI
```yaml
name: Node.js CI

on:
  push:
    branches: [main]
  pull_request:
    branches: [main]

jobs:
  build:
    runs-on: ubuntu-latest
    strategy:
      matrix:
        node-version: [18.x, 20.x]
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-node@v4
        with:
          node-version: ${{ matrix.node-version }}
          cache: 'npm'
      - run: npm ci
      - run: npm run lint
      - run: npm run build
      - run: npm test
```

---

## Makefile Templates

### Full Project Makefile
```makefile
.PHONY: help up down logs build test lint clean deploy

# Default target
help:
	@echo "Available commands:"
	@echo "  make up      - Start all services"
	@echo "  make down    - Stop all services"
	@echo "  make logs    - View logs"
	@echo "  make build   - Build Docker images"
	@echo "  make test    - Run tests"
	@echo "  make lint    - Run linters"
	@echo "  make clean   - Clean up"
	@echo "  make deploy  - Deploy to production"

# Development
up:
	docker-compose up -d --build

down:
	docker-compose down

logs:
	docker-compose logs -f

restart:
	docker-compose restart

# Building
build:
	docker build -t app:latest .

build-prod:
	docker build -t app:$(shell git rev-parse --short HEAD) --target production .

# Testing
test:
	docker-compose exec backend go test -v ./...

test-coverage:
	docker-compose exec backend go test -coverprofile=coverage.out ./...
	docker-compose exec backend go tool cover -html=coverage.out -o coverage.html

# Linting
lint:
	docker-compose exec backend golangci-lint run

# Database
db-migrate:
	docker-compose exec backend go run cmd/migrate/main.go up

db-rollback:
	docker-compose exec backend go run cmd/migrate/main.go down

db-reset:
	docker-compose down -v
	docker-compose up -d db
	sleep 5
	$(MAKE) db-migrate

# Cleanup
clean:
	docker-compose down -v --rmi local
	docker system prune -f

# Deployment
deploy:
	@echo "Deploying $(shell git rev-parse --short HEAD)..."
	# Add deployment commands
```
