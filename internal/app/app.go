package app

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/fatih/color"
	"github.com/ukrainskykirill/platform_common/pkg/closer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"

	"github.com/ukrainskykirill/auth/internal/config"
	"github.com/ukrainskykirill/auth/internal/interceptor"
	gaccess "github.com/ukrainskykirill/auth/pkg/access_v1"
	gauth "github.com/ukrainskykirill/auth/pkg/auth_v1"
	guser "github.com/ukrainskykirill/auth/pkg/user_v1"
)

type App struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run(ctx context.Context) error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	ctx, cancel := context.WithCancel(ctx)

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()

		err := a.runGRPCServer()
		if err != nil {
			fmt.Printf("run grpc service: %s", err)
		}
	}()

	go func() {
		defer wg.Done()

		err := a.serviceProvider.UserCreateConsumer(ctx).RunConsumer(ctx)
		if err != nil {
			fmt.Printf("run user create consumer: %s", err)
		}

	}()

	gracefulShutdown(ctx, cancel, wg)

	return nil
}

func gracefulShutdown(ctx context.Context, cancel context.CancelFunc, wg *sync.WaitGroup) {
	select {
	case <-ctx.Done():
		log.Println("terminating: context cancelled")
	case <-waitSignal():
		log.Println("terminating: via signal")
	}

	cancel()
	if wg != nil {
		wg.Wait()
	}
}

func waitSignal() chan os.Signal {
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	return sigterm
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initGRPCServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	err := config.LoadConfig()
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initGRPCServer(ctx context.Context) error {
	creds, err := credentials.NewServerTLSFromFile("service.pem", "service.key")
	if err != nil {
		log.Fatalf("failed to load TLS keys: %v", err)
	}

	a.grpcServer = grpc.NewServer(
		grpc.Creds(creds),
		grpc.UnaryInterceptor(interceptor.ValidateInterceptor),
	)

	reflection.Register(a.grpcServer)

	guser.RegisterUserV1Server(a.grpcServer, a.serviceProvider.UserAPI(ctx))
	gauth.RegisterAuthV1Server(a.grpcServer, a.serviceProvider.AuthAPI(ctx))
	gaccess.RegisterAccessV1Server(a.grpcServer, a.serviceProvider.AccessAPI(ctx))

	return nil
}

func (a *App) runGRPCServer() error {
	lis, err := net.Listen("tcp", a.serviceProvider.GRPCConfig().Address())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(color.GreenString("run server at %s", lis.Addr()))
	if err = a.grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}

	return nil
}
