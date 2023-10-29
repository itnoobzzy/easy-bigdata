package log

import (
	"easyCasbin/utils"
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"os"
	"path"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"easyCasbin/internal/conf"
)

var _ log.Logger = (*Logger)(nil)

type Logger struct {
	LogConf *conf.Log
	zl      *zap.Logger
}

func NewZapLogger(lconf *conf.Log) *Logger {
	zl, _ := zap.NewDevelopment()
	logger := &Logger{lconf, zl}
	zl, err := logger.GetZapLogger()
	if err != nil {
		panic(err)
	}
	logger.zl = zl
	return logger
}

// GetLogger 获取 zap 的 logger
func (z *Logger) GetZapLogger() (logger *zap.Logger, err error) {
	if ok, _ := utils.PathExists(z.LogConf.Director); !ok {
		fmt.Printf("create %v directory\n", z.LogConf.Director)
		err = os.Mkdir(z.LogConf.Director, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}

	cores := z.getZapCores()
	logger = zap.New(
		zapcore.NewTee(cores...))

	if z.LogConf.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger, nil
}

// zapEncodeLevel 根据 EncodeLevel 返回 zapcore.LevelEncoder
func (z *Logger) zapEncodeLevel() zapcore.LevelEncoder {
	switch {
	case z.LogConf.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		return zapcore.LowercaseLevelEncoder
	case z.LogConf.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		return zapcore.LowercaseColorLevelEncoder
	case z.LogConf.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		return zapcore.CapitalLevelEncoder
	case z.LogConf.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		return zapcore.CapitalColorLevelEncoder
	default:
		return zapcore.LowercaseLevelEncoder
	}
}

// transportLevel 根据字符串转化为 zapcore.Level
func (z *Logger) transportLevel() zapcore.Level {
	z.LogConf.Level = strings.ToLower(z.LogConf.Level)
	switch z.LogConf.Level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}

// getEncoder 获取 zapcore.Encoder
func (z *Logger) getEncoder() zapcore.Encoder {
	if z.LogConf.Format == "json" {
		return zapcore.NewJSONEncoder(z.getEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(z.getEncoderConfig())
}

// customTimeEncoder 自定义日志输出时间格式
func (z *Logger) customTimeEncoder(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
	encoder.AppendString(z.LogConf.Prefix + t.Format("2006/01/02 - 15:04:05.000"))
}

// getEncoderConfig 获取zapcore.EncoderConfig
func (z *Logger) getEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  z.LogConf.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    z.zapEncodeLevel(),
		EncodeTime:     z.customTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
}

// getEncoderCore 获取Encoder的 zapcore.Core
func (z *Logger) getEncoderCore(l zapcore.Level, level zap.LevelEnablerFunc) zapcore.Core {
	writer, err := z.getWriteSyncer(l.String())
	if err != nil {
		fmt.Printf("Get Write Sycner Failed err:%v", err.Error())
		return nil
	}
	return zapcore.NewCore(z.getEncoder(), writer, level)
}

// getWriteSyncer 获取 zapcore.WriteSyncer
func (z *Logger) getWriteSyncer(level string) (zapcore.WriteSyncer, error) {
	fileWriter, err := rotatelogs.New(
		path.Join(z.LogConf.Director, "%Y-%m-%d", level+".log"),
		rotatelogs.WithClock(rotatelogs.Local),
		rotatelogs.WithMaxAge(time.Duration(z.LogConf.MaxAge)*24*time.Hour), // 日志留存时间
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	//if z.LogConf.LogInConsole {
	//	return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	//}
	return zapcore.AddSync(fileWriter), err
}

// getLevelPriority 根据 zapcore.Level 获取 zap.LevelEnablerFunc
func (z *Logger) getLevelPriority(level zapcore.Level) zap.LevelEnablerFunc {
	switch level {
	case zapcore.DebugLevel:
		return func(level zapcore.Level) bool {
			return level == zap.DebugLevel
		}
	case zapcore.InfoLevel:
		return func(level zapcore.Level) bool {
			return level == zapcore.InfoLevel
		}
	case zapcore.WarnLevel:
		return func(level zapcore.Level) bool {
			return level == zapcore.WarnLevel
		}
	case zapcore.ErrorLevel:
		return func(level zapcore.Level) bool {
			return level == zapcore.ErrorLevel
		}
	case zapcore.DPanicLevel:
		return func(level zapcore.Level) bool {
			return level == zapcore.DPanicLevel
		}
	case zapcore.PanicLevel:
		return func(level zapcore.Level) bool {
			return level == zapcore.PanicLevel
		}
	case zapcore.FatalLevel:
		return func(level zapcore.Level) bool {
			return level == zapcore.DebugLevel
		}
	default:
		return func(level zapcore.Level) bool {
			return level == zap.DebugLevel
		}
	}
}

// getZapCores 根据配置文件的Level获取 []zapcore.Core
func (z *Logger) getZapCores() []zapcore.Core {
	cores := make([]zapcore.Core, 0, 7)
	for level := z.transportLevel(); level <= zapcore.FatalLevel; level++ {
		cores = append(cores, z.getEncoderCore(level, z.getLevelPriority(level)))
		pe := z.getEncoderConfig()
		consoleEncoder := zapcore.NewConsoleEncoder(pe)
		if z.LogConf.LogInConsole {
			cores = append(cores, zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), z.getLevelPriority(level)))
		}
	}
	return cores
}

func (z *Logger) Log(level log.Level, keyvals ...interface{}) error {
	keylen := len(keyvals)
	if keylen == 0 || keylen%2 != 0 {
		z.zl.Warn(fmt.Sprint("Keyvalues must appear in pairs: ", keyvals))
		return nil
	}

	data := make([]zap.Field, 0, (keylen/2)+1)
	for i := 0; i < keylen; i += 2 {
		data = append(data, zap.Any(fmt.Sprint(keyvals[i]), keyvals[i+1]))
	}

	switch level {
	case log.LevelDebug:
		z.zl.Debug("", data...)
	case log.LevelInfo:
		z.zl.Info("", data...)
	case log.LevelWarn:
		z.zl.Warn("", data...)
	case log.LevelError:
		z.zl.Error("", data...)
	case log.LevelFatal:
		z.zl.Fatal("", data...)
	}
	return nil
}
