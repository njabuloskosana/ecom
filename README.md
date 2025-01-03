<p align="center">
  <img src="./assests/company_logo.png" alt="Company Logo" width="200px">
</p>

# Isazi (Sophia) Software Development Onboarding : E-Commerce Platform

Welcome to the E-Commerce Platform repository. This monorepo hosts all components for a complete e-commerce solution, including backend services, database migrations, and a frontend UI. This documentation will guide you through setup, development workflows, and best practices.

## Tasks And Project Overview

## Excercise 1
[Go to Excercise 1](./activity1.md)
## Excercise 2
[Go to Excercise 2](./activity2.md)
## Excercise 3
[Go to Excercise 3](./activity3.md)
## Excercise 4
[Go to Excercise 4](./activity4.md)

## Bonus
[Go to Excercise 5](./activity5.md)

### Technologies

- **Go**: Backend API and services.
- **PostgreSQL**: Primary data store.
- **Vite + TypeScript**: Modern frontend stack.
- **Docker & Docker Compose**: Containerized local development and deployments.

### Features

- Modular services for authentication, user, product, cart, and order management.
- Secure JWT-based authentication.
- Automated database migrations.
- CI-ready testing and linting scripts.

## Repository Structure

```
.
├── bin
│   └── ecom # Compiled backend binary
├── cmd
│   ├── api
│   │   └── api.go # API server entrypoint
│   ├── main.go # Main CLI entrypoint
│   └── migrate
│       └── main.go # Migrate CLI entrypoint
├── configs
│   └── env.go # Environment variable configuration
├── db
│   ├── db.go # Database connection logic
│   └── migrations # SQL migration files
│       ├── 
│       └── 
├── docker-compose.yml # Docker services definition
├── Dockerfile # Dockerfile for backend
├── go.mod, go.sum # Go dependencies
├── LICENSE # Project license
├── Makefile # Common build/test tasks
├── migrate, migrate.tar.gz # Migration tools
├── README.md # This file
├── services
│   ├── auth
│   │   ├── jwt.go # JWT utilities
│   │   └── password.go # Password hashing
│   ├── cart
│   │   ├── routes.go # Cart endpoints
│   │   └── service.go # Cart business logic
│   ├── order
│   │   └── store.go # Order persistence
│   ├── product
│   │   ├── routes.go # Product endpoints
│   │   └── store.go # Product persistence
│   └── user
│       ├── routes.go # User endpoints
│       ├── routes_test.go # Integration tests
│       ├── routes_unit_test.go # Unit tests
│       └── store.go # User persistence
├── types
│   └── types.go # Shared data types
├── UI
│   └── ecom
│       ├── Dockerfile # Frontend Dockerfile
│       ├── env.d.ts # Frontend env typings
│       ├── index.html # Frontend entrypoint
│       ├── node_modules # Node dependencies (not committed)
│       ├── package.json # Node.js project config
│       ├── package-lock.json
│       ├── public # Public assets
│       ├── README.md # Frontend-specific instructions
│       ├── src # Frontend source (components, pages)
│       ├── tsconfig.*.json # TypeScript config files
│       └── vite.config.ts # Vite configuration
└── utils
    └── utils.go # Shared utility functions
```

## Getting Started

### Prerequisites

- Go (1.20+ recommended)
- Node.js & npm (LTS recommended)
- Docker & Docker Compose
- Make (Linux/Mac Os)

### Setup Instructions

1. **Clone the Repository**:
    ```sh
    git clone git@github.com:your-company/ecom-platform.git
    cd ecom-platform
    ```

2. **Environment Setup**:
    ```sh
    cp configs/env.example configs/env.go
    ```
    Update environment variables (DB credentials, JWT secrets, etc.) as needed.

3. **Install Dependencies**:
    - Backend:
        ```sh
        go mod tidy
        ```
    - Frontend:
        ```sh
        cd UI/ecom
        npm install
        cd ../..
        ```

4. **Run Migrations**:
    ```sh
    make migrate-up
    ```

5. **Start the Application**:
    - Using Docker:
        ```sh
        docker-compose up --build
        ```
        This will start the API server and database.
    - Locally (non-Docker):
        - Backend:
            ```sh
            make run
            ```
        - Frontend:
            ```sh
            cd UI/ecom
            npm run dev
            ```

    By default, the API runs on `http://localhost:8080` and the frontend on `http://localhost:3000`.

### Common Tasks

- **Build Backend**:
    ```sh
    make build
    ```
    Binary located in `./bin/ecom`

- **Test Backend**:
    ```sh
    make test
    ```

- **Lint Backend Code**:
    ```sh
    make lint
    ```

- **Frontend Development**:
    ```sh
    cd UI/ecom
    npm run dev
    ```

## Database Migrations

Migrations live in `db/migrations`.

- **Create a new migration**:
    ```sh
    migrate create -ext sql -dir db/migrations -seq create_new_table
    ```

- **Apply migrations**:
    ```sh
    make migrate-up
    ```

- **Rollback migrations**:
    ```sh
    make migrate-down
    ```

## Contributing

- **Branching & Pull Requests**: Follow the team's Git workflow.
- **Code Reviews**: All code changes must be reviewed and pass CI checks before merging.
- **Testing**: Ensure features have sufficient test coverage. Add or update tests as needed.
- **Documentation**: Keep the README and inline comments up to date.

## Troubleshooting

- **Service Not Starting**: Check environment variables or `docker-compose` logs.
- **Database Issues**: Verify that migrations have been applied. Check connection credentials.
- **Frontend Errors**: Reinstall dependencies (`rm -rf node_modules && npm install`).

## License

This project is licensed under the terms of the LICENSE file.

## Contact

For questions or support, contact the engineering team via internal channels or Slack (`#ecom-dev`).

### Additional Commands

- **Useful Commands**:
    ```sh
    docker build -t <container_name>
    docker restart <container_id>
    docker-compose up --build
    make migrate-up
    make run
    pg_restore -h localhost -p 5433 -U postgres -d postgres -W -v <file_name.mimetype>
    ```

**Note**: Check how Docker allows you to run everything on one network.
