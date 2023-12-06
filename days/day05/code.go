package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	p1 := Part01()
	p2 := Part02()
	fmt.Println("Part1 Answer:", p1)
	fmt.Println("Part2 Answer:", p2)
}

type Map []Transform

type Transform struct {
	Destination, Source, Length int
}

func parse() ([]int, [7]Map) {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	seeds := []int{}
	sc.Scan()
	fields := strings.Fields(sc.Text())
	for i := 1; i < len(fields); i++ {
		num, _ := strconv.Atoi(fields[i])
		seeds = append(seeds, num)
	}

	return seeds, parseMaps(sc)
}

func parseMaps(scanner *bufio.Scanner) [7]Map {
	maps := [7]Map{}
	t := -1
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if strings.Contains(line, "map:") {
			t++
			continue
		}

		fields := strings.Fields(line)
		dest, _ := strconv.Atoi(fields[0])
		src, _ := strconv.Atoi(fields[1])
		length, _ := strconv.Atoi(fields[2])
		maps[t] = append(maps[t], Transform{
			Destination: dest,
			Source:      src,
			Length:      length,
		})
	}
	return maps
}

func (m Map) Convert(from int) int {
	for _, t := range m {
		if from >= t.Source && from < t.Source+t.Length {
			return t.Destination + (from - t.Source)
		}
	}
	return from
}

func Part01() int {
	seeds, maps := parse()

	minLoc := 1<<31 - 1
	for _, s := range seeds {
		for _, m := range maps {
			s = m.Convert(s)
		}
		if s < minLoc {
			minLoc = s
		}
	}

	return minLoc
}

func (m Map) InverseConvert(to int) int {
	for _, t := range m {
		if to >= t.Destination && to < t.Destination+t.Length {
			return t.Source + (to - t.Destination)
		}
	}
	return to
}

func Part02() int {
	seeds, maps := parse()

	for loc := 0; ; loc++ {
		x := loc
		for m := len(maps) - 1; m >= 0; m-- {
			x = maps[m].InverseConvert(x)
		}

		for s := 0; s < len(seeds); s += 2 {
			if x >= seeds[s] && x < seeds[s]+seeds[s+1] {
				return loc
			}
		}
	}
}
