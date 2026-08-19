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
   cd taskflow
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