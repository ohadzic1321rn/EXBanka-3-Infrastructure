package models

type Permission struct {
	ID          uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string     `gorm:"uniqueIndex;not null" json:"name"`
	Description string     `json:"description"`
	SubjectType string     `gorm:"not null;default:employee;index" json:"subject_type"`
	Employees   []Employee `gorm:"many2many:employee_permissions;" json:"-"`
	Clients     []Client   `gorm:"many2many:client_permissions;" json:"-"`
}

const (
	PermissionSubjectEmployee = "employee"
	PermissionSubjectClient   = "client"
)

// Permission name constants
const (
	PermAdmin               = "admin"
	PermEmployeeCreate      = "employee.create"
	PermEmployeeRead        = "employee.read"
	PermEmployeeUpdate      = "employee.update"
	PermEmployeeActivate    = "employee.activate"
	PermEmployeePermissions = "employee.permissions"
	PermClientBasic         = "client.basic"
	PermClientTrading       = "client.trading"

	// Sprint 2: New employee permissions for role hierarchy
	PermEmployeeBankOperations   = "employee.bank_operations"
	PermEmployeeClientManagement = "employee.client_management"
	PermEmployeeStockTrading     = "employee.stock_trading"
	PermEmployeeNoLimits         = "employee.no_limits"
	PermEmployeeOTCTrading       = "employee.otc_trading"
	PermEmployeeFundManagement   = "employee.fund_management"
	PermEmployeeManageAll        = "employee.manage_all"

	// Sprint 2: New client permissions for role hierarchy
	PermClientBankOperations = "client.bank_operations"
	PermClientStockTrading   = "client.stock_trading"
	PermClientFundInvesting  = "client.fund_investing"
)

// DefaultPermissions are seeded on first run
var DefaultPermissions = []Permission{
	{Name: PermAdmin, Description: "Full administrative access", SubjectType: PermissionSubjectEmployee},
	{Name: PermEmployeeCreate, Description: "Can create new employees", SubjectType: PermissionSubjectEmployee},
	{Name: PermEmployeeRead, Description: "Can read employee data", SubjectType: PermissionSubjectEmployee},
	{Name: PermEmployeeUpdate, Description: "Can update employee data (non-admin targets only)", SubjectType: PermissionSubjectEmployee},
	{Name: PermEmployeeActivate, Description: "Can activate/deactivate employees", SubjectType: PermissionSubjectEmployee},
	{Name: PermEmployeePermissions, Description: "Can manage employee permissions", SubjectType: PermissionSubjectEmployee},
	{Name: PermClientBasic, Description: "Basic client role", SubjectType: PermissionSubjectClient},
	{Name: PermClientTrading, Description: "Trading-enabled client role", SubjectType: PermissionSubjectClient},

	// Sprint 2: New employee permissions for role hierarchy
	{Name: PermEmployeeBankOperations, Description: "Osnovno poslovanje banke (basic banking operations)", SubjectType: PermissionSubjectEmployee},
	{Name: PermEmployeeClientManagement, Description: "Upravljanje klijentima (client management)", SubjectType: PermissionSubjectEmployee},
	{Name: PermEmployeeStockTrading, Description: "Trgovina hartijama sa berze uz limite (stock trading with limits)", SubjectType: PermissionSubjectEmployee},
	{Name: PermEmployeeNoLimits, Description: "Bez limita (no trading limits)", SubjectType: PermissionSubjectEmployee},
	{Name: PermEmployeeOTCTrading, Description: "OTC trgovina (OTC trading)", SubjectType: PermissionSubjectEmployee},
	{Name: PermEmployeeFundManagement, Description: "Upravljanje fondovima i agentima (fund and agent management)", SubjectType: PermissionSubjectEmployee},
	{Name: PermEmployeeManageAll, Description: "Upravlja svim zaposlenima (manage all employees)", SubjectType: PermissionSubjectEmployee},

	// Sprint 2: New client permissions for role hierarchy
	{Name: PermClientBankOperations, Description: "Osnovno poslovanje banke - racuni, transferi, placanja (basic banking: accounts, transfers, payments)", SubjectType: PermissionSubjectClient},
	{Name: PermClientStockTrading, Description: "Trgovina hartijama sa berze i OTC (stock exchange and OTC trading)", SubjectType: PermissionSubjectClient},
	{Name: PermClientFundInvesting, Description: "Investiranje u fondove (fund investing)", SubjectType: PermissionSubjectClient},
}
