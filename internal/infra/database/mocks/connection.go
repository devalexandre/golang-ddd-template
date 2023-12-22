package mocks

import "context"

type MockDatabase struct {
	SelectFunc func(ctx context.Context, record interface{}, query string, params ...interface{}) error
	InsertFunc func(ctx context.Context, table string, data interface{}) error
	UpdateFunc func(ctx context.Context, table string, record interface{}) error
	DeleteFunc func(ctx context.Context, table string, record interface{}) error
	QueryFunc  func(ctx context.Context, record interface{}, query string, params ...interface{}) error
	CloseFunc  func() error
}

func (m MockDatabase) Select(ctx context.Context, record interface{}, query string, params ...interface{}) error {
	return m.SelectFunc(ctx, record, query, params...)
}

func (m MockDatabase) Insert(ctx context.Context, table string, data interface{}) error {
	return m.InsertFunc(ctx, table, data)
}

func (m MockDatabase) Update(ctx context.Context, table string, record interface{}) error {
	return m.UpdateFunc(ctx, table, record)
}

func (m MockDatabase) Delete(ctx context.Context, table string, record interface{}) error {
	return m.DeleteFunc(ctx, table, record)
}

func (m MockDatabase) Query(ctx context.Context, record interface{}, query string, params ...interface{}) error {
	return m.QueryFunc(ctx, record, query, params...)
}

func (m MockDatabase) Close() error {
	return m.CloseFunc()
}
