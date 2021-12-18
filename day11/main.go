package main

import (
	"fmt"
	"github.com/DylanMunyard/advent-of-code-2021/inputs"
)

var flashes = 0

func main() {
	for step := range [300]int{} {
		flashes = 0

		var flashed []inputs.Octopus
		for row, elements := range inputs.Arr {
			for col := range elements {
				inputs.Arr[row][col] = (inputs.Arr[row][col] + 1) % 10
				if inputs.Arr[row][col] == 0 {
					flashed = append(flashed, inputs.Octopus{R: row, C: col})
				}
			}
		}

		flashes += len(flashed)
		for _, octopus := range flashed {
			flash(octopus)
		}

		/*inputs.Print()
		fmt.Println()*/

		if flashes == 100 {
			fmt.Printf("Synchronicity achieved at step %d", step+1)
			break
		}
	}
}

func flash(point inputs.Octopus) {
	if inputs.Arr[point.R][point.C] != 0 {
		return
	}

	for _, octopus := range inputs.Adjacent(point.R, point.C) {
		if inputs.Arr[octopus.R][octopus.C] == 0 {
			continue
		}

		inputs.Arr[octopus.R][octopus.C] = (inputs.Arr[octopus.R][octopus.C] + 1) % 10
		if inputs.Arr[octopus.R][octopus.C] == 0 {
			flashes++
			flash(octopus)
		}
	}
}
