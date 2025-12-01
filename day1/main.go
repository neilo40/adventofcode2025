package main

import (
	"bufio"
	"container/ring"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/neilo40/adventofcode2025/helper"
)

func main() {
	helper.DownloadInput()
	part1And2()
}

func part1And2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	r := ring.New(100)
	for i := range 100 {
		r.Value = i
		r = r.Next()
	}

	for range 50 {
		r = r.Next() // start at pos 50
	}

	zeroCountPart1 := 0
	zeroCountPart2 := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		t := scanner.Text()
		dir := t[0]
		num, _ := strconv.Atoi(string(t[1:]))
		for range num {
			if dir == 'R' {
				r = r.Next()
			} else {
				r = r.Prev()
			}
			if r.Value == 0 {
				zeroCountPart2++
			}
		}

		if r.Value == 0 {
			zeroCountPart1++
		}
	}

	fmt.Println("part 1: ", zeroCountPart1)
	fmt.Println("part 2: ", zeroCountPart2)
}
