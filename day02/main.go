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
	f, _ := os.Open("day02/input-0")

	defer f.Close()

	var v, h int
	s := bufio.NewScanner(f)
	for s.Scan() {
		t := s.Text()
		switch {
		case strings.Contains(t, "forward"):
			n, _ := strconv.Atoi(t[len(t)-1:])
			h += n
		case strings.Contains(t, "up"):
			n, _ := strconv.Atoi(t[len(t)-1:])
			v -= n
		case strings.Contains(t, "down"):
			n, _ := strconv.Atoi(t[len(t)-1:])
			v += n

		}
	}
	fmt.Println(v*h)

}

func part2() {
	f, _ := os.Open("day02/input-0")

	defer f.Close()

	var v, h, aim int
	s := bufio.NewScanner(f)
	for s.Scan() {
		t := s.Text()
		switch {
		case strings.Contains(t, "forward"):
			n, _ := strconv.Atoi(t[len(t)-1:])
			h += n
			v += aim * n
		case strings.Contains(t, "up"):
			n, _ := strconv.Atoi(t[len(t)-1:])
			aim -= n
		case strings.Contains(t, "down"):
			n, _ := strconv.Atoi(t[len(t)-1:])
			aim += n

		}
	}
	fmt.Println(v*h)
}
