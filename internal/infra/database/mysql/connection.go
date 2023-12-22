package mysql

import (
	"context"
	"devalexandre/golang-ddd-template/internal/infra/database"

	"github.com/vingarcia/ksql"
	"github.com/vingarcia/ksql/adapters/kmysql"
)

type Database struct {
	DB database.IDB
}

func NewDatabase(ctx context.Context, connectionString string, config ksql.Config) (database.IDatabase, error) {
	conn, errDb := kmysql.New(ctx, connectionString, config)

	if errDb != nil {
		return nil, errDb
	}

	return &Database{
		DB: conn,
	}, nil
}

func (c *Database) Select(ctx context.Context, record interface{}, query string, params ...interface{}) error {
	return c.DB.QueryOne(ctx, record, query, params...)
}

func (c *Database) Insert(ctx context.Context, table string, data interface{}) error {
	tb := ksql.NewTable(table)
	return c.DB.Insert(ctx, tb, data)
}

func (c *Database) Update(ctx context.Context, table string, record interface{}) error {
	tb := ksql.NewTable(table)
	return c.DB.Patch(ctx, tb, record)
}

func (c *Database) Delete(ctx context.Context, table string, record interface{}) error {
	tb := ksql.NewTable(table)
	return c.DB.Delete(ctx, tb, record)
}

func (c *Database) Query(ctx context.Context, record interface{}, query string, params ...interface{}) error {
	return c.DB.Query(ctx, record, query, params...)
}

func (c *Database) Close() error {
	return c.DB.Close()
}
