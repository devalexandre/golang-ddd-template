package mocks

import (
	"context"

	"github.com/vingarcia/ksql"
)

type MockDb struct {
	QueryFunc    func(ctx context.Context, record interface{}, query string, params ...interface{}) error
	QueryOneFunc func(ctx context.Context, record interface{}, query string, params ...interface{}) error
	InsertFunc   func(ctx context.Context, table ksql.Table, record interface{}) error
	PatchFunc    func(ctx context.Context, table ksql.Table, record interface{}) error
	DeleteFunc   func(ctx context.Context, table ksql.Table, idOrRecord interface{}) error
	CloseFunc    func() error
}

func (m MockDb) Query(ctx context.Context, record interface{}, query string, params ...interface{}) error {
	return m.QueryFunc(ctx, record, query, params...)
}

func (m MockDb) QueryOne(ctx context.Context, record interface{}, query string, params ...interface{}) error {
	return m.QueryOneFunc(ctx, record, query, params...)
}

func (m MockDb) Insert(ctx context.Context, table ksql.Table, record interface{}) error {
	return m.InsertFunc(ctx, table, record)
}

func (m MockDb) Patch(ctx context.Context, table ksql.Table, record interface{}) error {
	return m.PatchFunc(ctx, table, record)
}

func (m MockDb) Delete(ctx context.Context, table ksql.Table, idOrRecord interface{}) error {
	return m.DeleteFunc(ctx, table, idOrRecord)
}

func (m MockDb) Close() error {
	return m.CloseFunc()
}
