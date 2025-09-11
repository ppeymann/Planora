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
	roompb "github.com/ppeymann/Planora.git/proto/room"
	"github.com/ppeymann/Planora/room/repository"
	"github.com/ppeymann/Planora/room/service"
	"github.com/ppeymann/Planora/room/transport"
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

	roomRepo := repository.NewRoomRepo(db, env.GetEnv("DATABASE", ""))
	if err := roomRepo.Migrate(); err != nil {
		log.Fatal("migration failed", err)
	}

	roomService := service.NewRoomServiceServer(roomRepo)

	// ======== NATS Connection =======
	nc, err := nats.Connect(env.GetEnv("NATS_CONNECTION", nats.DefaultURL))
	if err != nil {
		log.Fatal("failed to connection NATS", err)
	}
	defer nc.Close()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	g, ctx := errgroup.WithContext(ctx)

	// Run Nats subscriber
	g.Go(func() error {
		log.Println("starting NATS sub...")

		transport.RegisterRoomSubscriber(nc, roomService)

		<-ctx.Done()

		log.Println("shutting down NATS sub...")
		return nil
	})

	// Run gRPC server
	g.Go(func() error {
		port := env.GetEnv("ROOM_PORT", ":5003")
		lis, err := net.Listen("tcp", port)
		if err != nil {
			return err
		}

		grpcServer := grpc.NewServer()
		roompb.RegisterRoomServiceServer(grpcServer, roomService)

		go func() {
			<-ctx.Done()
			log.Println("shutting down gRPC server...")
			grpcServer.GracefulStop()
		}()

		log.Printf("gRPC server running on %s\n", port)
		return grpcServer.Serve(lis)
	})

	if err := g.Wait(); err != nil {
		log.Printf("server stopped with error: %v", err)
	} else {
		log.Println("server stopped gracefully")
	}
}
