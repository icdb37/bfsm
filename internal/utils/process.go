// Package utils 工具包
package utils

import (
	"context"
	"fmt"
)

// ProcessFunc 流程函数
type ProcessFunc[T any] func(context.Context, T) error

// ProcessAll 所有流程必须成功
func ProcessAll[T any](ctx context.Context, info T, ps ...ProcessFunc[T]) error {
	for _, p := range ps {
		if err := p(ctx, info); err != nil {
			return err
		}
	}
	return nil
}

// ProcessAny 任意流程成功即可
func ProcessAny[T any](ctx context.Context, info T, ps ...ProcessFunc[T]) error {
	for _, p := range ps {
		if err := p(ctx, info); err == nil {
			return nil
		}
	}
	return fmt.Errorf("all process failed")
}
