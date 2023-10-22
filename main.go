//go:build js && wasm
// +build js,wasm

package main

import (
	"syscall/js"
)

func main() {
	c := make(chan struct{}, 0)
	registerCallbacks()
	<-c
}

func registerCallbacks() {
	js.Global().Set("go_initGrid", js.FuncOf(initGrid))
}

// life game
const (
	row = 30
	col = 30
)

var grid [row][col]bool


func initGrid(_ js.Value, args []js.Value) interface{} {
    inputArray := args[0]

    var initGrid []interface{}

    for i := 0; i < inputArray.Length(); i++ {
        cell := inputArray.Index(i)
        row := cell.Index(0)
        col := cell.Index(1)

        jsRow := js.Global().Get("Array").New()
        jsRow.Call("push", row.Int())
        jsRow.Call("push", col.Int())

        initGrid = append(initGrid, jsRow)
		grid[row.Int()][col.Int()] = true
    }

    jsGrid := js.Global().Get("Array").New(initGrid)

    return jsGrid
}
