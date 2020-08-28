package entity

import "time"

type AdministratorRoles struct {
	ID          uint64                     `gorm:"primary_key" json:"id"`
	RoleName    string                     `gorm:"type:varchar(60);not null;default:'';comment:'角色名称'" json:"role_name"`
	Description string                     `gorm:"type:varchar(200);not null;default:'';comment:'描述'" json:"description"`
	Status      uint8                      `gorm:"type:tinyint(1);not null;default:1;comment:'状态(1正常，0无效)'" json:"status"`
	Permissions []AdministratorPermissions `gorm:"many2many:administrator_role_permission_relation" json:"permissions"`
	CreatedAt   time.Time                  `json:"created_at"`
	UpdatedAt   time.Time                  `json:"updated_at"`
}
type CreatedAt struct {
	CreatedAt string `json:"created_at"`
}

func (AdministratorRoles) TableName() string {
	return "administrator_roles"
}
