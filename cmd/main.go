package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"net"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/fatih/color"
	guser "github.com/ukrainskykirill/auth/pkg/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const grpcPort = 50052

var userRoles = [2]guser.UserRole{guser.UserRole_ADMIN, guser.UserRole_USER}

func getRandomRole() guser.UserRole {
	defaultRole := guser.UserRole_USER
	rand, err := rand.Int(rand.Reader, big.NewInt(2))
	if err != nil {
		return defaultRole
	}
	return userRoles[rand.Int64()]
}

type server struct {
	guser.UnimplementedUserV1Server
}

func (s *server) Create(ctx context.Context, req *guser.CreateRequest) (*guser.CreateResponse, error) {
	fmt.Println(color.BlueString("Create user: name - %+v, with ctx: %v", req, ctx))
	return &guser.CreateResponse{
		Id: gofakeit.Int64(),
	}, nil
}

func (s *server) Get(ctx context.Context, req *guser.GetRequest) (*guser.GetResponse, error) {
	fmt.Println(color.BlackString("Get user: id %d, with ctx: %v", req.Id, ctx))
	return &guser.GetResponse{
		Id:        gofakeit.Int64(),
		Name:      gofakeit.Name(),
		Email:     gofakeit.Email(),
		Role:      getRandomRole(),
		CreatedAt: timestamppb.New(time.Now()),
		UpdatedAt: timestamppb.New(time.Now()),
	}, nil
}

func (s *server) Delete(ctx context.Context, req *guser.DeleteRequest) (*emptypb.Empty, error) {
	fmt.Println(color.RedString("Delete user: id %d, with ctx: %v", req.Id, ctx))
	return &emptypb.Empty{}, nil
}

func (s *server) Update(ctx context.Context, req *guser.UpdateRequest) (*emptypb.Empty, error) {
	fmt.Println(color.WhiteString("Update user: id %+v, with ctx: %v", req, ctx))
	return &emptypb.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	guser.RegisterUserV1Server(s, &server{})
	fmt.Println(color.GreenString("run server at %s", lis.Addr()))
	if err = s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
