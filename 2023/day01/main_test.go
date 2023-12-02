package main

import (
	"testing"
)

func Test_part1(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{
			input: []string{
				"1abc2",
				"pqr3stu8vwx",
				"a1b2c3d4e5f",
				"treb7uchet",
			},
			expected: 142,
		},
	}
	for _, test := range tests {
		result := part1(test.input)

		if result != test.expected {
			t.Errorf("Expected %d, but got %d", test.expected, result)
		}
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{
			input: []string{
				"two1nine",
				"eightwothree",
				"abcone2threexyz",
				"xtwone3four",
				"4nineeightseven2",
				"zoneight234",
				"7pqrstsixteen",
			},
			expected: 281,
		},
	}
	for _, test := range tests {
		result := part2(test.input)

		if result != test.expected {
			t.Errorf("Expected %d, but got %d", test.expected, result)
		}
	}
}
