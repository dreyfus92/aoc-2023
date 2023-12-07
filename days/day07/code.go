package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
)

func main() {
	p1 := Part01()
	p2 := Part02()
	fmt.Println("Part1 Answer:", p1)
	fmt.Println("Part2 Answer:", p2)
}

var handTypes = map[int]string{
	1: "High Card",
	2: "One Pair",
	3: "Two Pair",
	4: "Three of a Kind",
	5: "Full House",
	6: "Four of a Kind",
	7: "Five of a Kind",
}

var cardValues = map[string]int{
	"2": 1,
	"3": 2,
	"4": 3,
	"5": 4,
	"6": 5,
	"7": 6,
	"8": 7,
	"9": 8,
	"T": 9,
	"J": 10,
	"Q": 11,
	"K": 12,
	"A": 13,
}

type Hand struct {
	hand string
	bid  int
}

func getHandType(hand string) int {
	var handValue int
	cardsInHand := map[string]int{}
	numJokers := 0

	for _, card := range hand {
		if card == 'J' {
			numJokers++
			continue
		}
		cardsInHand[string(card)]++
	}

	numOfSameCards := []int{}
	for _, value := range cardsInHand {
		if value > 0 {
			numOfSameCards = append(numOfSameCards, value)
		}
	}

	sort.Ints(numOfSameCards)

	if len(numOfSameCards) == 0 {
		handValue = 7

		return handValue
	}

	numOfSameCards[len(numOfSameCards)-1] += numJokers

	if slices.Contains(numOfSameCards, 5) {
		handValue = 7
	} else if slices.Contains(numOfSameCards, 4) {
		handValue = 6
	} else if slices.Contains(numOfSameCards, 3) && slices.Contains(numOfSameCards, 2) {
		handValue = 5
	} else if slices.Contains(numOfSameCards, 3) {
		handValue = 4
	} else if slices.Contains(numOfSameCards, 2) && len(numOfSameCards) == 3 {
		handValue = 3
	} else if slices.Contains(numOfSameCards, 2) {
		handValue = 2
	} else {
		handValue = 1
	}
	return handValue
}

func doesHandWin(hand1 string, hand2 string) bool {
	handWins := false
	hand1Value := getHandType(hand1)
	hand2Value := getHandType(hand2)
	if hand1Value > hand2Value {
		handWins = true
	} else if hand1Value < hand2Value {
		handWins = false
	} else {
		for i := 0; i < len(hand1); i++ {
			if cardValues[string(hand1[i])] > cardValues[string(hand2[i])] {
				handWins = true
				break
			} else if cardValues[string(hand1[i])] < cardValues[string(hand2[i])] {
				handWins = false
				break
			}
		}
	}
	return handWins
}

func Part01() int {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading file")
	}

	sc := bufio.NewScanner(file)

	hands := []Hand{}

	for sc.Scan() {
		line := sc.Text()
		bid, _ := strconv.Atoi(line[6:])

		newHand := struct {
			hand string
			bid  int
		}{
			hand: line[:5],
			bid:  bid,
		}

		hands = append(hands, newHand)
	}

	file.Close()

	sort.Slice(hands, func(i, j int) bool {
		return doesHandWin(hands[i].hand, hands[j].hand)
	})

	total := 0

	for i, hand := range hands {
		total += hand.bid * (len(hands) - i)
	}

	return total
}

var cardValuesPt2 = map[string]int{
	"J": 0,
	"2": 1,
	"3": 2,
	"4": 3,
	"5": 4,
	"6": 5,
	"7": 6,
	"8": 7,
	"9": 8,
	"T": 9,
	"Q": 11,
	"K": 12,
	"A": 13,
}

func doesHandWinPt2(hand1 string, hand2 string) bool {
	handWins := false
	hand1Value := getHandType(hand1)
	hand2Value := getHandType(hand2)
	if hand1Value > hand2Value {
		handWins = true
	} else if hand1Value < hand2Value {
		handWins = false
	} else {
		for i := 0; i < len(hand1); i++ {
			if cardValuesPt2[string(hand1[i])] > cardValuesPt2[string(hand2[i])] {
				handWins = true
				break
			} else if cardValuesPt2[string(hand1[i])] < cardValuesPt2[string(hand2[i])] {
				handWins = false
				break
			}
		}
	}
	return handWins
}

func Part02() int {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading file")
	}

	sc := bufio.NewScanner(file)

	hands := []Hand{}

	for sc.Scan() {
		line := sc.Text()
		bid, _ := strconv.Atoi(line[6:])

		newHand := struct {
			hand string
			bid  int
		}{
			hand: line[:5],
			bid:  bid,
		}

		hands = append(hands, newHand)
	}

	file.Close()

	sort.Slice(hands, func(i, j int) bool {
		return doesHandWinPt2(hands[i].hand, hands[j].hand)
	})

	total := 0

	for i, hand := range hands {
		total += hand.bid * (len(hands) - i)
	}

	return total
}
