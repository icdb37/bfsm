package sqlite

import (
	"github.com/icdb37/bfsm/internal/infra/logx"
	"github.com/icdb37/bfsm/internal/infra/store"
	"xorm.io/xorm"
)

type fUseFilter func(w *where)

type wfilter struct {
	fs []fUseFilter
}

func (w *wfilter) Eq(key string, val any) store.Filter {
	w.fs = append(w.fs, func(w *where) {
		w.Eq(key, val)
	})
	return w
}
func (w *wfilter) Ne(key string, val any) store.Filter {
	w.fs = append(w.fs, func(w *where) {
		w.Ne(key, val)
	})
	return w
}
func (w *wfilter) Gt(key string, val any) store.Filter {
	w.fs = append(w.fs, func(w *where) {
		w.Gt(key, val)
	})
	return w
}
func (w *wfilter) Gte(key string, val any) store.Filter {
	w.fs = append(w.fs, func(w *where) {
		w.Gte(key, val)
	})
	return w
}
func (w *wfilter) Lt(key string, val any) store.Filter {
	w.fs = append(w.fs, func(w *where) {
		w.Lt(key, val)
	})
	return w
}
func (w *wfilter) Lte(key string, val any) store.Filter {
	w.fs = append(w.fs, func(w *where) {
		w.Lte(key, val)
	})
	return w
}
func (w *wfilter) Between(key string, vMin, vMax any) store.Filter {
	w.fs = append(w.fs, func(w *where) {
		w.Between(key, vMin, vMax)
	})
	return w
}
func (w *wfilter) In(key string, val any) store.Filter {
	w.fs = append(w.fs, func(w *where) {
		w.In(key, val)
	})
	return w
}
func (w *wfilter) Nin(key string, val any) store.Filter {
	w.fs = append(w.fs, func(w *where) {
		w.Nin(key, val)
	})
	return w
}
func (w *wfilter) Regex(key string, val any) store.Filter {
	w.fs = append(w.fs, func(w *where) {
		w.Regex(key, val)
	})
	return w
}

func (w *wfilter) Or(val any) store.Filter {
	w.fs = append(w.fs, func(w *where) {
		w.Or(val)
	})
	return w
}
func (w *wfilter) And(val any) store.Filter {
	w.fs = append(w.fs, func(w *where) {
		w.And(val)
	})
	return w
}

// NewFilter 创建筛选器
func NewFilter() store.Filter {
	return &wfilter{}
}

// UseWhere 使用筛选器
func UseWhere(s *xorm.Session, f store.Filter) store.Filter {
	wf, ok := f.(*wfilter)
	if !ok {
		logx.Error("filter is not wfilter")
		return f
	}
	w := &where{sess: s}
	for _, uf := range wf.fs {
		uf(w)
	}
	return f
}
