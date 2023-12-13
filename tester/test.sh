#! /bin/bash
go build -tags=llvm16 ../.
./tinygo build -x -target=wasip2 -o main.wasm main.go
wasm-tools component embed -w wasi:cli/reactor ./wasi-cli/wit/ main.wasm -o embedded.wasm
wasm-tools component new embedded.wasm -o component.wasm --adapt ./wasi_snapshot_preview1.command.wasm
wasmtime --wasm component-model --env KEY0=VALUE0 --env KEY1=VALUE1 component.wasm arg1 arg2 arg3 arg4 arg5 
