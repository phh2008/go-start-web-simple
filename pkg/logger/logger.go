package logger

import (
	"com.gientech/equipment-data-collection/pkg/config"
	"github.com/google/wire"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"strings"
	"time"
)

var LoggerSet = wire.NewSet(InitLogger)

var log *zap.Logger = zap.L()
var slog *zap.SugaredLogger = log.Sugar()

var levelMap = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
}

func InitLogger(config *config.Config) *zap.Logger {
	log = newZapLog(config)
	slog = log.Sugar()
	zap.ReplaceGlobals(log)
	return log
}

func L() *zap.Logger {
	l := log
	return l
}

func S() *zap.SugaredLogger {
	s := slog
	return s
}

// getWriter
func getWriter(config *config.Config) io.Writer {
	return &lumberjack.Logger{
		Filename:   config.Viper.GetString("log.filename"),
		MaxSize:    config.Viper.GetInt("log.maxSize"), // megabytes
		MaxBackups: config.Viper.GetInt("log.maxBackups"),
		MaxAge:     config.Viper.GetInt("log.maxAge"), //days
		LocalTime:  config.Viper.GetBool("log.localTime"),
		Compress:   config.Viper.GetBool("log.compress"), // disabled by default
	}
}

func newZapLog(config *config.Config) *zap.Logger {
	// 设置日志格式
	//encoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
	encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:  "msg",
		LevelKey:    "level",
		EncodeLevel: zapcore.CapitalLevelEncoder,
		TimeKey:     "ts",
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		CallerKey:    "file",
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	})

	// 记录什么级别的日志
	level := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= levelMap[strings.ToLower(config.Viper.GetString("log.level"))]
	})

	// 获取 info、error日志文件的io.Writer 抽象 getWriter() 在下方实现
	writer := zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(getWriter(config)))
	// 如果info、debug、error分文件记录，就创建多个 writer
	// 最后创建具体的Logger
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(writer), level), // 可添加多个
	)
	// 需要传入 zap.AddCaller() 才会显示打日志点的文件名和行数
	return zap.New(core, zap.AddCaller())
}
