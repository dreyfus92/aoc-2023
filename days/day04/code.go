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

type ScratchCard struct {
	winning []int
	numbers []int
}

func parseNumbers(s string) []int {
	var numbers []int
	fields := strings.Fields(s)
	for _, f := range fields {
		num, err := strconv.Atoi(f)
		if err != nil {
			fmt.Println("Error parsing number:", err)
			continue
		}
		numbers = append(numbers, num)
	}
	return numbers
}

func calculatePoints(winning []int, numbers []int) int {
	points := 0
	matches := 0

	for _, num := range numbers {
		if contains(winning, num) {
			matches++
		}
	}

	if matches > 0 {
		points = 1 << (matches - 1)
	}

	return points
}

func contains(slice []int, item int) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func getInputAsScratchards(input string) []ScratchCard {
	file, err := os.Open(input)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	var scratchCards []ScratchCard

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		line := sc.Text()
		parts := strings.Split(line, "|")
		if len(parts) != 2 {
			fmt.Println("Error parsing line:", line)
		}

		winningNumbers := parseNumbers(parts[0])
		myNumbers := parseNumbers(parts[1])

		scratchCards = append(scratchCards, ScratchCard{
			winning: winningNumbers,
			numbers: myNumbers,
		})
	}

	if err := sc.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return scratchCards
}

func Part01() int {

	cards := getInputAsScratchards("input.txt")

	totalPoints := 0
	for _, card := range cards {
		totalPoints += calculatePoints(card.winning, card.numbers)
	}

	return totalPoints
}

func countTotalScratchCards(cards []ScratchCard) int {
	cardCounts := make([]int, len(cards))

	for i := range cardCounts {
		cardCounts[i] = 1
	}

	for i := 0; i < len(cards); i++ {
		matches := countMatches(cards[i].winning, cards[i].numbers)
		for j := 1; j <= matches && i+j < len(cardCounts); j++ {
			cardCounts[i+j] += cardCounts[i]
		}
	}

	total := 0

	for _, count := range cardCounts {
		total += count
	}

	return total

}

func countMatches(winning, numbers []int) int {
	matches := 0
	for _, num := range numbers {
		if contains(winning, num) {
			matches++
		}
	}
	return matches
}

func Part02() int {
	cards := getInputAsScratchards("input.txt")

	totalScratchards := countTotalScratchCards(cards)
	return totalScratchards
}
