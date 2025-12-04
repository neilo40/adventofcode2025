package main

import (
	"bufio"
	"log"
	"os"

	"github.com/neilo40/adventofcode2025/helper"
)

func main() {
	helper.DownloadInput()
	part1()
	part2()
}

func part1() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)

	sum := 0

	for scanner.Scan() {
		t := scanner.Text()
		firstDigit := 0
		firstDigitPos := len(t)
		for i := len(t) - 2; i >= 0; i-- {
			n := int(t[i] - '0') // convert ascii rune to int
			if n >= firstDigit {
				firstDigit = n
				firstDigitPos = i
			}
		}

		secondDigit := 0
		for i := firstDigitPos + 1; i < len(t); i++ {
			n := int(t[i] - '0') // convert ascii rune to int
			if n > secondDigit {
				secondDigit = n
			}
		}

		sum += (firstDigit * 10) + secondDigit
	}

	log.Println("part 1:", sum)
}

func part2() {
}
