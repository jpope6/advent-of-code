package main

import (
	"bufio"
	"fmt"
	"os"
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
	count := make([]int, len(input[0]))

	for _, line := range input {
		for i, num := range line {
			if num == '1' {
				count[i]++
			}
		}
	}

	half := len(input) / 2

	var gammaRate uint = 0
	var epsilonRate uint = 0
	for i, num := range count {
		index := len(count) - 1 - i
		if num > half {
			gammaRate |= 1 << index
		}

		if num < half {
			epsilonRate |= 1 << index
		}
	}

	return int(gammaRate * epsilonRate)
}
