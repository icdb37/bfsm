package cfpx

import (
	"testing"
	"time"
)

func TestUserConfig(t *testing.T) {
	s := &service{}
	s.load("./config.yaml")
	if len(s.itemFeature.Item) == 0 {
		t.Fatal("items is nil")
	}
	info := &demoPersion{
		Name:    " aaa \t",
		Age:     80,
		feature: "user",
	}
	nowTime := time.Now()
	if err := s.Process(info); err != nil {
		t.Fatal("process config more priority failed", "error", err)
	}
	if info.Name != " aaa" { //移除右边空白字符
		t.Fatal("fmtfn trim failed")
	}
	if diff := info.CreatedAt.Sub(nowTime); diff < 0 || diff > time.Minute {
		t.Fatal("fmtfn nowdt failed")
	}
}

func TestDefaultConfig(t *testing.T) {
	s := &service{}
	s.load("./config.yaml")
	if len(s.itemFeature.Item) == 0 {
		t.Fatal("items is nil")
	}
	info := &demoPersion{
		Name: " aaa \t",
		Age:  150,
	}
	nowTime := time.Now()
	if err := s.Process(info); err != nil {
		t.Fatal("process config more priority failed", "error", err)
	}
	if info.Name != "aaa" { //移除两端空白字符
		t.Fatal("fmtfn trim failed")
	}
	if diff := info.CreatedAt.Sub(nowTime); diff < 0 || diff > time.Minute {
		t.Fatal("fmtfn nowdt failed")
	}
}
