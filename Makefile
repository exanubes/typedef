.PHONY: build

build-wasm:
	GOOS=js GOARCH=wasm go build -C ./cmd/wasm -o ../../web/src/js/rpc/main.wasm
start-web:
	(cd web && pnpm start)

