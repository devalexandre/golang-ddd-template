package user_test

import (
	"context"
	"devalexandre/golang-ddd-template/internal/domain/user"
	"devalexandre/golang-ddd-template/internal/infra/database/mocks"
	"testing"
)

func TestRepository(t *testing.T) {

	// Create a new user
	t.Run("Create a new user", func(t *testing.T) {
		db := mocks.MockDatabase{
			InsertFunc: func(ctx context.Context, table string, record interface{}) error {
				return nil
			},
		}

		repository := user.Repository{
			DB: db,
		}

		user := user.User{
			ID:   "1",
			Name: "Alexandre",
		}

		err := repository.Insert(nil, &user)

		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})

	// Update a user
	t.Run("Update a user", func(t *testing.T) {
		db := mocks.MockDatabase{
			UpdateFunc: func(ctx context.Context, table string, record interface{}) error {
				return nil
			},
		}

		repository := user.Repository{
			DB: db,
		}

		user := user.User{
			ID:   "1",
			Name: "Alexandre",
		}

		err := repository.Update(nil, &user)

		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})

	// Delete a user
	t.Run("Delete a user", func(t *testing.T) {
		db := mocks.MockDatabase{
			DeleteFunc: func(ctx context.Context, table string, record interface{}) error {
				return nil
			},
		}

		repository := user.Repository{
			DB: db,
		}

		user := user.User{
			ID:   "1",
			Name: "Alexandre",
		}

		err := repository.Delete(nil, &user)

		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})

	// Find a user
	t.Run("Find a user", func(t *testing.T) {
		db := mocks.MockDatabase{
			SelectFunc: func(ctx context.Context, record interface{}, query string, params ...interface{}) error {
				return nil
			},
		}

		repository := user.Repository{
			DB: db,
		}

		_, err := repository.Find(nil, "1")

		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})

	// Find all users
	t.Run("Find all users", func(t *testing.T) {
		db := mocks.MockDatabase{
			QueryFunc: func(ctx context.Context, records interface{}, query string, params ...interface{}) error {
				return nil
			},
		}

		repository := user.Repository{
			DB: db,
		}

		_, err := repository.FindAll(nil)

		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})

}
