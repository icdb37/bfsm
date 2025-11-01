package cfpx

import (
	"testing"
	"time"

	"github.com/icdb37/bfsm/internal/infra/config"
)

func TestUserConfig(t *testing.T) {
	config.SetConfig(config.KeyCfpx, "./config.yaml")
	Init()
	info := &demoPersion{
		Name:    " aaa \t",
		Age:     80,
		feature: "user",
	}
	nowTime := time.Now()
	if err := Process(info); err != nil {
		t.Fatal("process config more priority failed", "error", err)
	}
	if info.Name != "aaa" { //移除右边空白字符
		t.Fatal("fmtfn trim failed")
	}
	if diff := info.CreatedAt.Sub(nowTime); diff < 0 || diff > time.Minute {
		t.Fatal("fmtfn nowdt failed")
	}
}

func TestDefaultConfig(t *testing.T) {
	config.SetConfig(config.KeyCfpx, "./config.yaml")
	Init()
	info := &demoPersion{
		Name:    " aaa \t",
		Age:     150,
		feature: "user",
	}
	nowTime := time.Now()
	if err := Process(info); err != nil {
		t.Fatal("process config more priority failed", "error", err)
	}
	if info.Name != "aaa" { //移除两端空白字符
		t.Fatal("fmtfn trim failed")
	}
	if diff := info.CreatedAt.Sub(nowTime); diff < 0 || diff > time.Minute {
		t.Fatal("fmtfn nowdt failed")
	}
}
