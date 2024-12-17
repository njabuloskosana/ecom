# Activity 4: **(Updating the ProductList.vue Compoonent) Updating Go api endpoint**

_Transmission decrypted... loading instructions..._

In this **cyberpunk** scenario, you need to update the way products are displayed on the UI. The problem is that our current method is not effcient enough in managing the displaying of fetched products from the database, currently we cannot restrict how many products we want at a time and we cannot filter through them. Your mission: **Implement pagination  and filtering to managing the retrieveing of prducts from the database**. Failure is not an option.

**Your Objective:**
1. **Update the ProductList.vue component to handle pagination**.
2. **Update the ProductList.vue component to handle filtering**.

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


