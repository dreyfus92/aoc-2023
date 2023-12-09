package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	p1 := Part01()
	p2 := Part02()
	fmt.Println("Part1 Answer:", p1)
	fmt.Println("Part2 Answer:", p2)
}

type Node struct {
	Left, Right string
}

// navigateNetwork navigates the network based on the instructions until it reaches "ZZZ".
func navigateNetwork(instructions string, network map[string]Node) int {
	current := "AAA"
	steps := 0

	for current != "ZZZ" {
		instruction := instructions[steps%len(instructions)]
		if instruction == 'L' {
			current = network[current].Left
		} else if instruction == 'R' {
			current = network[current].Right
		}
		steps++
	}

	return steps
}

func readInput(filename string) (string, map[string]Node, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	network := make(map[string]Node)
	var instructions string

	// Read the first line for instructions
	if scanner.Scan() {
		instructions = scanner.Text()
	}

	// Read the remaining lines for the network
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " = ")
		if len(parts) == 2 {
			node := strings.TrimSpace(parts[0])
			connections := strings.Trim(parts[1], "()")
			connParts := strings.Split(connections, ", ")
			if len(connParts) == 2 {
				network[node] = Node{strings.TrimSpace(connParts[0]), strings.TrimSpace(connParts[1])}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return "", nil, err
	}

	return instructions, network, nil
}

func Part01() int {
	instructions, network, err := readInput("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
	}

	steps := navigateNetwork(instructions, network)
	return steps
}

func complete(xs []int) bool {
	for _, x := range xs {
		if x == 0 {
			return false
		}
	}
	return true
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcmm(xs []int) int {
	lcm := func(a, b int) int { return a * b / gcd(a, b) }

	result := 1
	for _, n := range xs {
		result = lcm(result, n)
	}
	return result
}

func parse() (string, map[string]string, map[string]string) {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	instructions := scanner.Text()
	scanner.Scan()

	left := map[string]string{}
	right := map[string]string{}

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, ",", "")
		line = strings.ReplaceAll(line, ")", "")

		var h, l, r string
		_, _ = fmt.Sscanf(line, "%s = (%s %s", &h, &l, &r)

		left[h] = l
		right[h] = r
	}

	return instructions, left, right
}

func Part02() int {
	instructions, left, right := parse()

	curr := []string{}
	for k := range left {
		if k[2] == 'A' {
			curr = append(curr, k)
		}
	}

	periods := make([]int, len(curr))

	for i := 1; ; i++ {
		for j, k := range curr {

			switch instructions[(i-1)%len(instructions)] {
			case 'L':
				curr[j] = left[k]
			case 'R':
				curr[j] = right[k]
			}

			if curr[j][2] == 'Z' && periods[j] == 0 {
				periods[j] = i
			}
		}

		if complete(periods) {
			return lcmm(periods)
		}
	}
}
