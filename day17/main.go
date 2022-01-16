package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Coordinate struct {
	x int64
	y int64
}

type Velocity Coordinate

type TargetArea struct {
	tl Coordinate // top left
	tr Coordinate // top right
	bl Coordinate // bottom left
	br Coordinate // bottom right
}

func (t TargetArea) String() string {
	return fmt.Sprintf("(%v,%v)->(%v,%v) (%v,%v)->(%v,%v)", t.tl.x, t.tl.y, t.tr.x, t.tr.y, t.bl.x, t.bl.y, t.br.x, t.br.y)
}

func (t TargetArea) Within(c Coordinate) bool {
	if t.tl.x > c.x {
		return false
	}

	if t.tr.x < c.x {
		return false
	}

	if t.tl.y < c.y {
		return false
	}

	if t.bl.y > c.y {
		return false
	}

	return true
}

func (t TargetArea) WithinX(x int64) bool {
	if t.tl.x > x {
		return false
	}

	if t.tr.x < x {
		return false
	}

	return true
}

func (t TargetArea) WithinY(y int64) bool {
	if t.tl.y < y {
		return false
	}

	if t.bl.y > y {
		return false
	}

	return true
}

func (t TargetArea) OvershotY(y int64) bool {
	if t.bl.y > y {
		return true
	}

	return false
}

type WithinBounds func(c int64) bool

func (t TargetArea) Overshot(c Coordinate) bool {
	if t.tr.x < c.x {
		return true
	}

	if t.br.y > c.y {
		return true
	}

	return false
}

func CreateTarget(input string) TargetArea {
	inputs := strings.Split(input, ", ")
	xs := strings.Split(strings.Replace(inputs[0], "x=", "", 1), "..")
	ys := strings.Split(strings.Replace(inputs[1], "y=", "", 1), "..")
	x0, _ := strconv.ParseInt(xs[0], 10, 64)
	x1, _ := strconv.ParseInt(xs[1], 10, 64)
	y0, _ := strconv.ParseInt(ys[1], 10, 64)
	y1, _ := strconv.ParseInt(ys[0], 10, 64)

	return TargetArea{tl: Coordinate{x: x0, y: y0}, tr: Coordinate{x: x1, y: y0}, bl: Coordinate{x: x0, y: y1}, br: Coordinate{x: x1, y: y1}}
}

func step(c Coordinate, v Velocity, moves int) (Coordinate, Velocity) {
	newC := Coordinate{x: c.x + v.x, y: c.y + v.y}

	if v.x > 0 {
		v.x = v.x - 1
	}

	if v.x < 0 {
		v.x = v.x + 1
	}

	v.y = v.y - 1 /* -1 is darn gravity */

	fmt.Printf("#%v: (%v,%v)->(%v,%v) v: (%v,%v)\n", moves, c.x, c.y, newC.x, newC.y, v.x, v.y)
	return newC, v
}

func main() {
	/*t := CreateTarget("x=20..30, y=-10..-5")*/
	t := CreateTarget("x=111..161, y=-154..-101")

	Part1(t)
	fmt.Println()
	Part2(t)
}

func Part1(t TargetArea) {
	// A positive velocity will reach its max point at v! additive: (n*n + n) / 2 + 1 step for when it reaches 0.
	// A negative velocity reaches y = 0 at v!. It takes 2xv + 1 steps to reach y = 0,  at which point v = -starting v
	// e.g. for v = 9, it takes 19 steps to reach y = 0, and v will be v = -9
	// Therefore the lower right bound of y(r) will be overshot after v = y(r) because taking one step
	// after y = 0, will mean v = -v - 1, which overshoots y(r) by 1

	// The highest point of y is v! additive
	velocity := int64(math.Abs(float64(t.br.y))) - 1
	fmt.Printf("y: %v", factorial(1, velocity))
}

func Part2(t TargetArea) {
	var verticalX []int64                 // tracks how many x velocities result in vertical trajectories through the target area
	stepsCount := make(map[int64][]int64) // map of step counts, to their respective x velocities, eg. {1} = [20, 21, 23]
	var maxStep int64 = 0                 // keep track of the maximum steps an x velocity takes to reach the target area
	for x := t.tr.x; x >= 0; x-- {
		hit, steps := xStep(x, t.WithinX)
		if !hit {
			continue
		}

		max := steps[len(steps)-1]
		if max == x { // means x velocity is at zero within the target area
			verticalX = append(verticalX, x)
		}

		if maxStep < max {
			maxStep = max
		}

		for _, step := range steps {
			stepsCount[step] = append(stepsCount[step], x)
		}
	}

	var initialVelocities int64 = 0
	for y := int64(math.Abs(float64(t.br.y))) - 1; y >= t.br.y; y-- {
		hit, steps := yStep(0, y, t)
		if !hit {
			continue
		}

		seenX := make(map[int64]bool)
		for _, step := range steps {
			if xs, count := stepsCount[step]; count { // there is a matching x velocity
				for _, x := range xs {
					if _, seen := seenX[x]; seen { // don't double count x,y pairs
						continue
					}
					seenX[x] = true
					initialVelocities += 1
				}
			}

			for _, x := range verticalX { // check for x velocities that are at 0 in the target area
				if step < x { // they match any value of y greater or equal to the 0 x velocity
					continue
				}
				if _, seen := seenX[x]; seen { // don't double count x,y pairs
					continue
				}
				seenX[x] = true
				initialVelocities += 1
			}
		}
	}

	fmt.Printf("%v different initial velocities\n", initialVelocities)
}

// Perform x calculations incrementally, record any that fall within the x bounds of the target area
func xStep(n int64, boundsCheck WithinBounds) (bool, []int64) {
	var f int64 = 0
	var hits []int64
	i := n
	for {
		f += i
		if boundsCheck(f) {
			hits = append(hits, n-(i-1))
		}

		if n < 0 {
			i++ // go up
			if i > -1 {
				break
			}
		} else {
			i-- // go down
			if i < 1 {
				break
			}
		}

	}

	return len(hits) > 0, hits
}

// Perform y calculations incrementally, record any that fall within the y bounds of the target area
func yStep(y, v int64, t TargetArea) (bool, []int64) {
	var hits []int64
	var step int64 = 1
	for {
		y += v

		if t.WithinY(y) {
			hits = append(hits, step)
		}

		if t.OvershotY(y) {
			break
		}

		v-- // velocity decreases
		step++
	}

	return len(hits) > 0, hits
}

func factorial(k, n int64) int64 {
	if k == n {
		return n
	}

	return k + factorial(k+1, n)
}
