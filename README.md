# TaskFlow - Agile Team Productivity Suite

A self-hostable project management tool combining Kanban boards, sprint planning, and automated workflow customization.

## Tech Stack

- **Frontend**: Vue 3 + Vuetify + Pinia
- **Backend**: Go (Gin) + WebSocket
- **Database**: MongoDB
- **Auth**: JWT
- **Deployment**: Docker

## Key Features

- Drag-and-drop Kanban with WIP limits and swimlanes
- Customizable workflow automation engine (no-code rules)
- Sprint planning with velocity tracking and burndown charts
- Real-time collaboration with operational transforms
- Time tracking integrated with external APIs
- Advanced reporting with export to CSV/PDF/Excel

## Getting Started

### Prerequisites

- Go 1.19+
- Node.js (v16 or later)
- Docker and Docker Compose
- MongoDB

### Installation

1. Clone the repository
   ```bash
   git clone <repository-url>
   cd taskflow-agile-team-productivity-suite
   ```

2. Install backend dependencies
   ```bash
   cd backend
   go mod download
   ```

3. Install frontend dependencies
   ```bash
   cd ../frontend
   npm install
   ```

4. Set up environment variables
   Create a `.env` file in the backend directory based on `.env.example`.

5. Start the services
   ```bash
   # Start MongoDB via Docker Compose
   docker-compose up -d

   # Start the backend
   go run main.go

   # Start the frontend
   npm run serve
   ```

## API Endpoints

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/api/boards` | GET | List all boards |
| `/api/boards` | POST | Create a new board |
| `/api/boards/:id` | GET | Get a board by ID |
| `/api/tasks` | GET | List all tasks |
| `/api/tasks` | POST | Create a new task |
| `/api/tasks/:id` | PUT | Update a task |
| `/api/tasks/:id` | DELETE | Delete a task |
| `/ws` | WebSocket | Real-time updates |

## Real-time Collaboration

TaskFlow uses WebSocket for real-time updates. When a task is moved, updated, or created, all connected clients receive the update instantly.

### WebSocket Events

```json
{
  "type": "task-moved",
  "taskId": "task_123",
  "boardId": "board_456",
  "fromColumn": "todo",
  "toColumn": "in-progress"
}
```

## Deployment

### Local Development

```bash
docker-compose up -d
go run main.go
npm run serve
```

### Production

The application can be deployed as Docker containers or directly to a cloud provider.

## License

This project is licensed under the MIT License.

## Status (checkup 2026-08-18)
> Revisado na campanha de repo-checkup. Relatorio completo: `~/repo-checkup/reports/taskflow.md` (local do mantenedor, nao no repo).
- **Build/Install**: Backend Go — `go build ./...` RC=0; `go vet ./...` RC=0; `go.mod`/`go.sum` consistentes (`go mod tidy` sem mudancas). Frontend — `npm ci` + `npm run build` verdes (lockfile adicionado no checkup anterior).
- **Smoke test**: `go test ./...` -> 4 testes passando, incluindo `TestHealthCheck` (GET `/api/v1/health` -> 200 via httptest) e `TestBoardsListEndpoint` (GET `/api/v1/boards` -> 200); o handler de health/boards nao toca o DB.
- **Para rodar de ponta-a-ponta precisa de**: PostgreSQL alcancavel em `DATABASE_URL` (backend usa `lib/pq`/`sql.DB`; `initDB()` conecta de verdade ao rodar o server real). O smoke via httptest nao precisa. `docker-compose.yml` define mongodb+neo4j mas NAO bate com o codigo (usa Postgres).
- **Inconsistencias conhecidas (README vs codigo)**: README diz "Database: MongoDB" + "Auth: JWT", mas o backend usa PostgreSQL (`lib/pq`) e NAO tem JWT; README diz "Frontend: Vue 3 + Vuetify + Pinia", mas o frontend usa `vuex` (nao `pinia`) e Vuetify nao e importado em `src/main.js`/`App.vue`; README (install) diz `cd taskflow-agile-team-productivity-suite` (nome de dir incorreto; repo e `taskflow`); README referencia `.env.example` que NAO existe; `docker-compose.yml` (mongodb+neo4j) nao bate com o backend (Postgres).
- **Seguranca**: secret scan nenhum segredo real; `docker-compose.yml` tem credencial padrao fraca `NEO4J_AUTH=neo4j/secret` (nao e segredo real exposto; recomenda mover para `.env`). 12 vulns npm no toolchain do Vue CLI 5 — NAO corrigidas porque exigiria migracao para Vite ou `npm audit fix --force` (breaking); decisao humana.
- **Estado resumido**: build verde (Go + frontend) + smoke (health/boards 200 via httptest, sem DB); runtime real precisa de PostgreSQL; contradicoes de docs (Postgres vs Mongo, Vuex vs Pinia, JWT, nome do dir, `.env.example`) e 12 vulns npm de toolchain nao remediadas (decisao humana).

