.PHONY: start-web build-wasm build-cli build-linux build-macos build-wasm build-web build-web-ci

# Build-time variables (can be overridden via make VAR=value)
VERSION ?= dev
COMMIT_SHA ?= unknown
BUILD_TIME ?= unknown

# Construct ldflags for version injection
LDFLAGS_CLI := -X main.Version=$(VERSION) -X main.CommitSHA=$(COMMIT_SHA) -X main.BuildTime=$(BUILD_TIME)
LDFLAGS_WASM := -X main.Version=$(VERSION) -X main.CommitSHA=$(COMMIT_SHA) -X main.BuildTime=$(BUILD_TIME)
LDFLAGS_RPC := -X main.Version=$(VERSION) -X main.CommitSHA=$(COMMIT_SHA) -X main.BuildTime=$(BUILD_TIME)

build-wasm-dev:
	CGO_ENABLED=0 GOOS=js GOARCH=wasm go build -C ./cmd/wasm -o ../../client/web/src/js/rpc/main.wasm

start-web:
	(cd web && pnpm start)

build-native:
	go build -ldflags "$(LDFLAGS_CLI)" -o dist/cli/cli ./cmd/cli

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "$(LDFLAGS_CLI)" -o dist/cli/linux/amd64/cli ./cmd/cli
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags "$(LDFLAGS_CLI)" -o dist/cli/linux/arm64/cli ./cmd/cli

build-macos:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "$(LDFLAGS_CLI)" -o dist/cli/darwin/amd64/cli ./cmd/cli
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags "$(LDFLAGS_CLI)" -o dist/cli/darwin/arm64/cli ./cmd/cli

build-wasm:
	CGO_ENABLED=0 GOOS=js GOARCH=wasm go build -ldflags "$(LDFLAGS_WASM)" -o dist/wasm/main.wasm ./cmd/wasm

build-rpc:
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags "$(LDFLAGS_RPC)" -o dist/rpc/typedef-rpc ./cmd/rpc
