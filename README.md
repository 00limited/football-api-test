# football-api

REST API backend for managing football teams, players, matches, results, and reports using Go, Gin, GORM, and JWT authentication.

## Requirements

- Go 1.25.0+
- PostgreSQL (default runtime database)

## Setup

1. Copy the example environment file:
   ```bash
   cp .env.example .env
   ```
2. Update database and JWT settings in `.env`.
3. Install dependencies:
   ```bash
   go mod tidy
   ```
4. Run the server:
   ```bash
   go run ./cmd/main.go
   ```

> `DB_DRIVER=sqlite` with `DB_DSN=football_api.db` is also supported for quick local development.

## API

### Auth
- `POST /api/v1/auth/register`
- `POST /api/v1/auth/login`

### Teams
- `GET /api/v1/teams`
- `POST /api/v1/teams`
- `GET /api/v1/teams/:id`
- `PUT /api/v1/teams/:id`
- `DELETE /api/v1/teams/:id`
- `POST /api/v1/teams/:id/logo`

### Players
- `GET /api/v1/teams/:teamId/players`
- `POST /api/v1/teams/:teamId/players`
- `GET /api/v1/players/:id`
- `PUT /api/v1/players/:id`
- `DELETE /api/v1/players/:id`

### Matches
- `GET /api/v1/matches`
- `POST /api/v1/matches`
- `GET /api/v1/matches/:id`
- `PUT /api/v1/matches/:id`
- `DELETE /api/v1/matches/:id`
- `POST /api/v1/matches/:id/result`
- `GET /api/v1/matches/:id/result`

### Reports
- `GET /api/v1/reports/matches`
- `GET /api/v1/reports/matches/:id`

## Notes

- GORM auto-migration runs on startup.
- Delete operations use GORM soft delete.
- Match results can only be submitted once and automatically set the match to `FINISHED`.
- Match scores are calculated from submitted goals.
- Protected routes require `Authorization: ******
