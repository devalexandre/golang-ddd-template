package mysql_test

import (
	"context"
	"testing"

	"devalexandre/golang-ddd-template/internal/infra/database/mysql/mocks"

	mysql "devalexandre/golang-ddd-template/internal/infra/database/mysql"

	"github.com/testcontainers/testcontainers-go"
	msql "github.com/testcontainers/testcontainers-go/modules/mysql"
	"github.com/vingarcia/ksql"
)

type UserTest struct {
	Name string `ksql:"name"`
}

func TestDatabase(t *testing.T) {

	t.Run("Shoud be able to connect to database", func(t *testing.T) {
		ctx := context.Background()

		mysqlContainer, err := msql.RunContainer(ctx,
			testcontainers.WithImage("mysql"),
			// msql.WithConfigFile(filepath.Join("testdata", "my_8.cnf")),
			msql.WithDatabase("foo"),
			msql.WithUsername("root"),
			msql.WithPassword("password"),
			// msql.WithScripts(filepath.Join("testdata", "schema.sql")),
		)
		if err != nil {
			panic(err)
		}

		t.Cleanup(func() {
			if err := mysqlContainer.Terminate(ctx); err != nil {
				t.Fatalf("failed to terminate mysqlContainer: %s", err)
			}
		})

		connStr, err := mysqlContainer.ConnectionString(ctx)

		if err != nil {
			t.Errorf("Error getting connection string: %s", err)
		}

		db, err := mysql.NewDatabase(context.Background(), connStr, ksql.Config{})

		if err != nil {
			t.Errorf("Error connecting to database: %s", err)
		}

		if db == nil {
			t.Errorf("Database is nil")
		}

	})

	t.Run("Shoud be able to insert data", func(t *testing.T) {

		mockDB := mocks.MockDb{
			InsertFunc: func(ctx context.Context, table ksql.Table, data interface{}) error {
				return nil
			},
		}

		db := &mysql.Database{
			DB: mockDB,
		}

		err := db.Insert(context.Background(), "users", map[string]interface{}{
			"name": "Alexandre",
		})

		if err != nil {
			t.Errorf("Error inserting data: %s", err)
		}
	})

	t.Run("Shoud be able to update data", func(t *testing.T) {

		mockDB := mocks.MockDb{
			PatchFunc: func(ctx context.Context, table ksql.Table, data interface{}) error {
				return nil
			},
		}

		db := &mysql.Database{
			DB: mockDB,
		}

		err := db.Update(context.Background(), "users", map[string]interface{}{
			"name": "Alexandre",
		})

		if err != nil {
			t.Errorf("Error updating data: %s", err)
		}
	})

	t.Run("Shoud be able to delete data", func(t *testing.T) {

		mockDB := mocks.MockDb{
			DeleteFunc: func(ctx context.Context, table ksql.Table, idOrRecord interface{}) error {
				return nil
			},
		}

		db := &mysql.Database{
			DB: mockDB,
		}

		err := db.Delete(context.Background(), "users", map[string]interface{}{
			"name": "Alexandre",
		})

		if err != nil {
			t.Errorf("Error deleting data: %s", err)
		}
	})

	t.Run("Shoud be able to query data", func(t *testing.T) {

		mockDB := mocks.MockDb{
			QueryFunc: func(ctx context.Context, record interface{}, query string, params ...interface{}) error {
				return nil
			},
		}

		db := &mysql.Database{
			DB: mockDB,
		}

		var user UserTest

		err := db.Query(context.Background(), user, "SELECT * FROM users WHERE name = $1", "Alexandre", 1)

		if err != nil {
			t.Errorf("Error querying data: %s", err)
		}
	})

}
