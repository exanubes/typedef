package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/exanubes/typedef/internal/drivers/cli"
	"golang.design/x/clipboard"
)

func main() {
	args := os.Args[1:]
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	err := clipboard.Init()

	if err != nil {
		fmt.Println("Clipboard error: ", err.Error())
		fmt.Println("Clipboard input will not be supported")
	}

	err = cli.Start(ctx, args)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)

	}
	os.Exit(0)
}
