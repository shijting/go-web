package logger

import (
	"github.com/shijting/go-web/src/boot"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// 初始化Logger
func Init() {
	logfile := "runtime/logs/" + viper.GetString("logger.logfile")
	writeSyncer := getLogWriter(logfile,
		viper.GetInt("logger.max_size"),
		viper.GetInt("logger.max_backups"),
		viper.GetInt("logger.max_age"),
	)
	encoder := getEncoder()
	var l = new(zapcore.Level)
	err := l.UnmarshalText([]byte(viper.GetString("logger.level")))
	if err != nil {
		boot.ErrNotify <- err
	}
	core := zapcore.NewCore(encoder, writeSyncer, l)

	Logger := zap.New(core, zap.AddCaller())
	// 替换
	zap.ReplaceGlobals(Logger)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}
