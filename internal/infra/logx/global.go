package logx

import "go.uber.org/zap"

var (
	log = newZapLogger(zap.NewNop().
		WithOptions(zap.AddCallerSkip(2)).
		Sugar())
	emptyWithLog = log.With()
)

// Debug print debug log
func Debug(msg string, kv ...any) {
	log.Debug(msg, kv...)
}

// Info print info log
func Info(msg string, kv ...any) {
	log.Info(msg, kv...)
}

// Warn print warn log
func Warn(msg string, kv ...any) {
	log.Warn(msg, kv...)
}

// Error print error log
func Error(msg string, kv ...any) {
	log.Error(msg, kv...)
}

// Fatal print fatal log and exit
func Fatal(msg string, kv ...any) {
	log.Fatal(msg, kv...)
}

func With(kv ...any) Logger {
	return log.With(kv...)
}

// Flush flush log buffer to disk
func Flush() error {
	return log.Flush()
}

func IsDebugEnabled() bool {
	return log.IsDebugEnabled()
}
