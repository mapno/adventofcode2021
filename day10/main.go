package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	part1()
	part2()
}

func part1() {
	f, _ := os.Open("day10/input-0")
	defer f.Close()

	s := bufio.NewScanner(f)
	var lines []string
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	values := map[rune]int{
		' ': 0,
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}

	var sum int
	for _, line := range lines {
		corrupted := findCorrupted(line)
		sum += values[corrupted]
	}

	fmt.Println("Part 1:", sum)
}

func findCorrupted(l string) rune {
	stack := make([]rune, 0)
	for _, c := range l {
		switch c {
		case '<':
			stack = append(stack, '>')
		case '[':
			stack = append(stack, ']')
		case '(':
			stack = append(stack, ')')
		case '{':
			stack = append(stack, '}')
		default:
			if len(stack) > 0 && stack[len(stack)-1] == c {
				stack = stack[:len(stack)-1]
			} else {
				return c
			}
		}
	}
	return ' '
}

func part2() {
	f, _ := os.Open("day10/input-0")
	defer f.Close()

	s := bufio.NewScanner(f)
	var lines []string
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	values := map[rune]int{
		' ': 0,
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}

	var scores []int
	for _, line := range lines {
		remaining := findIncomplete(line)
		var sum int
		for i := len(remaining) - 1; i >= 0; i-- {
			sum *= 5
			sum += values[remaining[i]]
		}
		if sum > 0 {
			scores = append(scores, sum)
		}
	}

	sort.Ints(scores)

	l := len(scores)
	fmt.Println("Part 2:", scores[l/2])
}

func findIncomplete(l string) []rune {
	stack := make([]rune, 0)
	for _, c := range l {
		switch c {
		case '<':
			stack = append(stack, '>')
		case '[':
			stack = append(stack, ']')
		case '(':
			stack = append(stack, ')')
		case '{':
			stack = append(stack, '}')
		default:
			if len(stack) > 0 && stack[len(stack)-1] == c {
				stack = stack[:len(stack)-1]
			} else {
				return make([]rune, 0)
			}
		}
	}
	return stack
}
