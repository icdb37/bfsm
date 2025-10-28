package sqlite

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/icdb37/bfsm/internal/infra/store"
	"xorm.io/builder"
)

const (
	tagKey = "where"
	tagLen = 3
)

const (
	opEq  = "eq"
	opNe  = "ne"
	opIn  = "in"
	opNin = "nin"
	opLt  = "lt"
	opLte = "lte"
	opGt  = "gt"
	opGte = "gte"
	// 正则匹配
	opRegex = "regex"
	// 区间查找
	opBetween = "between"
	opRange   = "range"
	// and子节点
	opAnd = "and"
	// or子节点
	opOr = "or"
)

type tagWhere struct {
	Op        string
	Field     string
	Omitempty bool
}

func parseTag(strField, strTag string) *tagWhere {
	t := &tagWhere{}
	parts := strings.Split(strTag, ",")
	partSize := len(parts)
	if len(parts) == 0 {
		return nil
	}
	for i, size := 0, tagLen-partSize; i < size; i++ {
		parts = append(parts, "")
	}
	t.Op = parts[0]
	t.Field = parseTagField(strField, parts[1])
	t.Omitempty = strings.TrimSpace(parts[2]) == "omitempty"
	return t
}

func parseTagField(name, tag string) string {
	tag = strings.TrimSpace(tag)
	if tag == "-" {
		return ""
	}
	if tag != "" {
		return tag
	}
	vb := []byte{}
	underline := false
	for _, v := range []byte(name) {
		if v >= 'A' && v <= 'Z' {
			if underline {
				vb = append(vb, '_')
			}
			vb = append(vb, 'a'+v-'A')
			underline = false
		} else {
			vb = append(vb, v)
			underline = true
		}
	}
	return string(vb)
}

// Unmarshal 从结构体中解析出Where条件
func Unmarshal(v any) store.Filter {
	w := &wfilter{}
	w.And(newAndCond(v))
	return w
}

type parseCond func(key string, pv *reflect.Value) []builder.Cond

type leafCond struct {
	op  string
	key string
	val any
}

func (c *leafCond) WriteTo(writer builder.Writer) error {
	writer.Write([]byte(fmt.Sprintf("%s %s ?", c.key, c.op)))
	writer.Append(c.val)
	return nil
}
func (c *leafCond) And(...builder.Cond) builder.Cond {
	return nil
}
func (c *leafCond) Or(...builder.Cond) builder.Cond {
	return nil
}
func (c *leafCond) IsValid() bool {
	return true
}

func newCond(op, key string, pv *reflect.Value) builder.Cond {
	switch op {
	case opEq:
		op = "="
	case opNe:
		op = "<>"
	case opIn:
		op = "in"
	case opNin:
		op = "not in"
	case opLt:
		op = "<"
	case opLte:
		op = "<="
	case opGt:
		op = ">"
	case opGte:
		op = ">="
	case opRegex:
		return &leafCond{
			op:  "like",
			key: key,
			val: fmt.Sprintf("%%%v%%", pv.Interface()),
		}
	}
	return &leafCond{
		op:  op,
		key: key,
		val: pv.Interface(),
	}
}

type betweenCond struct {
	key  string
	vMin any
	vMax any
}

func (c *betweenCond) WriteTo(writer builder.Writer) error {
	writer.Write([]byte(fmt.Sprintf("%s between ? and ?", c.key)))
	writer.Append(c.vMin, c.vMax)
	return nil
}
func (c *betweenCond) And(...builder.Cond) builder.Cond {
	return nil
}
func (c *betweenCond) Or(...builder.Cond) builder.Cond {
	return nil
}
func (c *betweenCond) IsValid() bool {
	return true
}
func newBetweenCond(key string, pv *reflect.Value) builder.Cond {
	if pv.Kind() == reflect.Ptr {
		tvv := pv.Elem()
		pv = &tvv
	}
	c := &betweenCond{
		key: key,
	}
	switch pv.Kind() {
	case reflect.Array, reflect.Slice:
		size := pv.Len()
		for i := 0; i < size && i < 2; i++ {
			if i == 0 {
				c.vMin = pv.Index(i).Interface()
			} else {
				c.vMax = pv.Index(i).Interface()
			}
		}
	case reflect.Map:
		// 遍历 map，取前两个 key 对应的 value 作为区间边界
		iter := pv.MapRange()
		for iter.Next() {
			key := strings.ToLower(fmt.Sprint(iter.Key().Interface()))
			switch key {
			case "min", "vmin", "start", "start_at", "started_at", "beg", "begin":
				c.vMin = iter.Value().Interface()
			case "max", "vmax", "end", "ended_at", "end_at", "stop", "stop_at", "stopped_at":
				c.vMax = iter.Value().Interface()
			}
		}
	case reflect.Struct:
		for i, size := 0, pv.NumField(); i < size; i++ {
			f := pv.Type().Field(i)
			tagVal := f.Tag.Get(tagKey)
			tvv := pv.Field(i)
			tvi := tvv.Interface()
			tag := parseTag(f.Name, tagVal)
			if tag == nil {
				continue
			}
			if tvv.IsZero() && tag.Omitempty {
				continue
			}
			switch tag.Op {
			case opLt, opLte:
				c.vMax = tvi
			case opGt, opGte:
				c.vMin = tvi
			}
		}
	}
	return c
}

