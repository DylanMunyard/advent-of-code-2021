package main

import (
	"day12/inputs"
	"fmt"
	"strings"
)

var Caves = inputs.Map()

type allowSmall struct {
	name  string
	count int
}

type Path struct {
	visited map[string]inputs.Cave
	path    []string
	small   allowSmall
}

func copyMap(from map[string]inputs.Cave) map[string]inputs.Cave {
	to := map[string]inputs.Cave{}
	for k, v := range from {
		to[k] = v
	}
	return to
}

func copyArray(from []string) []string {
	to := make([]string, len(from))
	copy(to, from)
	return to
}

func StartPath(caveName string, path Path) []Path {
	var paths []Path
	cave := Caves[caveName]
	for _, c := range cave.Connects {
		singleSmallPaths := Path{small: allowSmall{"", 0}, path: copyArray(append(path.path, c)), visited: copyMap(path.visited)}
		singleSmallPaths.visited[c] = Caves[c]
		paths = append(paths, FindPath(c, singleSmallPaths)...)

		if !inputs.IsBig(c) {
			// If small traverse again and allow it to be visited twice
			revisitSmallPaths := Path{small: allowSmall{c, 1}, path: copyArray(append(path.path, c)), visited: copyMap(path.visited)}
			revisitSmallPaths.visited[c] = Caves[c]
			paths = append(paths, FindPath(c, revisitSmallPaths)...)
		}
	}

	return paths
}

func FindPath(caveName string, path Path) []Path {
	cave := Caves[caveName]
	var paths []Path
	for _, c := range cave.Connects {
		if c == "start" {
			continue
		}
		if c == "end" {
			if path.small.name == "" || path.small.count >= 2 {
				path := Path{small: path.small, path: path.path, visited: copyMap(path.visited)}
				path.path = copyArray(append(path.path, c))
				paths = append(paths, path)
			}
			continue
		}

		value, key := path.visited[c]
		if !inputs.IsBig(c) && path.small.name == "" && !key {
			// If small traverse again and allow it to be visited twice
			revisitSmallPaths := Path{small: allowSmall{c, 1}, path: copyArray(path.path), visited: copyMap(path.visited)}
			revisitSmallPaths.path = append(path.path, c)
			revisitSmallPaths.visited[c] = Caves[c]
			paths = append(paths, FindPath(c, revisitSmallPaths)...)
		}

		singleSmallPaths := Path{small: path.small, path: copyArray(path.path), visited: copyMap(path.visited)}
		singleSmallPaths.path = append(singleSmallPaths.path, c)

		if c == singleSmallPaths.small.name && singleSmallPaths.small.count >= 2 {
			continue
		} else if c == singleSmallPaths.small.name && singleSmallPaths.small.count < 2 {
			singleSmallPaths.small.count++
		} else if key && !value.Big {
			continue
		}
		singleSmallPaths.visited[c] = Caves[c]
		paths = append(paths, FindPath(c, singleSmallPaths)...)
	}
	return paths
}

func main() {
	paths := StartPath("start", Path{path: []string{"start"}, visited: map[string]inputs.Cave{"start": Caves["start"]}})

	for i, p := range paths {
		fmt.Printf("%d %s\n", i+1, strings.Join(p.path, ","))
	}
	fmt.Printf("Part 1 = %d\n", len(paths))
}
