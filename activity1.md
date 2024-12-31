# Activity 1: **(Running the E-Commerce Store) Running the API , Frontend and Worker **

_Transmission decrypted... loading instructions..._

In this **cyberpunk** scenario, the data smuggler's Docker pipeline has encountered an issue. Your mission: **Run the applications independently**. Failure is not an option.

**Toolkit:**
- **Local Terminal**
- **Install dependencies and run the applications using your local terminal**
- **Install Docker and Docker Hub on your system**

**Objective:**
1. **Run the PostgreSQL container:**
   ```bash
   docker run --name ecom-db-1 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=0845635040 -e POSTGRES_DB=postgres -p 5434:5432 -d postgres:14
   ```
2. **Access the database container using a database client (e.g., DBeaver or pgAdmin).**
3. **Seed the database with the command.**
 ```bash
   make migrate-up
   ```
4. **Need to add the notify listener function into the database. the function is found in the db folder**

5. **Navigate to the frontend directory and execute the following commands:**
   ```bash
   npm install
   npm run dev
   ```

6. **Navigate to the api directory and execute the following commands:**
   ```bash
   go build
   go run main.go
   ```

7. **Navigate to the worker directory and execute the following commands:**
   ```bash
   go build
   go run main.go
   ```

---

### Important Notes

**Useful Commands:**

```bash
# List all running containers
docker ps

# Restart a specific container
docker restart <container_id>

# Build and start containers using Docker Compose
docker-compose up --build

# Access the ecom-db-1 container
docker exec -it ecom-db-1 bash

# Access PostgreSQL within the container
psql -U postgres

# Change the PostgreSQL user password
ALTER USER postgres PASSWORD '<password>';

# View logs of a specific container
docker logs <container_name>

# Run database migrations
make migrate-up

# Build the Go application
go build

# Run the Go application
go run
```

