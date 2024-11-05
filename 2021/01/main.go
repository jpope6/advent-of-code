package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
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
	var ans int = part2(input)
	fmt.Println(ans)
	fmt.Printf("Time taken: %d ms\n", time.Since(start))
}

func part1(input []string) int {
	curr, next := 0, 1

	count := 0
	for next < len(input) {
		currValue, _ := strconv.Atoi(input[curr])
		nextValue, _ := strconv.Atoi(input[next])

		if currValue < nextValue {
			count++
		}

		curr++
		next++
	}

	return count
}

func part2(input []string) int {
	first, second, third := 0, 1, 2

	prevSum := math.MaxInt
	count := 0
	for third < len(input) {
		firstVal, _ := strconv.Atoi(input[first])
		secondVal, _ := strconv.Atoi(input[second])
		thirdVal, _ := strconv.Atoi(input[third])

		sum := firstVal + secondVal + thirdVal
		if prevSum < sum {
			count++
		}
		prevSum = sum

		first++
		second++
		third++
	}

	return count
}
