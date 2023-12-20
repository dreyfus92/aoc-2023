package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
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
	steps := strings.Split(strings.TrimSpace(string(input)), ",")
	sum := 0
	for _, s := range steps {
		var h byte
		for _, b := range []byte(s) {
			h = (h + b) * 17
		}
		sum += int(h)
	}
	return sum
}

type lens struct {
	label string
	value int
}

func Part02(input []byte) int {
	steps := strings.Split(strings.TrimSpace(string(input)), ",")
	boxes := make([][]lens, 256)
stepLoop:
	for _, s := range steps {
		var h byte
		label, valueStr, add := strings.Cut(s, "=")
		if !add {
			label = strings.TrimSuffix(s, "-")
		}
		for _, b := range []byte(label) {
			h = (h + b) * 17
		}
		lenses := boxes[h]
		if add {
			value, _ := strconv.Atoi(valueStr)
			for i := 0; i < len(lenses); i++ {
				if lenses[i].label == label {
					lenses[i].value = value
					continue stepLoop
				}
			}
			boxes[h] = append(lenses, lens{label, value})
			continue stepLoop
		}
		for i := 0; i < len(lenses); i++ {
			if lenses[i].label == label {
				copy(lenses[i:], lenses[i+1:])
				boxes[h] = lenses[:len(lenses)-1]
				continue stepLoop
			}
		}
	}
	sum := 0
	for i, lenses := range boxes {
		for j, l := range lenses {
			sum += (i + 1) * (j + 1) * l.value
		}
	}
	return sum
}
