package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func part1(input []string) int {
	lines := make([][]int, 0)

	for _, line := range input {
		s := []int{}

		for _, num := range strings.Fields(line) {
			n, _ := strconv.Atoi(num)
			s = append(s, n)
		}

		lines = append(lines, s)
	}

	ans := 0

	for _, line := range lines {
		val := 0
		originalLine := line

		for !isAllZeroes(line) {
			line = findNextRow(line)
			val += line[len(line)-1]
		}

		val += originalLine[len(originalLine)-1]
		ans += val
	}

	return ans
}

func part2(input []string) int {
	lines := make([][]int, 0)

	for _, line := range input {
		s := []int{}

		for _, num := range strings.Fields(line) {
			n, _ := strconv.Atoi(num)
			s = append(s, n)
		}

		lines = append(lines, s)
	}

	ans := 0

	for _, line := range lines {
		val := 0
		originalLine := line

		for !isAllZeroes(line) {
			line = findNextRow2(line)
			val += line[0]
		}

		val += originalLine[0]
		ans += val
	}

	return ans
}

func findNextRow(line []int) []int {
	curr := make([]int, len(line)-1)
	for i := 1; i < len(line); i++ {
		curr[i-1] = line[i] - line[i-1]
	}

	return curr
}

func findNextRow2(line []int) []int {
	curr := make([]int, len(line)-1)
	for i := 1; i < len(line); i++ {
		curr[i-1] = line[i-1] - line[i]
	}

	return curr
}

func isAllZeroes(arr []int) bool {
	for _, num := range arr {
		if num != 0 {
			return false
		}
	}

	return true
}
