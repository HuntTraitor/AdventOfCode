package main 

import(
	"unicode"
	"strconv"
)

func Part2(lines []string) int {
	total := 0
	for _, line := range lines {
		var first string
		var last string
		first_checked := false
		for i, char := range line {
			if unicode.IsDigit(char) {
				if !first_checked {
					first = string(char)
					first_checked = true
				}
				last = string(char)
			}

			switch char {
			case 'o':
				temp := line[i:]
				if CheckWord("one", temp) {
					if !first_checked {
						first = "1"
						first_checked = true
					}
					last = "1"
				}
			
			case 't':
				temp := line[i:]
				if CheckWord("two", temp) {
					if !first_checked {
						first = "2"
						first_checked = true
					}
					last = "2"
				} else if CheckWord("three", temp) {
					if !first_checked {
						first = "3"
						first_checked = true
					}
					last = "3"
				}
			
			case 'f':
				temp := line[i:]
				if CheckWord("four", temp) {
					if !first_checked {
						first = "4"
						first_checked = true
					}
					last = "4"
				} else if CheckWord("five", temp) {
					if !first_checked {
						first = "5"
						first_checked = true
					}
					last = "5"
				}
	
			case 's':
				temp := line[i:]
				if CheckWord("six", temp) {
					if !first_checked {
						first = "6"
						first_checked = true
					}
					last = "6"
				} else if CheckWord("seven", temp) {
					if !first_checked {
						first = "7"
						first_checked = true
					}
					last = "7"
				}
	
			case 'e':
				temp := line[i:]
				if CheckWord("eight", temp) {
					if !first_checked {
						first = "8"
						first_checked = true
					}
					last = "8"
				}
	
			case 'n':
				temp := line[i:]
				if CheckWord("nine", temp) {
					if !first_checked {
						first = "9"
						first_checked = true
					}
					last = "9"
				}
			}
		}
		temp_total, _ := strconv.Atoi(first + last)
		total += temp_total
	}
	return total
}

func CheckWord(word string, arr string) bool {
	for i := range arr {
		if !(word[i] == arr[i]) {
			return false
		}
		if i == len(word)-1 {
			return true
		}
	}
	return false
}