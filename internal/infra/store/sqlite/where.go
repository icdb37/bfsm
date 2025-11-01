package sqlite

import (
	"fmt"
	"reflect"

	"xorm.io/xorm"

	"github.com/icdb37/bfsm/internal/infra/store"
)

type where struct {
	sess *xorm.Session
}

func (w *where) In(key string, val any) store.Filter {
	if w.skip(val) {
		return w
	}
	reflectVal := reflect.ValueOf(val)
	if reflectVal.Kind() == reflect.Slice {
		vs := []any{}
		for i := 0; i < reflectVal.Len(); i++ {
			vs = append(vs, reflectVal.Index(i).Interface())
		}
		w.sess.In(key, vs...)
	} else {
		w.sess.In(key, val)
	}
	return w
}

func (w *where) Nin(key string, val any) store.Filter {
	if w.skip(val) {
		return w
	}
	reflectVal := reflect.ValueOf(val)
	if reflectVal.Kind() == reflect.Slice {
		vs := []any{}
		for i := 0; i < reflectVal.Len(); i++ {
			vs = append(vs, reflectVal.Index(i).Interface())
		}
		w.sess.NotIn(key, vs...)
	} else {
		w.sess.NotIn(key, val)
	}
	return w
}

func (w *where) Eq(key string, val any) store.Filter {
	if w.skip(val) {
		return w
	}
	w.sess.Where(key+" = ?", val)
	return w
}

func (w *where) Ne(key string, val any) store.Filter {
	if w.skip(val) {
		return w
	}
	w.sess.Where(key+" <> ?", val)
	return w
}

func (w *where) Gt(key string, val any) store.Filter {
	if w.skip(val) {
		return w
	}
	w.sess.Where(key+" > ?", val)
	return w
}

func (w *where) Gte(key string, val any) store.Filter {
	if w.skip(val) {
		return w
	}
	w.sess.Where(key+" >= ?", val)
	return w
}
func (w *where) Lt(key string, val any) store.Filter {
	if w.skip(val) {
		return w
	}
	w.sess.Where(key+" < ?", val)
	return w
}

func (w *where) Lte(key string, val any) store.Filter {
	if w.skip(val) {
		return w
	}
	w.sess.Where(key+" <= ?", val)
	return w
}

func (w *where) Regex(key string, val any) store.Filter {
	if w.skip(val) {
		return w
	}
	val = fmt.Sprintf("%%%v%%", val)
	w.sess.Where(key+" like ?", val)
	return w
}

func (w *where) Between(key string, vMin, vMax any) store.Filter {
	w.sess.Where(key+" between ? and ?", vMin, vMax)
	return w
}

func (w *where) Or(val any) store.Filter {
	if w.skip(val) {
		return w
	}
	w.sess.Or(val)
	return w
}

func (w *where) And(val any) store.Filter {
	if w.skip(val) {
		return w
	}
	w.sess.And(val)
	return w
}
func (w *where) skip(v any) bool {
	if v == nil {
		return true
	}
	vv := reflect.ValueOf(v)
	if vv.Kind() == reflect.Ptr {
		vv = vv.Elem()
	}
	if vv.IsZero() {
		return true
	}
	return false
}
