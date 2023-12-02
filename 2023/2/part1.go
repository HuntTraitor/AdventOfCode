package main

import (
	"unicode"
	"strconv"
	"regexp"
)

func Part1(lines []string ) int {
	total := 0
	for _, line := range lines {
		red := 0
		green := 0
		blue := 0
		id := GetID(line)
		passed := true

		for i := 8; i < len(line); i++ {
			if line[i] == ';' || i == len(line)-1 {
				if !CheckValidity(red, green, blue) && passed == true {
					passed = false
				}
				red = 0
				green = 0
				blue = 0
			}

			if unicode.IsDigit(rune(line[i])) {
				number := string(line[i])
				temp := 1

				for {
					if unicode.IsDigit(rune(line[i+temp])) {
						number += string(line[i+temp])
						i++
					} else {
						break
					}
				}

				finalnum, _ := strconv.Atoi(number)
				switch line[i+2] {
				case 'r':
					red += finalnum
				case 'g':
					green += finalnum
				case 'b':
					blue += finalnum
				}
			}
		}
		if passed == true {
			total += id
		}
	}
	return total
}

func CheckValidity(red int, green int, blue int) bool {
	if red > 12 {
		return false
	}
	if green > 13 {
		return false
	}
	if blue > 14 {
		return false
	}
	return true
}

func GetID(input string) int {
	pattern := regexp.MustCompile(`Game (\d+)`)
	match := pattern.FindStringSubmatch(input)
	id, _ := strconv.Atoi(match[1])
	return id
}