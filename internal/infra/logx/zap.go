package logx

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func newRotateSink(opt *Options) zap.Sink {
	lj := &lumberjack.Logger{
		Filename:   opt.File,
		MaxSize:    opt.Size,
		MaxAge:     opt.Age,
		MaxBackups: opt.Backups,
		Compress:   opt.Compress,
	}

	flushInterval := opt.FlushInterval
	if flushInterval <= 0 {
		flushInterval = 5 * time.Second
	}

	syncer := &zapcore.BufferedWriteSyncer{
		WS:            zapcore.AddSync(lj),
		FlushInterval: flushInterval,
	}

	return &rotateSink{
		syncer,
		lj,
	}
}

func resolveZapOutputs(opt *Options) ([]string, error) {
	outputs := make([]string, 0, 2)

	// 日志文件分割
	if opt.File != "" {
		err := zap.RegisterSink("rotate", func(url *url.URL) (zap.Sink, error) {
			return newRotateSink(opt), nil
		})
		if err != nil {
			return nil, fmt.Errorf("register rotate zap.Sink error: %v", err)
		}

		outputs = append(outputs, fmt.Sprintf("rotate:%s", opt.File))
	}

	if opt.StdOut {
		outputs = append(outputs, "stdout")
	}

	return outputs, nil
}

func parseLevel(level string) zapcore.Level {
	level = strings.ToLower(level)
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info", "INFO":
		return zapcore.InfoLevel
	case "warn", "WARN":
		return zapcore.WarnLevel
	case "error", "ERROR":
		return zapcore.ErrorLevel
	case "fatal", "FATAL":
		return zapcore.FatalLevel
	}
	return zapcore.InfoLevel
}

func newZapSugared(opt *Options) (*zap.SugaredLogger, error) {
	outputs, err := resolveZapOutputs(opt)
	if err != nil {
		return nil, err
	}

	config := zap.NewProductionConfig()
	config.OutputPaths = outputs
	config.Level = zap.NewAtomicLevelAt(parseLevel(opt.Level))
	config.EncoderConfig.TimeKey = "time"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.DisableStacktrace = !opt.Stacktrace

	var buildOpts []zap.Option
	if opt.Caller {
		buildOpts = append(buildOpts, zap.AddCallerSkip(2), zap.WithCaller(true))
	} else {
		buildOpts = append(buildOpts, zap.WithCaller(false))
	}

	// 当开启debug和明确指定禁用采样时取消采样
	if opt.DisableSampling || parseLevel(opt.Level) == zapcore.DebugLevel {
		config.Sampling = nil
	}

	zl, err := config.Build(buildOpts...)
	if err != nil {
		return nil, fmt.Errorf("build zap.Log error: %v", err)
	}

	return zl.Sugar(), nil
}

type rotateSink struct {
	*zapcore.BufferedWriteSyncer
	lj *lumberjack.Logger
}

func (r *rotateSink) Close() error {
	err1 := r.Stop()
	err2 := r.lj.Close()

	if err1 != nil {
		return err1
	}

	return err2
}

type zapLogger struct {
	sugar      *zap.SugaredLogger
	callerSkip int
}

func newZapLogger(sugared *zap.SugaredLogger) Logger {
	return &zapLogger{
		sugar:      sugared,
		callerSkip: 2,
	}
}

func (zl *zapLogger) Debug(msg string, kv ...any) {
	zl.sugar.Debugw(msg, kv...)
}

func (zl *zapLogger) Info(msg string, kv ...any) {
	zl.sugar.Infow(msg, kv...)
}

func (zl *zapLogger) Warn(msg string, kv ...any) {
	zl.sugar.Warnw(msg, kv...)
}

func (zl *zapLogger) Error(msg string, kv ...any) {
	zl.sugar.Errorw(msg, kv...)
}

func (zl *zapLogger) Fatal(msg string, kv ...any) {
	zl.sugar.Fatalw(msg, kv...)
}

func (zl *zapLogger) With(kv ...any) Logger {
	if zl.callerSkip == 1 {
		if len(kv) == 0 {
			return zl
		}

		return &zapLogger{
			sugar:      zl.sugar.With(kv...),
			callerSkip: 1,
		}
	} else {
		return &zapLogger{
			sugar:      zl.sugar.WithOptions(zap.AddCallerSkip(-1)).With(kv...),
			callerSkip: 1,
		}
	}
}

func (zl *zapLogger) Flush() error {
	return zl.sugar.Sync()
}

func (zl *zapLogger) IsDebugEnabled() bool {
	return zl.sugar.Level() == zapcore.DebugLevel
}
