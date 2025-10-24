package cfpx

import (
	"encoding/base64"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/icdb37/bfsm/internal/utils"
)

func fmtfnTriml(pv *reflect.Value) {
	if pv.Kind() != reflect.String || !pv.CanSet() {
		return
	}
	tv := pv.String()
	tv = strings.TrimLeft(tv, " \t\r\n")
	pv.SetString(tv)
}

func fmtfnTrimr(pv *reflect.Value) {
	if pv.Kind() != reflect.String || !pv.CanSet() {
		return
	}
	tv := pv.String()
	tv = strings.TrimRight(tv, " \t\r\n")
	pv.SetString(tv)
}

func fmtfnTrim(pv *reflect.Value) {
	if pv.Kind() != reflect.String || !pv.CanSet() {
		return
	}
	tv := pv.String()
	tv = strings.TrimSpace(tv)
	pv.SetString(tv)
}
func fmtfnLower(pv *reflect.Value) {
	if pv.Kind() != reflect.String || !pv.CanSet() {
		return
	}
	tv := pv.String()
	tv = strings.ToLower(tv)
	pv.SetString(tv)
}

func fmtfnUpper(pv *reflect.Value) {
	if pv.Kind() != reflect.String || !pv.CanSet() {
		return
	}
	tv := pv.String()
	tv = strings.ToUpper(tv)
	pv.SetString(tv)
}

func fmtfnNowdt(pv *reflect.Value) {
	if pv.Type() != typeTime || !pv.CanSet() {
		return
	}
	tv := time.Now()
	pv.Set(reflect.ValueOf(tv))
}
func parseValue[T comparable](s string, pv *T) error {
	if pv == nil {
		return nil
	}
	switch v := any(pv).(type) {
	case *int64:
		val, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return err
		}
		*v = val
	case *uint64:
		val, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			return err
		}
		*v = val
	case *float64:
		val, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return err
		}
		*v = val
	case *int:
		val, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return err
		}
		*v = int(val)
	default:
		return fmt.Errorf("not support %v convert to numerical value", *pv)
	}
	return nil
}

func parseValues[T comparable](is string, pvs ...*T) error {
	ss := strings.Split(is, ",")
	if len(ss) == 0 {
		return fmt.Errorf("param empty")
	}
	for i := 0; i < len(pvs) && i < len(ss); i++ {
		sv := strings.TrimSpace(ss[i])
		if err := parseValue(sv, pvs[i]); err != nil {
			return err
		}
	}
	return nil
}

func checkEq(pv *reflect.Value, pf *field, pn *node) error {
	if !pv.Comparable() || pn.val == "" {
		return nil
	}
	if pn.val != fmt.Sprint(pv.Interface()) {
		return fmt.Errorf("%s 等于 %s", pn.desc, pn.val)
	}
	return nil
}

func checkLt(pv *reflect.Value, pf *field, pn *node) error {
	if !pv.Comparable() || pn.val == "" {
		return nil
	}
	if pv.Kind() == reflect.Pointer {
		tv := pv.Elem()
		pv = &tv
	}
	switch pv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		var min int64
		if err := parseValue(pn.val, &min); err != nil {
			return nil
		}
		tv := pv.Int()
		if tv >= min {
			return fmt.Errorf("%s 必须小于 %d", pn.desc, min)
		}
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		var min uint64
		if err := parseValue(pn.val, &min); err != nil {
			return nil
		}
		tv := pv.Uint()
		if tv >= min {
			return fmt.Errorf("%s 必须小于 %d", pn.desc, min)
		}
	case reflect.Float32, reflect.Float64:
		var min float64
		if err := parseValue(pn.val, &min); err != nil {
			return nil
		}
		tv := pv.Float()
		if tv >= min {
			return fmt.Errorf("%s 必须小于 %f", pn.desc, min)
		}
	case reflect.String:
		var min int
		if err := parseValue(pn.val, &min); err != nil {
			return nil
		}
		tv := pv.String()
		if len(tv) >= min {
			return fmt.Errorf("%s 长度必须小于 %d", pn.desc, min)
		}
	case reflect.Slice, reflect.Array:
		for i, size := 0, pv.Len(); i < size; i++ {
			tv := pv.Index(i)
			if err := checkLt(&tv, pf, pn); err != nil {
				return err
			}
		}
	}
	return nil
}

