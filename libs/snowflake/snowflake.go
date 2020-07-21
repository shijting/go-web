package snowflake

import (
	"fmt"
	"github.com/shijting/go-web/boot"
	"github.com/shijting/go-web/pkg/snowflake"
	"github.com/spf13/viper"
	"time"
)

var Ids chan int64

// 每天生成number个唯一id
func Init(number int) {

	GenIds(time.Now().Format("2006-01-02"), number)
	ticker := time.NewTicker(24 * time.Hour)
	go func(t *time.Ticker) {
		for {
			<-t.C
			fmt.Println("定时生成")
			GenIds(time.Now().Format("2006-01-02"), number)
		}
	}(ticker)
}

func GenIds(startTime string, number int) {
	st, err := time.Parse("2006-01-02", startTime)
	if err != nil {
		boot.ErrNotify <- err
	}
	Ids = make(chan int64, number)

	// 雪花算法
	go func() {
		sf, err := snowflake.NewSnowflake(st, int64(viper.GetInt("machine_id")))
		if err != nil {
			boot.ErrNotify <- err
		}
		for i := 0; i < number; i++ {
			Ids <- sf.GenID()
		}
		close(Ids)
	}()
}
