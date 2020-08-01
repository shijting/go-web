package entity

import "time"

type AdministratorPermissions struct {
	Id             int                        `gorm:"primary_key" json:"id"`
	PermissionName string                     `gorm:"type:varchar(60);not null;default:'';comment:'菜单名称'" json:"permission_name"`
	UniqueKey      string                     `gorm:"type:varchar(50);not null;unique_index:unq_key;comment:'唯一标识字段，与vue路由name一致'" json:"unique_key"`
	Method         string                     `gorm:"type:varchar(20);not null;comment:'http请求方法'" json:"method"`
	Url            string                     `gorm:"type:varchar(200);not null;comment:'http路由'" json:"url"`
	Pid            int                        `gorm:"type:int(10);not null;default 0;index:idx_pid;comment:'父级菜单id，0为顶级菜单'" json:"pid"`
	Description    string                     `gorm:"type:varchar(200);not null;default:'';comment:'描述'" json:"description"`
	CreatedAt      time.Time                  `json:"created_at"`
	UpdatedAt      time.Time                  `json:"updated_at"`
	Children       []AdministratorPermissions `json:"children,omitempty"`
}

func (AdministratorPermissions) TableName() string {
	return "administrator_permissions"
}