func checkLte(pv *reflect.Value, pf *field, pn *node) error {
	if !pv.Comparable() || pn.val == "" {
		return nil
	}
	if pv.Kind() == reflect.Pointer {
		tv := pv.Elem()
		pv = &tv
	}
	switch pv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		var min int64
		if err := parseValue(pn.val, &min); err != nil {
			return nil
		}
		tv := pv.Int()
		if tv > min {
			return fmt.Errorf("%s 必须小于等于 %d", pn.desc, min)
		}
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		var min uint64
		if err := parseValue(pn.val, &min); err != nil {
			return nil
		}
		tv := pv.Uint()
		if tv > min {
			return fmt.Errorf("%s 必须小于等于 %d", pn.desc, min)
		}
	case reflect.Float32, reflect.Float64:
		var min float64
		if err := parseValue(pn.val, &min); err != nil {
			return nil
		}
		tv := pv.Float()
		if tv > min {
			return fmt.Errorf("%s 必须小于等于 %f", pn.desc, min)
		}
	case reflect.String:
		var min int
		if err := parseValue(pn.val, &min); err != nil {
			return nil
		}
		tv := pv.String()
		if len(tv) >= min {
			return fmt.Errorf("%s 长度必须小于等于 %d", pn.desc, min)
		}
	case reflect.Slice, reflect.Array:
		for i, size := 0, pv.Len(); i < size; i++ {
			tv := pv.Index(i)
			if err := checkLte(&tv, pf, pn); err != nil {
				return err
			}
		}
	}
	return nil
}

func checkGt(pv *reflect.Value, pf *field, pn *node) error {
	if !pv.Comparable() || pn.val == "" {
		return nil
	}
	if pv.Kind() == reflect.Pointer {
		tv := pv.Elem()
		pv = &tv
	}
	switch pv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		var max int64
		if err := parseValue(pn.val, &max); err != nil {
			return nil
		}
		tv := pv.Int()
		if tv <= max {
			return fmt.Errorf("%s 必须大于 %d", pn.desc, max)
		}
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		var max uint64
		if err := parseValue(pn.val, &max); err != nil {
			return nil
		}
		tv := pv.Uint()
		if tv <= max {
			return fmt.Errorf("%s 必须大于 %d", pn.desc, max)
		}
	case reflect.Float32, reflect.Float64:
		var max float64
		if err := parseValue(pn.val, &max); err != nil {
			return nil
		}
		tv := pv.Float()
		if tv <= max {
			return fmt.Errorf("%s 必须大于 %f", pn.desc, max)
		}
	case reflect.String:
		var max int
		if err := parseValue(pn.val, &max); err != nil {
			return nil
		}
		tv := pv.String()
		if len(tv) <= max {
			return fmt.Errorf("%s 长度必须大于 %d", pn.desc, max)
		}
	case reflect.Slice, reflect.Array:
		for i, size := 0, pv.Len(); i < size; i++ {
			tv := pv.Index(i)
			if err := checkGt(&tv, pf, pn); err != nil {
				return err
			}
		}
	}
	return nil
}

func checkGte(pv *reflect.Value, pf *field, pn *node) error {
	if !pv.Comparable() || pn.val == "" {
		return nil
	}
	if pv.Kind() == reflect.Pointer {
		tv := pv.Elem()
		pv = &tv
	}
	switch pv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		var max int64
		if err := parseValue(pn.val, &max); err != nil {
			return nil
		}
		tv := pv.Int()
		if tv < max {
			return fmt.Errorf("%s 必须大于等于 %d", pn.desc, max)
		}
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		var max uint64
		if err := parseValue(pn.val, &max); err != nil {
			return nil
		}
		tv := pv.Uint()
		if tv < max {
			return fmt.Errorf("%s 必须大于等于 %d", pn.desc, max)
		}
	case reflect.Float32, reflect.Float64:
		var max float64
		if err := parseValue(pn.val, &max); err != nil {
			return nil
		}
		tv := pv.Float()
		if tv < max {
			return fmt.Errorf("%s 必须大于等于 %f", pn.desc, max)
		}
	case reflect.String:
		var max int
		if err := parseValue(pn.val, &max); err != nil {
			return nil
		}
		tv := pv.String()
		if len(tv) <= max {
			return fmt.Errorf("%s 长度必须大于等于 %d", pn.desc, max)
		}
	case reflect.Slice, reflect.Array:
		for i, size := 0, pv.Len(); i < size; i++ {
			tv := pv.Index(i)
			if err := checkGte(&tv, pf, pn); err != nil {
				return err
			}
		}
	}
	return nil
}

