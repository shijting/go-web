package main

import (
	"context"
	"fmt"
	"github.com/shijting/go-web/boot"
	"github.com/shijting/go-web/boot/logger"
	"github.com/shijting/go-web/boot/mysql"
	"github.com/shijting/go-web/boot/redis"
	"github.com/shijting/go-web/boot/setup"
	_ "github.com/shijting/go-web/config"
	"github.com/shijting/go-web/libs/snowflake"
	"github.com/spf13/viper"
	"go.uber.org/zap"
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
	// 雪花算法生成唯一id
	snowflake.Init(50000)
	fmt.Println()
	// 加载路由
	r := setup.Init()
	serv := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("port")),
		Handler: r,
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
