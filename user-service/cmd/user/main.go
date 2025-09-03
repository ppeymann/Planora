package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/nats-io/nats.go"
	"github.com/ppeymann/Planora.git/pkg/env"
	userpb "github.com/ppeymann/Planora.git/proto/user"
	"github.com/ppeymann/Planora/user/repository"
	"github.com/ppeymann/Planora/user/service"
	"github.com/ppeymann/Planora/user/transport"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// ======= DB Connection ======
	dsn := pg.Open(env.GetEnv("DSN", ""))
	db, err := gorm.Open(dsn, &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		log.Fatal("failed to connect to DB:", err)
	}

	userRepo := repository.NewUserRepo(db, env.GetEnv("DATABASE", ""))
	if err := userRepo.Migrate(); err != nil {
		log.Fatal("migration failed", err)
	}

	userService := service.NewUserServiceServer(userRepo)

	// ======== NATS Connection =======
	nc, err := nats.Connect(env.GetEnv("NATS_CONNECTION", nats.DefaultURL))
	if err != nil {
		log.Fatal("failed to connection NATS", err)
	}
	defer nc.Close()

	// ======== Context + errgroup ======
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	g, ctx := errgroup.WithContext(ctx)

	// Run Nats subscriber
	g.Go(func() error {
		log.Println("starting NATS sub...")

		transport.RegisterUserSubscriber(nc, userService)

		<-ctx.Done()

		log.Println("shutting down NATS sub...")
		return nil
	})

	// Run gRPC server
	g.Go(func() error {
		port := env.GetEnv("USER_PORT", ":5001")
		lis, err := net.Listen("tcp", port)
		if err != nil {
			return err
		}

		grpcServer := grpc.NewServer()
		userpb.RegisterUserServiceServer(grpcServer, userService)

		go func() {
			<-ctx.Done()
			log.Println("shutting down gRPC server...")
			grpcServer.GracefulStop()
		}()

		log.Printf("gRPC server running on %s\n", port)
		return grpcServer.Serve(lis)
	})

	// wait for all goroutines
	if err := g.Wait(); err != nil {
		log.Printf("server stopped with error: %v", err)
	} else {
		log.Println("server stopped gracefully")
	}

}
