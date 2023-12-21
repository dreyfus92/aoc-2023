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

const (
	none = iota
	right
	down
	up
	left
)

type id struct {
	x, y, dir, count int
}
type node struct {
	id
	cost int
}

func parseInput(input string) ([][]int, int, int) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	w, h := len(lines[0]), len(lines)
	grid := make([][]int, h)
	for y := range grid {
		grid[y] = make([]int, w)
		for x := range grid[y] {
			grid[y][x] = int(lines[y][x]) - '0'
		}
	}
	return grid, w, h
}

func Part01(input []byte) int {
	grid, w, h := parseInput(string(input))
	pq := []node{{}}
	visited := make(map[id]int)
	visited[id{}] = 0

	add := func(n node, newDir, oldDir, oldCount int) {
		if oldDir == newDir {
			n.count = oldCount + 1
		}
		if n.count > 3 {
			return
		}
		if cost, found := visited[n.id]; found && cost <= n.cost {
			return
		}
		visited[n.id] = n.cost
		pq = append(pq, node{})
		for i := len(pq) - 2; i >= 0; i-- {
			if n.cost > pq[i].cost {
				pq[i+1] = n
				return
			}
			pq[i+1] = pq[i]
		}
		pq[0] = n
	}

	cost := 0
	for len(pq) > 0 {
		n := pq[0]
		pq = pq[1:]
		if n.x == w-1 && n.y == h-1 {
			cost = n.cost
			break
		}
		if n.x+1 < w && n.dir != left {
			nn := node{id{n.x + 1, n.y, right, 1}, n.cost + grid[n.y][n.x+1]}
			add(nn, right, n.dir, n.count)
		}
		if n.y+1 < h && n.dir != up {
			nn := node{id{n.x, n.y + 1, down, 1}, n.cost + grid[n.y+1][n.x]}
			add(nn, down, n.dir, n.count)
		}
		if n.x > 0 && n.dir != right {
			nn := node{id{n.x - 1, n.y, left, 1}, n.cost + grid[n.y][n.x-1]}
			add(nn, left, n.dir, n.count)
		}
		if n.y > 0 && n.dir != down {
			nn := node{id{n.x, n.y - 1, up, 1}, n.cost + grid[n.y-1][n.x]}
			add(nn, up, n.dir, n.count)
		}
	}

	return cost
}

func Part02(input []byte) int {
	grid, w, h := parseInput(string(input))
	pq := []node{{}}
	visited := make(map[id]int)
	visited[id{}] = 0

	add := func(n node, newDir, oldDir, oldCount int) {
		if oldDir != none && oldDir != newDir && oldCount < 4 {
			return
		}
		if oldDir == newDir {
			n.count = oldCount + 1
		}
		if n.count > 10 {
			return
		}
		if cost, found := visited[n.id]; found && cost <= n.cost {
			return
		}
		visited[n.id] = n.cost
		pq = append(pq, node{})
		for i := len(pq) - 2; i >= 0; i-- {
			if n.cost > pq[i].cost {
				pq[i+1] = n
				return
			}
			pq[i+1] = pq[i]
		}
		pq[0] = n
	}

	cost := 0
	for len(pq) > 0 {
		n := pq[0]
		pq = pq[1:]
		if n.x == w-1 && n.y == h-1 && n.count >= 4 {
			cost = n.cost
			break
		}
		if n.x+1 < w && n.dir != left {
			nn := node{id{n.x + 1, n.y, right, 1}, n.cost + grid[n.y][n.x+1]}
			add(nn, right, n.dir, n.count)
		}
		if n.y+1 < h && n.dir != up {
			nn := node{id{n.x, n.y + 1, down, 1}, n.cost + grid[n.y+1][n.x]}
			add(nn, down, n.dir, n.count)
		}
		if n.x > 0 && n.dir != right {
			nn := node{id{n.x - 1, n.y, left, 1}, n.cost + grid[n.y][n.x-1]}
			add(nn, left, n.dir, n.count)
		}
		if n.y > 0 && n.dir != down {
			nn := node{id{n.x, n.y - 1, up, 1}, n.cost + grid[n.y-1][n.x]}
			add(nn, up, n.dir, n.count)
		}
	}
	return cost
}
