package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	f, _ := os.Open("day08/input-0")
	defer f.Close()

	s := bufio.NewScanner(f)

	var inputs [][14]string
	for s.Scan() {
		var input [14]string
		fmt.Fscanf(
			strings.NewReader(s.Text()),
			"%s %s %s %s %s %s %s %s %s %s | %s %s %s %s",
			&input[0], &input[1], &input[2], &input[3], &input[4], &input[5], &input[6],
			&input[7], &input[8], &input[9], &input[10], &input[11], &input[12], &input[13],
		)
		inputs = append(inputs, input)
	}

	var count int
	for _, input := range inputs {
		for i := 10; i < 14; i++ {
			l := len(input[i])
			if l == 2 || l == 3 || l == 4 || l == 7 {
				count++
			}
		}
	}

	fmt.Println(count)
}

func part2() {
	f, _ := os.Open("day08/input-0")
	defer f.Close()

	s := bufio.NewScanner(f)

	var inputs [][2][]string
	for s.Scan() {
		patterns := make([]string, 10)
		outputs := make([]string, 4)
		fmt.Fscanf(
			strings.NewReader(s.Text()),
			"%s %s %s %s %s %s %s %s %s %s | %s %s %s %s",
			&patterns[0], &patterns[1], &patterns[2], &patterns[3], &patterns[4], &patterns[5], &patterns[6], &patterns[7], &patterns[8], &patterns[9],
			&outputs[0], &outputs[1], &outputs[2], &outputs[3],
		)
		sort.Slice(patterns, func(i, j int) bool {
			return len(patterns[i]) < len(patterns[j])
		})
		input := [2][]string{patterns, outputs}
		inputs = append(inputs, input)
	}

	var total int
	for _, input := range inputs {
		// Unique segement numbers
		one, seven, four, eight := sortString(input[0][0]), sortString(input[0][1]), sortString(input[0][2]), sortString(input[0][9])

		// Get 6 segment numbers
		var six, zero, nine string
		for i := 6; i < 9; i++ {
			for _, c := range one {
				if !strings.Contains(input[0][i], string(c)) {
					six = sortString(input[0][i])
					input[0][i] = eight
					break
				}
			}
			for _, c := range four {
				if !strings.Contains(input[0][i], string(c)) {
					zero = sortString(input[0][i])
					input[0][i] = eight
					break
				}
			}
		}
		for i := 6; i < 9; i++ {
			for _, c := range four {
				if !strings.Contains(input[0][i], string(c)) {
					zero = sortString(input[0][i])
					input[0][i] = eight
					break
				}
			}
		}
		for i := 6; i < 9; i++ {
			if input[0][i] != eight {
				nine = sortString(input[0][i])
				break
			}
		}

		// Get 5 segment numbers
		var two, three, five string
		for i := 3; i < 6; i++ {
			var cOne, cFour int
			for _, c := range one {
				if strings.Contains(input[0][i], string(c)) {
					cOne++
				}
			}
			for _, c := range four {
				if strings.Contains(input[0][i], string(c)) {
					cFour++
				}
			}
			if cOne == 2 {
				three = sortString(input[0][i])
			} else if cOne != 2 && cFour == 2 {
				two = sortString(input[0][i])
			} else {
				five = sortString(input[0][i])
			}
		}

		m := map[string]string{one: "1", two: "2", three: "3", four: "4", five: "5", six: "6", seven: "7", eight: "8", nine: "9", zero: "0"}
		var sCount string
		for _, o := range input[1] {
			sCount += m[sortString(o)]
		}
		count, _ := strconv.Atoi(sCount)
		total += count
	}
	fmt.Println(total)
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func sortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}
