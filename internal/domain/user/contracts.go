package user

import (
	"context"
	"devalexandre/golang-ddd-template/internal/infra/database"
)

type IRepository interface {
	Insert(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, user *User) error
	Find(ctx context.Context, id string) (*User, error)
	FindAll(ctx context.Context) ([]*User, error)
}

type Repository struct {
	DB database.IDatabase
}

type User struct {
	ID   string
	Name string
}
