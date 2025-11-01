// Package enum 枚举
package enum

type StatusCode uint16

const (
	StatusCodeUndefined StatusCode = 0     // 未定义
	StatusCodeCompleted StatusCode = 10000 // 已完成
)
