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
	f, _ := os.Open("day09/input-0")
	defer f.Close()

	s := bufio.NewScanner(f)
	var grid [][]int
	var visited [][]bool
	for s.Scan() {
		var row []int
		for _, c := range s.Text() {
			row = append(row, int(c-'0'))
		}
		grid = append(grid, row)
		visited = append(visited, make([]bool, len(row)))
	}

	for i, row := range grid {
		for j := range row {
			dfs1(grid, visited, i, j)
		}
	}

	fmt.Println(sum)

}

func dfs1(grid [][]int, visited [][]bool, i, j int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || visited[i][j] {
		return
	}

	if (i-1 < 0 || grid[i][j] < grid[i-1][j]) &&
		(i+1 >= len(grid) || grid[i][j] < grid[i+1][j]) &&
		(j-1 < 0 || grid[i][j] < grid[i][j-1]) &&
		(j+1 >= len(grid[0]) || grid[i][j] < grid[i][j+1]) {
		sum += grid[i][j] + 1
	}
	visited[i][j] = true
	dfs1(grid, visited, i-1, j)
	dfs1(grid, visited, i+1, j)
	dfs1(grid, visited, i, j-1)
	dfs1(grid, visited, i, j+1)
}

func part2() {
	f, _ := os.Open("day09/input-0")
	defer f.Close()

	s := bufio.NewScanner(f)
	var grid [][]int
	var visited [][]bool
	for s.Scan() {
		var row []int
		for _, c := range s.Text() {
			row = append(row, int(c-'0'))
		}
		grid = append(grid, row)
		visited = append(visited, make([]bool, len(row)))
	}

	var basins []int
	for i, row := range grid {
		for j := range row {
			size = 0
			dfs2(grid, i, j)
			if size > 0 {
				basins = append(basins, size)
			}
		}
	}

	sort.Ints(basins)
	l := len(basins)

	fmt.Println(basins[l-1] * basins[l-2] * basins[l-3])
}

func dfs2(grid [][]int, i, j int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] >= 9 {
		return
	}

	size++
	grid[i][j] = 9
	dfs2(grid, i-1, j)
	dfs2(grid, i+1, j)
	dfs2(grid, i, j-1)
	dfs2(grid, i, j+1)
}
