package main

import (
	"fmt"

	"github.com/bytecodealliance/wasmtime-go"
)

type Runner struct {
	extension *Extension
}

func NewRunner(extension *Extension) *Runner {
	return &Runner{
		extension: extension,
	}
}

func (r *Runner) Run() (string, error) {
	engine := wasmtime.NewEngine()

	store := wasmtime.NewStore(engine)
	wasiConfig := wasmtime.NewWasiConfig()
	wasiConfig.InheritStdout()
	store.SetWasi(wasiConfig)

	module, err := wasmtime.NewModule(store.Engine, r.extension.Code)
	if err != nil {
		return "", err
	}

	linker := wasmtime.NewLinker(engine)
	err = linker.DefineWasi()
	if err != nil {
		return "", err
	}

	instance, err := linker.Instantiate(store, module)
	if err != nil {
		return "", err
	}

	mem := instance.GetExport(store, "memory").Memory()

	input := "hello from runner!"
	inputLength := len(input)

	inputPtr, err := allocate(instance, store, inputLength-1)
	if err != nil {
		return "", err
	}

	memBytes := mem.UnsafeData(store)[inputPtr:]
	copy(memBytes, input)

	run := instance.GetFunc(store, "run")
	if run == nil {
		return "", fmt.Errorf("run not a function")
	}
	ptr, err := run.Call(store, inputPtr, inputLength)
	if err != nil {
		return "", err
	}

	data := mem.UnsafeData(store)[int(ptr.(int32)):int(ptr.(int32)+32)]

	return string(data), nil
}

func allocate(instance *wasmtime.Instance, store *wasmtime.Store, length int) (int32, error) {
	run := instance.GetFunc(store, "allocate")
	if run == nil {
		return 0, fmt.Errorf("allocate not a function")
	}
	ptr, err := run.Call(store, length)
	if err != nil {
		return 0, err
	}

	return ptr.(int32), nil
}
