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
	var colorMap = map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	sum := 0

	for i, s := range input {
		s := strings.TrimPrefix(s, "Game "+fmt.Sprint(i+1)+": ")
		var gameInfo []string = strings.Split(s, "; ")
		valid := true

		for _, round := range gameInfo {
			var colorInfo []string = strings.Split(round, ", ")

			for _, colors := range colorInfo {
				numStr, color, _ := strings.Cut(colors, " ")
				num, _ := strconv.Atoi(numStr)

				if num > colorMap[color] {
					valid = false
					break
				}
			}
		}

		if valid {
			sum += i + 1
		} else {
			valid = true
		}
	}

	return sum
}

func part2(input []string) int {
	colorMap := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
	sum := 0

	for i, s := range input {
		s := strings.TrimPrefix(s, "Game "+fmt.Sprint(i+1)+": ")
		var gameInfo []string = strings.Split(s, "; ")

		for _, round := range gameInfo {
			var colorInfo []string = strings.Split(round, ", ")

			for _, colors := range colorInfo {
				numStr, color, _ := strings.Cut(colors, " ")
				num, _ := strconv.Atoi(numStr)

				colorMap[color] = max(colorMap[color], num)
			}
		}

		product := 1
		for color := range colorMap {
			product *= colorMap[color]
			colorMap[color] = 0
		}
		sum += product
	}

	return sum
}
