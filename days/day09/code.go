package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	seqs := parse("input.txt")

	a1 := Part01(seqs)
	a2 := Part02(seqs)
	fmt.Println("Part1 Answer:", a1)
	fmt.Println("Part2 Answer:", a2)
}

func parse(filename string) [][]int {
	input, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	sequences := make([][]int, 0, 256)
	for _, l := range strings.Split(string(input), "\n") {
		if l == "" {
			continue
		}

		sequence := make([]int, 0, strings.Count(l, " ")+1)
		for _, str := range strings.Split(l, " ") {
			n, err := strconv.Atoi(str)
			if err != nil {
				panic(err)
			}
			sequence = append(sequence, n)
		}

		sequences = append(sequences, sequence)
	}

	return sequences
}

func Part01(mainSequences [][]int) int {
	pt1 := 0
	diffs := make([]int, 0, len(mainSequences[0])-1)

	for _, s := range mainSequences {
		pt1 += s[len(s)-1]
		hasNonZero := 1

		for hasNonZero != 0 {
			diffs = diffs[:0]
			hasNonZero = 0
			for i := 0; i < len(s)-1; i++ {
				diffs = append(diffs, s[i+1]-s[i])
				hasNonZero |= diffs[i]
			}

			pt1 += diffs[len(diffs)-1]
			s = diffs
		}
	}

	return pt1
}

func Part02(seqs [][]int) int {
	for i := 0; i < len(seqs); i++ {
		slices.Reverse(seqs[i])
	}
	return Part01(seqs)
}
