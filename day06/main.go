package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	f, _ := os.Open("day06/input-0")
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Scan()
	sNums := strings.Split(s.Text(), ",")

	var lanterns []int
	for _, sNum := range sNums {
		num, _ := strconv.Atoi(sNum)
		lanterns = append(lanterns, num)
	}

	for days := 0; days < 80; days++ {
		for i := range lanterns {
			lanterns[i]--
			if lanterns[i] < 0 {
				lanterns = append(lanterns, 8)
				lanterns[i] = 6
			}

		}
	}

	fmt.Println("Total:", len(lanterns))
}

func part2() {
	f, _ := os.Open("day06/input-0")
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Scan()
	sNums := strings.Split(s.Text(), ",")

	lanterns := make([]int, 9)
	for _, sNum := range sNums {
		num, _ := strconv.Atoi(sNum)
		lanterns[num]++
	}

	for days := 0; days < 256; days++ {
		new := make([]int, 9)
		for i := 8; i > 0; i-- {
			new[i-1] = lanterns[i]
		}
		new[8] = lanterns[0]
		new[6] += new[8]
		lanterns = new
	}

	var total int
	for _, l := range lanterns {
		total += l
	}

	fmt.Println("Total:", total)
}
