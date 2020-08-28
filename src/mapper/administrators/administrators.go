package administrators

import (
	"github.com/shijting/go-web/src/boot/mysql"
	"github.com/shijting/go-web/src/entity"
)

// CURD
func GetOneById(id int64) (res *entity.Administrators, err error) {
	res = new(entity.Administrators)
	err = mysql.GetMysqlInstance().Where("id = ?", id).First(res).Error
	return
}
func Create(data *entity.AdministratorInsert) error {
	return mysql.GetMysqlInstance().Create(data).Error
}
func Update(id int64, data map[string]interface{}) error {
	model := &entity.Administrators{ID: uint64(id)}
	return mysql.GetMysqlInstance().Model(model).Update(data).Error
}
