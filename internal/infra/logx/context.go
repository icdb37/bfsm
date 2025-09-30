package logx

import "context"

type contextKey struct{}

var key contextKey

func FromContext(ctx context.Context) Logger {
	local, ok := ctx.Value(key).(Logger)
	if !ok {
		// 解决堆栈打印问题
		local = emptyWithLog
	}
	return local
}

func WithContext(ctx context.Context, logger Logger) context.Context {
	return context.WithValue(ctx, key, logger)
}

func CopyContext(from, to context.Context) context.Context {
	v := from.Value(key)
	if v != nil {
		return context.WithValue(to, key, v)
	} else {
		return to
	}
}
