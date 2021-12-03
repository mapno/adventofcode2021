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
	f, _ := os.Open("day03/input-0")
	defer f.Close()

	var inputLines []string
	s := bufio.NewScanner(f)
	for s.Scan() {
		inputLines = append(inputLines, s.Text())
	}

	mostFreq := mostFrequentBits(inputLines)

	gamma, epsilon := buildNums(mostFreq)

	fmt.Println(gamma * epsilon)
}

func part2() {
	f, _ := os.Open("day03/input-0")
	defer f.Close()

	var inputLines []string
	s := bufio.NewScanner(f)
	for s.Scan() {
		inputLines = append(inputLines, s.Text())
	}

	mostFreq := mostFrequentBits(inputLines)

	filteredO2 := filterLines(0, mostFreq, inputLines, false)
	filteredCO2 := filterLines(0, mostFreq, inputLines, true)

	O2, _ := buildNums(mostFrequentBits(filteredO2))
	CO2, _ := buildNums(mostFrequentBits(filteredCO2))

	fmt.Println(O2 * CO2)
}

func buildNums(b []int) (int64, int64) {
	var gamma, epsilon string
	for i := 0; i < 12; i++ {
		if b[i] == 0 {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}

	gammaInt, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonInt, _ := strconv.ParseInt(epsilon, 2, 64)
	return gammaInt, epsilonInt
}

func filterLines(idx int, freq []int, lines []string, reverse bool) []string {
	if len(lines) == 1 {
		return lines
	}

	var filtered []string
	for _, l := range lines {
		b, _ := strconv.Atoi(string(l[idx]))
		if (b == freq[idx]) == !reverse {
			filtered = append(filtered, l)
		}
	}

	return filterLines(idx+1, mostFrequentBits(filtered), filtered, reverse)
}

func mostFrequentBits(inputLines []string) []int {
	freq := make([][]int, 12)
	for i := 0; i < 12; i++ {
		freq[i] = make([]int, 2)
	}

	for _, l := range inputLines {
		for i, v := range l {
			b, _ := strconv.Atoi(string(v))
			freq[i][b]++
		}
	}

	mostFreq := make([]int, 12)
	for i := 0; i < 12; i++ {
		if freq[i][0] > freq[i][1] {
			mostFreq[i] = 0
		} else {
			mostFreq[i] = 1
		}
	}
	return mostFreq
}
