package main

import (
	"day15/inputs"
	"fmt"
)

func main() {
	// Part 1

	grid := buildGrid(inputs.Input, 1)
	grid.visitCells()
	fmt.Printf("Part 1: %d", grid.cave[grid.rows-1][grid.cols-1].distance)
	fmt.Println()

	// Part 2
	grid = buildGrid(inputs.Input, 5)
	grid.visitCells()
	fmt.Printf("Part 2: %d", grid.cave[grid.rows-1][grid.cols-1].distance)
}
