package main

import (
	"log"
	"time"

	"EXBanka/internal/config"
	"EXBanka/internal/database"
	"EXBanka/internal/models"
	"EXBanka/internal/util"
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

	// Skip if admin already exists
	var existing models.Employee
	if result := db.Where("email = ?", "admin@bank.com").First(&existing); result.Error == nil {
		log.Println("Admin already exists, skipping seed")
		return
	}

	salt, err := util.GenerateSalt()
	if err != nil {
		log.Fatalf("Failed to generate salt: %v", err)
	}

	hashedPwd, err := util.HashPassword("Admin123!", salt)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	var allPerms []models.Permission
	if err := db.Find(&allPerms).Error; err != nil {
		log.Fatalf("Failed to fetch permissions: %v", err)
	}

	admin := models.Employee{
		Ime:           "Admin",
		Prezime:       "User",
		DatumRodjenja: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
		Pol:           "M",
		Email:         "admin@bank.com",
		BrojTelefona:  "0601234567",
		Adresa:        "Admin Street 1",
		Username:      "admin",
		Password:      hashedPwd,
		SaltPassword:  salt,
		Pozicija:      "Administrator",
		Departman:     "IT",
		Aktivan:       true,
		Permissions:   allPerms,
	}

	if err := db.Create(&admin).Error; err != nil {
		log.Fatalf("Failed to create admin: %v", err)
	}

	log.Println("===================================")
	log.Println("Admin user created successfully!")
	log.Println("  Email:    admin@bank.com")
	log.Println("  Password: Admin123!")
	log.Println("===================================")
}
