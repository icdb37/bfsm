package cfpx

import (
	"reflect"
	"testing"
	"time"

	"github.com/icdb37/bfsm/internal/infra/config"
)

func TestParseValue(t *testing.T) {
	iv1 := 111
	iv2 := 0
	ivs := "111"
	if err := parseValue(ivs, &iv2); err != nil || iv1 != iv2 {
		t.Fatalf("convert string to int fail")
	}
	fv1 := float64(3.14)
	fv2 := float64(0)
	fvs := "3.14"
	if err := parseValue(fvs, &fv2); err != nil || fv1-fv2 > 0.0001 {
		t.Fatalf("convert string to int fail")
	}
}

func TestCheckLt(t *testing.T) {
	pn := &Param{
		Code: "age",
		Desc: "年龄",
		Val:  "70",
	}
	v1 := 20
	tv := reflect.ValueOf(&v1)
	if err := checkLt(&tv, pn); err != nil {
		t.Fatal(err)
	}
	v1 = 100
	if err := checkLt(&tv, pn); err == nil {
		t.Fatal("100 < 70 check fail")
	}
	v2 := []int{10, 20, 30}
	tv = reflect.ValueOf(&v2)
	if err := checkLt(&tv, pn); err != nil {
		t.Fatal(err)
	}
	v2 = []int{30, 40, 80}
	tv = reflect.ValueOf(&v2)
	if err := checkLt(&tv, pn); err == nil {
		t.Fatal("age [30, 40, 80] must be less 70 check fail")
	}
}

type demoPersion struct {
	Name      string    `cfpx:"field=code:name|desc:名称,fmtfn=trim"`
	Age       int       `cfpx:"field=code:age|desc:年龄,check=lt:100"`
	CreatedAt time.Time `cfpx:"field=code:created_at|desc:创建,fmtfn=nowdt"`
	feature   string
}

func (d *demoPersion) GetFeature() string {
	return d.feature
}

func TestFmtfnOps(t *testing.T) {
	config.SetConfig(config.KeyCfpx, "./config.yaml")
	Init()
	pService.load("./config.yaml")
	info := &demoPersion{
		Name:    " aaa \t",
		Age:     120,
		feature: "user",
	}
	nowTime := time.Now()
	if err := Process(info); err != nil {
		t.Fatal("process fail", "error", err)
	}
	if info.Name != "aaa" {
		t.Fatal("fmtfn trim failed")
	}
	if diff := info.CreatedAt.Sub(nowTime); diff < 0 || diff > time.Minute {
		t.Fatal("fmtfn nowdt failed")
	}
}
