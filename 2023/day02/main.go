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
	var red, green, blue int = 12, 13, 14
	sum := 0

	for i, s := range input {
		s := strings.TrimPrefix(s, "Game "+fmt.Sprint(i+1)+": ")
		var gameInfo []string = strings.Split(s, "; ")
		valid := true

		for _, round := range gameInfo {
			var colorInfo []string = strings.Split(round, ", ")

			for _, colors := range colorInfo {
				var info []string = strings.Split(colors, " ")
				color := info[1]
				num, err := strconv.Atoi(info[0])
				if err != nil {
					fmt.Println("Error converting string to int:", err)
					fmt.Println(colorInfo)
					return 0
				}

				switch color {
				case "red":
					if num > red {
						valid = false
						break
					}
				case "green":
					if num > green {
						valid = false
						break
					}
				case "blue":
					if num > blue {
						valid = false
						break
					}
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
	red, green, blue := 0, 0, 0
	sum := 0

	for i, s := range input {
		s := strings.TrimPrefix(s, "Game "+fmt.Sprint(i+1)+": ")
		var gameInfo []string = strings.Split(s, "; ")

		for _, round := range gameInfo {
			var colorInfo []string = strings.Split(round, ", ")

			for _, colors := range colorInfo {
				var info []string = strings.Split(colors, " ")
				color := info[1]
				num, err := strconv.Atoi(info[0])
				if err != nil {
					fmt.Println("Error converting string to int:", err)
					fmt.Println(colorInfo)
					return 0
				}

				switch color {
				case "red":
					red = max(red, num)
				case "green":
					green = max(green, num)
				case "blue":
					blue = max(blue, num)
				}
			}
		}

		sum += (red * green * blue)
		red, green, blue = 0, 0, 0
	}

	return sum
}
