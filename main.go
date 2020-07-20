package main

import (
	"context"
	"fmt"
	"github.com/shijting/go-web/boot"
	"github.com/shijting/go-web/boot/logger"
	"github.com/shijting/go-web/boot/mysql"
	"github.com/shijting/go-web/boot/setup"
	_ "github.com/shijting/go-web/config"
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
	r := setup.Init()
	// mysql初始化
	mysql.Init()
	defer mysql.Close()
	serv := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("app.port")),
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
