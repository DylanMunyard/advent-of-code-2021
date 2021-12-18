package inputs

import "fmt"

var Arr = [][]int{
	{6, 3, 1, 8, 1, 8, 5, 7, 3, 2},
	{1, 1, 2, 2, 6, 8, 7, 1, 3, 5},
	{5, 1, 7, 3, 2, 3, 7, 6, 7, 6},
	{8, 7, 5, 4, 3, 6, 2, 6, 1, 2},
	{5, 7, 1, 8, 4, 7, 4, 6, 6, 6},
	{8, 4, 4, 3, 6, 5, 4, 1, 3, 7},
	{1, 2, 4, 7, 6, 3, 4, 3, 4, 6},
	{1, 4, 4, 6, 5, 1, 4, 5, 8, 5},
	{6, 7, 1, 7, 2, 8, 8, 2, 6, 7},
	{1, 7, 2, 7, 8, 7, 1, 2, 2, 8}}

type Octopus struct {
	R int
	C int
}

func Print() {
	for _, elements := range Arr {
		for _, element := range elements {
			fmt.Printf("%d ", element)
		}
		fmt.Println()
	}
}

func Adjacent(r, c int) []Octopus {
	var octopi []Octopus
	if c+1 < len(Arr[r]) {
		octopi = append(octopi, Octopus{r, c + 1})

		if r-1 >= 0 {
			octopi = append(octopi, Octopus{r - 1, c + 1})
		}
		if r+1 < len(Arr) {
			octopi = append(octopi, Octopus{r + 1, c + 1})
		}
	}

	if c-1 >= 0 {
		octopi = append(octopi, Octopus{r, c - 1})

		if r-1 >= 0 {
			octopi = append(octopi, Octopus{r - 1, c - 1})
		}
		if r+1 < len(Arr) {
			octopi = append(octopi, Octopus{r + 1, c - 1})
		}
	}

	if r-1 >= 0 {
		octopi = append(octopi, Octopus{r - 1, c})
	}

	if r+1 < len(Arr) {
		octopi = append(octopi, Octopus{r + 1, c})
	}

	return octopi
}
