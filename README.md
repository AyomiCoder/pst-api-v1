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
