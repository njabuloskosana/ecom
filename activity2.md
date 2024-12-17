# Activity 2: **(Running the Migrations) Seed the Postgresql Database**

_Transmission decrypted... loading instructions..._

In this **cyberpunk** scenario, the data smuggler's database is empty. The root cause? No data has been added into the database Your mission: **Seed the database by running the migrations**. Failure is not an option.

**Your Toolkit:**
- **Use your local terminal to perform container updates**

**Your Objective:**
1. **The makefile has commands to run migrations , open the makefile and find out which command to use to seed the database** (in the main directory).
2. **Verify success** by checking the UI, you should be abe to see the list of products.

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
   ALTER USER postgres PASSWORD <password>
   docker logs <container_name>
   make migrate-up
   go build
   go run

