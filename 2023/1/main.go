package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
)

func main() {
	lines := LoadInput("input.txt")

	fmt.Printf("Part 1: %d\n", Part1(lines))
	fmt.Printf("Part 2: %d\n", Part2(lines))
}

func LoadInput(file_path string) []string {
	var arr []string
	file, err := os.Open(file_path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		arr = append(arr, scanner.Text())
	}
	return arr
}