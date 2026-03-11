package models

type Permission struct {
	ID          uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string     `gorm:"uniqueIndex;not null" json:"name"`
	Description string     `json:"description"`
	Employees   []Employee `gorm:"many2many:employee_permissions;" json:"-"`
}

// Permission name constants
const (
	PermAdmin               = "admin"
	PermEmployeeCreate      = "employee.create"
	PermEmployeeRead        = "employee.read"
	PermEmployeeUpdate      = "employee.update"
	PermEmployeeActivate    = "employee.activate"
	PermEmployeePermissions = "employee.permissions"
)

// DefaultPermissions are seeded on first run
var DefaultPermissions = []Permission{
	{Name: PermAdmin, Description: "Full administrative access"},
	{Name: PermEmployeeCreate, Description: "Can create new employees"},
	{Name: PermEmployeeRead, Description: "Can read employee data"},
	{Name: PermEmployeeUpdate, Description: "Can update employee data (non-admin targets only)"},
	{Name: PermEmployeeActivate, Description: "Can activate/deactivate employees"},
	{Name: PermEmployeePermissions, Description: "Can manage employee permissions"},
}
