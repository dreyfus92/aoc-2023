package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Number struct {
	Value  int
	XPos   int
	YPos   int
	Length int
}

type Symbol struct {
	Value   rune
	XPos    int
	YPos    int
	Numbers []int
}

func main() {
	p1 := Part01()
	p2 := Part02()
	fmt.Println("Part1 Answer:", p1)
	fmt.Println("Part2 Answer:", p2)
}

func parseInput(input []string) ([]Number, []Symbol) {
	nums := []Number{}
	symbols := []Symbol{}
	num := ""
	for y, s := range input {
		s = s + "."
		for x, c := range s {
			if unicode.IsDigit(c) {
				num += string(c)
			} else {
				if len(num) > 0 {
					l := len(num)
					xPos := x - l
					v, err := strconv.Atoi(num)
					if err != nil {
						fmt.Println(err)
					}

					nums = append(nums, Number{
						Value:  v,
						XPos:   xPos,
						YPos:   y,
						Length: l,
					})
					num = ""
				}

				if c != '.' {
					symbols = append(symbols, Symbol{
						Value:   c,
						XPos:    x,
						YPos:    y,
						Numbers: []int{},
					})
				}
			}
		}
	}
	return nums, symbols
}

func getSchematic() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	var schematic []string

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		schematic = append(schematic, strings.TrimSpace(sc.Text()))
	}

	if err := sc.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	return schematic
}

func Part01() int {

	schematic := getSchematic()

	sum := 0

	nums, symbols := parseInput(schematic)

	for _, n := range nums {
		minX := n.XPos - 1
		maxX := n.XPos + n.Length
		minY := n.YPos - 1
		maxY := n.YPos + 1

		for i, s := range symbols {
			if (s.XPos >= minX && s.XPos <= maxX) &&
				(s.YPos >= minY && s.YPos <= maxY) {

				symbols[i].Numbers = append(s.Numbers, n.Value)
				sum += n.Value
				break
			}
		}
	}

	return sum
}

func Part02() int {

	schematic := getSchematic()

	nums, symbols := parseInput(schematic)

	sum := 0

	for _, n := range nums {
		minX := n.XPos - 1
		maxX := n.XPos + n.Length
		minY := n.YPos - 1
		maxY := n.YPos + 1

		for i, s := range symbols {
			if (s.XPos >= minX && s.XPos <= maxX) &&
				(s.YPos >= minY && s.YPos <= maxY) {

				symbols[i].Numbers = append(s.Numbers, n.Value)
				sum += n.Value
				break
			}
		}
	}

	gearSum := 0
	for _, s := range symbols {
		if s.Value != '*' {
			continue
		}

		if len(s.Numbers) != 2 {
			continue
		}

		gearSum += s.Numbers[0] * s.Numbers[1]
	}

	return gearSum
}
