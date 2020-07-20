package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/shijting/go-web/boot"
	"github.com/spf13/viper"
	"sync"
)

var db *gorm.DB

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
	db.DB().SetMaxIdleConns(viper.GetInt("mysql.max_idle_conns"))
	db.DB().SetMaxOpenConns(viper.GetInt("mysql.max_open_conns"))
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
