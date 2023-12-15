package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {

	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	p1 := Part01(input)
	p2 := Part02(input)
	fmt.Println("Part1 Answer:", p1)
	fmt.Println("Part2 Answer:", p2)
}

var grid [][]byte
var visited []bool = make([]bool, 20000)
var distances []int = make([]int, 20000)

func xyToIndex(x, y int) int {
	return y*len(grid[0]) + x
}

func dfs(x, y, d int) {
	if visited[xyToIndex(x, y)] {
		return
	}

	visited[xyToIndex(x, y)] = true
	distances[xyToIndex(x, y)] = d

	if grid[y][x] == '-' {
		if x > 0 {
			dfs(x-1, y, d+1)
		}

		if x < len(grid[y])-1 {
			dfs(x+1, y, d+1)
		}
	}

	if grid[y][x] == '|' {
		if y > 0 {
			dfs(x, y-1, d+1)
		}

		if y < len(grid)-1 {
			dfs(x, y+1, d+1)
		}
	}

	if grid[y][x] == 'L' {
		if x < len(grid[y])-1 {
			dfs(x+1, y, d+1)
		}

		if y > 0 {
			dfs(x, y-1, d+1)
		}
	}

	if grid[y][x] == 'J' {
		if x > 0 {
			dfs(x-1, y, d+1)
		}

		if y > 0 {
			dfs(x, y-1, d+1)
		}

	}

	if grid[y][x] == '7' {
		if x > 0 {
			dfs(x-1, y, d+1)
		}

		if y < len(grid)-1 {
			dfs(x, y+1, d+1)
		}

	}

	if grid[y][x] == 'F' {
		if x < len(grid[y])-1 {
			dfs(x+1, y, d+1)
		}

		if y < len(grid)-1 {
			dfs(x, y+1, d+1)
		}
	}
}

func Part01(input []byte) int {
	var b = bufio.NewReader(bytes.NewBuffer(input))

	var startx, starty int

	grid = grid[:0]

	var y int
	for {
		line, err := b.ReadBytes('\n')
		if err != nil {
			break
		}

		if startx == 0 && starty == 0 {
			for x, c := range line {
				if c == 'S' {
					startx = x
					starty = y
					line[x] = '-'
					break
				}
			}
		}

		grid = append(grid, line[:len(line)-1])

		y++
	}

	dfs(startx, starty, 0)

	max := 0
	for _, d := range distances {
		if d > max {
			max = d
		}
	}

	return max/2 + 1
}

func Part02(input []byte) int {
	var b = bufio.NewReader(bytes.NewBuffer(input))

	grid = grid[:0]

	var startx, starty int

	var y int
	for {
		line, err := b.ReadBytes('\n')
		if err != nil {
			break
		}

		if startx == 0 && starty == 0 {
			for x, c := range line {
				if c == 'S' {
					startx = x
					starty = y
					line[x] = '-'
					break
				}
			}
		}

		grid = append(grid, line[:len(line)-1])

		y++
	}

	dfs(startx, starty, 0)

	count := 0
	var state bool
	var lastChar byte
	for y, line := range grid {
		for x := range line {
			v := visited[xyToIndex(x, y)]
			if v {
				switch grid[y][x] {
				case '|':
					state = !state
				case '-':
				default:
					if lastChar == '.' {
						lastChar = grid[y][x]
						state = !state
					} else {
						if (lastChar == 'J' && grid[y][x] == 'L') ||
							(lastChar == 'L' && grid[y][x] == 'J') ||
							(lastChar == '7' && grid[y][x] == 'F') ||
							(lastChar == 'F' && grid[y][x] == '7') {
							state = !state
						}
						lastChar = '.'
					}

				}
			} else {
				lastChar = '.'
				if state {
					count++
				}
			}
		}
		state = false
	}

	return count
}
