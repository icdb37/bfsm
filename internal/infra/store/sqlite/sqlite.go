// Package sqlite sqlite 存储
package sqlite

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	_ "modernc.org/sqlite"
	"xorm.io/xorm"

	"github.com/icdb37/bfsm/internal/infra/config"
	"github.com/icdb37/bfsm/internal/infra/logx"
	"github.com/icdb37/bfsm/internal/infra/store"
)

/*
  xorm.Session 在执行操作语句之后就会被清理
*/

var gDbe *xorm.Engine

func Init() error {
	if gDbe != nil {
		return nil
	}
	strCfg := config.GetDatabase()
	cfg := &store.Config{}
	if err := json.Unmarshal([]byte(strCfg), cfg); err != nil {
		logx.Warn("parse json sqlite config invalid", "config", strCfg)
		cfg.Name = "ymzy"
		cfg.Debug = true
		cfg.Path = filepath.Join("bfsm", "data")
	} else {
		if pos := strings.LastIndex(cfg.Name, "."); pos != -1 {
			cfg.Name = cfg.Name[:pos]
		}
	}
	filename := filepath.Join(cfg.Path, cfg.Name+".db")
	os.MkdirAll(filepath.Dir(filename), os.ModePerm) //nolint
	var err error
	gDbe, err = xorm.NewEngine("sqlite", filename)
	if err != nil {
		return err
	}
	gDbe.ShowSQL(cfg.Debug)
	gDbe.Charset("utf8mb4")
	store.Init(NewFilter, Unmarshal, New, SessionProcess)
	return nil
}

// Close 关闭数据库
func Close() error {
	if gDbe == nil {
		return nil
	}
	err := gDbe.Close()
	gDbe = nil
	return err
}

type sqlite struct {
	db    *xorm.Engine
	table string
}

// NewSqliteStore 创建 sqlite 存储
func New(v store.TableNamer) (store.Tabler, error) {
	s := &sqlite{db: gDbe, table: v.TableName()}
	if err := s.CreateTable(context.Background(), v); err != nil {
		return nil, err
	}
	return s, nil
}

// TableName 获取表名
func (s *sqlite) TableName() string {
	return s.table
}

// Total 统计数量
func (s *sqlite) Total(ctx context.Context, f store.Filter) (int64, error) {
	sess := s.db.Table(s.table)
	UseWhere(sess, f)
	var total int64
	total, err := sess.Count()
	if err != nil {
		return total, err
	}
	return total, nil
}

// Search 查询数据
func (s *sqlite) Search(ctx context.Context, f store.Filter, p store.Pager, v any) (int64, error) {
	sess := s.db.Table(s.table)
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
func (s *sqlite) Query(ctx context.Context, f store.Filter, v any) (err error) {
	sess := s.db.Table(s.table)
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
func (s *sqlite) Insert(ctx context.Context, vs ...any) error {
	if len(vs) == 0 {
		return nil
	}
	_, err := s.db.Table(s.table).InsertMulti(&vs)
	return err
}

// Upsert 更新数据
func (s *sqlite) Upsert(ctx context.Context, f store.Filter, v any) error {
	if v == nil {
		return nil
	}
	sess := s.db.Table(s.table)
	UseWhere(sess, f)
	count, err := sess.Update(v)
	if err != nil {
		return err
	}
	if count == 1 {
		return nil
	}
	sess = s.db.Table(s.table)
	UseWhere(sess, f)
	if _, err = sess.Insert(v); err != nil {
		return err
	}
	return err
}

// Update 修改数据
func (s *sqlite) Update(ctx context.Context, f store.Filter, v any) error {
	if v == nil {
		return nil
	}
	sess := s.db.Table(s.table)
	UseWhere(sess, f)
	_, err := sess.Update(v)
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlite) NewFilter() store.Filter {
	return &where{sess: s.db.Table(s.table)}
}

// Delete 删除数据
func (s *sqlite) Delete(ctx context.Context, f store.Filter) error {
	sess := s.db.Table(s.table)
	UseWhere(sess, f)
	if _, err := sess.Delete(); err != nil {
		return err
	}
	return nil
}

// CreateIndex 创建索引
func (s *sqlite) CreateIndex(ctx context.Context, keys ...string) error {
	return nil
}

// CreateTable 	创建表
func (s *sqlite) CreateTable(ctx context.Context, v any) error {
	if err := s.db.Sync2(v); err != nil {
		return err
	}
	return nil
}

func (s *sqlite) withPage(sess *xorm.Session, p store.Pager) {
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
