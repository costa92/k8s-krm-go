package log

import (
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	gormLogger "gorm.io/gorm/logger"

	krtlog "github.com/go-kratos/kratos/v2/log"
)

// Field is an alias for zapcore.Field
type Field = zapcore.Field

type Logger interface {
	Debugf(format string, args ...any)
	Debugw(msg string, keyVals ...any)
	Infof(format string, args ...any)
	Infow(msg string, keyVals ...any)
	Warnf(format string, args ...any)
	Warnw(msg string, keyVals ...any)
	Errorf(format string, args ...any)
	Errorw(err error, msg string, keyVals ...any)
	Panicf(format string, args ...any)
	Panicw(msg string, keyVals ...any)
	Fatalf(format string, args ...any)
	Fatalw(msg string, keyVals ...any)
	With(fields ...Field) Logger
	AddCallerSkip(skip int) Logger
	Sync()

	krtlog.Logger
	gormLogger.Interface
}

type zapLogger struct {
	z    *zap.Logger
	opts *Options
}

var _ Logger = (*zapLogger)(nil)

var (
	mu sync.Mutex
	// DefaultLogger is the default logger
	std = NewLogger(NewOptions())
)

func Init(opts *Options) {
	mu.Lock()
	defer mu.Unlock()

	std = NewLogger(opts)
}

func NewLogger(opts *Options) *zapLogger {
	if opts == nil {
		opts = NewOptions()
	}

	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(opts.Level)); err != nil {
		zapLevel = zapcore.InfoLevel
	}

	// 创建一个默认的 encoder 配置
	encoderConfig := zap.NewProductionEncoderConfig()
	// 自定义 MessageKey 和 TimeKey
	encoderConfig.MessageKey = "message"
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(time.Format("2006-01-02 15:04:05.000"))
	}

	encoderConfig.EncodeDuration = func(d time.Duration, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendFloat64(float64(d) / float64(time.Millisecond))
	}

	// where the log will be written
	if opts.Format == "console" && opts.EnableColor {
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	cfg := &zap.Config{
		Level:         zap.NewAtomicLevelAt(zapLevel),
		DisableCaller: opts.DisableCaller,
		Encoding:      opts.Format,
		EncoderConfig: encoderConfig,

		OutputPaths:      opts.OutputPaths,
		ErrorOutputPaths: []string{"stderr"},
	}

	z, err := cfg.Build(zap.AddStacktrace(zap.PanicLevel), zap.AddCallerSkip(2))
	if err != nil {
		panic(err)
	}
	logger := &zapLogger{z: z, opts: opts}
	// 	把标准库的 log.Logger 的 info 级别的输出重定向到 zap.Logger
	zap.RedirectStdLog(z)
	return logger
}

func Default() Logger {
	return std
}

func (l *zapLogger) Options() *Options {
	return l.opts
}

// Sync 调用底层 zap.Logger 的 Sync 方法，将缓存中的日志刷新到磁盘文件中. 主程序需要在退出前调用 Sync.

func Sync() {
	std.Sync()
}

func (l *zapLogger) Sync() {
	_ = l.z.Sync()
}

func (l *zapLogger) Debugf(format string, args ...any) {
	l.z.Sugar().Debugf(format, args...)
}

func Debugf(format string, args ...any) {
	std.Debugf(format, args...)
}

func (l *zapLogger) Debugw(msg string, keyVals ...any) {
	l.z.Sugar().Debugw(msg, keyVals...)
}

func Debugw(msg string, keyVals ...any) {
	std.Debugw(msg, keyVals...)
}

func (l *zapLogger) Infof(format string, args ...any) {
	l.z.Sugar().Infof(format, args...)
}

func Infof(format string, args ...any) {
	std.Infof(format, args...)
}

func (l *zapLogger) Infow(msg string, keyVals ...any) {
	l.z.Sugar().Infow(msg, keyVals...)
}

func Infow(msg string, keyVals ...any) {
	std.Infow(msg, keyVals...)
}

func (l *zapLogger) Warnf(format string, args ...any) {
	l.z.Sugar().Warnf(format, args...)
}

func Warnf(format string, args ...any) {
	std.Warnf(format, args...)
}

func (l *zapLogger) Warnw(msg string, keyVals ...any) {
	l.z.Sugar().Warnw(msg, keyVals...)
}

func Warnw(msg string, keyVals ...any) {
	std.Warnw(msg, keyVals...)
}

func (l *zapLogger) Errorf(format string, args ...any) {
	l.z.Sugar().Errorf(format, args...)
}

func Errorf(format string, args ...any) {
	std.Errorf(format, args...)
}

func (l *zapLogger) Errorw(err error, msg string, keyVals ...any) {
	l.z.Sugar().Errorw(msg, append(keyVals, "error", err)...)
}

func Errorw(err error, msg string, keyVals ...any) {
	std.Errorw(err, msg, keyVals...)
}

func (l *zapLogger) Panicf(format string, args ...any) {
	l.z.Sugar().Panicf(format, args...)
}

func Panicf(format string, args ...any) {
	std.Panicf(format, args...)
}

func (l *zapLogger) Panicw(msg string, keyVals ...any) {
	l.z.Sugar().Panicw(msg, keyVals...)
}

func Panicw(msg string, keyVals ...any) {
	std.Panicw(msg, keyVals...)
}

func (l *zapLogger) Fatalf(format string, args ...any) {
	l.z.Sugar().Fatalf(format, args...)
}

func Fatalf(format string, args ...any) {
	std.Fatalf(format, args...)
}

func (l *zapLogger) Fatalw(msg string, keyVals ...any) {
	l.z.Sugar().Fatalw(msg, keyVals...)
}

func Fatalw(msg string, keyVals ...any) {
	std.Fatalw(msg, keyVals...)
}

func (l *zapLogger) With(keyVals ...Field) Logger {
	if len(keyVals) == 0 {
		return l
	}

	lc := l.clone()
	lc.z = l.z.With(keyVals...)

	return lc
}

func AddCallerSkip(skip int) Logger {
	return std.AddCallerSkip(skip)
}

// AddCallerSkip increases the number of callers skipped by caller annotation
// (as enabled by the AddCaller option). When building wrappers around the
// Logger and SugaredLogger, supplying this Option prevents zap from always
// reporting the wrapper code as the caller.
func (l *zapLogger) AddCallerSkip(skip int) Logger {
	lc := l.clone()
	lc.z = lc.z.WithOptions(zap.AddCallerSkip(skip))
	return lc
}

// clone 深度拷贝 zapLogger.
func (l *zapLogger) clone() *zapLogger {
	copied := *l
	return &copied
}
