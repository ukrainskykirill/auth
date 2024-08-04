package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/fatih/color"
	pool "github.com/jackc/pgx/v5/pgxpool"
	userApi "github.com/ukrainskykirill/auth/internal/api/user"
	"github.com/ukrainskykirill/auth/internal/config"
	userRepo "github.com/ukrainskykirill/auth/internal/repository/user"
	userServ "github.com/ukrainskykirill/auth/internal/service/user"
	guser "github.com/ukrainskykirill/auth/pkg/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

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

	repository := userRepo.NewUserRepository(pgxPool)
	service := userServ.NewServ(repository)
	api := userApi.NewImplementation(service)

	s := grpc.NewServer()
	reflection.Register(s)
	guser.RegisterUserV1Server(s, api)

	fmt.Println(color.GreenString("run server at %s", lis.Addr()))

	if err = s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
