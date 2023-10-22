package main

// import (
// 	"fmt"
// 	"time"
// )

// // const (
// // 	row = 30
// // 	col = 30
// // )

// // var grid [row][col]bool

// // type CellState struct {
// // 	X     int  `json:"x"`
// // 	Y     int  `json:"y"`
// // 	Alive bool `json:"alive"`
// // }

// // wasmでlife gameを実装する. wasmで実装するためにはmain関数が必要

// func life_main() {
// 	initGrid()
// 	for {
// 		printGrid()
// 		updateGrid()
// 		time.Sleep(1 * time.Second)
// 	}

// }

// // func initGrid() {
// // 	rand.Seed(time.Now().UnixNano()) // ランダムシードを設定

// // 	for i := 0; i < row; i++ {
// // 		for j := 0; j < col; j++ {
// // 			// ランダムにセルを生または死に設定
// // 			grid[i][j] = rand.Intn(2) == 1
// // 		}
// // 	}
// // }

// // ライフゲームのルールに沿ってグリットの状態を更新する
// func updateGrid() {
// 	var newGrid [row][col]bool
// 	for i := 0; i < row; i++ {
// 		for j := 0; j < col; j++ {
// 			newGrid[i][j] = false
// 		}
// 	}
// 	for i := 0; i < row; i++ {
// 		for j := 0; j < col; j++ {
// 			//周囲の生きたセルの数を数える
// 			count := 0
// 			for x := i - 1; x <= i+1; x++ {
// 				for y := j - 1; y <= j+1; y++ {
// 					if x >= 0 && x < row && y >= 0 && y < col && grid[x][y] {
// 						count++
// 					}
// 				}
// 			}
// 			//自分自身は数えない
// 			if grid[i][j] {
// 				count--
// 			}
// 			//ルールに沿って次の状態を決定する
// 			if grid[i][j] && (count == 2 || count == 3) {
// 				newGrid[i][j] = true
// 			} else if !grid[i][j] && count == 3 {
// 				newGrid[i][j] = true
// 			}
// 		}
// 	}
// 	//グリットの状態を更新する
// 	grid = newGrid
// }

// // グリットの状態をconsoleに出力する
// func printGrid() {
// 	for i := 0; i < row; i++ {
// 		for j := 0; j < col; j++ {
// 			if grid[i][j] {
// 				fmt.Print("o")
// 			} else {
// 				fmt.Print("□")
// 			}
// 		}
// 		fmt.Println()
// 	}
// }
