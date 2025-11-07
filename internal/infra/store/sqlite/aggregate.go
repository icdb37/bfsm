package sqlite

import (
	"context"
	"reflect"

	"github.com/icdb37/bfsm/internal/infra/store"
	"xorm.io/xorm"
)

type aggregate struct {
	db      *xorm.Engine
	sqlStmt string
	sqlArgs []any
}

// Aggregate 聚合数据
func (a *aggregate) Aggregate(ctx context.Context, f store.Filter, v any) (err error) {
	sess := a.db.SQL(a.sqlStmt, a.sqlArgs...)
	UseWhere(sess, f)
	vv := reflect.ValueOf(v)
	vv = vv.Elem()
	if vv.Kind() == reflect.Slice {
		if err = sess.Find(v); err != nil {
			return err
		}
		return nil
	}
	if _, err = sess.Get(v); err != nil {
		return err
	}
	return nil
}
