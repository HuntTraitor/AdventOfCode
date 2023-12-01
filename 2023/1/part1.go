package main

import(
	"unicode"
	"strings"
	"strconv"
)

func Part1(lines []string) int {
	total := 0
	for _, line := range lines {
		var first string
		var last string
		for _, char := range line {
			if !unicode.IsDigit(char) {
				line = strings.Replace(line, string(char), "", -1)
			}
		}
		first = string(line[0])
		last = string(line[len(line)-1])
		temp_total, _ := strconv.Atoi(first + last)
		total += temp_total
	}
	return total
}