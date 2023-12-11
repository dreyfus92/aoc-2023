package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	p1 := Part01()
	p2 := Part02()
	fmt.Println("Part1 Answer:", p1)
	fmt.Println("Part2 Answer:", p2)
}

func parseLine(line string) []int {
	var nums []int
	for _, numStr := range strings.Fields(line) {
		var num int
		fmt.Sscanf(numStr, "%d", &num)
		nums = append(nums, num)
	}
	return nums
}

func nextValueInHistory(history []int) int {
	for {
		differences := make([]int, len(history)-1)
		allZeroes := true

		for i := 0; i < len(history)-1; i++ {
			differences[i] = history[i+1] - history[i]
			if differences[i] != 0 {
				allZeroes = false
			}
		}

		if allZeroes {
			return history[len(history)-1] + differences[0]
		}

		history = differences
	}
}

// sumOfExtrapolations calculates the sum of extrapolated values for each history.
func sumOfExtrapolations(histories [][]int) int {
	sum := 0
	for _, history := range histories {
		sum += nextValueInHistory(history)
	}
	return sum
}

func Part01() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var histories [][]int
	for scanner.Scan() {
		line := scanner.Text()
		histories = append(histories, parseLine(line))
	}

	return sumOfExtrapolations(histories)
}

func Part02() int {
	return 0
}
