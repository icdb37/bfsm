// Package cfpx 自定义字段处理器
package cfpx

import (
	"reflect"
	"strings"
	"time"
)

const (
	keyFmtfn = "fmtfn" // 数据格式化
	keyCheck = "check" // 数据校验
)

type fFmtfn = func(pv *reflect.Value)
type fCheck = func(pv *reflect.Value, pf *field, pn *node) error

const (
	OpFmtfnTriml = "triml" // 左去空格
	OpFmtfnTrimr = "trimr" // 右去空格
	OpFmtfnTrim  = "trim"  // 去空格
	OpFmtfnLower = "lower" // 小写
	OpFmtfnUpper = "upper" // 大写
	OpFmtfnNowdt = "nowdt" // 当前时间
)

var (
	typeTime = reflect.TypeOf(time.Time{})
	fmtfnOps = map[string]fFmtfn{
		OpFmtfnTriml: fmtfnTriml,
		OpFmtfnTrimr: fmtfnTrimr,
		OpFmtfnTrim:  fmtfnTrim,
		OpFmtfnLower: fmtfnLower,
		OpFmtfnUpper: fmtfnUpper,
		OpFmtfnNowdt: fmtfnNowdt,
	}
)

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

/*
示例：
 {tag},cfpx:"fmtfn=trim|lower,check=in:1,2,3|minlen:3"
*/

type field struct {
	name string
}

type node struct {
	name string
	desc string
	opx  string
	val  string
}

func parseCheckNode(s string, beg, end int) (ns []*node) {
	if beg == -1 {
		return
	}
	s = s[beg+len(keyCheck) : end]
	if s[beg] != ' ' && s[beg] != '=' {
		return
	}
	s = strings.Trim(s, " =")
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return
	}
	fields := strings.Split(s, "|")
	for _, field := range fields {
		field = strings.TrimSpace(field)
		if len(field) == 0 {
			continue
		}
		kvs := strings.Split(field, ":")
		pn := &node{}
		for i, pnv := range []*string{&pn.opx, &pn.val} {
			if i >= len(kvs) {
				break
			}
			*pnv = kvs[i]
		}
		pn.opx = strings.ToLower(pn.opx)
		if _, ok := checkOps[pn.opx]; !ok {
			continue
		}
		ns = append(ns, pn)
	}
	return ns
}

func parseFmtfnNode(s string, beg, end int) (ns []*node) {
	if beg == -1 {
		return
	}
	s = s[beg+len(keyFmtfn) : end]
	if s[beg] != ' ' && s[beg] != '=' {
		return
	}
	s = strings.Trim(s, " =")
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return
	}
	fields := strings.Split(s, "|")
	for _, field := range fields {
		field = strings.TrimSpace(field)
		if len(field) == 0 {
			continue
		}
		kvs := strings.Split(field, ":")
		pn := &node{}
		for i, pnv := range []*string{&pn.opx, &pn.val} {
			if i >= len(kvs) {
				break
			}
			*pnv = kvs[i]
		}
		pn.opx = strings.ToLower(pn.opx)
		if _, ok := fmtfnOps[pn.opx]; !ok {
			continue
		}
		ns = append(ns, pn)
	}
	return ns
}

func parseNodes(tag string) (fmtfns []*node, checks []*node) {
	name := ""
	desc := ""
	if pos := strings.Index(tag, ","); pos != -1 {
		name = strings.TrimSpace(tag[:pos])
		tag = tag[pos+1:]
		if pos = strings.Index(name, "="); pos != -1 {
			desc = strings.TrimSpace(name[pos+1:])
			name = strings.TrimSpace(name[:pos])
		} else {
			desc = name
		}
	}
	tag = strings.TrimSpace(tag)
	size := len(tag)
	begFmtfn := strings.Index(tag, keyFmtfn)
	begCheck := strings.Index(tag, keyCheck)
	endFmtfn, endCheck := size, size
	if begCheck != -1 && begFmtfn != -1 && begCheck > begFmtfn {
		endFmtfn = begCheck
	}
	if begCheck != -1 && begFmtfn != -1 && begFmtfn > begCheck {
		endCheck = begFmtfn
	}
	fmtfns = parseFmtfnNode(tag, begFmtfn, endFmtfn)
	checks = parseCheckNode(tag, begCheck, endCheck)
	for _, n := range fmtfns {
		n.name = name
		n.desc = desc
	}
	for _, n := range checks {
		n.name = name
		n.desc = desc
	}
	return
}

// Process 数据格式化，字段校验
func Process(param any) error {
	v := reflect.ValueOf(param)
	if v.Kind() != reflect.Ptr {
		return nil
	}
	v = v.Elem()
	if v.Kind() != reflect.Struct {
		return nil
	}
	// 遍历结构体字段
	for i := 0; i < v.NumField(); i++ {
		iv := v.Field(i)
		f := v.Type().Field(i)
		if f.Anonymous {
			continue
		}
		// 检查是否有 cfpx 标签
		tag := f.Tag.Get("cfpx")
		if tag == "" {
			continue
		}
		tag = strings.TrimSpace(tag)
		fmtfns, checks := parseNodes(tag)
		for _, n := range fmtfns {
			pf, ok := fmtfnOps[n.opx]
			if !ok {
				continue
			}
			pf(&iv)
		}
		for _, n := range checks {
			pf, ok := checkOps[n.opx]
			if !ok {
				continue
			}
			if err := pf(&iv, &field{name: f.Name}, n); err != nil {
				return err
			}
		}
	}
	return nil
}
