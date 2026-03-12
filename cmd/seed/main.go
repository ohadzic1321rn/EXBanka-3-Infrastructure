package main

import (
	"log"

	"EXBanka/internal/config"
	"EXBanka/internal/database"
)

func main() {
	cfg := config.Load()

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	if err := database.Migrate(db); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	if err := database.SeedPermissions(db); err != nil {
		log.Fatalf("Failed to seed permissions: %v", err)
	}

	if err := database.SeedDefaultAdmin(db); err != nil {
		log.Fatalf("Failed to seed default admin: %v", err)
	}

	log.Println("===================================")
	log.Println("Admin user created successfully!")
	log.Println("  Email:    admin@bank.com")
	log.Println("  Password: Admin123!")
	log.Println("===================================")
}
