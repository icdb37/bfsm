// Package fetchx 获取器
package fetchx

import "context"

// GetFunc 是一个获取器接口
type GetFunc[T any] = func(context.Context) T

// SetFunc 是一个设置器函数
type SetFunc[T any] = func(context.Context, T)

// Featcher 是一个获取器接口
type Featcher[T any] interface {
	Fetch(context.Context) T
}

// NewFeatcher 创建一个新的获取器
func NewFeatcher[T any](fg GetFunc[T], fs SetFunc[T]) Featcher[T] {
	return &feat[T]{
		fg: fg,
		fs: fs,
	}
}

type feat[T any] struct {
	fg GetFunc[T]
	fs SetFunc[T]
	ve bool
	vd T
}

// Fetch 获取值
func (f *feat[T]) Fetch(ctx context.Context) T {
	if f.ve {
		return f.vd
	}
	f.vd = f.fg(ctx)
	return f.vd
}
