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
	grid          gridBool
)

type gridBool [30][30]bool

func main() {
	c := make(chan struct{}, 0)
	registerCallbacks()
	<-c
}

func registerCallbacks() {
	js.Global().Set("go_createGrid", js.FuncOf(createGrid))
	js.Global().Set("go_cellClickHandler", js.FuncOf(cellClickHandler))
	// js.Global().Get("document").Call("getElementById", "runButton").Call("addEventListener", "click", js.FuncOf(runButtonClicked))
	// js.Global().Set("go_updateGrid", js.FuncOf(updateGrid))
	// js.Global().Set("go_clearAllCellColors", js.FuncOf(clearAllCellColors))
	// js.Global().Set("go_colorgrid", js.FuncOf(colerGrid))
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

func cellClickHandler(this js.Value, args []js.Value) interface{} {
	cell := args[0]
	cell.Get("classList").Call("toggle", "clicked")
	row := cell.Index(0).Int()
	col := cell.Index(1).Int()

	grid[row][col] = !grid[row][col]
	return nil
}

func runButtonClicked(_ js.Value, _ []js.Value) interface{} {
	UpdateGrid(grid)
	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			cellElem := js.Global().Get("document").Call("querySelector", ".cell[data-row='"+strconv.Itoa(i)+"'][data-col='"+strconv.Itoa(j)+"']")
			if grid[i][j] {
				cellElem.Get("classList").Call("add", "clicked")
			} else {
				cellElem.Get("classList").Call("remove", "clicked")
			}
		}
	}
	return nil
}

func UpdateGrid(grid gridBool) gridBool {
	var newGrid gridBool
	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			newGrid[i][j] = false
		}
	}
	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			//周囲の生きたセルの数を数える
			count := 0
			for x := i - 1; x <= i+1; x++ {
				for y := j - 1; y <= j+1; y++ {
					if x >= 0 && x < numRows && y >= 0 && y < numCols && grid[x][y] {
						count++
					}
				}
			}
			//自分自身は数えない
			if grid[i][j] {
				count--
			}
			//ルールに沿って次の状態を決定する
			if grid[i][j] && (count == 2 || count == 3) {
				newGrid[i][j] = true
			} else if !grid[i][j] && count == 3 {
				newGrid[i][j] = true
			}
		}
	}
	//グリットの状態を更新する
	return newGrid
}

// func clearAllCellColors(_ js.Value, _ []js.Value) interface{} {
// 	document := js.Global().Get("document")
// 	cellElems := document.Call("querySelectorAll", ".cell.clicked")

// 	for i := 0; i < cellElems.Length(); i++ {
// 		removeClickedCells(cellElems.Index(i))
// 	}

// 	return nil
// }

// func addClickedCells(cell js.Value) {
// 	cell.Get("classList").Call("add", "clicked")
// }

// func removeClickedCells(cell js.Value) {
// 	cell.Get("classList").Call("remove", "clicked")
// }
