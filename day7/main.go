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

	beams := make([]bool, 141)
	nextBeams := make([]bool, 141)
	splitCount := 0

	for scanner.Scan() {
		t := scanner.Text()
		for i, r := range t {
			if r == 'S' {
				nextBeams[i] = true
			}
			if r == '^' && beams[i] {
				nextBeams[i+1] = true
				nextBeams[i-1] = true
				splitCount++
			}
			if r == '.' && beams[i] {
				nextBeams[i] = true
			}
		}
		beams = nextBeams
		nextBeams = make([]bool, 141)
	}

	log.Println(splitCount)
}

func part2() {
}
