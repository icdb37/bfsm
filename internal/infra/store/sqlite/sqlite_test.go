package sqlite

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/icdb37/bfsm/internal/infra/config"
	"github.com/icdb37/bfsm/internal/infra/store"
)

var cfg = &store.Config{
	Name:  "test",
	Debug: true,
	Path:  "./testdata",
}

var count int

func setup(t *testing.T) {
	cfg.Name += fmt.Sprintf("%d", count)
	count++
	jsonCfg, err := json.Marshal(cfg)
	if err != nil {
		t.Fatal("init sqlite config failed", err)
	}
	config.SetConfig(config.KeyDatabase, string(jsonCfg))
}
func teardown(t *testing.T) {
	Close()
	if err := os.RemoveAll(cfg.Path); err != nil {
		t.Fatal("teardown sqlite failed", err)
	}
}

type demo struct {
	Xid       uint32    `xorm:"pk autoincr 'xid' comment('自增主键')"`
	ID        string    `xorm:"varchar(36) not null unique 'id' comment('唯一标识')"`
	Name      string    `xorm:"varchar(20) not null unique 'name' comment('用户名称')"`
	Age       uint8     `xorm:"tinyint(3) not null default 18 'age' comment('用户年龄')"`
	Address   string    `xorm:"varchar(255) not null default '' 'address' comment('用户地址')"`
	CreatedAt time.Time `xorm:"created 'created_at' comment('创建时间')"`
}

func (d *demo) TableName() string {
	return "demo"
}

func TestSqlite(t *testing.T) {
	teardown(t)
	setup(t)
	defer teardown(t)
	if err := Init(); err != nil {
		t.Fatal(err)
	}
	info := &demo{
		ID:      uuid.NewString(),
		Name:    "test",
		Age:     20,
		Address: "SH",
	}
	col, err := New(info)
	if err != nil {
		t.Fatal(err)
	}
	if col == nil {
		t.Fatal("new sqlite failed")
	}
	ctx := context.Background()
	if err := col.Insert(ctx, info); err != nil {
		t.Fatal("insert info failed", "error", err)
	}
	infos := []*demo{}
	pf := &store.PageFilter{}
	qf := store.NewFilter()
	if _, err := col.Search(ctx, qf, pf, &infos); err != nil {
		t.Fatal("query infos failed", "error", err)
	}
	if len(infos) != 1 || infos[0].Name != info.Name {
		t.Fatal("query infos failed", "infos", infos)
	}
	infos = []*demo{}
	qf.Eq("id", info.ID).Regex("name", "te")
	if _, err := col.Search(ctx, qf, pf, &infos); err != nil {
		t.Fatal("query infos failed", "error", err)
	}
	if len(infos) != 1 || infos[0].Name != info.Name {
		t.Fatal("query infos failed", "infos", infos)
	}
	if err := col.Delete(ctx, qf); err != nil {
		t.Fatal("delete info failed", "error", err)
	}
	qf.Or(map[string]interface{}{"age": []int{20, 30}, "address": "BJ"})
	infos = []*demo{}
	if _, err := col.Search(ctx, qf, pf, &infos); err != nil {
		t.Fatal("query infos failed", "error", err)
	}
	if len(infos) != 0 {
		t.Fatal("query infos failed", "infos", infos)
	}
}

func TestSqliteQuery(t *testing.T) {
	teardown(t)
	setup(t)
	defer teardown(t)
	if err := Init(); err != nil {
		t.Fatal(err)
	}
	info := &demo{
		ID:      uuid.NewString(),
		Name:    "test",
		Age:     20,
		Address: "SH",
	}
	col, err := New(info)
	if err != nil {
		t.Fatal(err)
	}
	if col == nil {
		t.Fatal("new sqlite failed")
	}
	ctx := context.Background()
	if err := col.Insert(ctx, info); err != nil {
		t.Fatal("insert info failed", "error", err)
	}
	data := demo{}
	qf := store.NewFilter()
	if err := col.Query(ctx, qf, &data); err != nil {
		t.Fatal("query infos failed", "error", err)
	}
	if data.Name != info.Name {
		t.Fatal("query infos failed", "data", data)
	}
}
