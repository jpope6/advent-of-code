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

func part1(input []string) int {
	sum := 0

	for i, s := range input {
		s := strings.TrimPrefix(s, "Card "+fmt.Sprint(i+1)+": ")
		winningNumsStr, myNumsStr, _ := strings.Cut(s, "|")
		winningNums := strings.Fields(winningNumsStr)
		myNums := strings.Fields(myNumsStr)

		score := 0

		for _, num := range winningNums {
			for _, n := range myNums {
				if num == n {
					if score == 0 {
						score = 1
					} else {
						score *= 2
					}
				}
			}
		}
		sum += score
		score = 0
	}

	return sum
}

func part2(input []string) int {
	countMap := make(map[int]int)

	// Initialize each card to 1
	for i := 0; i < len(input); i++ {
		countMap[i]++
	}

	for i := 0; i < len(input); i++ {
		s := strings.TrimPrefix(input[i], "Card "+fmt.Sprint(i+1)+": ")
		winningNumsStr, myNumsStr, _ := strings.Cut(s, "|")
		winningNums := strings.Fields(winningNumsStr)
		myNums := strings.Fields(myNumsStr)

		index := i + 1

		for _, num := range winningNums {
			for _, n := range myNums {
				if num == n {
					countMap[index] += countMap[i]
					index++
				}
			}
		}
	}

	sum := 0
	for _, value := range countMap {
		sum += value
	}

	return sum
}
