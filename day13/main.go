package main

import (
	"day13/inputs"
	"fmt"
)

func main() {
	instructions := inputs.FoldInstructions
	dots := inputs.Inputs
	var max inputs.Dot

	for i, instruction := range instructions {
		split := inputs.AsMap(instruction.Horizontal, instruction.FoldAlong, dots)
		fmt.Printf("(horiz=%t,fold=%d) lower=%d,higher=%d\n", instruction.Horizontal, instruction.FoldAlong, len(split.Lower), len(split.Higher))

		for _, dot := range split.Higher {
			foldedDot := inputs.Dot{X: dot.X, Y: instruction.FoldAlong*2 - dot.Y}
			if !instruction.Horizontal {
				foldedDot.X = instruction.FoldAlong*2 - dot.X
				foldedDot.Y = dot.Y
			}

			if _, k := split.Lower[inputs.DotToString(&foldedDot)]; !k {
				split.Lower[inputs.DotToString(&foldedDot)] = foldedDot
			}
		}

		fmt.Printf("#%d = lower=%d\n", i+1, len(split.Lower))
		dots, max = inputs.LowerToDots(split.Lower)

		for y := range make([]int, max.Y+1) {
			for x := range make([]int, max.X+1) {
				printDot := inputs.Dot{X: x, Y: y}
				if _, k := split.Lower[inputs.DotToString(&printDot)]; k {
					fmt.Printf("#")
				} else {
					fmt.Printf(".")
				}
			}
			fmt.Println()
		}
	}
}
