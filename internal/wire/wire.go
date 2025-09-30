// Package wire -
package wire

// gSC - 全局服务缓存
var gSC = &wCache{
	servers: make(map[any]any),
}

type wKey[T any] struct{}

type wCache struct {
	servers map[any]any
}

// Provide 注册服务
func Provide[T any](g func() T) {
	gSC.servers[wKey[T]{}] = g()
}

// Resolve 获取服务
func Resolve[T any]() T {
	return gSC.servers[wKey[T]{}].(T)
}
