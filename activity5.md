# Activity 1: **(Running the E-Commerce Store) Running the docker containers**

_Transmission decrypted... loading instructions..._

In this **cyberpunk** scenario, the data smuggler's docker pipeline has glitched. The root cause? An outdated password hidden deep inside the steel-and-chrome bowels of your container image. Your mission: **update the password and rebuild the containers**. Failure is not an option.

**Your Toolkit:**
- **Docker CLI**
- **Use your local terminal to perform container updates**

**Your Objective:**
1. **Go to each .env in the go/api , go/worker , ./ecom and complete the instructions in those files**
1. **Run the project with docker ps --build** (in the main directory).
2. **Check the container logs to fix the database error** .
3. **Fix the errors and rebuild the containers**.
4. **Verify success** by logging into the ecom-db-1 container.

---

### Important things to note

1. **Useful Commands:**

   ```bash
   # Use these docker commands to interact with the list of containers
    docker ps
   docker restart <container_id> 
   docker-compose up --build
   docker exec -it ecom-db-1 bash
   psql -U postgres
   ALTER USER postgres PASSWORD <password> ;
   docker logs <container_name>
   make migrate-up
   go build
   go run

