package sqlite

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"

	"github.com/icdb37/bfsm/internal/infra/store"
)

func TestSession(t *testing.T) {
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
	err = SessionProcess(ctx, &store.SessionStatement{
		Repo: col,
		Process: func(ctx context.Context, tab store.Tabler) error {
			return tab.Insert(ctx, info)
		},
	}, &store.SessionStatement{
		Repo: col,
		Process: func(ctx context.Context, tab store.Tabler) error {
			info.Age = 30
			info.Address = "BJ"
			return tab.Update(ctx, store.NewFilter().Eq("id", info.ID), info)
		},
	}, &store.SessionStatement{
		Repo: col,
		Process: func(ctx context.Context, tab store.Tabler) error {
			dbInfo := &demo{}
			if err := tab.Query(ctx, store.NewFilter().Eq("id", info.ID), dbInfo); err != nil {
				t.Fatal("query info failed", "error", err)
			}
			if dbInfo.Age != 30 || dbInfo.Address != "BJ" {
				t.Fatal("query info failed", "dbInfo", dbInfo)
			}
			return nil
		},
	}, &store.SessionStatement{
		Repo: col,
		Process: func(ctx context.Context, tab store.Tabler) error {
			return fmt.Errorf("rollback demo")
		},
	})
	if err != nil && err.Error() != "rollback demo" {
		t.Fatal("session process failed", "error", err)
	}
	dbInfo := &demo{}
	if err := col.Query(ctx, store.NewFilter().Eq("id", info.ID), dbInfo); err != nil {
		t.Fatal("query info failed", "error", err)
	}
	if dbInfo.Age != 0 || dbInfo.Address != "" { //出错回滚了
		t.Fatal("query info failed", "dbInfo", dbInfo)
	}
}
