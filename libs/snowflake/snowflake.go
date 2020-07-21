package snowflake

import (
	"github.com/shijting/go-web/pkg/snowflake"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"time"
)

var Ids chan int64

// 每天生成number个唯一id
func Init(number int) {
	st := viper.GetString("start_time")
	GenIds(st, number)
	ticker := time.NewTicker(24 * time.Hour)
	go func(t *time.Ticker, st string) {
		for {
			<-t.C
			GenIds(st, number)
		}
	}(ticker, st)
}

func GenIds(startTime string, number int) {
	st, err := time.Parse("2006-01-02", startTime)
	if err != nil {
		zap.L().Error("", zap.Error(err))
	}
	Ids = make(chan int64, number)
	go func() {
		defer close(Ids)
		sf, err := snowflake.NewSnowflake(st, int64(viper.GetInt("machine_id")))
		if err != nil {
			zap.L().Error("", zap.Error(err))
		}
		for i := 0; i < number; i++ {
			Ids <- sf.GenID()
		}
	}()
}
