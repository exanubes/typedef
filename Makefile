.PHONY: start-web build-wasm build-cli build-linux build-macos build-wasm build-web build-web-ci build-rpc-linux build-rpc-macos test-homebrew-formula test-homebrew-install test-homebrew-formula-ci test-homebrew-install-ci

# Build-time variables (can be overridden via make VAR=value)
VERSION ?= dev
COMMIT_SHA ?= unknown
BUILD_TIME ?= unknown

BUILD_FLAGS := -X main.Version=$(VERSION) -X main.CommitSHA=$(COMMIT_SHA) -X main.BuildTime=$(BUILD_TIME)

build-wasm-dev:
	CGO_ENABLED=0 GOOS=js GOARCH=wasm go build -C ./cmd/wasm -o ../../client/web/src/js/rpc/main.wasm

start-web:
	(cd web && pnpm start)

build-native:
	go build -ldflags "$(BUILD_FLAGS_CLI)" -o dist/cli/cli ./cmd/cli

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "$(BUILD_FLAGS)" -o dist/cli/linux/amd64/typedef-cli ./cmd/cli
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags "$(BUILD_FLAGS)" -o dist/cli/linux/arm64/typedef-cli ./cmd/cli

build-macos:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "$(BUILD_FLAGS)" -o dist/cli/darwin/amd64/typedef-cli ./cmd/cli
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags "$(BUILD_FLAGS)" -o dist/cli/darwin/arm64/typedef-cli ./cmd/cli

build-wasm:
	CGO_ENABLED=0 GOOS=js GOARCH=wasm go build -ldflags "$(BUILD_FLAGS)" -o dist/wasm/main.wasm ./cmd/wasm

build-rpc-linux:
	mkdir -p dist/rpc/linux/amd64 dist/rpc/linux/arm64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "$(BUILD_FLAGS)" -o dist/rpc/linux/amd64/typedef-rpc ./cmd/rpc
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags "$(BUILD_FLAGS)" -o dist/rpc/linux/arm64/typedef-rpc ./cmd/rpc

build-rpc-macos:
	mkdir -p dist/rpc/darwin/amd64 dist/rpc/darwin/arm64
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "$(BUILD_FLAGS)" -o dist/rpc/darwin/amd64/typedef-rpc ./cmd/rpc
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags "$(BUILD_FLAGS)" -o dist/rpc/darwin/arm64/typedef-rpc ./cmd/rpc

build-cli: build-linux build-macos

test-homebrew-formula: build-cli
	@echo "Testing Homebrew formula..."
	chmod +x client/homebrew/scripts/test-formula-local.sh
	./client/homebrew/scripts/test-formula-local.sh $(VERSION)

test-homebrew-install: build-cli
	@echo "Testing Homebrew formula with installation..."
	chmod +x client/homebrew/scripts/test-formula-local.sh
	./client/homebrew/scripts/test-formula-local.sh $(VERSION) --install

test-homebrew-formula-ci:
	@echo "Testing Homebrew formula generation with pre-packaged binaries (CI mode)..."
	chmod +x client/homebrew/scripts/test-formula-ci.sh
	./client/homebrew/scripts/test-formula-ci.sh $(VERSION) ./release

test-homebrew-install-ci:
	@echo "Testing Homebrew formula with pre-packaged binaries (CI mode)..."
	chmod +x client/homebrew/scripts/test-formula-ci.sh
	./client/homebrew/scripts/test-formula-ci.sh $(VERSION) ./release --install
