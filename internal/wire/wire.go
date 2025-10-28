// Package wire -
package wire

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/icdb37/bfsm/internal/infra/logx"
)

// gSC - 全局服务缓存
var gSC = &wCache{
	servers: make(map[string]map[any]any),
}

type wKey[T any] struct{}

type wCache struct {
	servers map[string]map[any]any
}

// Provide 注册服务
func Provide[T any](g func() T) {
	ProvideName("", g)
}

// ProvideName 注册服务
func ProvideName[T any](f string, g func() T) {
	if _, ok := gSC.servers[f]; !ok {
		gSC.servers[f] = make(map[any]any)
	}
	gSC.servers[f][wKey[T]{}] = g()
}

// Resolve 获取服务
func Resolve[T any]() T {
	return ResolveName[T]("")
}

// ResolveName 获取服务
func ResolveName[T any](f string) T {
	if _, ok := gSC.servers[f]; !ok {
		gSC.servers[f] = make(map[any]any)
	}
	return gSC.servers[f][wKey[T]{}].(T)
}

type Starter interface {
	Start(ctx context.Context) error
}

type Stopper interface {
	Stop(ctx context.Context) error
}

// Start 启动服务
func Start(ctx context.Context) {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	var starts []Starter
	var stops []Stopper
	for _, s := range gSC.servers {
		if st, ok := s[wKey[Starter]{}].(Starter); ok {
			starts = append(starts, st)
		}
		if sp, ok := s[wKey[Stopper]{}].(Stopper); ok {
			stops = append(stops, sp)
		}
	}
	gSC.servers = make(map[string]map[any]any)
	for _, st := range starts {
		if err := st.Start(ctx); err != nil {
			logx.Fatal("start server failed", "error", err)
		}
	}
	<-sigint
	logx.Info("signal interrupt, will stop all servers")
	for _, sp := range stops {
		if err := sp.Stop(ctx); err != nil {
			logx.Error("stop server failed", "error", err)
		}
	}
	time.Sleep(time.Second * 2)
	logx.Info("all servers stopped")
}
