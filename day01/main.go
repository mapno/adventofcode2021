package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	part1()
	part2()
}

func part1() {
	f, _ := os.Open("day01/input-0")

	defer f.Close()

	var prev, c int
	s := bufio.NewScanner(f)
	for s.Scan() {
		d, _ := strconv.Atoi(s.Text())
		if d > prev && prev != 0 {
			c++
		}
		prev = d
	}
	fmt.Println(c)
}

func part2() {
	f, _ := os.Open("day01/input-0")
	defer f.Close()

	var prev1, prev2, prev3, c int
	s := bufio.NewScanner(f)
	for s.Scan() {
		d, _ := strconv.Atoi(s.Text())

		if prev1 != 0 && prev1+prev2+prev3 < prev2+prev3+d {
			c++
		}

		prev1 = prev2
		prev2 = prev3
		prev3 = d
	}
	fmt.Println(c)
}
