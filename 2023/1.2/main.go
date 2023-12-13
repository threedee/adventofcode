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

	var forward, backward int
	result := 0

	for scanner.Scan() {
		runes := []rune(scanner.Text())
		forward, backward = -1, -1
		for i := 0; i < len(runes); i++ {
			j := len(runes) - 1 - i
			if forward == -1 {
				forward = checkDigit(runes, i)
			}
			if backward == -1 {
				backward = checkDigit(runes, j)
			}
			if forward != -1 && backward != -1 {
				break
			}

		}
		result += forward*10 + backward
		fmt.Printf("%s - %d%d = %d\n", scanner.Text(), forward, backward, result)
	}
	fmt.Printf("Result: %d\n", result)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func checkDigit(runes []rune, i int) int {
	result := -1
	var err error
	if unicode.IsDigit(runes[i]) {
		result, err = strconv.Atoi(string(runes[i]))
		if err != nil {
			log.Fatal(err)
		}
		return result
	}
	if i < len(runes)-2 {
		if string(runes[i:i+3]) == "one" {
			return 1
		} else if string(runes[i:i+3]) == "six" {
			return 6
		} else if string(runes[i:i+3]) == "two" {
			return 2
		}
	}
	if i < len(runes)-4 {
		if string(runes[i:i+5]) == "three" {
			return 3
		} else if string(runes[i:i+5]) == "seven" {
			return 7
		} else if string(runes[i:i+5]) == "eight" {
			return 8
		}
	}
	if i < len(runes)-3 {
		if string(runes[i:i+4]) == "four" {
			return 4
		} else if string(runes[i:i+4]) == "five" {
			return 5
		} else if string(runes[i:i+4]) == "nine" {
			return 9
		}
	}

	return result
}
