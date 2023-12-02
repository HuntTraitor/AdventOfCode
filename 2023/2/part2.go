package main

import (
	"unicode"
	"strconv"
)

func Part2(lines []string ) int {
	total := 0
	for _, line := range lines {
		red := 0
		green := 0
		blue := 0
		max_red := 0
		max_green := 0
		max_blue := 0
		for i := 8; i < len(line); i++ {
			if line[i] == ';' || i == len(line)-1 {
				if red > max_red {
					max_red = red
				}

				if green > max_green {
					max_green = green
				}
				
				if blue > max_blue {
					max_blue = blue
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
		total += max_red * max_green * max_blue
	}
	return total
}