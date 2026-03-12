package database

import (
	"fmt"
	"log"
	"time"

	"EXBanka/internal/models"
	"EXBanka/internal/util"

	"gorm.io/gorm"
)

// SeedDefaultAdmin creates the default admin user on a fresh database.
// The operation is idempotent and safely skips creation if the user exists.
func SeedDefaultAdmin(db *gorm.DB) error {
	var existing models.Employee
	if result := db.Where("email = ?", "admin@bank.com").First(&existing); result.Error == nil {
		log.Println("Admin already exists, skipping seed")
		return nil
	} else if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return fmt.Errorf("failed to check existing admin: %w", result.Error)
	}

	salt, err := util.GenerateSalt()
	if err != nil {
		return fmt.Errorf("failed to generate salt: %w", err)
	}

	hashedPwd, err := util.HashPassword("Admin123!", salt)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	var allPerms []models.Permission
	if err := db.Find(&allPerms).Error; err != nil {
		return fmt.Errorf("failed to fetch permissions: %w", err)
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
		return fmt.Errorf("failed to create admin: %w", err)
	}

	log.Println("Default admin user created successfully")
	return nil
}
