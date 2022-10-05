package main

import "os"

func main() {
	code, err := os.ReadFile("/workspace/wasmlib/target/wasm32-wasi/debug/wasmlib.wasm")
	if err != nil {
		panic(err)
	}

	extension := &Extension{
		Name: "test",
		Code: code,
	}

	runner := NewRunner(extension)

	result, err := runner.Run()
	if err != nil {
		panic(err)
	}

	println(result)
}
