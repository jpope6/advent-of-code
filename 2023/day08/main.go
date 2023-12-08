package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input []string

	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	var ans int = part2(input)
	fmt.Println(ans)
}

type LeftRight struct {
	L string
	R string
}

func part1(input []string) int {
	instructions := input[0]
	destMap := make(map[string]LeftRight)

	for _, line := range input[2:] {
		s := strings.Fields(line)

		lr := LeftRight{
			L: s[2][1 : len(s[2])-1],
			R: s[3][:len(s[3])-1],
		}

		destMap[s[0]] = lr
	}

	current := "AAA"
	idx := 0
	count := 0

	for current != "ZZZ" {
		instruction := instructions[idx]
		if instruction == 'L' {
			current = destMap[current].L
		} else {
			current = destMap[current].R
		}
		count++
		idx = (idx + 1) % len(instructions)
	}

	return count
}

func part2(input []string) int {
	instructions := input[0]
	destMap := make(map[string]LeftRight)
	startNodes := []string{}

	for _, line := range input[2:] {
		s := strings.Fields(line)

		lr := LeftRight{
			L: s[2][1 : len(s[2])-1],
			R: s[3][:len(s[3])-1],
		}

		destMap[s[0]] = lr

		if strings.HasSuffix(s[0], "A") {
			startNodes = append(startNodes, s[0])
		}
	}

	ans := 1

	for _, node := range startNodes {
		idx := 0
		count := 0
		current := node
		for !strings.HasSuffix(current, "Z") {
			instruction := instructions[idx]
			if instruction == 'L' {
				current = destMap[current].L
			} else {
				current = destMap[current].R
			}
			count++
			idx = (idx + 1) % len(instructions)
		}

		ans = LCM(ans, count)
	}

	return ans
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
