package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
	"sync"
)

var rdb *redis.Client

func Init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", viper.GetString("redis.host"), viper.GetInt("redis.port")),
		Password: viper.GetString("redis.auth"), // no password set
		DB:       0,                             // use default DB
	})

	_, err := rdb.Ping().Result()
	if err != nil {
		log.Println(err)
		zap.L().Error("connect to redis is failed ", zap.Error(err))
		rdb = nil
	}
}
func Reload() (err error) {
	var lock sync.Mutex
	lock.Lock()
	defer lock.Unlock()
	if rdb != nil {
		if err := rdb.Close(); err != nil {
			return err
		}
		rdb = nil
	}
	Init()
	return err
}
func Close() {
	if rdb != nil {
		rdb.Close()
	}
}
func GetRedisInstance() *redis.Client {
	return rdb
}
