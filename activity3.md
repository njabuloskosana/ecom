# Activity 3: **(Updating the products endpoint) Updating Go api endpoint**

_Transmission decrypted... loading instructions..._

In this **cyberpunk** scenario, you need to update the way products are fetched from the backend to the UI. The problem is that our current method is not effcient enough in managing retrieveing products from the database, currently we cannot restrict how many products we want at a time. Your mission: **Implement pagination to managing the retrieveing of prducts from the database**. Failure is not an option.

**Your Objective:**
1. **Update the products endpoint to handle pagination parameters**.
2. **Update the products endpoint to handle filtering parameters**.
3. **Update the products endpoint to incoporate caching**.

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


