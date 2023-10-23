package main

import (
	"syscall/js"
)


var (
    numRows       = 30
    numCols       = 30
    gridContainer js.Value
    coloredCells  []js.Value
	cells 		  [][]bool
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
	targetCell := args[0]
	row := targetCell.Get("getAttribute").Get("data-row")
	col := targetCell.Get("getAttribute").Get("data-col")
	targetCell.Get("classList").Call("add", "clicked")
	// cells[row.Int()][col.Int()] = true
	coloredCells = append(coloredCells, js.ValueOf([]interface{}{row, col}))
	return nil
}

// func runButtonClicked(_ js.Value, _ []js.Value) interface{} {
// println(coloredCells) //[4/4]0x43c040
//     js.Global().Get("console").Call()

//     return nil
// }

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
            row := cell.Index(0).String()
            col := cell.Index(1).String()

            cellElem := js.Global().Get("document").Call("querySelector", ".cell[data-row='" + row + "'][data-col='" + col + "']")
            if cellElem.Truthy() {
                cellElem.Get("classList").Call("add", "clicked")
            } else {
                js.Global().Get("console").Call("log", "セルが見つかりません - 行:", row, "列:", col)
            }

            return nil
		}
func clearAllCellColors(_ js.Value, _ []js.Value) interface{} {
	document := js.Global().Get("document")
	cellElems := document.Call("querySelectorAll", ".cell.clicked")

	for i := 0; i < cellElems.Length(); i++ {
		cellElem := cellElems.Index(i)
		cellElem.Get("classList").Call("remove", "clicked")
	}

	return nil
}
