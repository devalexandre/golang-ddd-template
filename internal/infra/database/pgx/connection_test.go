package pgx_test

import (
	"context"
	"testing"
	"time"

	"devalexandre/golang-ddd-template/internal/infra/database/mocks"
	"devalexandre/golang-ddd-template/internal/infra/database/pgx"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"github.com/vingarcia/ksql"
)

type UserTest struct {
	Name string `ksql:"name"`
}

func TestDatabase(t *testing.T) {

	t.Run("Shoud be able to connect to database", func(t *testing.T) {
		ctx := context.Background()
		pgContainer, err := postgres.RunContainer(ctx,
			testcontainers.WithImage("postgres:15.3-alpine"),
			// postgres.WithInitScripts(filepath.Join(".", "testdata", "init-db.sql")),
			postgres.WithDatabase("test-db"),
			postgres.WithUsername("postgres"),
			postgres.WithPassword("postgres"),
			testcontainers.WithWaitStrategy(
				wait.ForLog("database system is ready to accept connections").
					WithOccurrence(2).WithStartupTimeout(5*time.Second)),
		)
		if err != nil {
			t.Fatal(err)
		}

		t.Cleanup(func() {
			if err := pgContainer.Terminate(ctx); err != nil {
				t.Fatalf("failed to terminate pgContainer: %s", err)
			}
		})

		connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")

		if err != nil {
			t.Errorf("Error getting connection string: %s", err)
		}

		db, err := pgx.NewDatabase(context.Background(), connStr, ksql.Config{})

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

		db := &pgx.Database{
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

		db := &pgx.Database{
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

		db := &pgx.Database{
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

		db := &pgx.Database{
			DB: mockDB,
		}

		var user UserTest

		err := db.Query(context.Background(), user, "SELECT * FROM users WHERE name = $1", "Alexandre", 1)

		if err != nil {
			t.Errorf("Error querying data: %s", err)
		}
	})

}
