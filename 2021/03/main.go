package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
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

	start := time.Now()
	var ans int = part1(input)
	fmt.Println(ans)
	fmt.Printf("Time taken: %d ms\n", time.Since(start))
}

func part1(input []string) int {
	count := [5]int{}

	for _, line := range input {
		for i, num := range line {
			if num == rune(1) {
				count[i]++
			}
		}
	}

	return 0
}

func part2(input []string) int {
	x, y, aim := 0, 0, 0

	for _, line := range input {
		instructions := strings.Split(line, " ")

		val, _ := strconv.Atoi(instructions[1])
		switch instructions[0] {
		case "forward":
			x += val
			y += aim * val
		case "down":
			aim += val
		case "up":
			aim -= val
		}
	}

	return x * y
}
