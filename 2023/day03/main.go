package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

// Pair represents a 2D coordinate pair.
type Pair struct {
	X, Y int
}

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

	directions := []Pair{
		{-1, 0},  // above
		{1, 0},   // below
		{0, -1},  // left
		{0, 1},   // right
		{-1, -1}, // top-left
		{-1, 1},  // top-right
		{1, -1},  // bottom-left
		{1, 1},   // bottom-right
	}

	for i, line := range input {
		numStr, isPartNum := "", false
		for j, symbol := range line {
			if !unicode.IsDigit(rune(symbol)) {
				if isPartNum {
					num, _ := strconv.Atoi(numStr)
					sum += num
				}

				numStr = ""
				isPartNum = false
				continue
			}

			rows := len(input)
			cols := len(input[0])

			numStr += string(symbol)

			for _, dir := range directions {
				newI, newJ := i+dir.X, j+dir.Y

				if newI >= 0 && newI < rows && newJ >= 0 && newJ < cols {
					if isSymbol(input[newI][newJ]) {
						isPartNum = true
					}
				}
			}
		}

		if isPartNum {
			num, _ := strconv.Atoi(numStr)
			sum += num
		}
	}

	return sum
}

func part2(input []string) int {
	sum := 0

	directions := []Pair{
		{-1, 0},  // above
		{1, 0},   // below
		{0, -1},  // left
		{0, 1},   // right
		{-1, -1}, // top-left
		{-1, 1},  // top-right
		{1, -1},  // bottom-left
		{1, 1},   // bottom-right
	}

	numMap := make(map[Pair]int)

	for i, line := range input {
		for j, symbol := range line {
			if symbol != '*' {
				continue
			}

			rows := len(input)
			cols := len(input[0])
			numCount := 0
			numStr := ""
			product := 1

			for _, dir := range directions {
				newI, newJ := i+dir.X, j+dir.Y

				if newI >= 0 && newI < rows && newJ >= 0 && newJ < cols {
					if unicode.IsDigit(rune(input[newI][newJ])) {
						if _, exists := numMap[Pair{newI, newJ}]; exists {
							continue
						}

						numCount++

						// Get the start of the num
						for newJ >= 0 && unicode.IsDigit(rune(input[newI][newJ])) {
							newJ--
						}

						// Get the entire num
						for newJ+1 < len(input[newI]) && unicode.IsDigit(rune(input[newI][newJ+1])) {
							newJ++
							numStr += string(input[newI][newJ])

							pair := Pair{newI, newJ}
							numMap[pair]++
						}

						num, _ := strconv.Atoi(numStr)
						numStr = ""
						product *= num
					}
				}
			}

			if numCount == 2 {
				sum += product
			}

			numStr = ""
			product = 1
			numCount = 0

		}
	}

	return sum
}

func isSymbol(char byte) bool {
	return char != '.' && !unicode.IsDigit(rune(char))
}
