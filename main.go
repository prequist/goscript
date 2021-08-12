package main

import (
	"fmt"
	"io/ioutil"

	"github.com/wasmerio/wasmer-go/wasmer"
)

func main() {
	wasmBytes, _ := ioutil.ReadFile("./build/optimized.wasm")

	engine := wasmer.NewEngine()
	store := wasmer.NewStore(engine)

	module, _ := wasmer.NewModule(store, wasmBytes)

	importObject := wasmer.NewImportObject()
	instance, _ := wasmer.NewInstance(module, importObject)

	add, _ := instance.Exports.GetFunction("pluralize")
	result, _ := add("children", 20.0)

	fmt.Println(result)
}
