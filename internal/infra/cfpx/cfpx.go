// Package cfpx 自定义字段处理器
package cfpx

import (
	"reflect"
	"time"
)

/*
示例：
 {tag},cfpx:"field=code:{tag}|desc:测试,fmtfn=trim|lower,check=in:1,2,3|minlen:3"
*/

const (
	keyFmtfn = "fmtfn" // 数据格式化
	keyCheck = "check" // 数据校验
	keyField = "field" // 字段信息
)

// 数据描述
const (
	OpFieldCode = "code"
	OpFieldDesc = "desc"
)

// 数据格式化操作
const (
	OpFmtfnTriml = "triml" // 左去空格
	OpFmtfnTrimr = "trimr" // 右去空格
	OpFmtfnTrim  = "trim"  // 去空格
	OpFmtfnLower = "lower" // 小写
	OpFmtfnUpper = "upper" // 大写
	OpFmtfnNowdt = "nowdt" // 当前时间
)

// 数据校验操作
const (
	OpCheckEq    = "eq"    // 等于
	OpCheckLt    = "lt"    // 小于
	OpCheckLte   = "lte"   // 小于等于
	OpCheckGt    = "gt"    // 大于
	OpCheckGte   = "gte"   // 大于等于
	OpCheckIn    = "in"    // 包含
	OpCheckRegex = "regex" // 正则匹配，值使用base64格式
	OpCheckRange = "range" // 范围
)

type fFmtfn = func(pv *reflect.Value)
type fCheck = func(pv *reflect.Value, pn *Param) error

var (
	typeTime = reflect.TypeOf(time.Time{})
	fmtfnOps = map[string]fFmtfn{
		OpFmtfnTrim:  fmtfnTrim,
		OpFmtfnTriml: fmtfnTriml,
		OpFmtfnTrimr: fmtfnTrimr,
		OpFmtfnLower: fmtfnLower,
		OpFmtfnUpper: fmtfnUpper,
		OpFmtfnNowdt: fmtfnNowdt,
	}
)

var (
	checkOps = map[string]fCheck{
		OpCheckEq:    checkEq,
		OpCheckLt:    checkLt,
		OpCheckLte:   checkLte,
		OpCheckGt:    checkGt,
		OpCheckGte:   checkGte,
		OpCheckIn:    checkIn,
		OpCheckRegex: checkRegex,
		OpCheckRange: checkRange,
	}
)

// Unit 处理单元
type Unit struct {
	Kind string `json:"kind" yaml:"kind"`
	Op   string `json:"op" yaml:"op"`
	Val  string `json:"val" yaml:"val"`
}

// Elem 处理节点
type Elem struct {
	Code    string  `json:"code" yaml:"code"`
	Desc    string  `json:"desc" yaml:"desc"`
	Process []*Unit `json:"process" yaml:"process"`
}

type Param struct {
	Code string `json:"code" yaml:"code"`
	Desc string `json:"desc" yaml:"desc"`
	Val  string `json:"val" yaml:"val"`
}

// Process 数据格式化，字段校验
func Process(param Featurer) error {
	return pService.Process(param)
}

func processFmtfn(pv *reflect.Value, pn *Elem) error {
	if pn == nil {
		return nil
	}
	for _, p := range pn.Process {
		if p.Kind != keyFmtfn {
			continue
		}
		pf, ok := fmtfnOps[p.Op]
		if !ok {
			continue
		}
		pf(pv)
	}
	return nil
}

func processCheck(pv *reflect.Value, pn *Elem) error {
	if pn == nil {
		return nil
	}
	for _, p := range pn.Process {
		if p.Kind != keyCheck {
			continue
		}
		pf, ok := checkOps[p.Op]
		if !ok {
			continue
		}
		if err := pf(pv, &Param{Code: pn.Code, Desc: pn.Desc, Val: p.Val}); err != nil {
			return err
		}
	}
	return nil
}
