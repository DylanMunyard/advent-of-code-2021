package main

import (
	"day14/inputs"
	"fmt"
	"sort"
	"strings"
)

// ApplyRules makes insertions between pairs according to the rules
func ApplyRules(pairs map[string]int, elements map[string]int, ch chan map[string]int) {
	newPairs := make(map[string]int)
	for s, i := range pairs {
		newPairs[s] = i
	}
	// pairs map counts instances of the pair in the template
	// e.g. NCNBCHB -> {"NC": 1, "CN": 1, "NB": 1, "BC": 1, "CH": 1, "HB": 1}
	for p, instances := range pairs {
		if rule, i := inputs.Rules[p]; i {
			// We have a rule insertion between this pair
			if _, k := elements[rule]; k {
				elements[rule] = elements[rule] + instances // increase element count by the pair count
			} else {
				elements[rule] = 1
			}

			// given rule insertion I split the pair AB into AI BI
			pairElements := strings.Split(p, "")
			pair1 := fmt.Sprintf("%s%s", pairElements[0], rule)
			pair2 := fmt.Sprintf("%s%s", rule, pairElements[1])

			// Splitting a pair will mean we have <pair instance count> of the new pairs
			pairInstances := instances
			if existingPair, hasNewPair := newPairs[pair1]; hasNewPair {
				pairInstances = existingPair + instances
			}
			newPairs[pair1] = pairInstances

			pairInstances = instances
			if existingPair, hasNewPair := newPairs[pair2]; hasNewPair {
				pairInstances = existingPair + instances
			}
			newPairs[pair2] = pairInstances
			// decrease # instances of the pair by its original count
			// leaves only the pairs that were split in this round
			newPairs[p] = newPairs[p] - instances
		}
	}

	ch <- newPairs
}

func main() {
	// split the string into pairs
	input := strings.Split(inputs.Inputs, "")
	num := len(input)

	pairs := make(map[string]int)    // pair and their count
	elements := make(map[string]int) // count of elements

	for i, e := range input {
		// Take the starting template and create counts of individual elements
		if _, k := elements[e]; k {
			elements[e] = elements[e] + 1
		} else {
			elements[e] = 1
		}

		if i+1 == num {
			continue
		}

		// Take the starting template and create counts of pairs
		pair := fmt.Sprintf("%s%s", e, input[i+1])
		if _, k := pairs[pair]; k {
			pairs[pair] = pairs[pair] + 1
		} else {
			pairs[pair] = 1
		}
	}

	totalPairs := 0
	for _, i := range pairs {
		totalPairs += i
	}
	fmt.Printf("Total: %d\n", totalPairs)

	ch := make(chan map[string]int)
	for _ = range make([]int, 40) {
		go ApplyRules(pairs, elements, ch)
		newPairs := <-ch
		pairs = make(map[string]int)
		for s, i := range newPairs {
			if i <= 0 {
				continue
			}
			pairs[s] = i
		}
	}

	elementCount := make([]int, len(elements))
	elementTick := 0
	for _, i := range elements {
		elementCount[elementTick] = i
		elementTick = elementTick + 1
	}

	sort.Ints(elementCount)
	fmt.Printf("%d - %d = %d", elementCount[len(elementCount)-1], elementCount[0], elementCount[len(elementCount)-1]-elementCount[0])
}
