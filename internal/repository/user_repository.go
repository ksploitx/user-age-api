package repository

import (
	"context"
	"database/sql"
	"time"

	db "github.com/ksploitx/user-age-api/db/sqlc"
)

type UserRepository interface {
	Create(ctx context.Context, name string, dob time.Time) (db.User, error)
	GetByID(ctx context.Context, id int32) (db.User, error)
	Update(ctx context.Context, id int32, name string, dob time.Time) (db.User, error)
	Delete(ctx context.Context, id int32) error
	List(ctx context.Context) ([]db.User, error)
}

type userRepository struct {
	queries *db.Queries
}

func NewUserRepository(database *sql.DB) UserRepository {
	return &userRepository{queries: db.New(database)}
}

func (r *userRepository) Create(ctx context.Context, name string, dob time.Time) (db.User, error) {
	return r.queries.CreateUser(ctx, db.CreateUserParams{
		Name: name,
		Dob:  dob,
	})
}

func (r *userRepository) GetByID(ctx context.Context, id int32) (db.User, error) {
	return r.queries.GetUserByID(ctx, id)
}

func (r *userRepository) Update(ctx context.Context, id int32, name string, dob time.Time) (db.User, error) {
	return r.queries.UpdateUser(ctx, db.UpdateUserParams{
		ID:   id,
		Name: name,
		Dob:  dob,
	})
}

func (r *userRepository) Delete(ctx context.Context, id int32) error {
	return r.queries.DeleteUser(ctx, id)
}

func (r *userRepository) List(ctx context.Context) ([]db.User, error) {
	return r.queries.ListUsers(ctx)
}
