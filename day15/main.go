package main

import "fmt"

var Grid = [][]int{
	{1, 1, 6, 3, 7, 5, 1, 7, 4, 2},
	{1, 3, 8, 1, 3, 7, 3, 6, 7, 2},
	{2, 1, 3, 6, 5, 1, 1, 3, 2, 8},
	{3, 6, 9, 4, 9, 3, 1, 5, 6, 9},
	{7, 4, 6, 3, 4, 1, 7, 1, 1, 1},
	{1, 3, 1, 9, 1, 2, 8, 1, 3, 7},
	{1, 3, 5, 9, 9, 1, 2, 4, 2, 1},
	{3, 1, 2, 5, 4, 2, 1, 6, 3, 9},
	{1, 2, 9, 3, 1, 3, 8, 5, 2, 1},
	{2, 3, 1, 1, 9, 4, 4, 5, 8, 1},
}

type RiskFunc func(risk int)

/*
 find a path to the bottom right of the grid starting at row r and column c.
  rows and cols are the nxm dimensions of the grid.
*/
func walkGrid(r int, c int, rows int, cols int, risk int, riskFn RiskFunc) {
	if r == rows-1 {
		// at the bottom, go right
		for i := c; i < cols; i++ {
			risk += Grid[r][i]
		}
		riskFn(risk)
		return
	}

	if c == cols-1 {
		// at the very right, go down
		for i := r; i < rows; i++ {
			risk += Grid[i][c]
		}
		riskFn(risk)
		return
	}

	// add current grid cell to path
	if r > 0 || c > 0 {
		risk = risk + Grid[r][c]
	}

	// recursively walkGrid down
	walkGrid(r+1, c, rows, cols, risk, riskFn)

	// recursively walkGrid right
	walkGrid(r, c+1, rows, cols, risk, riskFn)
}

func walkRisk() <-chan int {
	risks := make(chan int)

	go func() {
		defer close(risks)
		walkGrid(0, 0, len(Grid[0]), len(Grid), 0, func(risk int) {
			risks <- risk
		})
	}()

	return risks
}

func main() {
	risks := walkRisk()
	lowestRisk := -1
	for risk := range risks {
		if lowestRisk == -1 || risk < lowestRisk {
			fmt.Printf("old risk: %d, new risk: %d\n", lowestRisk, risk)
			lowestRisk = risk
		}
	}
}
