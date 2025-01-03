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
├── activity1.md
├── activity2.md
├── activity3.md
├── activity4.md
├── activity5.md
├── assests
│   └── company_logo.png  # Company logo image
├── bin
│   └── ecom  # Compiled backend binary
├── docker-compose.yml  # Docker Compose configuration
├── Dockerfile  # Dockerfile for building the backend service
├── go
│   ├── api
│   │   ├── api  # API handlers
│   │   ├── configs  # Configuration files
│   │   ├── main.go  # Main entry point for the API
│   │   ├── services  # Business logic and services
│   │   ├── types  # Type definitions
│   │   ├── utils  # Utility functions
│   │   └── web  # Web server related files
│   ├── db
│   │   ├── db.go  # Database connection and setup
│   │   ├── functions  # Database functions
│   │   ├── migrate  # Migration scripts
│   │   └── migrations  # Migration files
│   └── worker
│       ├── client  # Worker client code
│       ├── common  # Common utilities for workers
│       ├── configs  # Worker configuration files
│       ├── convert  # Data conversion utilities
│       ├── db  # Database interactions for workers
│       ├── Dockerfile  # Dockerfile for building the worker service
│       ├── go.mod  # Go module file for worker
│       ├── go.sum  # Go dependencies for worker
│       ├── handler  # Handlers for worker tasks
│       ├── loader  # Data loading utilities
│       ├── main.go  # Main entry point for the worker
│       ├── util  # Utility functions for workers
│       ├── worker  # Worker implementation
│       └── worker-secrets.json  # Secrets for worker
├── go.mod  # Go module file for the project
├── go.sum  # Go dependencies for the project
├── LICENSE  # License file
├── Makefile  # Makefile for build automation
├── README.md  # Project documentation
└── web
    └── frontend
        ├── Dockerfile  # Dockerfile for building the frontend
        ├── env.d.ts  # TypeScript environment definitions
        ├── index.html  # Main HTML file for the frontend
        ├── node_modules  # Node.js dependencies
        ├── package.json  # Node.js project metadata
        ├── package-lock.json  # Lockfile for Node.js dependencies
        ├── public  # Public assets for the frontend
        ├── README.md  # Frontend documentation
        ├── src  # Source code for the frontend
        ├── tsconfig.app.json  # TypeScript configuration for the app
        ├── tsconfig.json  # Base TypeScript configuration
        ├── tsconfig.node.json  # TypeScript configuration for Node.js
        └── vite.config.ts  # Vite configuration file

27 directories, 33 files
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
        cd web/frontend
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
            go build , go run main.go
            ```
        - Frontend:
            ```sh
            cd web/frontend
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
    cd web/frontend
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
