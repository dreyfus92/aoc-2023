package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	p1 := Part01()
	p2 := Part02()
	fmt.Println("Part1 Answer:", p1)
	fmt.Println("Part2 Answer:", p2)
}

type Race struct {
	timeLimit      int
	recordDistance int
}

func calculateWays(timeLimit int, recordDistance int) int {
	waysToWin := 0
	for buttonPressTime := 0; buttonPressTime < timeLimit; buttonPressTime++ {
		speed := buttonPressTime
		travelTime := timeLimit - buttonPressTime
		distanceTravelled := speed * travelTime
		if distanceTravelled > recordDistance {
			waysToWin++
		}
	}
	return waysToWin
}

func parseLine(line string) ([]int, error) {
	parts := strings.Fields(line)
	var intValues []int
	var err error

	for _, part := range parts {
		var intValue int
		intValue, err = strconv.Atoi(part)
		if err != nil {
			continue
		}
		intValues = append(intValues, intValue)
	}
	return intValues, err
}

func Part01() int {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	sc.Scan()
	timeLimits, err := parseLine(sc.Text())
	if err != nil {
		panic(err)
	}

	sc.Scan()
	recordDistances, err := parseLine(sc.Text())
	if err != nil {
		panic(err)
	}
	result := 1
	for i := range timeLimits {
		ways := calculateWays(timeLimits[i], recordDistances[i])
		result *= ways
	}
	return result
}

func parseSingleNumber(line string) (int, error) {
	filtered := strings.Map(func(r rune) rune {
		if unicode.IsDigit(r) {
			return r
		}
		return -1
	}, line)
	return strconv.Atoi(filtered)
}

func Part02() int {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	sc.Scan()
	timeLimit, err := parseSingleNumber(sc.Text())
	if err != nil {
		panic(err)
	}

	sc.Scan()
	recordDistance, err := parseSingleNumber(sc.Text())
	if err != nil {
		panic(err)
	}

	waysToWin := calculateWays(timeLimit, recordDistance)

	return waysToWin
}
