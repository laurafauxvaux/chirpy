# Chirpy

Chirpy is a small Twitter-like backend API written in Go as part of the Boot.dev backend curriculum, designed to practice building an HTTP server backed by PostgreSQL, with authentication, authorization, database migrations, and third-party webhooks.


## Features

- Create and update users
- Password hashing with Argon2id
- JWT access-token authentication
- Refresh-token creation and revocation, with access-token renewal
- Create, retrieve, filter, sort, and delete chirps
- Limits chirps to 140 characters and filters designated profanities
- User authorization: authenticated users can only update their own account and delete their own chirps.
- PostgreSQL persistence
- Go Database code generated with SQLC
- Database migrations with Goose
- Chirpy Red membership upgrades through authenticated Polka webhooks
- Basic file-server visit metrics and health-check endpoints


## Technologies used

- Go
- PostgreSQL
- SQLC
- Goose

## How To Run Locally

### Prerequisites

- Go 1.25 or later
- PostgreSQL
- [Goose](https://github.com/pressly/goose) for database migrations
- SQLC

### Setup

Clone the repository:

```bash
git clone https://github.com/laurafauxvaux/chirpy.git
cd chirpy
```
Create a PostgreSQL database for Chirpy, then create a `.env` file at the root of the project:
```env
DB_URL=postgres://<user>:<password>@localhost:5432/chirpy
PLATFORM=dev
SECRET=<your-jwt-secret>
POLKA_KEY=<your-polka-api-key>
```

Apply the database migrations:
```bash
goose -dir sql/schema postgres "$DB_URL" up
```

Download the Go dependencies:
```bash
go mod download
```

Run the server:
```bash
go run .
```

The API will be available at http://localhost:8080

### Development
Database tables changes should be applied with a new Goose migration after adding file in `sql/schema`. At the root of the project:
```bash
goose -dir sql/schema postgres "$DB_URL" up
```
SQL queries are managed with SQLC. After modifying queries in `sql/queries`, (re) generate the database code with:
```bash
sqlc generate
```

Run the tests with:
```bash
go test ./...
```

## API

Some of the main endpoints implemented in the project:

| Method | Endpoint | Auth | Description |
| --- | --- | --- | --- |
| POST | `/api/users` | None | Create a user |
| PUT | `/api/users` | Bearer access token | Update the authenticated user |
| POST | `/api/login` | None | Log in and receive access and refresh tokens |
| POST | `/api/refresh` | Bearer refresh token | Refresh an access token |
| POST | `/api/revoke` | Bearer refresh token | Revoke a refresh token |
| POST | `/api/chirps` | Bearer access token | Create a chirp |
| GET | `/api/chirps` | None | Get chirps, with optional author filtering and sorting |
| GET | `/api/chirps/{chirpID}` | None | Get a chirp by ID |
| DELETE | `/api/chirps/{chirpID}` | Bearer access token | Delete one of the authenticated user's chirps |
| POST | `/api/polka/webhooks` | API key | Process Chirpy Red membership webhooks |
| GET | `/api/healthz` | None | Health check |

### Filtering chirps

Chirps can be filtered by author:

`GET /api/chirps?author_id=<user-id>`

### Sorting chirps

Chirps can be sorted by creation date:

`GET /api/chirps?sort=asc`

`GET /api/chirps?sort=desc`

## What I practiced

This project gave me hands-on practice with:

- HTTP routing and handlers in Go
- REST-style API design
- JSON request and response handling
- PostgreSQL and SQL
- SQLC-generated database code
- Database migrations
- Authentication vs. authorization
- Password hashing
- JWT and refresh-token flows
- Webhooks and API-key authentication
- Query parameters for filtering and sorting
- Error handling and HTTP status codes