func newAndCond(v any) builder.Cond {
	vv := reflect.ValueOf(v)
	if vv.Kind() == reflect.Ptr {
		vv = vv.Elem()
	}
	conds := parseAndCond("", &vv)
	return builder.And(conds...)
}

func parseAndCond(key string, pv *reflect.Value) []builder.Cond {
	conds := []builder.Cond{}
	if pv.Kind() == reflect.Ptr {
		tvv := pv.Elem()
		pv = &tvv
	}
	switch pv.Kind() {
	case reflect.Array, reflect.Slice:
		conds = parseCondSlice(key, pv, parseAndCond)
	case reflect.Map:
		conds = parseCondMap(pv, parseAndCond)
	case reflect.Struct:
		conds = append(conds, builder.And(parseCondStruct(key, pv)...))
	case reflect.Pointer:
		if pv.IsNil() {
			return nil
		}
		tvv := pv.Elem()
		pv = &tvv
		conds = parseAndCond(key, pv)
	}
	return conds
}
func parseCondSlice(key string, pv *reflect.Value, fSubParse parseCond) []builder.Cond {
	conds := []builder.Cond{}
	if key != "" {
		conds = append(conds, &leafCond{
			op:  opIn,
			key: key,
			val: pv.Interface(),
		})
		return conds
	}
	size := pv.Len()
	for i := 0; i < size; i++ {
		f := pv.Type().Field(i)
		tagVal := f.Tag.Get(tagKey)
		tvv := pv.Field(i)
		tag := parseTag(f.Name, tagVal)
		if tag == nil {
			continue
		}
		if tvv.IsZero() && tag.Omitempty {
			continue
		}
		tkey := key
		if tag.Field != "" {
			tkey = tag.Field
		}
		conds = append(conds, fSubParse(tkey, &tvv)...)
	}
	return conds
}
func parseCondMap(pv *reflect.Value, fSubParse parseCond) []builder.Cond {
	conds := []builder.Cond{}
	iter := pv.MapRange()
	for iter.Next() {
		tkey := strings.ToLower(fmt.Sprint(iter.Key().Interface()))
		tvv := iter.Value()
		conds = append(conds, fSubParse(tkey, &tvv)...)
	}
	return conds
}
func parseCondStruct(key string, pv *reflect.Value) []builder.Cond {
	conds := []builder.Cond{}
	for i, size := 0, pv.NumField(); i < size; i++ {
		f := pv.Type().Field(i)
		tagVal := f.Tag.Get(tagKey)
		tvv := pv.Field(i)
		tag := parseTag(f.Name, tagVal)
		if tag == nil {
			continue
		}
		if tag.Op == "" || tag.Op == "-" || tag.Field == "-" {
			conds = append(conds, parseCondStruct(key, &tvv)...)
			continue
		}
		if tvv.IsZero() && tag.Omitempty {
			continue
		}
		tkey := key
		if tag.Field != "" {
			tkey = tag.Field
		}
		switch tag.Op {
		case opBetween, opRange:
			conds = append(conds, newBetweenCond(tkey, &tvv))
		case opAnd:
			conds = append(conds, parseAndCond(tkey, &tvv)...)
		case opOr:
			conds = append(conds, parseOrCond(tkey, &tvv)...)
		default:
			conds = append(conds, newCond(tag.Op, tkey, &tvv))
		}
	}
	return conds
}

func parseOrCond(key string, pv *reflect.Value) []builder.Cond {
	conds := []builder.Cond{}
	if pv.Kind() == reflect.Ptr {
		tvv := pv.Elem()
		pv = &tvv
	}
	switch pv.Kind() {
	case reflect.Array, reflect.Slice:
		conds = parseCondSlice(key, pv, parseOrCond)
	case reflect.Map:
		conds = parseCondMap(pv, parseOrCond)
	case reflect.Struct:
		conds = append(conds, builder.Or(parseCondStruct(key, pv)...))
	case reflect.Pointer:
		if pv.IsNil() {
			return nil
		}
		tvv := pv.Elem()
		pv = &tvv
		conds = parseOrCond(key, pv)
	}
	return conds
}
