package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var forward, backward string
	result := 0

	for scanner.Scan() {
		forward, backward = "", ""
		for _, rune := range scanner.Text() {
			if unicode.IsDigit(rune) {
				forward = string(rune)
				break
			}
		}
		runes := []rune(scanner.Text())
		for i := len(runes) - 1; i >= 0; i-- {
			if unicode.IsDigit(runes[i]) {
				backward = string(runes[i])
				break
			}
		}
		code, err := strconv.Atoi(string(forward) + string(backward))
		if err != nil {
			log.Fatal(err)
		}
		result += code
	}
	fmt.Printf("Result: %d\n", result)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
