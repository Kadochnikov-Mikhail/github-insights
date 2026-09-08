# GitHub Insights

A full-stack web application for analyzing GitHub profiles and repository activity.

Enter a GitHub username to get an overview of their repositories, total stars, and programming language distribution. Each analysis is also stored in PostgreSQL, allowing the backend to keep historical snapshots for future analytics.

## Features

- GitHub profile analysis
- Repository count
- Total stars across repositories
- Programming language distribution
- Historical insight snapshots stored in PostgreSQL
- REST API built with Go and Fiber
- React frontend with TypeScript
- PostgreSQL persistence
- Dockerized development and production environment
- Nginx reverse proxy
- Automated tests
- GitHub Actions CI
- Go vet checks
- Frontend linting and production build

## Tech Stack

### Frontend

- React 19
- TypeScript
- Vite
- CSS
- pnpm

### Backend

- Go 1.26.2
- Fiber v2
- PostgreSQL
- `database/sql`
- GitHub REST API

### DevOps

- Docker
- Docker Compose
- Nginx
- GitHub Actions

## Architecture

```text
                    ┌───────────────┐
                    │    Browser    │
                    └───────┬───────┘
                            │
                            ▼
                    ┌───────────────┐
                    │     Nginx     │
                    │  React + Proxy│
                    └───────┬───────┘
                            │
                            ▼
                    ┌───────────────┐
                    │  Go / Fiber   │
                    │    Backend    │
                    └───────┬───────┘
                            │
                 ┌──────────┴──────────┐
                 ▼                     ▼
        ┌────────────────┐    ┌────────────────┐
        │   GitHub API   │    │   PostgreSQL   │
        └────────────────┘    └────────────────┘
```

The frontend communicates with the backend through Nginx.

Requests to `/api/*` are proxied to the Go backend, while the backend communicates with the GitHub API and PostgreSQL.

## Project Structure

```text
github-insights/
├── backend/
│   ├── cmd/
│   │   └── api/
│   │       └── main.go
│   ├── internal/
│   │   ├── apperror/
│   │   ├── client/
│   │   ├── config/
│   │   ├── database/
│   │   ├── handlers/
│   │   ├── middleware/
│   │   ├── models/
│   │   ├── repository/
│   │   ├── routes/
│   │   └── services/
│   ├── migrations/
│   ├── Dockerfile
│   ├── .env.example
│   ├── go.mod
│   └── go.sum
│
├── frontend/
│   ├── src/
│   │   ├── components/
│   │   ├── App.tsx
│   │   ├── App.css
│   │   └── main.tsx
│   ├── Dockerfile
│   ├── nginx.conf
│   ├── package.json
│   └── pnpm-lock.yaml
│
├── .github/
│   └── workflows/
│       └── ci.yml
│
├── docker-compose.yml
└── README.md
```

## API

### Health Check

```http
GET /health
```

Returns the current API health status.

### Get GitHub User

```http
GET /github?user=<username>
```

Returns GitHub user information.

### Get GitHub Insights

```http
GET /github/insights?user=<username>
```

Analyzes the user's repositories and returns:

```json
{
  "username": "octocat",
  "repositories": 8,
  "total_stars": 120,
  "languages": {
    "Go": 3,
    "TypeScript": 2,
    "JavaScript": 2,
    "Python": 1
  },
  "created_at": "2026-01-01T12:00:00Z"
}
```

### Get Insights History

```http
GET /github/insights/history?user=<username>
```

Returns previously stored insight snapshots for a GitHub username.

## Running with Docker

The easiest way to run the complete application locally is Docker Compose.

### Requirements

- Docker
- Docker Compose

### Start the application

From the project root:

```bash
docker compose up --build
```

The application will start:

- Frontend: `http://localhost`
- Backend: `http://localhost:3000`
- PostgreSQL: `localhost:5432`

Open the frontend in your browser:

```text
http://localhost
```

### Stop the application

```bash
docker compose down
```

To remove the PostgreSQL data volume as well:

```bash
docker compose down -v
```

## Environment Variables

Backend configuration is provided through environment variables.

Example:

```env
PORT=3000
GITHUB_API_URL=https://api.github.com
DATABASE_URL=postgres://admin:password@localhost:5432/github_insights?sslmode=disable
JWT_SECRET=
```

For local development, create:

```text
backend/.env
```

based on:

```text
backend/.env.example
```

The `.env` file is intentionally excluded from Git.

## Running Backend Locally

Go 1.26.2 is required.

From the backend directory:

```bash
cd backend
go mod download
go run ./cmd/api
```

The API will be available at:

```text
http://localhost:3000
```

Run backend tests:

```bash
go test ./...
```

Run tests with coverage:

```bash
go test ./... -coverprofile=coverage.out
```

Run static analysis:

```bash
go vet ./...
```

## Running Frontend Locally

Node.js and pnpm are required.

From the frontend directory:

```bash
cd frontend
pnpm install
pnpm dev
```

Run linting:

```bash
pnpm lint
```

Build the production frontend:

```bash
pnpm build
```

## Testing

The backend includes unit and integration tests covering:

- GitHub API client
- Services
- HTTP handlers
- Middleware
- Repository layer
- PostgreSQL persistence

The CI pipeline runs:

```text
Backend
├── PostgreSQL service
├── Database schema
├── go test ./...
├── test coverage
└── go vet ./...

Frontend
├── pnpm install
├── ESLint
└── production build
```

## CI

GitHub Actions automatically validates pushes and pull requests targeting:

- `master`
- `develop`

The CI pipeline checks both the backend and frontend before changes are considered ready.

## Database

PostgreSQL stores GitHub insight snapshots.

Current schema:

```text
github_insights
├── id
├── username
├── repositories
├── total_stars
└── created_at
```

Every successful insight analysis creates a new snapshot in the database.

## Development Approach

The project is structured around separation of responsibilities:

```text
HTTP Request
     ↓
Handler
     ↓
Service
     ↓
Client / Repository
     ↓
External API / Database
```

Interfaces are used for dependencies such as the GitHub client and repository, making the business logic easier to test and mock.

Error handling is centralized through Fiber's custom error handler.

## Current Status

**Beta v1 — functional**

The core application is complete and runs locally through Docker Compose.

Current focus:

- Production deployment
- CI/CD deployment pipeline
- Public demo
- GitHub repository polish

## Roadmap

### v2

Planned features include:

- User authentication
- JWT-based authorization
- User profiles
- Saved GitHub profiles
- Historical charts
- Repository growth analytics
- Additional GitHub statistics

## License

This project is currently intended as a portfolio project.
