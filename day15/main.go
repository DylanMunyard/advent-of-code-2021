package main

import "fmt"

/*
 find a path to the bottom right of the grid starting at row r and column c.
  rows and cols are the nxm dimensions of the grid.
*/
func walk(grid [][]int, r int, c int, rows int, cols int, path []int) [][]int {
	var paths [][]int

	if r == rows-1 {
		// at the bottom, go right
		for i := c; i < cols; i++ {
			path = append(path, grid[r][i])
		}
		return append(paths, path)
	}

	if c == cols-1 {
		// at the very right, go down
		for i := r; i < rows; i++ {
			path = append(path, grid[i][c])
		}
		return append(paths, path)
	}

	// add current grid cell to path
	path = append(path, grid[r][c])

	// recursively walk down
	paths = append(paths, walk(grid, r+1, c, rows, cols, path)...)

	// recursively walk right
	paths = append(paths, walk(grid, r, c+1, rows, cols, path)...)

	return paths
}

func main() {
	grid := [][]int{
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

	var path []int
	visited := walk(grid, 0, 0, len(grid[0]), len(grid), path)
	fmt.Printf("%d", len(visited))
}
