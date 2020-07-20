package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/shijting/go-web/boot/mysql"
	"github.com/shijting/go-web/boot/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
)

//type configStruct struct {
//	App    *App    `mapstructure:"app"`
//	Logger *Logger `mapstructure:"logger"`
//	Mysql  *Mysql  `mapstructure:"mysql"`
//	Redis  *Redis  `mapstructure:"redis"`
//}
type configStruct struct {
	Name    string `mapstructure:"name"`
	Mode    string `mapstructure:"mode"`
	Port    string `mapstructure:"port"`
	*Logger `mapstructure:"logger"`
	*Mysql  `mapstructure:"mysql"`
	*Redis  `mapstructure:"redis"`
}

type Logger struct {
	Level      string `mapstructure:"level"`
	LogFile    int    `mapstructure:"logfile"`
	MaxSize    string `mapstructure:"max_size"`
	MaxAge     string `mapstructure:"max_age"`
	MaxBackups string `mapstructure:"max_backups"`
}

type Mysql struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	Dbname       string `mapstructure:"dbname"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
}
type Redis struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
	Auth string `mapstructure:"auth"`
}

var Config = new(configStruct)

func init() {
	//Config.Mysql = new(Mysql)
	//Config.Logger = new(Logger)
	//Config.Redis = new(Redis)
	fmt.Println("配置加载中...")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config/")
	if err := viper.Unmarshal(Config); err != nil {
		log.Fatal(err)
	}
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	fmt.Println("host", Config)
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		if err := viper.Unmarshal(Config); err != nil {
			log.Fatal(err)
		}
		fmt.Println("host", Config.Mysql.Host)
		// todo 重新获取mysql/redis连接对象
		err := mysql.Reload()
		if err != nil {
			zap.L().Error("mysql 重置失败:", zap.Error(err))
		}
		err = redis.Reload()
		if err != nil {
			zap.L().Error("redis 重置失败:", zap.Error(err))
		}
	})

}
