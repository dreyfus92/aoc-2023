package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"helpers"
)

func main() {
	p1 := Part01()
	p2 := Part02()
	fmt.Println("Part1 Answer:", p1)
	fmt.Println("Part2 Answer:", p2)
}

func Part01() int {
	dataInput, err := helpers.GetInput("01")
	if err != nil {
		os.Exit(1)
	}

	pattern := "[a-zA-Z]"
	re := regexp.MustCompile(pattern)
	cleanInput := re.ReplaceAllString(dataInput, "")

	inputArr := strings.Fields(cleanInput)

	sum := 0
	for _, v := range inputArr {
		num, _ := strconv.Atoi(string(v[0]) + string(v[len(v)-1]))
		sum += (num)
	}

	return sum
}

func Part02() int {
	dataInput, err := helpers.GetInput("01")
	if err != nil {
		os.Exit(1)
	}

	numPair := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	cleanInput := dataInput
	for k, v := range numPair {
		safeInsertIdx := 2
		rep := k[:safeInsertIdx] + v + k[safeInsertIdx:]
		cleanInput = strings.ReplaceAll(cleanInput, k, rep)
	}

	// Remove any characters from the text field with regex
	pattern := "[a-zA-Z]"
	re := regexp.MustCompile(pattern)
	cleanInput = re.ReplaceAllString(cleanInput, "")

	inputArr := strings.Fields(cleanInput)

	sum := 0
	for _, v := range inputArr {
		num, _ := strconv.Atoi(string(v[0]) + string(v[len(v)-1]))
		sum += num
	}

	return sum
}
