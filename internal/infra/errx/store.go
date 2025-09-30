// Package errx 业务异常
package errx

import "fmt"

// ErrStore 业务异常存储
type ErrStore struct {
	Table   string `json:"table"`
	Query   string `json:"where"`
	Message string `json:"message"`
}

// Error 实现 error 接口
func (e *ErrStore) Error() string {
	return e.Message
}

// Store 存储业务异常
func Store(format string, args ...any) *ErrStore {
	return &ErrStore{
		Message: fmt.Sprintf(format, args...),
	}
}
