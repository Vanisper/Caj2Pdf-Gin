package main

import (
	"caj2pdf/md5"
	"syscall/js"
)

func main() {
	c := make(chan struct{})
	registerCallbacks()
	<-c
}

func registerCallbacks() {
	// js.Global().Set("sayHello", js.FuncOf(sayHello))
	/* js.Global().Get("document").
	Call("getElementById", "brightness").
	Call("addEventListener", "change", brightnessCb) */
	js.Global().Set("getFileMd5", js.FuncOf(md5.GetFileMd5))
	js.Global().Set("wasmMd5Add", js.FuncOf(md5.WasmMd5Add))
	js.Global().Set("wasmMd5End", js.FuncOf(md5.WasmMd5End))
}
