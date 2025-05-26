# PST API

A RESTful API built with Go and Gin framework for managing posts and user authentication.

## Features

- User authentication (register/login)
- JWT-based authentication
- CRUD operations for posts
- Protected routes
- Database integration

## Prerequisites

- Go 1.x
- PostgreSQL
- Air (for hot reloading during development)

## Environment Variables

Create a `.env` file in the root directory with the following variables:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_username
DB_PASSWORD=your_password
DB_NAME=your_database_name
JWT_SECRET=your_jwt_secret
```

## Installation

1. Clone the repository
2. Install dependencies:
   ```bash
   go mod download
   ```
3. Set up your environment variables
4. Run the application:
   ```bash
   go run main.go
   ```
   
For development with hot reloading:
```bash
air
```

## API Endpoints

### Authentication

#### Register User
- **POST** `/auth/register`
- **Body:**
  ```json
  {
    "username": "string",
    "email": "string",
    "password": "string"
  }
  ```

#### Login
- **POST** `/auth/login`
- **Body:**
  ```json
  {
    "email": "string",
    "password": "string"
  }
  ```

### Posts (Protected Routes)

#### Create Post
- **POST** `/api/posts`
- **Headers:** `Authorization: Bearer <token>`
- **Body:**
  ```json
  {
    "title": "string",
    "content": "string"
  }
  ```

#### Get All Posts
- **GET** `/api/posts`
- **Headers:** `Authorization: Bearer <token>`

#### Get Single Post
- **GET** `/api/posts/:id`
- **Headers:** `Authorization: Bearer <token>`

#### Update Post
- **PUT** `/api/posts/:id`
- **Headers:** `Authorization: Bearer <token>`
- **Body:**
  ```json
  {
    "title": "string",
    "content": "string"
  }
  ```

#### Delete Post
- **DELETE** `/api/posts/:id`
- **Headers:** `Authorization: Bearer <token>`

## Data Models

### User
```go
type User struct {
    ID        int       `json:"id"`
    Username  string    `json:"username"`
    Email     string    `json:"email"`
    Password  string    `json:"-"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
```

### Post
```go
type Post struct {
    ID        int       `json:"id"`
    Title     string    `json:"title"`
    Content   string    `json:"content"`
    AuthorID  int       `json:"author_id"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
```

## Project Structure

```
.
├── db/           # Database configuration and initialization
├── handlers/     # Request handlers
├── middleware/   # Custom middleware (auth, etc.)
├── models/       # Data models
├── routes/       # Route definitions
├── main.go       # Application entry point
└── .env          # Environment variables
```

## Error Handling

The API uses standard HTTP status codes:
- 200: Success
- 201: Created
- 400: Bad Request
- 401: Unauthorized
- 403: Forbidden
- 404: Not Found
- 500: Internal Server Error

## Security

- Passwords are hashed before storage
- JWT tokens are used for authentication
- Protected routes require valid JWT token
- Environment variables for sensitive data 