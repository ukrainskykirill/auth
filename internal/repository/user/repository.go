package user

import (
	"context"
	"fmt"
	"github.com/fatih/color"
	"github.com/ukrainskykirill/auth/internal/client/db"
	"github.com/ukrainskykirill/auth/internal/model"
	"github.com/ukrainskykirill/auth/internal/repository"
	"github.com/ukrainskykirill/auth/internal/repository/user/converter"
	modelRepo "github.com/ukrainskykirill/auth/internal/repository/user/model"
	"time"
)

const (
	userRepo     = "user_repository"
	createRepoFn = userRepo + "." + "Create"
	deleteRepoFn = userRepo + "." + "Delete"
	updateRepoFn = userRepo + "." + "Update"
	getRepoFn    = userRepo + "." + "Get"
	isExistByID  = userRepo + "." + "IsExistByID"
)

type repo struct {
	db db.Client
}

func NewUserRepository(db db.Client) repository.UserRepository {
	return &repo{db: db}
}

func (r *repo) IsExistByID(ctx context.Context, userID int64) (bool, error) {
	var isExist bool

	sql := `SELECT EXISTS(SELECT 1 FROM users WHERE id = $1);`

	q := db.Query{
		Name:     isExistByID,
		QueryRaw: sql,
	}

	err := r.db.DB().QueryRowContext(
		ctx,
		q,
		userID,
	).Scan(&isExist)
	if err != nil {
		return false, err
	}

	return isExist, nil
}

func (r *repo) Create(ctx context.Context, user *model.UserIn) (int64, error) {
	var userID int64

	q := db.Query{
		Name:     createRepoFn,
		QueryRaw: `INSERT INTO users (name, email, role, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;`,
	}

	_, err := r.db.DB().ExecContext(
		ctx,
		q,
		user.Name, user.Email, user.Role, user.Password, time.Now(), time.Now(),
	)
	if err != nil {
		return 0, err
	}
	fmt.Println(color.BlueString("Create user: name - %+v, with ctx: %v", user, ctx))
	return userID, nil
}

func (r *repo) Delete(ctx context.Context, userID int64) error {

	q := db.Query{
		Name:     deleteRepoFn,
		QueryRaw: `DELETE FROM users WHERE id = $1;`,
	}

	_, err := r.db.DB().ExecContext(
		ctx,
		q,
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

	q := db.Query{
		Name:     updateRepoFn,
		QueryRaw: sql,
	}

	_, err := r.db.DB().ExecContext(
		ctx,
		q,
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

	q := db.Query{
		Name:     getRepoFn,
		QueryRaw: `SELECT name, email, role, created_at, updated_at FROM users WHERE id = $1;`,
	}

	err := r.db.DB().ScanOneContext(
		ctx,
		&userRow,
		q,
		userID,
	)
	if err != nil {
		return &model.User{}, err
	}

	userRow.ID = userID
	fmt.Println(color.BlackString("Get user: id %d, with ctx: %v", userID, ctx))
	return converter.ToUserFromRepo(userRow), nil
}
