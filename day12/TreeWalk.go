package main

import (
	"fmt"
	"golang.org/x/tour/tree"
	"strconv"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int, top bool) {
	if t.Left != nil {
		Walk(t.Left, ch, false)
	}

	ch <- t.Value

	if t.Right != nil {
		Walk(t.Right, ch, false)
	}

	if top {
		close(ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	st1 := ""
	st2 := ""

	ch1 := make(chan int)
	go Walk(t1, ch1, true)

	ch2 := make(chan int)
	go Walk(t2, ch2, true)

	for value := range ch1 {
		st1 += strconv.Itoa(value)
	}

	for value := range ch2 {
		st2 += strconv.Itoa(value)
	}

	return st1 == st2
}

func main() {
	fmt.Printf("%t\n", Same(tree.New(1), tree.New(1)))
	fmt.Printf("%t\n", Same(tree.New(1), tree.New(2)))
}
