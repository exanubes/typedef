.PHONY: start-web build-wasm build-cli build-linux build-macos build-wasm build-web build-web-ci

build-wasm-dev:
	CGO_ENABLED=0 GOOS=js GOARCH=wasm go build -C ./cmd/wasm -o ../../web/src/js/rpc/main.wasm

start-web:
	(cd web && pnpm start)

build-native:
	go build -o dist/cli/cli ./cmd/cli

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/cli/linux/amd64/cli ./cmd/cli
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o dist/cli/linux/arm64/cli ./cmd/cli

build-macos:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o dist/cli/darwin/amd64/cli ./cmd/cli
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o dist/cli/darwin/arm64/cli ./cmd/cli

build-wasm:
	CGO_ENABLED=0 GOOS=js GOARCH=wasm go build -o dist/wasm/main.wasm ./cmd/wasm

