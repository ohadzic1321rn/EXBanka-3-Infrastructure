package repository

import (
	"EXBanka/internal/models"

	"gorm.io/gorm"
)

type PermissionRepository struct {
	db *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) *PermissionRepository {
	return &PermissionRepository{db: db}
}

func (r *PermissionRepository) FindAll() ([]models.Permission, error) {
	var perms []models.Permission
	err := r.db.Find(&perms).Error
	return perms, err
}

func (r *PermissionRepository) FindByName(name string) (*models.Permission, error) {
	var perm models.Permission
	err := r.db.Where("name = ?", name).First(&perm).Error
	if err != nil {
		return nil, err
	}
	return &perm, nil
}

func (r *PermissionRepository) FindByNames(names []string) ([]models.Permission, error) {
	var perms []models.Permission
	err := r.db.Where("name IN ?", names).Find(&perms).Error
	return perms, err
}
