package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
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

func part1(arr []string) int {
	sum := 0

	for _, s := range arr {
		var num int

		for _, char := range s {
			if unicode.IsDigit(char) {
				num = int(char - '0')
				break
			}
		}

		for i := len(s) - 1; i >= 0; i-- {
			if unicode.IsDigit(rune(s[i])) {
				num = num*10 + int(s[i]-'0')
				break
			}
		}

		sum += num
	}

	return sum
}

func part2(input []string) int {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	sum := 0

	for _, s := range input {
		lowest, highest := len(s), -1
		lowestVal, highestVal := 0, 0

		for key := range m {
			firstIndex := strings.Index(s, key)
			lastIndex := strings.LastIndex(s, key)

			if firstIndex != -1 && firstIndex < lowest {
				lowestVal = m[key]
				lowest = firstIndex
			}

			if lastIndex != -1 && lastIndex > highest {
				highestVal = m[key]
				highest = lastIndex
			}
		}

		var num int = lowestVal
		for i, char := range s {
			if unicode.IsDigit(char) {
				if i < lowest {
					num = int(char - '0')
				}

				break
			}
		}

		num = num*10 + highestVal
		for i := len(s) - 1; i >= 0; i-- {
			if unicode.IsDigit(rune(s[i])) {
				if i > highest {
					num -= highestVal
					num += int(s[i] - '0')
				}

				break
			}
		}

		sum += num
	}

	return sum
}
