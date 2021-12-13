package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	X = 120
	Y = 121
)

func main() {
	f, _ := os.Open("day13/input-0")
	defer f.Close()

	grid := make([][]bool, 895)
	for i := range grid {
		grid[i] = make([]bool, 1308)
	}

	s := bufio.NewScanner(f)
	for s.Scan() {
		if s.Text() == "" {
			break
		}
		var x, y int
		fmt.Fscanf(strings.NewReader(s.Text()), "%d,%d", &x, &y)
		grid[y][x] = true
	}

	var folds []string
	for s.Scan() {
		l := s.Text()
		fold := l[11:]
		folds = append(folds, fold)
	}

	for _, f := range folds {
		grid = fold(grid, f)
	}

	var c int
	for _, r := range grid {
		for _, v := range r {
			if v {
				c++
			}
		}
	}

	for _, r := range grid {
		fmt.Println(r)
	}
}

func fold(g [][]bool, fold string) [][]bool {
	f, _ := strconv.Atoi(fold[2:])
	var ng [][]bool
	switch fold[0] {
	case X:
		ng = hGrid(g, f)
		hFold(g, ng, f)
	case Y:
		ng = vGrid(g, f)
		vFold(g, ng, f)
	}

	return ng
}

func vGrid(g [][]bool, f int) [][]bool {
	newGrid := make([][]bool, f)
	for i := range newGrid {
		newGrid[i] = make([]bool, len(g[i]))
	}
	return newGrid
}

func hGrid(g [][]bool, f int) [][]bool {
	newGrid := make([][]bool, len(g))
	for i := range newGrid {
		newGrid[i] = make([]bool, f)
	}
	return newGrid
}

func vFold(g, newGrid [][]bool, f int) {
	for y := 0; y < f; y++ {
		for x := 0; x < len(newGrid[y]); x++ {
			if 2*f-y < len(g) {
				newGrid[y][x] = g[y][x] || g[2*f-y][x]
			} else {
				newGrid[y][x] = g[y][x]
			}
		}
	}
}

func hFold(g, newGrid [][]bool, f int) {
	for y := 0; y < len(newGrid); y++ {
		for x := 0; x < f; x++ {
			if 2*f-x < len(g[y]) {
				newGrid[y][x] = g[y][x] || g[y][2*f-x]
			} else {
				newGrid[y][x] = g[y][x]
			}
		}
	}
}
