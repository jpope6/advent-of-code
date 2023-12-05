package main

import (
	"bufio"
	"fmt"
	"math"
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
	// seeds := []int{3082872446, 316680412, 2769223903, 74043323, 4131958457,
	// 	99539464, 109726392, 353536902, 619902767, 648714498,
	// 	3762874676, 148318192, 1545670780, 343889780, 4259893555,
	// 	6139816, 3980757676, 20172062, 2199623551, 196958359}
	seeds := []int{79, 14, 55, 13}

	sToS := [][]int{}
	sToF := [][]int{}
	fToW := [][]int{}
	wToL := [][]int{}
	lToT := [][]int{}
	tToH := [][]int{}
	hToL := [][]int{}

	for i, line := range input {
		switch line {
		case "seed-to-soil map:":
			i = parseInput(input[i+1:], &sToS)
		case "soil-to-fertilizer map:":
			i = parseInput(input[i+1:], &sToF)
		case "fertilizer-to-water map:":
			i = parseInput(input[i+1:], &fToW)
		case "water-to-light map:":
			i = parseInput(input[i+1:], &wToL)
		case "light-to-temperature map:":
			i = parseInput(input[i+1:], &lToT)
		case "temperature-to-humidity map:":
			i = parseInput(input[i+1:], &tToH)
		case "humidity-to-location map:":
			i = parseInput(input[i+1:], &hToL)
		}
	}

	smallest := math.MaxInt
	for _, seed := range seeds {
		soil := findOutput(seed, &sToS)
		fert := findOutput(soil, &sToF)
		water := findOutput(fert, &fToW)
		light := findOutput(water, &wToL)
		temp := findOutput(light, &lToT)
		hum := findOutput(temp, &tToH)
		loc := findOutput(hum, &hToL)

		smallest = min(smallest, loc)
	}

	return smallest
}

func part2(input []string) int {
	seeds := []int{3082872446, 316680412, 2769223903, 74043323, 4131958457,
		99539464, 109726392, 353536902, 619902767, 648714498,
		3762874676, 148318192, 1545670780, 343889780, 4259893555,
		6139816, 3980757676, 20172062, 2199623551, 196958359}
	// seeds := []int{79, 14, 55, 13}

	sToS := [][]int{}
	sToF := [][]int{}
	fToW := [][]int{}
	wToL := [][]int{}
	lToT := [][]int{}
	tToH := [][]int{}
	hToL := [][]int{}

	for i, line := range input {
		switch line {
		case "seed-to-soil map:":
			i = parseInput(input[i+1:], &sToS)
		case "soil-to-fertilizer map:":
			i = parseInput(input[i+1:], &sToF)
		case "fertilizer-to-water map:":
			i = parseInput(input[i+1:], &fToW)
		case "water-to-light map:":
			i = parseInput(input[i+1:], &wToL)
		case "light-to-temperature map:":
			i = parseInput(input[i+1:], &lToT)
		case "temperature-to-humidity map:":
			i = parseInput(input[i+1:], &tToH)
		case "humidity-to-location map:":
			i = parseInput(input[i+1:], &hToL)
		}
	}

	smallest := math.MaxInt
	for i := 0; i < len(seeds); i += 2 {
		seed := seeds[i]
		r := seeds[i+1]

		for j := seed; j < seed+r; j++ {
			soil := findOutput(j, &sToS)
			fert := findOutput(soil, &sToF)
			water := findOutput(fert, &fToW)
			light := findOutput(water, &wToL)
			temp := findOutput(light, &lToT)
			hum := findOutput(temp, &tToH)
			loc := findOutput(hum, &hToL)

			smallest = min(smallest, loc)
		}
	}

	return smallest
}

func parseInput(input []string, inputArray *[][]int) int {
	i := 0
	for i < len(input) && input[i] != "" {
		s := strings.Fields(input[i])
		i++

		dest, _ := strconv.Atoi(s[0])
		src, _ := strconv.Atoi(s[1])
		range_len, _ := strconv.Atoi(s[2])

		arr := []int{dest, src, range_len}

		*inputArray = append(*inputArray, arr)
	}
	return i
}

func findOutput(num int, inputArray *[][]int) int {
	output := num

	for _, line := range *inputArray {
		dest := line[0]
		src := line[1]
		range_len := line[2]

		if num >= src && num < src+range_len {
			offset := num - src
			output = dest + offset
			return output
		}
	}

	return output
}
