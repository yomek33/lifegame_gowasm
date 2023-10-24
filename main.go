package main

import (
	"strconv"
	"syscall/js"
)

var (
	numRows       = 30
	numCols       = 30
	gridContainer js.Value
	cellElems     js.Value
)

func main() {
	c := make(chan struct{}, 0)
	registerCallbacks()
	<-c
}

func registerCallbacks() {
	js.Global().Set("go_createGrid", js.FuncOf(createGrid))
	js.Global().Set("go_cellClickHandler", js.FuncOf(cellClickHandler))
	// js.Global().Get("document").Call("getElementById", "runButton").Call("addEventListener", "click", js.FuncOf(runButtonClicked))
	js.Global().Set("go_updateGrid", js.FuncOf(updateGrid))
	js.Global().Set("go_clearAllCellColors", js.FuncOf(clearAllCellColors))
	js.Global().Set("go_colorgrid", js.FuncOf(colerGrid))
}

func createGrid(_ js.Value, _ []js.Value) interface{} {
	gridContainer = js.Global().Get("document").Call("getElementById", "grid-container")
	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			cell := js.Global().Get("document").Call("createElement", "div")
			cell.Set("className", "cell")
			cell.Call("setAttribute", "data-row", i)
			cell.Call("setAttribute", "data-col", j)
			gridContainer.Call("appendChild", cell)
		}
	}
	return nil
}

func cellClickHandler(_ js.Value, args []js.Value) interface{} {
	addClickedCells(args[0])
	return nil
}

func updateGrid(this js.Value, _ []js.Value) interface{} {
	clearAllCellColors(this, nil)

	if len(cells) > 0 {
		for _, innerArray := range cells {
			for i := range innerArray {
				js.Global().Call("colerGrid", js.ValueOf([]interface{}{i, innerArray[i]}))
			}
		}
	}

	return nil
}

func colerGrid(_ js.Value, args []js.Value) interface{} {
	cell := args[0]
	row := strconv.Itoa(cell.Index(0).Int())
	col := strconv.Itoa(cell.Index(1).Int())

	cellElem := js.Global().Get("document").Call("querySelector", ".cell[data-row='"+row+"'][data-col='"+col+"']")
	if cellElem.Truthy() {
		addClickedCells(cellElem)
	} else {
		js.Global().Get("console").Call("log", "セルが見つかりません - 行:", row, "列:", col)
	}

	return nil
}
func clearAllCellColors(_ js.Value, _ []js.Value) interface{} {
	document := js.Global().Get("document")
	cellElems := document.Call("querySelectorAll", ".cell.clicked")

	for i := 0; i < cellElems.Length(); i++ {
		removeClickedCells(cellElems.Index(i))
	}

	return nil
}
func addClickedCells(elem js.Value) {
	elem.Get("classList").Call("add", "clicked")
}

func removeClickedCells(elem js.Value) {
	elem.Get("classList").Call("remove", "clicked")
}
