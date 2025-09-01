package main

import (
	"fmt"
	"log"
	"os"
	"time"

	kitLog "github.com/go-kit/log"
	"github.com/nats-io/nats.go"
	"github.com/ppeymann/Planora.git/pkg/env"
	"github.com/ppeymann/Planora/gateway/server"
)

func main() {
	now := time.Now().UTC()

	base := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC).Unix()
	start := time.Date(now.Year(), now.Month(), now.Day(), 7, 35, 0, 0, time.UTC).Unix()
	end := time.Date(now.Year(), now.Month(), now.Day(), 23, 30, 0, 0, time.UTC).Unix()

	fmt.Println("date:", base, "start:", start, "end:", end)

	// Connect to nats
	nc, err := nats.Connect(env.GetEnv("NATS_CONNECTION", ""))
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// configuration logger
	var logger kitLog.Logger
	logger = kitLog.NewJSONLogger(kitLog.NewSyncWriter(os.Stderr))
	logger = kitLog.With(logger, "ts", kitLog.DefaultTimestampUTC)

	sl := kitLog.With(logger, "component", "http")

	// Server instance
	svr := server.NewServer(sl)

	// =======  SERVICE  ========

	// listen and serve...
	svr.Listen()

}
