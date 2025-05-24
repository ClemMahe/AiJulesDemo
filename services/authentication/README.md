# Authentication Service

This service is responsible for handling user authentication, registration, and other account-related functionalities.
It can be built and run using the Dockerfile in this directory and orchestrated via the main `docker-compose.yml` in the parent `services` directory.

## API Endpoints

### 1. Register User
- **URL:** `/auth/register`
- **Method:** `POST`
- **Description:** Registers a new user.
- **Request Body (JSON):**
  ```json
  {
    "username": "testuser",
    "password": "password123"
  }
  ```
- **Success Response (201 Created):**
  ```json
  {
    "message": "User registered successfully"
  }
  ```
- **Error Responses:**
  - `400 Bad Request`: Invalid request payload, or username/password missing.
  - `405 Method Not Allowed`: If not a POST request.
  - `409 Conflict`: User already exists.
  - `500 Internal Server Error`: Error during password hashing or other server-side issue.

### 2. Login User
- **URL:** `/auth/login`
- **Method:** `POST`
- **Description:** Logs in an existing user.
- **Request Body (JSON):**
  ```json
  {
    "username": "testuser",
    "password": "password123"
  }
  ```
- **Success Response (200 OK):**
  ```json
  {
    "message": "Login successful"
  }
  ```
- **Error Responses:**
  - `400 Bad Request`: Invalid request payload, or username/password missing.
  - `401 Unauthorized`: Invalid username or password.
  - `405 Method Not Allowed`: If not a POST request.
```
