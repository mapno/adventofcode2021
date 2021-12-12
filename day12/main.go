package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	routes int
	nodes  map[string]int
	adj    map[string][]string
)

func main() {
	part1()
	part2()
}

func part1() {
	routes = 0
	getInput()

	dfs("start", func(n string) bool {
		return n[0] > 96 && nodes[n] == 1
	})

	fmt.Println("Part 1:", routes)
}

func part2() {
	routes = 0
	getInput()

	dfs("start", func(n string) bool {
		if n == "start" && nodes[n] == 1 {
			return true
		}

		if n[0] > 96 {
			for c, v := range nodes {
				if c[0] > 96 && v == 2 && nodes[n] == 1 {
					return true
				}
			}
		}

		return n[0] > 96 && nodes[n] == 2
	})

	fmt.Println("Part 2:", routes)
}

func dfs(n string, shouldReturn func(n string) bool) {
	if n == "end" {
		routes++
		return
	}

	if shouldReturn(n) {
		return
	}

	nodes[n]++
	for _, v := range adj[n] {
		dfs(v, shouldReturn)
	}
	nodes[n]--
}

func getInput() {
	f, _ := os.Open("day12/input-0")
	defer f.Close()

	nodes = make(map[string]int)
	adj = make(map[string][]string)

	s := bufio.NewScanner(f)
	for s.Scan() {
		s := strings.Split(s.Text(), "-")
		nodes[s[0]], nodes[s[1]] = 0, 0
		adj[s[0]] = append(adj[s[0]], s[1])
		adj[s[1]] = append(adj[s[1]], s[0])
	}
}
