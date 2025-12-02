package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/neilo40/adventofcode2025/helper"
)

func main() {
	helper.DownloadInput()
	part1()
	part2()
}

func part1() {
	partBoth(isInvalidPart1)
}

func part2() {
	partBoth(isInvalidPart2)
}

func partBoth(pred func(int) bool) {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	t := scanner.Text()
	ranges := strings.Split(t, ",")
	sum := 0
	for _, r := range ranges {
		parts := strings.Split(r, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])
		for i := start; i <= end; i++ {
			if pred(i) {
				sum += i
			}
		}
	}

	log.Printf("Sum: %d\n", sum)
}

func isInvalidPart1(num int) bool {
	numStr := fmt.Sprintf("%d", num)
	return numStr[0:len(numStr)/2] == numStr[len(numStr)/2:]
}

func isInvalidPart2(num int) bool {
	numStr := fmt.Sprintf("%d", num)
	subStrLen := len(numStr) / 2
	for ssl := subStrLen; ssl > 0; ssl-- {
		parts, err := splitStringNParts(numStr, ssl)
		if err != nil {
			continue
		}

		allSame := true
		for _, p := range parts[1:] {
			if p != parts[0] {
				allSame = false
				break
			}
		}

		if allSame {
			return true
		}
	}
	return false
}

func splitStringNParts(s string, n int) ([]string, error) {
	lenParts := 0
	parts := make([]string, 0)

	for i := 0; i <= len(s)-n; i += n {
		parts = append(parts, s[i:i+n])
		lenParts += n
	}

	if lenParts != len(s) {
		return nil, errors.New("string not evenly divisible")
	} else {
		return parts, nil
	}
}
