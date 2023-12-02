package main

import (
	"fmt"
	"os"
	"log"
	"strings"
)

func main() {
	lines := LoadInput("input.txt")
	fmt.Println(Part1(lines))
	fmt.Println(Part2(lines))
}

func LoadInput(file_path string) []string {
	file, err := os.ReadFile(file_path)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(file), "\n")
}

