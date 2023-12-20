package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, _ := os.Open("input.txt")
	input, _ := io.ReadAll(f)
	p1 := Part01(input)
	p2 := Part02(input)
	fmt.Println("Part1 Answer:", p1)
	fmt.Println("Part2 Answer:", p2)
}

func findSymLine(grid []string) int {
lineLoop:
	for i := 1; i < len(grid); i++ {
		n := min(i, len(grid)-i)
		for j := 0; j < n; j++ {
			if grid[i-j-1] != grid[i+j] {
				continue lineLoop
			}
		}
		return i
	}
	return 0
}

func transpose(grid []string) []string {
	t := make([]string, len(grid[0]))
	for _, s := range grid {
		for j := range s {
			t[j] += s[j : j+1]
		}
	}
	return t
}

func Part01(input []byte) int {
	lines := strings.Split(strings.TrimSpace(string(input))+"\n", "\n")
	var grid []string
	sum := 0
	for _, line := range lines {
		if line == "" {
			i := findSymLine(grid)
			if i > 0 {
				sum += 100 * i
			}
			sum += findSymLine(transpose(grid))
			grid = grid[:0]
			continue
		}
		grid = append(grid, line)
	}
	return sum
}

func count2Smudges(s, t string) int {
	n := 0
	for i := 0; i < len(s); i++ {
		if s[i] != t[i] {
			n++
			if n > 1 {
				return n
			}
		}
	}
	return n
}

func findSmudgeSymLine(grid []string) int {
lineLoop:
	for i := 1; i < len(grid); i++ {
		n := min(i, len(grid)-i)
		smudges := 0
		for j := 0; j < n; j++ {
			smudges += count2Smudges(grid[i-j-1], grid[i+j])
			if smudges > 1 {
				continue lineLoop
			}
		}
		if smudges == 1 {
			return i
		}
	}
	return 0
}

func Part02(input []byte) int {
	lines := strings.Split(strings.TrimSpace(string(input))+"\n", "\n")
	var grid []string
	sum := 0
	for _, line := range lines {
		if line == "" {
			i := findSmudgeSymLine(grid)
			if i > 0 {
				sum += 100 * i
			}
			sum += findSmudgeSymLine(transpose(grid))
			grid = grid[:0]
			continue
		}
		grid = append(grid, line)
	}
	return sum
}