func checkIn(pv *reflect.Value, pf *field, pn *node) error {
	if !pv.Comparable() || pn.val == "" {
		return nil
	}
	if pv.Kind() == reflect.Pointer {
		tv := pv.Elem()
		pv = &tv
	}
	vs := strings.Split(pn.val, ",")
	switch pv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		vis := utils.PmakeX[int64](len(vs))
		if err := parseValues(pn.val, vis...); err != nil {
			return nil
		}
		tv := pv.Int()
		if !utils.Pcontain(vis, &tv) {
			return fmt.Errorf("%s 必须属于 %s", pn.desc, pn.val)
		}
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		vis := utils.PmakeX[uint64](len(vs))
		if err := parseValues(pn.val, vis...); err != nil {
			return nil
		}
		tv := pv.Uint()
		if !utils.Pcontain(vis, &tv) {
			return fmt.Errorf("%s 必须属于 %s", pn.desc, pn.val)
		}
	case reflect.String:
		utils.SstrTrims(vs)
		tv := pv.String()
		if !utils.Contain(vs, tv) {
			return fmt.Errorf("%s 必须属于 %s", pn.desc, pn.val)
		}
	case reflect.Array, reflect.Slice:
		for i, size := 0, pv.Len(); i < size; i++ {
			tv := pv.Index(i)
			if err := checkIn(&tv, pf, pn); err != nil {
				return err
			}
		}
	}
	return nil
}

func checkRegex(pv *reflect.Value, pf *field, pn *node) error {
	if pv.Kind() == reflect.Pointer {
		tv := pv.Elem()
		pv = &tv
	}
	switch pv.Kind() {
	case reflect.String:
		pattern, err := base64.StdEncoding.DecodeString(pn.val)
		if err == nil {
			pattern = []byte(pn.val)
		}
		re, err := regexp.Compile(string(pattern))
		if err != nil {
			return nil
		}
		if !re.MatchString(fmt.Sprint(pv.Interface())) {
			return fmt.Errorf("%s 校验失败", pn.desc)
		}
	case reflect.Array, reflect.Slice:
		for i, size := 0, pv.Len(); i < size; i++ {
			tv := pv.Index(i)
			if err := checkRegex(&tv, pf, pn); err != nil {
				return err
			}
		}
	}

	return nil
}

func checkRange(pv *reflect.Value, pf *field, pn *node) error {
	if !pv.Comparable() || pn.val == "" {
		return nil
	}
	if pv.Kind() == reflect.Pointer {
		tv := pv.Elem()
		pv = &tv
	}
	switch pv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		var min, max int64
		if err := parseValues(pn.val, &min, &max); err != nil {
			return nil
		}
		if min > max {
			max = min
		}
		tv := pv.Int()
		if tv < min || tv > max {
			return fmt.Errorf("%s 必须在 [%d,%d] 之间", pn.desc, min, max)
		}
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		var min, max uint64
		if err := parseValues(pn.val, &min, &max); err != nil {
			return nil
		}
		if min > max {
			max = min
		}
		tv := pv.Uint()
		if tv < min || tv > max {
			return fmt.Errorf("%s 必须在 [%d,%d] 之间", pn.desc, min, max)
		}
	case reflect.Float32, reflect.Float64:
		var min, max float64
		if err := parseValues(pn.val, &min, &max); err != nil {
			return nil
		}
		if min > max {
			max = min
		}
		tv := pv.Float()
		if tv < min || tv > max {
			return fmt.Errorf("%s 必须在 [%f,%f] 之间", pn.desc, min, max)
		}
	case reflect.String:
		var min, max int
		if err := parseValues(pn.val, &min, &max); err != nil {
			return nil
		}
		if min > max {
			max = min
		}
		v := pv.String()
		if len(v) < min || len(v) > max {
			return fmt.Errorf("%s 长度必须在 [%d,%d] 之间", pn.desc, min, max)
		}
	case reflect.Array, reflect.Slice:
		for i, size := 0, pv.Len(); i < size; i++ {
			tv := pv.Index(i)
			if err := checkRange(&tv, pf, pn); err != nil {
				return err
			}
		}
	}
	return nil
}
