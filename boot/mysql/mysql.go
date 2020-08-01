package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/shijting/go-web/boot"
	"github.com/shijting/go-web/entity"
	"github.com/spf13/viper"
	"sync"
	"time"
)

var db *gorm.DB

type MigrationInterface interface {
	Migration()
}

func Reload() (err error) {
	var lock sync.Mutex
	lock.Lock()
	defer lock.Unlock()
	if db != nil {
		err = db.Close()
		if err != nil {
			return
		}
		db = nil

	}
	Init()
	fmt.Println("mysql 重载成功")
	return
}

func Init() {
	var err error
	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.dbname"),
	)
	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		db = nil
		boot.ErrNotify <- err
		return
	}
	db.SingularTable(true)
	db.DB().SetConnMaxLifetime(20 * time.Second)
	db.DB().SetMaxIdleConns(viper.GetInt("mysql.max_idle_conns"))
	db.DB().SetMaxOpenConns(viper.GetInt("mysql.max_open_conns"))
	Migration()
	fmt.Println("mysql 配置成功")
}
func GetMysqlInstance() *gorm.DB {
	return db
}
func Close() {
	if db != nil {
		db.Close()
	}
}

func Migration() {
	db.
		Set("gorm:table_options", "ENGINE=InnoDB").
		Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&entity.Administrators{}, &entity.AdministratorRoles{}, &entity.AdministratorPermissions{})

}
