package sqlite_test

import (
	"context"
	"devalexandre/golang-ddd-template/internal/infra/database/sqlite"
	"devalexandre/golang-ddd-template/internal/infra/database/mocks"
	"os"
	"testing"

	"github.com/vingarcia/ksql"
)

type UserTest struct {
	Name string `ksql:"name"`
}

func TestDatabase(t *testing.T) {

	t.Run("Shoud be able to connect to database", func(t *testing.T) {

		bdUrl := "testdata"
		db, err := sqlite.NewDatabase(context.Background(), bdUrl, ksql.Config{})

		if err != nil {
			t.Errorf("Error connecting to database: %s", err)
		}

		if db == nil {
			t.Errorf("Database is nil")
		}

		// remove database file
		defer func() {
			err = db.Close()
			if err != nil {
				t.Errorf("Error closing database: %s", err)
			}
			os.Remove(bdUrl)
		}()

	})

	t.Run("Shoud be able to insert data", func(t *testing.T) {

		mockDB := mocks.MockDb{
			InsertFunc: func(ctx context.Context, table ksql.Table, data interface{}) error {
				return nil
			},
		}

		db := &sqlite.Database{
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

		db := &sqlite.Database{
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

		db := &sqlite.Database{
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

		db := &sqlite.Database{
			DB: mockDB,
		}

		var user UserTest

		err := db.Query(context.Background(), user, "SELECT * FROM users WHERE name = $1", "Alexandre", 1)

		if err != nil {
			t.Errorf("Error querying data: %s", err)
		}
	})

}
