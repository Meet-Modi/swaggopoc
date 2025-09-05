#!/bin/bash

echo "Testing Swaggo POC API..."

# Start the server in background
go run main.go &
SERVER_PID=$!

# Wait for server to start
sleep 3

echo "Testing health endpoint..."
curl -s http://localhost:8080/api/v1/health

echo -e "\n\nTesting get all users..."
curl -s http://localhost:8080/api/v1/users

echo -e "\n\nTesting create user..."
curl -s -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name": "Test User", "email": "test@example.com"}'

echo -e "\n\nTesting get user by ID..."
curl -s http://localhost:8080/api/v1/users/1

echo -e "\n\nSwagger UI available at: http://localhost:8080/swagger/index.html"

# Kill the server
kill $SERVER_PID

echo -e "\n\nTest completed!"
