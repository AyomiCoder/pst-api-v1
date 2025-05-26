package main

import (
	"godb/db"
	"godb/routes"
	"log"
)

func main() {
	// Initialize database
	db.Init()
	defer db.DB.Close()

	// Initialize database tables
	if err := db.InitTables(); err != nil {
		log.Fatal("Error initializing tables:", err)
	}

	// Setup routes
	r := routes.SetupRouter()

	// Start server
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Error starting server:", err)
	}
}