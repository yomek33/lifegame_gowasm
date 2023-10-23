//go:build js && wasm
// +build js,wasm

package main

import "syscall/js"



func main() {
	c := make(chan struct{}, 0)
	registerCallbacks()
	<-c
}

func registerCallbacks() {
	js.Global().Set("go_initGrid", js.FuncOf(initGridfromJS))
}

// life game
const (
	row = 30
	col = 30
)

var grid [row][col]bool


func initGridfromJS(_ js.Value, args []js.Value) interface{} {
    inputArray := args[0]

    var updatedGridtoJS []interface{}

    for i := 0; i < inputArray.Length(); i++ {
        cell := inputArray.Index(i)
        row := cell.Index(0)
        col := cell.Index(1)
		grid[row.Int()][col.Int()] = true
    }

	updatedGrid :=UpdateGrid(grid)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if updatedGrid[i][j] {
				jsRow := js.Global().Get("Array").New()
				jsRow.Call("push", i)
				jsRow.Call("push", j)
				updatedGridtoJS = append(updatedGridtoJS, jsRow)
			}
		}
	}

	
    jsGrid := js.Global().Get("Array").New(updatedGridtoJS)

    return jsGrid
}

func UpdateGrid(grid [row][col]bool) [row][col]bool {
	var newGrid [row][col]bool
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			newGrid[i][j] = false
		}
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			//周囲の生きたセルの数を数える
			count := 0
			for x := i - 1; x <= i+1; x++ {
				for y := j - 1; y <= j+1; y++ {
					if x >= 0 && x < row && y >= 0 && y < col && grid[x][y] {
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