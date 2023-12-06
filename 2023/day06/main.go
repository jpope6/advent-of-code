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
	ans := 1

	times := strings.Fields(input[0])[1:]
	distances := strings.Fields(input[1])[1:]

	for i := 0; i < len(times); i++ {
		count := 0
		distance, _ := strconv.Atoi(distances[i])
		time, _ := strconv.Atoi(times[i])
		for j := 0; j < time; j++ {
			d := j * (time - j)

			if d > distance {
				count++
			}
		}
		ans *= count
		fmt.Println(count)
	}

	return ans
}

func part2(input []string) int {
	times := strings.Fields(input[0])[1:]
	distances := strings.Fields(input[1])[1:]
	t := strings.Join(times, "")
	d := strings.Join(distances, "")

	count := 0
	time, _ := strconv.Atoi(t)
	distance, _ := strconv.Atoi(d)

	for j := 0; j < time; j++ {
		d := j * (time - j)

		if d > distance {
			count++
		}
	}

	return count
}
