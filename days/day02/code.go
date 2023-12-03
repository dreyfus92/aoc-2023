package main

import (
	"fmt"
	"helpers"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	p1 := Part01()
	p2 := Part02()
	fmt.Println("Part1 Answer:", p1)
	fmt.Println("Part2 Answer:", p2)
}

type CubeCounts struct {
	Red, Green, Blue int
}

type Game struct {
	ID   int
	Sets []CubeCounts
}

// parseGame parses a single game's data from a string.
func parseGame(gameStr string) Game {
	var game Game
	parts := strings.Split(gameStr, ": ")
	game.ID, _ = strconv.Atoi(strings.TrimPrefix(parts[0], "Game "))

	setRegex := regexp.MustCompile(`(\d+) red|(\d+) green|(\d+) blue`)
	for _, set := range strings.Split(parts[1], "; ") {
		var counts CubeCounts
		matches := setRegex.FindAllStringSubmatch(set, -1)
		for _, match := range matches {
			if match[1] != "" {
				counts.Red, _ = strconv.Atoi(match[1])
			} else if match[2] != "" {
				counts.Green, _ = strconv.Atoi(match[2])
			} else if match[3] != "" {
				counts.Blue, _ = strconv.Atoi(match[3])
			}
		}
		game.Sets = append(game.Sets, counts)
	}
	return game
}

// isGamePossible checks if a game is possible given the cube limits.
func isGamePossible(game Game, limits CubeCounts) bool {
	for _, set := range game.Sets {
		if set.Red > limits.Red || set.Green > limits.Green || set.Blue > limits.Blue {
			return false
		}
	}
	return true
}

func Part01() int {
	dataInput, err := helpers.GetInput("02")
	if err != nil {
		os.Exit(1)
	}

	games := strings.Split(dataInput, "\n")
	cubeLimits := CubeCounts{Red: 12, Green: 13, Blue: 14}
	sumOfPossibleGameIDs := 0

	for _, gameStr := range games {
		game := parseGame(gameStr)
		if isGamePossible(game, cubeLimits) {
			sumOfPossibleGameIDs += game.ID
		}
	}

	return sumOfPossibleGameIDs
}

func minCubesNeeded(game Game) CubeCounts {
	var minCubes CubeCounts
	for _, set := range game.Sets {
		if set.Red > minCubes.Red {
			minCubes.Red = set.Red
		}
		if set.Green > minCubes.Green {
			minCubes.Green = set.Green
		}
		if set.Blue > minCubes.Blue {
			minCubes.Blue = set.Blue
		}
	}
	return minCubes
}

func cubeSetPower(cubes CubeCounts) int {
	return cubes.Red * cubes.Green * cubes.Blue
}

func Part02() int {
	dataInput, err := helpers.GetInput("02")
	if err != nil {
		os.Exit(1)
	}

	games := strings.Split(dataInput, "\n")
	totalPower := 0

	for _, gameStr := range games {
		game := parseGame(gameStr)
		minCubes := minCubesNeeded(game)
		power := cubeSetPower(minCubes)
		totalPower += power
	}

	return totalPower
}
