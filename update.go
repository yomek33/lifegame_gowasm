package main

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
