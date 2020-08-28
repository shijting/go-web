package entity

import "time"

type Administrators struct {
	ID            uint64               `gorm:"primary_key" json:"id"`
	Name          string               `gorm:"type:varchar(60);not null;default:'';comment:'姓名'" json:"name"`
	Email         string               `gorm:"type:varchar(100);unique_index:unq_email;not null;comment:'email(登录账户)'" json:"email"`
	Password      string               `gorm:"type:varchar(100);not null;comment:'密码'" json:"password"`
	LastLoginDate time.Time            `gorm:"column:last_login_date" json:"last_login_date"`
	LastLoginIp   string               `gorm:"type:varchar(15);not null;default:'';comment:'最后登录ip'" json:"last_login_ip"`
	Status        uint8                `gorm:"type:tinyint(1);not null;default:1;comment:'状态：(1:正常，0无效)'" json:"status"`
	Roles         []AdministratorRoles `gorm:"many2many:administrator_roles_relation" json:"roles"`
	CreatedAt     time.Time            `gorm:"column:created_at" json:"created_at" format:"YYYY-mm-dd"`
	UpdatedAt     time.Time            `gorm:"column:updated_at" json:"updated_at"`
}

func (Administrators) TableName() string {
	return "administrators"
}

type AdministratorInsert struct {
	Name          string    `gorm:"column:name" json:"name" binding:"required,gte=2,lte=6"`
	Email         string    `gorm:"column:email" json:"email" binding:"required,email"`
	Password      string    `gorm:"column:password" json:"password" binding:"required"`
	LastLoginDate time.Time `gorm:"column:last_login_date" json:"last_login_date"`
	LastLoginIp   string    `gorm:"column:last_login_ip" json:"last_login_ip"`
	Status        uint8     `gorm:"column:status" json:"status"`
	RoleId        uint64    `json:"role_id"`
}

// 管理员- 角色 多对多关联
type AdministratorsDetail struct {
	ID            uint64               `json:"id"`
	Name          string               `json:"name"`
	Email         string               `json:"email"`
	Password      string               `json:"password"`
	LastLoginDate time.Time            `json:"last_login_date"`
	LastLoginIp   string               `json:"last_login_ip"`
	Status        uint8                `json:"status"`
	Roles         []AdministratorRoles `gorm:"many2many:administrator_role_relation;association_foreignkey:id;foreignkey:administrator_id" json:"roles"`
	CreatedAt     time.Time            `json:"created_at"`
	UpdatedAt     time.Time            `json:"updated_at"`
}
