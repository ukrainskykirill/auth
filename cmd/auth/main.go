package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"github.com/ukrainskykirill/auth/internal/config"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pool "github.com/jackc/pgx/v5/pgxpool"

	"github.com/fatih/color"
	guser "github.com/ukrainskykirill/auth/pkg/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type user struct {
	ID        int64
	Name      string
	Email     string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func getRole(role string) guser.UserRole {
	switch role {
	case guser.UserRole_USER.String():
		return guser.UserRole_USER
	case guser.UserRole_ADMIN.String():
		return guser.UserRole_ADMIN
	default:
		return guser.UserRole_UNKNOW
	}
}

func validatePassword(pass, confPass string) error {
	if pass != confPass {
		return fmt.Errorf("password does not match")
	}
	return nil
}

func validateEmail(email string) error {
	if !strings.Contains(email, "@") {
		return fmt.Errorf("invalid email")
	}
	return nil

}

type server struct {
	guser.UnimplementedUserV1Server
	pg *pool.Pool
}

func (s *server) Create(ctx context.Context, req *guser.CreateRequest) (*guser.CreateResponse, error) {
	err := validateEmail(req.Email)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err = validatePassword(req.Password, req.PasswordConfirm)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	tx, err := s.pg.Begin(ctx)
	defer func() {
		if err := tx.Rollback(ctx); err != nil {
			log.Printf("rollback error: %v", err)
		}
	}()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var userID int64
	err = tx.QueryRow(
		ctx,
		`INSERT INTO users (name, email, role, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;`,
		req.Name, req.Email, req.Role, password, time.Now(), time.Now(),
	).Scan(&userID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	fmt.Println(color.BlueString("Create user: name - %+v, with ctx: %v", req, ctx))
	return &guser.CreateResponse{
		Id: userID,
	}, nil
}

func (s *server) Get(ctx context.Context, req *guser.GetRequest) (*guser.GetResponse, error) {
	tx, err := s.pg.Begin(ctx)
	defer func() {
		if err := tx.Rollback(ctx); err != nil {
			log.Printf("rollback error: %v", err)
		}
	}()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var userRow user
	err = tx.QueryRow(
		ctx,
		`SELECT name, email, role, created_at, updated_at FROM users WHERE id = $1;`,
		req.Id,
	).Scan(&userRow.Name, &userRow.Email, &userRow.Role, &userRow.CreatedAt, &userRow.UpdatedAt)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	fmt.Println(color.BlackString("Get user: id %d, with ctx: %v", req.Id, ctx))
	return &guser.GetResponse{
		Id:        userRow.ID,
		Name:      userRow.Name,
		Email:     userRow.Email,
		Role:      getRole(userRow.Role),
		CreatedAt: timestamppb.New(userRow.CreatedAt),
		UpdatedAt: timestamppb.New(userRow.UpdatedAt),
	}, nil
}

func (s *server) Delete(ctx context.Context, req *guser.DeleteRequest) (*emptypb.Empty, error) {
	tx, err := s.pg.Begin(ctx)
	defer func() {
		if err := tx.Rollback(ctx); err != nil {
			log.Printf("rollback error: %v", err)
		}
	}()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	_, err = tx.Exec(
		ctx,
		`DELETE FROM users WHERE id = $1;`,
		req.Id,
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	fmt.Println(color.RedString("Delete user: id %d, with ctx: %v", req.Id, ctx))
	return &emptypb.Empty{}, nil
}

func (s *server) Update(ctx context.Context, req *guser.UpdateRequest) (*emptypb.Empty, error) {
	var sql string
	var args []interface{}
	paramIndex := 1

	sql = `UPDATE users SET `

	if req.Name != nil {
		sql += fmt.Sprintf("name = $%d", paramIndex)
		args = append(args, req.Name.Value)
		paramIndex++
	}

	if req.Email != nil {
		err := validateEmail(req.Email.Value)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		if len(args) > 0 {
			sql += ", "
		}
		sql += fmt.Sprintf("email = $%d", paramIndex)
		args = append(args, req.Email.Value)
		paramIndex++
	}

	if req.Role != nil {
		if len(args) > 0 {
			sql += ", "
		}
		sql += fmt.Sprintf("role = $%d", paramIndex)
		args = append(args, req.Role)
		paramIndex++
	}

	if len(args) > 0 {
		sql += ", updated_at = $" + fmt.Sprintf("%d", paramIndex)
		args = append(args, time.Now())
		paramIndex++
	}

	sql += " WHERE id = $" + fmt.Sprintf("%d;", paramIndex)
	args = append(args, req.Id)

	tx, err := s.pg.Begin(ctx)
	defer func() {
		if err := tx.Rollback(ctx); err != nil {
			log.Printf("rollback error: %v", err)
		}
	}()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	_, err = tx.Exec(
		ctx,
		sql,
		args...,
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	fmt.Println(color.WhiteString("Update user: id %+v, with ctx: %v", req, ctx))
	return &emptypb.Empty{}, nil
}

func main() {
	appConf, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	pgxPool, err := pool.New(context.Background(), appConf.DB.URL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer pgxPool.Close()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", appConf.GRPC.Port))
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	guser.RegisterUserV1Server(s, &server{
		pg: pgxPool,
	})

	fmt.Println(color.GreenString("run server at %s", lis.Addr()))

	if err = s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
