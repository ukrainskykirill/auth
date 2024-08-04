package user

import (
	"context"
	"fmt"
	"github.com/fatih/color"
	pool "github.com/jackc/pgx/v5/pgxpool"
	"github.com/ukrainskykirill/auth/internal/model"
	"github.com/ukrainskykirill/auth/internal/repository"
	"github.com/ukrainskykirill/auth/internal/repository/user/converter"
	modelRepo "github.com/ukrainskykirill/auth/internal/repository/user/model"
	"time"
)

type repo struct {
	db *pool.Pool
}

func NewUserRepository(db *pool.Pool) repository.UserRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, user *model.UserIn) (int64, error) {
	var userID int64
	err := r.db.QueryRow(
		ctx,
		`INSERT INTO users (name, email, role, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;`,
		user.Name, user.Email, user.Role, user.Password, time.Now(), time.Now(),
	).Scan(&userID)
	if err != nil {
		return 0, err
	}
	fmt.Println(color.BlueString("Create user: name - %+v, with ctx: %v", user, ctx))
	return userID, nil
}

func (r *repo) Delete(ctx context.Context, userID int64) error {
	_, err := r.db.Exec(
		ctx,
		`DELETE FROM users WHERE id = $1;`,
		userID,
	)
	if err != nil {
		return err
	}

	fmt.Println(color.RedString("Delete user: id %d, with ctx: %v", userID, ctx))
	return nil
}

func (r *repo) Update(ctx context.Context, user *model.UserInUpdate) error {
	var sql string
	var args []interface{}
	paramIndex := 1

	sql = `UPDATE users SET `

	if len(user.Name) != 0 {
		sql += fmt.Sprintf("name = $%d", paramIndex)
		args = append(args, user.Name)
		paramIndex++
	}

	if len(user.Email) != 0 {
		if len(args) > 0 {
			sql += ", "
		}
		sql += fmt.Sprintf("email = $%d", paramIndex)
		args = append(args, user.Email)
		paramIndex++
	}

	if len(user.Role) != 0 {
		if len(args) > 0 {
			sql += ", "
		}
		sql += fmt.Sprintf("role = $%d", paramIndex)
		args = append(args, user.Role)
		paramIndex++
	}

	if len(args) > 0 {
		sql += ", updated_at = $" + fmt.Sprintf("%d", paramIndex)
		args = append(args, time.Now())
		paramIndex++
	}

	sql += " WHERE id = $" + fmt.Sprintf("%d;", paramIndex)
	args = append(args, user.ID)

	_, err := r.db.Exec(
		ctx,
		sql,
		args...,
	)
	if err != nil {
		return err
	}

	fmt.Println(color.WhiteString("Update user: id %+v, with ctx: %v", user, ctx))
	return nil
}

func (r *repo) Get(ctx context.Context, userID int64) (*model.User, error) {
	var userRow modelRepo.RepoUser
	err := r.db.QueryRow(
		ctx,
		`SELECT name, email, role, created_at, updated_at FROM users WHERE id = $1;`,
		userID,
	).Scan(&userRow.Name, &userRow.Email, &userRow.Role, &userRow.CreatedAt, &userRow.UpdatedAt)
	if err != nil {
		return &model.User{}, err
	}

	userRow.ID = userID
	fmt.Println(color.BlackString("Get user: id %d, with ctx: %v", userID, ctx))
	return converter.ToUserFromRepo(userRow), nil
}
