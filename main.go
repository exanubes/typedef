package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/exanubes/typedef/internal/drivers/cli"
)

func main() {
	args := os.Args[1:]
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	err := cli.Start(ctx, args)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)

	}
	os.Exit(0)
}
