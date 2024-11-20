package user

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/jackc/pgx/v4"
	"github.com/ukrainskykirill/platform_common/pkg/db"

	prError "github.com/ukrainskykirill/auth/internal/error"
	"github.com/ukrainskykirill/auth/internal/model"
	"github.com/ukrainskykirill/auth/internal/repository"
	"github.com/ukrainskykirill/auth/internal/repository/user/converter"
	modelRepo "github.com/ukrainskykirill/auth/internal/repository/user/model"
)

const (
	userRepo          = "user_repository"
	createRepoFn      = userRepo + "." + "Create"
	deleteRepoFn      = userRepo + "." + "Delete"
	updateRepoFn      = userRepo + "." + "Update"
	getRepoFn         = userRepo + "." + "Get"
	getPasswordByName = userRepo + "." + "GetPasswordByName"
)

type repo struct {
	db db.Client
}

func NewUserRepository(db db.Client) repository.UserRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, user *model.UserIn) (int64, error) {
	var userID int64

	q := db.Query{
		Name:     createRepoFn,
		QueryRaw: `INSERT INTO users (name, email, role, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;`,
	}

	err := r.db.DB().QueryRowContext(
		ctx,
		q,
		user.Name, user.Email, user.Role, user.Password, time.Now(), time.Now(),
	).Scan(&userID)
	if err != nil {
		return 0, prError.ErrNameNotUnique
	}

	fmt.Println(color.BlueString("Create user: name - %+v, with ctx: %v", user, ctx))
	return userID, nil
}

func (r *repo) Delete(ctx context.Context, userID int64) error {

	q := db.Query{
		Name:     deleteRepoFn,
		QueryRaw: `DELETE FROM users WHERE id = $1;`,
	}

	tag, err := r.db.DB().ExecContext(
		ctx,
		q,
		userID,
	)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return prError.ErrUserNotFound
	}

	fmt.Println(color.RedString("Delete user: id %d, with ctx: %v", userID, ctx))
	return nil
}

func (r *repo) Update(ctx context.Context, user *model.UserInUpdate) error {
	var rowSQL string
	var args []interface{}
	paramIndex := 1

	rowSQL = `UPDATE users SET `

	if len(user.Name) != 0 {
		rowSQL += fmt.Sprintf("name = $%d", paramIndex)
		args = append(args, user.Name)
		paramIndex++
	}

	if len(user.Email) != 0 {
		if len(args) > 0 {
			rowSQL += ", "
		}
		rowSQL += fmt.Sprintf("email = $%d", paramIndex)
		args = append(args, user.Email)
		paramIndex++
	}

	if len(user.Role) != 0 {
		if len(args) > 0 {
			rowSQL += ", "
		}
		rowSQL += fmt.Sprintf("role = $%d", paramIndex)
		args = append(args, user.Role)
		paramIndex++
	}

	if len(args) > 0 {
		rowSQL += ", updated_at = $" + fmt.Sprintf("%d", paramIndex)
		args = append(args, time.Now())
		paramIndex++
	}

	rowSQL += " WHERE id = $" + fmt.Sprintf("%d;", paramIndex)
	args = append(args, user.ID)

	q := db.Query{
		Name:     updateRepoFn,
		QueryRaw: rowSQL,
	}

	tag, err := r.db.DB().ExecContext(
		ctx,
		q,
		args...,
	)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return prError.ErrUserNotFound
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

	err := r.db.DB().QueryRowContext(
		ctx,
		q,
		userID,
	).Scan(&userRow.Name, &userRow.Email, &userRow.Role, &userRow.CreatedAt, &userRow.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return &model.User{}, prError.ErrUserNotFound
		}
		return &model.User{}, err
	}

	userRow.ID = userID
	fmt.Println(color.BlackString("Get user: id %d, with ctx: %v", userID, ctx))
	return converter.ToUserFromRepo(userRow), nil
}

func (r *repo) GetUserAuthInfo(ctx context.Context, name string) (*model.UserAuthInfo, error) {
	var userAuthInfo modelRepo.RepoUserAuthInfo

	q := db.Query{
		Name:     getPasswordByName,
		QueryRaw: `SELECT role, password FROM users WHERE name = $1;`,
	}

	err := r.db.DB().QueryRowContext(
		ctx,
		q,
		name,
	).Scan(&userAuthInfo.Role, &userAuthInfo.Password)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return &model.UserAuthInfo{}, prError.ErrUserNotFound
		}
		return &model.UserAuthInfo{}, err
	}

	return converter.ToUserAuthInfoFromRepo(userAuthInfo), nil
}
