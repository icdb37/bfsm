// Package errx 业务异常
package errx

import "fmt"

// ErrStore 业务异常存储
type ErrStore struct {
	Table   string `json:"table,omitempty"`
	Query   string `json:"where,omitempty"`
	Message string `json:"message,omitempty"`
}

// Error 实现 error 接口
func (e *ErrStore) Error() string {
	return e.Message
}

// NewStore 存储业务异常
func NewStore(format string, args ...any) *ErrStore {
	return &ErrStore{
		Message: fmt.Sprintf(format, args...),
	}
}

// ErrCfpx cfpx校验异常
type ErrCfpx struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message,omitempty"`
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
	Message string `json:"message,omitempty"`
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
	Message string `json:"message,omitempty"`
}

// Error 实现 error 接口
func (e *ErrStatus) Error() string {
	return e.Message
}

// NewStatus 创建状态异常
func NewStatus(format string, args ...any) *ErrStatus {
	return &ErrStatus{
		Message: fmt.Sprintf(format, args...),
	}
}

type ErrParam struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message,omitempty"`
}

// Error 实现 error 接口
func (e *ErrParam) Error() string {
	return e.Message
}

func NewParam(field string, format string, args ...any) *ErrParam {
	return &ErrParam{
		Field:   field,
		Message: fmt.Sprintf(format, args...),
	}
}
