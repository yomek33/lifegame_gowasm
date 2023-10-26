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
	js.Global().Get("document").Call("getElementById", "runButton").Call("addEventListener", "click", js.FuncOf(runUpdateGrid))
	js.Global().Get("document").Call("getElementById", "reset").Call("addEventListener", "click", js.FuncOf(clearAllCellColors))
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
	row, _ := strconv.Atoi(cell.Call("getAttribute", "data-row").String())
	col, _ := strconv.Atoi(cell.Call("getAttribute", "data-col").String())
	grid[row][col] = !grid[row][col]
	return nil
}

func runUpdateGrid(_ js.Value, _ []js.Value) interface{} {
	grid = UpdateGrid(grid)
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

func clearAllCellColors(_ js.Value, _ []js.Value) interface{} {
	document := js.Global().Get("document")
	cellElems := document.Call("querySelectorAll", ".cell.clicked")

	for i := 0; i < cellElems.Length(); i++ {
		cellElems.Index(i).Get("classList").Call("remove", "clicked")
	}
	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			grid[i][j] = false
		}
	}
	return nil
}
