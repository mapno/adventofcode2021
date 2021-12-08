package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	f, _ := os.Open("day07/input-0")
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Scan()
	inputStr := strings.Split(s.Text(), ",")

	var pos []int
	for _, pStr := range inputStr {
		num, _ := strconv.Atoi(pStr)
		pos = append(pos, num)
	}

	min := math.MaxInt64
	for i := 0; i < len(pos); i++ {
		var count int
		for _, p := range pos {
			count += int(math.Abs(float64(p - i)))
		}
		if count < min {
			min = count
		}
	}

	fmt.Println(min)
}

func part2() {
	f, _ := os.Open("day07/input-0")
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Scan()
	inputStr := strings.Split(s.Text(), ",")

	var pos []int
	for _, pStr := range inputStr {
		num, _ := strconv.Atoi(pStr)
		pos = append(pos, num)
	}

	min := math.MaxInt64
	for i := 0; i < len(pos); i++ {
		var count int
		for _, p := range pos {
			dist := int(math.Abs(float64(p - i)))
			for j := 0; j < dist; j++ {
				count += 1 + j
			}
		}
		if count < min {
			min = count
		}
	}

	fmt.Println(min)
}
