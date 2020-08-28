package main

import (
	"context"
	"fmt"
	"github.com/shijting/go-web/src/boot"
	"github.com/shijting/go-web/src/boot/logger"
	"github.com/shijting/go-web/src/boot/mysql"
	"github.com/shijting/go-web/src/boot/redis"
	_ "github.com/shijting/go-web/src/config"
	"github.com/shijting/go-web/src/libs/validator"
	"github.com/shijting/go-web/src/routes"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	logger.Init()
	defer zap.L().Sync()

	// mysql初始化
	mysql.Init()
	defer mysql.Close()
	// 初始化redis
	redis.Init()
	defer redis.Close()

	err := validator.InitTrans("zh")
	if err != nil {
		log.Fatal(err)
	}
	// 雪花算法生成唯一id
	//snowflake.Init(10000)
	// 加载路由
	r := routes.Init()
	serv := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("port")),
		Handler: r.Engine,
	}
	go func() {
		if err := serv.ListenAndServe(); err != nil {
			boot.ErrNotify <- err
		}
	}()
	go func() {
		sig := make(chan os.Signal)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
		boot.ErrNotify <- fmt.Errorf("%s", <-sig)
	}()
	errMsg := <-boot.ErrNotify
	fmt.Println(errMsg)
	zap.L().Info(errMsg.Error())
	// 3s后关机
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	fmt.Println("正在关机...")
	defer cancel()
	<-ctx.Done()
	serv.Shutdown(ctx)
}
