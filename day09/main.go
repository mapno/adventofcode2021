package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var sum, size int

func main() {
	part1()
	part2()
}

func part1() {
	grid := getInput()

	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[i]))
	}

	dfsPart1(grid, visited, 0, 0)

	fmt.Println("Part 1:", sum)
}

func dfsPart1(grid [][]int, visited [][]bool, i, j int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || visited[i][j] {
		return
	}

	// check if adjacent points are all greater than current point
	if (i-1 < 0 || grid[i][j] < grid[i-1][j]) &&
		(i+1 >= len(grid) || grid[i][j] < grid[i+1][j]) &&
		(j-1 < 0 || grid[i][j] < grid[i][j-1]) &&
		(j+1 >= len(grid[0]) || grid[i][j] < grid[i][j+1]) {
		sum += grid[i][j] + 1
	}

	visited[i][j] = true
	dfsPart1(grid, visited, i-1, j) // visit up
	dfsPart1(grid, visited, i+1, j) // visit down
	dfsPart1(grid, visited, i, j-1) // visit right
	dfsPart1(grid, visited, i, j+1) // visit left
}

func part2() {
	grid := getInput()

	var basins []int
	for i, row := range grid {
		for j := range row {
			size = 0
			dfsPart2(grid, i, j)
			if size > 0 {
				basins = append(basins, size)
			}
		}
	}

	sort.Ints(basins)
	l := len(basins)

	fmt.Println("Part 2:", basins[l-1]*basins[l-2]*basins[l-3])
}

func dfsPart2(grid [][]int, i, j int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] >= 9 {
		return
	}

	size++
	grid[i][j] = 9
	dfsPart2(grid, i-1, j)
	dfsPart2(grid, i+1, j)
	dfsPart2(grid, i, j-1)
	dfsPart2(grid, i, j+1)
}

func getInput() [][]int {
	f, _ := os.Open("day09/input-0")
	defer f.Close()

	s := bufio.NewScanner(f)
	var grid [][]int
	for s.Scan() {
		var row []int
		for _, c := range s.Text() {
			row = append(row, int(c-'0'))
		}
		grid = append(grid, row)
	}

	return grid
}
