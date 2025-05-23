# Go Microservices Project

This project contains three microservices: Authentication, GameLogic, and Rest.
The Go module for these services is located in the `services` directory.
Each service has its own Dockerfile and can be orchestrated using the `docker-compose.yml` file located in the `services` directory.

## Structure
- `/services/authentication`: Handles user authentication. (Dockerfile inside)
- `/services/gamelogic`: Manages game logic. (Dockerfile inside)
- `/services/rest`: Provides a RESTful API. (Dockerfile inside)
- `/services/go.mod`: Go module file for all services.
- `/services/docker-compose.yml`: Docker Compose file to build and run all services.

## Running with Docker
Navigate to the `services` directory and run:
```bash
docker-compose up --build
```
This will build the images for each service and start them.
- Authentication service will be available on port 8080.
- GameLogic service will be available on port 8081.
- Rest service will be available on port 8082.
