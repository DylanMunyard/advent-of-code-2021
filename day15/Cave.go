package main

import (
	"container/heap"
	"math"
)

type Cell struct {
	value      int
	distance   int
	neighbours []*Cell
	visited    bool
	x          int
	y          int
}

type Row []*Cell

type Cave []Row

type Grid struct {
	rows int
	cols int
	cave Cave
}

func buildGrid(input [][]int, repeater int) (grid Grid) {
	rows := len(input)
	cols := len(input[0])
	initialDistance := rows * cols * 10
	grid.rows = rows * repeater
	grid.cols = cols * repeater

	for r := 0; r < rows; r++ {
		for repeatRow := 0; repeatRow < repeater; repeatRow++ {
			grid.cave = append(grid.cave, make(Row, grid.cols))
		}
	}

	for r := 0; r < rows; r++ {
		for repeatRow := 0; repeatRow < repeater; repeatRow++ {
			for c := 0; c < cols; c++ {
				value := input[r][c]

				for repeatCol := 0; repeatCol < repeater; repeatCol++ {
					row := rows*repeatRow + r
					col := cols*repeatCol + c
					repeatedValue := value + repeatRow + repeatCol

					if repeatedValue > 9 {
						repeatedValue = repeatedValue - 9
					}

					grid.cave[row][col] = &Cell{
						value:    repeatedValue,
						distance: initialDistance,
						visited:  false,
						x:        row,
						y:        col,
					}
				}
			}
		}
	}

	grid.cave[0][0].distance = 0
	calculateNeighbours(&grid)

	return grid
}

func calculateNeighbours(grid *Grid) {
	for r, row := range grid.cave {
		for c, col := range row {
			indices := [][]int{
				{r, c + 1},
				{r, c - 1},
				{r - 1, c},
				{r + 1, c},
			}

			for _, cell := range indices {
				r := cell[0]
				c := cell[1]

				if r < 0 || c < 0 {
					continue
				}

				if r > grid.rows-1 || c > grid.cols-1 {
					continue
				}

				col.neighbours = append(col.neighbours, grid.cave[r][c])
			}
		}
	}
}

func (grid *Grid) visitCells() {
	pq := make(PriorityQueue, grid.rows*grid.cols)
	for r, row := range grid.cave {
		for c, cell := range row {
			i := r*grid.rows + c
			pq[i] = &Item{
				value:    cell,
				priority: cell.distance,
				index:    i,
			}
		}
	}
	heap.Init(&pq)

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)

		for _, neighbour := range item.value.neighbours {
			if neighbour.visited {
				continue
			}

			neighbour.distance = int(math.Min(
				float64(item.value.distance+neighbour.value),
				float64(neighbour.distance)))

			pq.update(neighbour, neighbour.distance)
		}
		item.value.visited = true
	}
}
