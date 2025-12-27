package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/exanubes/typedef/internal/drivers/cli"
)

// Build-time variables (injected via -ldflags)
var (
	Version   = "dev"
	CommitSHA = "unknown"
	BuildTime = "unknown"
)

func main() {
	args := os.Args[1:]
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	err := cli.Start(ctx, args)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)

	}
	os.Exit(0)
}
