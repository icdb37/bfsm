package sqlite

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/icdb37/bfsm/internal/infra/store"
)

type rangeX[T any] struct {
	Beg T `where:"gte,-,omitempty"`
	End T `where:"lte,-,omitempty"`
}

func newRangeX[T any](beg T, end T) *rangeX[T] {
	return &rangeX[T]{
		Beg: beg,
		End: end,
	}
}

type demoQuery struct {
	Code      string             `where:"eq,code,omitempty"`
	Status    []string           `where:"in,status,omitempty"`
	Name      string             `where:"regex,name,omitempty"`
	Age       *rangeX[int]       `where:"or,age,omitempty"`
	CreatedAt *rangeX[time.Time] `where:"range,created_at,omitempty"`
}

func TestParseFieldTag(t *testing.T) {
	if parseTagField("Code", "") != "code" {
		t.Fatal("parseFieldTag 'Code' failed")
	}
	if parseTagField("ID", "") != "id" {
		t.Fatal("parseFieldTag 'ID' failed")
	}
	if parseTagField("CreatedAt", "") != "created_at" {
		t.Fatal("parseFieldTag 'CreatedAt' failed")
	}
}

func TestUnmarshal(t *testing.T) {
	createdBeg := time.Time{}.AddDate(2000, 1, 1)
	cratedEnd := time.Time{}.AddDate(2050, 1, 1)
	q := &demoQuery{
		Name:      "t",
		Age:       newRangeX(10, 30),
		CreatedAt: newRangeX(createdBeg, cratedEnd),
	}
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
	qf = Unmarshal(q)
	infos = []*demo{}
	if _, err := col.Search(ctx, qf, pf, &infos); err != nil {
		t.Fatal("query infos failed", "error", err)
	}
	if len(infos) != 1 || infos[0].Name != info.Name {
		t.Fatal("query infos failed", "infos", infos)
	}
}
