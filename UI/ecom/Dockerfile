# Dockerfile for Vue.js UI

# Step 1: Build the Vue.js application
FROM node:18-alpine AS builder
WORKDIR /app

# Copy package.json and package-lock.json
COPY package.json package-lock.json ./
RUN npm install

# Copy the rest of the application code
COPY . .

# Build the Vue.js application
RUN npm run build

# Step 2: Serve the static files using nginx
FROM nginx:alpine
COPY --from=builder /app/dist /usr/share/nginx/html

# Expose the Nginx port
EXPOSE 80

