package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	part1()
	part2()
}

func part1() {
	f, _ := os.Open("day11/input-0")
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

	var c int
	for it := 0; it < 100; it++ {
		for _, r := range grid {
			for j := range r {
				r[j]++
			}
		}
		flashed := make([][]bool, len(grid))
		for i := range grid {
			flashed[i] = make([]bool, len(grid[i]))
		}

		for {
			exit := true
			for i, r := range grid {
				for j := range r {
					if r[j] > 9 {
						flash(grid, i, j, flashed)
						c++
						exit = false
					}
				}
			}

			if exit {
				break
			}
		}
	}

	fmt.Println("Part 1:", c)
}

func part2() {
	f, _ := os.Open("day11/input-0")
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

	for it := 1; ; it++ {
		for _, r := range grid {
			for j := range r {
				r[j]++
			}
		}
		flashed := make([][]bool, len(grid))
		for i := range grid {
			flashed[i] = make([]bool, len(grid[i]))
		}

		for {
			exit := true
			for i, r := range grid {
				for j := range r {
					if r[j] > 9 {
						flash(grid, i, j, flashed)
						exit = false
					}
				}
			}

			var n int
			for _, r := range grid {
				for _, c := range r {
					if c == 0 {
						n++
					}
				}
			}

			if n == (len(grid) * len(grid[0])) {
				fmt.Println("Part 2:", it)
				return
			}

			if exit {
				break
			}
		}
	}
}

func flash(grid [][]int, i, j int, f [][]bool) {
	grid[i][j] = 0
	f[i][j] = true
	if i > 0 {
		if !f[i-1][j] {
			grid[i-1][j]++
		}
	}
	if i < len(grid)-1 {
		if !f[i+1][j] {
			grid[i+1][j]++
		}
	}
	if j > 0 {
		if !f[i][j-1] {
			grid[i][j-1]++
		}
	}
	if j < len(grid[i])-1 {
		if !f[i][j+1] {
			grid[i][j+1]++
		}
	}
	if i > 0 && j > 0 {
		if !f[i-1][j-1] {
			grid[i-1][j-1]++
		}
	}
	if i > 0 && j < len(grid[i])-1 {
		if !f[i-1][j+1] {
			grid[i-1][j+1]++
		}
	}
	if i < len(grid)-1 && j > 0 {
		if !f[i+1][j-1] {
			grid[i+1][j-1]++
		}
	}
	if i < len(grid)-1 && j < len(grid[i])-1 {
		if !f[i+1][j+1] {
			grid[i+1][j+1]++
		}
	}
}
