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
func countMatches(springs string, ranges []int) int {
	if len(ranges) == 0 {
		if strings.Count(springs, "#") > 0 {
			return 0
		}
		return 1
	}
	if springs == "" {
		return 0
	}
	s := springs[0]
	if s == '.' {
		return countMatches(springs[1:], ranges)
	}
	sum := 0
	if s == '?' {
		sum = countMatches(springs[1:], ranges)
	}
	r := ranges[0]
	if len(springs) < r || strings.Count(springs[:r], ".") > 0 {
		return sum
	}
	springs = springs[r:]
	if len(springs) > 0 {
		if springs[0] == '#' {
			return sum
		}
		springs = springs[1:]
	}
	return sum + countMatches(springs, ranges[1:])
}

func Part01(input []byte) int {
	sum := 0
	for _, line := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		springs, rangesStr, _ := strings.Cut(line, " ")
		var ranges []int
		for _, s := range strings.Split(rangesStr, ",") {
			v, _ := strconv.Atoi(s)
			ranges = append(ranges, v)
		}
		sum += countMatches(springs, ranges)
	}
	return sum
}

func countMatchesMemo(si, ri int, springs string, ranges []int, memo [][]int) int {
	if m := memo[ri][si]; m >= 0 {
		return m
	}
	// Memoize every other return with the given si/ri.
	if ri >= len(ranges) {
		if strings.Count(springs[si:], "#") > 0 {
			memo[ri][si] = 0
			return 0
		}
		memo[ri][si] = 1
		return 1
	}
	if si >= len(springs) {
		memo[ri][si] = 0
		return 0
	}
	s := springs[si]
	if s == '.' {
		count := countMatchesMemo(si+1, ri, springs, ranges, memo)
		memo[ri][si] = count
		return count
	}
	sum := 0
	if s == '?' {
		sum = countMatchesMemo(si+1, ri, springs, ranges, memo)
	}
	r := ranges[ri]
	sir := si + r
	if len(springs) < sir || strings.Count(springs[si:sir], ".") > 0 {
		memo[ri][si] = sum
		return sum
	}
	if sir < len(springs) {
		if springs[sir] == '#' {
			memo[ri][si] = sum
			return sum
		}
		sir++
	}
	count := sum + countMatchesMemo(sir, ri+1, springs, ranges, memo)
	memo[ri][si] = count
	return count
}

func Part02(input []byte) int {
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	sum := 0
	for _, line := range lines {
		springs, rangesStr, _ := strings.Cut(line, " ")
		var ranges []int
		for _, s := range strings.Split(rangesStr, ",") {
			v, _ := strconv.Atoi(s)
			ranges = append(ranges, v)
		}
		// "Unfold" the input to be 5x the input.
		{
			var springs0 []string
			var ranges0 []int
			for i := 0; i < 5; i++ {
				springs0 = append(springs0, springs)
				ranges0 = append(ranges0, ranges...)
			}
			// Unfolding springs separate input copies by "?".
			springs = strings.Join(springs0, "?")
			ranges = ranges0
		}
		// Oversize memoized result table to remember end indexes as well.
		memo := make([][]int, len(ranges)+1)
		for i := range memo {
			memo[i] = make([]int, len(springs)+1)
			for j := 0; j < len(springs)+1; j++ {
				memo[i][j] = -1 // We want to remember 0s as well.
			}
		}
		sum += countMatchesMemo(0, 0, springs, ranges, memo)
	}
	return sum
}
