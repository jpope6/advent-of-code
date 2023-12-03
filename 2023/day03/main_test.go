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
				"467..114..",
				"...*......",
				"..35..633.",
				"......#...",
				"617*......",
				".....+.58.",
				"..592.....",
				"......755.",
				"...$.*....",
				".664.598..",
			},
			expected: 4361,
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
				"467..114..",
				"...*......",
				"..35..633.",
				"......#...",
				"617*......",
				".....+.58.",
				"..592.....",
				"......755.",
				"...$.*....",
				".664.598..",
			},
			expected: 467835,
		},
	}
	for _, test := range tests {
		result := part2(test.input)

		if result != test.expected {
			t.Errorf("Expected %d, but got %d", test.expected, result)
		}
	}
}
