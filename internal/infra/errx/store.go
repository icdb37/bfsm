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

// ErrCfpx cfpx校验异常
type ErrCfpx struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// NewCfpx 创建cfpx校验异常
func NewCfpx(field string, format string, args ...any) *ErrCfpx {
	return &ErrCfpx{
		Field:   field,
		Message: fmt.Sprintf(format, args...),
	}
}

// Error 实现 error 接口
func (e *ErrCfpx) Error() string {
	return e.Message
}

// ErrNexist 资源不存在异常
type ErrNexist struct {
	Message string `json:"message"`
}

// Error 实现 error 接口
func (e *ErrNexist) Error() string {
	return e.Message
}

// NewNexist 创建资源不存在异常
func NewNexist(format string, args ...any) *ErrNexist {
	return &ErrNexist{
		Message: fmt.Sprintf(format, args...),
	}
}

// ErrStatus 状态异常
type ErrStatus struct {
	Message string `json:"message"`
}

// Error 实现 error 接口
func (e *ErrStatus) Error() string {
	return e.Message
}

// NewErrStatus 创建状态异常
func NewErrStatus(format string, args ...any) *ErrStatus {
	return &ErrStatus{
		Message: fmt.Sprintf(format, args...),
	}
}
