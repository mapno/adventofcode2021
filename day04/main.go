package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	nums, boards := getInput()
	part1(nums, boards)
	part2(nums, boards)
}

func part1(nums []int, boards [][][]int) {
	winner, lastNum := winnerLoop(nums, boards)
	sum := unmarkedSum(boards[winner])

	fmt.Println("Part 1:", sum*lastNum)
}

func part2(nums []int, boards [][][]int) {
	// Eliminate all boards but one
	for _, n := range nums {
		if len(boards) == 1 {
			break
		}
		markAll(boards, n)

		for { // Remove all boards that are completed in this loop
			if complete, n := checkAll(boards); complete {
				boards = append(boards[:n], boards[n+1:]...)
				continue
			}
			break
		}
	}

	_, lastNum := winnerLoop(nums, boards)
	sum := unmarkedSum(boards[0])

	fmt.Println("Part 2:", sum*lastNum)
}

func getInput() ([]int, [][][]int) {
	f, _ := os.Open("day04/input-0")
	defer f.Close()

	s := bufio.NewScanner(f)

	var nums []int
	s.Scan()
	sNums := strings.Split(s.Text(), ",")
	for _, sNum := range sNums {
		num, _ := strconv.Atoi(sNum)
		nums = append(nums, num)
	}

	var boards [][][]int
	for s.Scan() {
		board := make([][]int, 5)
		for i := 0; i < 5; i++ {
			board[i] = make([]int, 5)
		}

		for i := 0; i < 5; i++ {
			s.Scan()
			fmt.Fscanf(
				strings.NewReader(s.Text()),
				"%d %d %d %d %d",
				&board[i][0], &board[i][1], &board[i][2], &board[i][3], &board[i][4])

		}
		boards = append(boards, board)
	}
	return nums, boards
}

func winnerLoop(nums []int, boards [][][]int) (int, int) {
	var winner, lastNum int
	for _, n := range nums {
		lastNum = n
		markAll(boards, n)
		if complete, n := checkAll(boards); complete {
			winner = n
			break
		}
	}
	return winner, lastNum
}

func unmarkedSum(board [][]int) int {
	var sum int
	for _, c := range board {
		for _, n := range c {
			if n != 100 {
				sum += n
			}
		}
	}
	return sum
}

func markAll(boards [][][]int, n int) {
	for _, b := range boards {
		mark(b, n)
	}
}

func mark(board [][]int, n int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if board[i][j] == n {
				board[i][j] = 100
			}
		}
	}
}

func checkAll(boards [][][]int) (bool, int) {
	for i, b := range boards {
		if check(b) {
			return true, i
		}
	}
	return false, 0
}

func check(board [][]int) bool {
	// check rows
	for i := 0; i < 5; i++ {
		if board[i][0] == 100 && board[i][1] == 100 && board[i][2] == 100 && board[i][3] == 100 && board[i][4] == 100 {
			return true
		}
	}
	// check cols
	for j := 0; j < 5; j++ {
		if board[0][j] == 100 && board[1][j] == 100 && board[2][j] == 100 && board[3][j] == 100 && board[4][j] == 100 {
			return true
		}
	}
	return false
}
