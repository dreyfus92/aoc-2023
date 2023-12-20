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

func Part01(input []byte) int {
	const left = 1
	const right = 2
	const up = 4
	const down = 8
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	w, h := int8(len(lines[0])), int8(len(lines))
	beams := make([][]byte, h)
	for i := range beams {
		beams[i] = make([]byte, w)
	}
	beams[0][0] = right
	var qxyd []int8
	qxyd = append(qxyd, 0, 0, right)
	for len(qxyd) > 0 {
		x, y, d := qxyd[0], qxyd[1], qxyd[2]
		qxyd = qxyd[3:]
		c := lines[y][x]
		if c == '.' && d == left || c == '/' && d == down || c == '\\' && d == up || c == '-' && (d == left || d == up || d == down) {
			if x > 0 && beams[y][x-1]&left == 0 {
				qxyd = append(qxyd, x-1, y, left)
				beams[y][x-1] |= left
			}
		}
		if c == '.' && d == right || c == '/' && d == up || c == '\\' && d == down || c == '-' && (d == right || d == up || d == down) {
			if x < w-1 && beams[y][x+1]&right == 0 {
				qxyd = append(qxyd, x+1, y, right)
				beams[y][x+1] |= right
			}
		}
		if c == '.' && d == up || c == '/' && d == right || c == '\\' && d == left || c == '|' && (d == up || d == left || d == right) {
			if y > 0 && beams[y-1][x]&up == 0 {
				qxyd = append(qxyd, x, y-1, up)
				beams[y-1][x] |= up
			}
		}
		if c == '.' && d == down || c == '/' && d == left || c == '\\' && d == right || c == '|' && (d == down || d == left || d == right) {
			if y < h-1 && beams[y+1][x]&down == 0 {
				qxyd = append(qxyd, x, y+1, down)
				beams[y+1][x] |= down
			}
		}
	}
	sum := 0
	for _, bs := range beams {
		for _, c := range bs {
			if c > 0 {
				sum++
			}
		}
	}
	return sum
}

func Part02(input []byte) int {
	const left = 1
	const right = 2
	const up = 4
	const down = 8
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	w, h := int8(len(lines[0])), int8(len(lines))
	var sxyd []int8
	for i := int8(0); i < w; i++ {
		sxyd = append(sxyd, i, 0, down)
		sxyd = append(sxyd, i, h-1, up)
	}
	for i := int8(0); i < h; i++ {
		sxyd = append(sxyd, 0, i, right)
		sxyd = append(sxyd, w-1, i, left)
	}
	maxSum := 0
	for si := 0; si < len(sxyd); si += 3 {
		beams := make([][]byte, h)
		for i := range beams {
			beams[i] = make([]byte, w)
		}
		beams[sxyd[si+1]][sxyd[si]] = byte(sxyd[si+2])
		var qxyd []int8
		qxyd = append(qxyd, sxyd[si], sxyd[si+1], sxyd[si+2])
		for len(qxyd) > 0 {
			x, y, d := qxyd[0], qxyd[1], qxyd[2]
			qxyd = qxyd[3:]
			c := lines[y][x]
			if c == '.' && d == left || c == '/' && d == down || c == '\\' && d == up || c == '-' && (d == left || d == up || d == down) {
				if x > 0 && beams[y][x-1]&left == 0 {
					qxyd = append(qxyd, x-1, y, left)
					beams[y][x-1] |= left
				}
			}
			if c == '.' && d == right || c == '/' && d == up || c == '\\' && d == down || c == '-' && (d == right || d == up || d == down) {
				if x < w-1 && beams[y][x+1]&right == 0 {
					qxyd = append(qxyd, x+1, y, right)
					beams[y][x+1] |= right
				}
			}
			if c == '.' && d == up || c == '/' && d == right || c == '\\' && d == left || c == '|' && (d == up || d == left || d == right) {
				if y > 0 && beams[y-1][x]&up == 0 {
					qxyd = append(qxyd, x, y-1, up)
					beams[y-1][x] |= up
				}
			}
			if c == '.' && d == down || c == '/' && d == left || c == '\\' && d == right || c == '|' && (d == down || d == left || d == right) {
				if y < h-1 && beams[y+1][x]&down == 0 {
					qxyd = append(qxyd, x, y+1, down)
					beams[y+1][x] |= down
				}
			}
		}
		sum := 0
		for _, bs := range beams {
			for _, c := range bs {
				if c > 0 {
					sum++
				}
			}
		}
		maxSum = max(maxSum, sum)
	}
	return maxSum
}
