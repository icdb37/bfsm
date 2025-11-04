package sqlite

import (
	"context"
	"reflect"
	"strings"

	"xorm.io/xorm"

	"github.com/icdb37/bfsm/internal/infra/store"
)

// 事务

type session struct {
	sess  *xorm.Session
	table string
}

// TableName 获取表名
func (s *session) TableName() string {
	return s.table
}

// Total 统计数量
func (s *session) Total(ctx context.Context, f store.Filter) (int64, error) {
	sess := s.sess.Table(s.table)
	UseWhere(sess, f)
	var total int64
	total, err := sess.Count()
	if err != nil {
		return total, err
	}
	return total, nil
}

// Search 查询数据
func (s *session) Search(ctx context.Context, f store.Filter, p store.Pager, v any) (int64, error) {
	sess := s.sess.Table(s.table)
	if f != nil {
		UseWhere(sess, f)
	}
	if p != nil {
		s.withPage(sess, p)
	}
	total, err := sess.FindAndCount(v)
	if err != nil {
		return 0, err
	}
	return total, nil
}

// Query 查询数据
func (s *session) Query(ctx context.Context, f store.Filter, v any) (err error) {
	sess := s.sess.Table(s.table)
	if wg, ok := f.(store.Wgetter); ok {
		sess.Where(wg.GetxWhere())
	} else if f != nil {
		UseWhere(sess, f)
	}
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

// Insert 插入数据
func (s *session) Insert(ctx context.Context, vs ...any) error {
	if len(vs) == 0 {
		return nil
	}
	for _, v := range vs {
		if _, err := s.sess.Table(s.table).Insert(v); err != nil {
			return err
		}
	}
	return nil
}

// Upsert 更新数据
func (s *session) Upsert(ctx context.Context, f store.Filter, v any) error {
	if v == nil {
		return nil
	}
	sess := s.sess.Table(s.table)
	UseWhere(sess, f)
	count, err := sess.Update(v)
	if err != nil {
		return err
	}
	if count == 1 {
		return nil
	}
	sess = s.sess.Table(s.table)
	UseWhere(sess, f)
	if _, err = sess.Insert(v); err != nil {
		return err
	}
	return err
}

// Update 修改数据
func (s *session) Update(ctx context.Context, f store.Filter, v any) error {
	if v == nil {
		return nil
	}
	sess := s.sess.Table(s.table)
	UseWhere(sess, f)
	_, err := sess.Update(v)
	if err != nil {
		return err
	}
	return nil
}

func (s *session) NewFilter() store.Filter {
	return &where{sess: s.sess.Table(s.table)}
}

// Delete 删除数据
func (s *session) Delete(ctx context.Context, f store.Filter) error {
	sess := s.sess.Table(s.table)
	UseWhere(sess, f)
	if _, err := sess.Delete(); err != nil {
		return err
	}
	return nil
}

func (s *session) withPage(sess *xorm.Session, p store.Pager) {
	for _, sort := range p.GetSorts() {
		if strings.HasPrefix(sort, "-") {
			sess.Desc(sort[1:])
		} else if strings.HasPrefix(sort, "+") {
			sess.Asc(sort[1:])
		} else {
			sess.Asc(sort)
		}
	}
	page, size := p.GetPageIndex(), p.GetPageSize()
	if size == 0 {
		size = 1
	}
	sess.Limit(size, page*size)
}

// SessionProcess 事务
func SessionProcess(ctx context.Context, fs ...*store.SessionStatement) error {
	_, err := gDbe.Transaction(func(sess *xorm.Session) (interface{}, error) {
		for _, f := range fs {
			if err := f.Process(ctx, &session{sess: sess, table: f.Repo.TableName()}); err != nil {
				sess.Rollback() //nolint
				return nil, err
			}
		}
		sess.Commit() //nolint
		return nil, nil
	})
	return err
}
