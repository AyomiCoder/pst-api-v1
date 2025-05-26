package db

import (
	"fmt"
)

func InitTables() error {
	// Create users table
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(255) NOT NULL UNIQUE,
		email VARCHAR(255) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := DB.Exec(createUsersTable)
	if err != nil {
		return fmt.Errorf("error creating users table: %v", err)
	}

	// Create posts table
	createPostsTable := `
	CREATE TABLE IF NOT EXISTS posts (
		id SERIAL PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		content TEXT NOT NULL,
		author_id INTEGER NOT NULL REFERENCES users(id),
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	);`

	_, err = DB.Exec(createPostsTable)
	if err != nil {
		return fmt.Errorf("error creating posts table: %v", err)
	}

	return nil
} 