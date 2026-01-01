package main

import (
	"log"

	"github.com/exanubes/typedef/internal/infrastructure/version"
	"github.com/exanubes/typedef/internal/rpc"
)

// Build-time variables (injected via -ldflags)
var (
	Version   = "dev"
	CommitSHA = "unknown"
	BuildTime = "unknown"
)

func main() {
	v := version.New()
	if v.Selected() {
		v.Print(map[string]string{
			"version":    Version,
			"commit_sha": CommitSHA,
			"build_time": BuildTime,
		})
		return
	}

	router := rpc.NewRouter()
	server := rpc.NewServer(router)
	if err := server.Start(); err != nil {
		log.Fatal("RPC Server failed: ", err)
	}
}
