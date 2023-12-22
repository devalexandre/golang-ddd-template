package user

import (
	"context"
	"devalexandre/golang-ddd-template/internal/infra/database"
)

func NewRepository(db database.IDatabase) IRepository {
	return &Repository{
		DB: db,
	}
}

func (r *Repository) Insert(ctx context.Context, user *User) error {
	return r.DB.Insert(ctx, "users", user)
}

func (r *Repository) Update(ctx context.Context, user *User) error {
	return r.DB.Update(ctx, "users", user)
}

func (r *Repository) Delete(ctx context.Context, user *User) error {
	return r.DB.Delete(ctx, "users", user)
}

func (r *Repository) Find(ctx context.Context, id string) (*User, error) {
	user := &User{}
	err := r.DB.Select(ctx, user, "SELECT * FROM users WHERE id = $1 LIMIT 1", id)
	return user, err
}

func (r *Repository) FindAll(ctx context.Context) ([]*User, error) {
	var users []*User
	err := r.DB.Query(ctx, &users, "SELECT * FROM users")
	return users, err
}
