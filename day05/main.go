package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	f, _ := os.Open("day05/input-0")
	defer f.Close()

	s := bufio.NewScanner(f)

	var vents [][2][2]int
	for s.Scan() {
		var vent [2][2]int
		fmt.Fscanf(
			strings.NewReader(s.Text()),
			"%d,%d -> %d,%d",
			&vent[0][0], &vent[0][1], &vent[1][0], &vent[1][1],
		)
		if vent[0][0] > vent[1][0] || vent[0][1] > vent[1][1] {
			vent[0], vent[1] = vent[1], vent[0]
		}
		vents = append(vents, vent)
	}

	var hvVents [][2][2]int
	for _, v := range vents {
		if v[0][0] == v[1][0] || v[0][1] == v[1][1] {
			hvVents = append(hvVents, v)
		}
	}

	grid := make([][]int, 1000)
	for i := range grid {
		grid[i] = make([]int, 1000)
	}

	for _, v := range hvVents {
		for i := v[0][0]; i <= v[1][0]; i++ {
			for j := v[0][1]; j <= v[1][1]; j++ {
				grid[j][i]++
			}
		}
	}

	var total int
	for _, c := range grid {
		for _, v := range c {
			if v > 1 {
				total++
			}
		}
	}

	fmt.Println("Total:", total)
}

func part2() {
	f, _ := os.Open("day05/input-0")
	defer f.Close()

	s := bufio.NewScanner(f)

	var vents [][2][2]int
	for s.Scan() {
		var vent [2][2]int
		fmt.Fscanf(
			strings.NewReader(s.Text()),
			"%d,%d -> %d,%d",
			&vent[0][0], &vent[0][1], &vent[1][0], &vent[1][1],
		)
		if vent[0][0] > vent[1][0] || vent[0][1] > vent[1][1] {
			vent[0], vent[1] = vent[1], vent[0]
		}
		vents = append(vents, vent)
	}

	var hvVents [][2][2]int
	for _, v := range vents {
		if v[0][0] == v[1][0] || v[0][1] == v[1][1] {
			hvVents = append(hvVents, v)
		}
	}

	var diagonalVents [][2][2]int
	for _, v := range vents {
		if (v[1][0] - v[0][0]) == (v[1][1] - v[0][1]) {
			diagonalVents = append(diagonalVents, v)
		}
	}

	var weirdDiagonalVents [][2][2]int
	for _, v := range vents {
		var notWeird bool
		for _, dVent := range diagonalVents {
			if v == dVent {
				notWeird = true
				break
			}
		}
		if math.Abs(float64(v[1][0]-v[0][0])) == math.Abs(float64(v[1][1]-v[0][1])) && !notWeird {
			weirdDiagonalVents = append(weirdDiagonalVents, v)
		}
	}

	grid := make([][]int, 1000)
	for i := range grid {
		grid[i] = make([]int, 1000)
	}

	for _, v := range hvVents {
		for i := v[0][0]; i <= v[1][0]; i++ {
			for j := v[0][1]; j <= v[1][1]; j++ {
				grid[j][i]++
			}
		}
	}

	for _, v := range diagonalVents {
		var i int
		for x := v[0][0]; x <= v[1][0]; x++ {
			j := v[0][1] + i
			grid[j][x]++
			i++
		}
	}

	for _, v := range weirdDiagonalVents {
		if v[0][0] > v[1][0] {
			v[0], v[1] = v[1], v[0]
		}
		var i int
		for x := v[0][0]; x <= v[1][0]; x++ {
			j := v[0][1] - i
			if j < 0 {
				continue
			}
			grid[j][x]++
			i++
		}
	}

	var total int
	for _, c := range grid {
		for _, v := range c {
			if v > 1 {
				total++
			}
		}
	}

	fmt.Println("Total:", total)

}
