package pgx

import (
	"context"

	"github.com/vingarcia/ksql"
)

type IDatabase interface {
	Select(ctx context.Context, record interface{}, query string, params ...interface{}) error
	Insert(ctx context.Context, table string, data interface{}) error
	Update(ctx context.Context, table string, record interface{}) error
	Delete(ctx context.Context, table string, record interface{}) error
	Query(ctx context.Context, record interface{}, query string, params ...interface{}) error
	Close() error
}

type IDB interface {
	Query(ctx context.Context, record interface{}, query string, params ...interface{}) error
	QueryOne(ctx context.Context, record interface{}, query string, params ...interface{}) error
	Insert(ctx context.Context, table ksql.Table, record interface{}) error
	Patch(ctx context.Context, table ksql.Table, record interface{}) error
	Delete(ctx context.Context, table ksql.Table, idOrRecord interface{}) error
	Close() error
}
