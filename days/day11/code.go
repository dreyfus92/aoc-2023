package main

import (
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

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

var galaxyList [500]uint64
var columnWithGalaxy [140]int

func getSumOfLengths(offset int, input []byte) int {
	var lineWithGalaxy bool
	var lineOffset = 0
	var y int
	var x int
	var galaxyCount int
	for _, c := range input {
		switch c {
		case '\n':
			if !lineWithGalaxy {
				lineOffset += 1
			}
			y++
			x = -1
			lineWithGalaxy = false
		case '#':
			galaxyList[galaxyCount] = uint64(y+lineOffset*(offset-1))<<32 | uint64(x)
			lineWithGalaxy = true
			columnWithGalaxy[x] = 1
			galaxyCount++
		}
		x++
	}

	var colWithGalaxy [140]int

	cum := 0
	for i, c := range columnWithGalaxy {
		if c == 0 {
			cum++
		}
		colWithGalaxy[i] = cum
	}

	for i := 0; i < galaxyCount; i++ {
		g := galaxyList[i]
		x := int(g & 0xFFFFFF)
		columnOffset := colWithGalaxy[x]
		galaxyList[i] += uint64(columnOffset * (offset - 1))
	}

	sum := 0
	for i := 0; i < galaxyCount; i++ {
		x, y := int(galaxyList[i]&0xFFFFFF), int(galaxyList[i]>>32)
		for j := i + 1; j < galaxyCount; j++ {
			x1, y1 := int(galaxyList[j]&0xFFFFFF), int(galaxyList[j]>>32)

			// Calculate distance between i and j
			var distance = abs(x-x1) + abs(y-y1)
			sum += distance
		}
	}

	return sum
}

func Part01(input []byte) int {
	sum := getSumOfLengths(2, input)
	return sum
}

func Part02(input []byte) int {
	sum := getSumOfLengths(1000000, input)
	return sum
}
