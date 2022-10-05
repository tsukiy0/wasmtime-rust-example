build-wasmlib:
	@cd wasmlib && cargo wasi build
	@npx wasm2wat /workspace/wasmlib/target/wasm32-wasi/debug/wasmlib.wasm -o /workspace/wasmlib/target/wasm32-wasi/debug/wasmlib.wat

run:
	@go run .