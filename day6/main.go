package main

import (
	"bufio"
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
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)

	operands1 := make([]int, 0)
	operands2 := make([]int, 0)
	operands3 := make([]int, 0)
	operands4 := make([]int, 0)

	scanner.Scan()
	t := scanner.Text()
	parts := strings.Fields(t)
	for _, p := range parts {
		num, _ := strconv.Atoi(p)
		operands1 = append(operands1, num)
	}

	scanner.Scan()
	t = scanner.Text()
	parts = strings.Fields(t)
	for _, p := range parts {
		num, _ := strconv.Atoi(p)
		operands2 = append(operands2, num)
	}

	scanner.Scan()
	t = scanner.Text()
	parts = strings.Fields(t)
	for _, p := range parts {
		num, _ := strconv.Atoi(p)
		operands3 = append(operands3, num)
	}

	scanner.Scan()
	t = scanner.Text()
	parts = strings.Fields(t)
	for _, p := range parts {
		num, _ := strconv.Atoi(p)
		operands4 = append(operands4, num)
	}

	sum := 0
	scanner.Scan()
	t = scanner.Text()
	parts = strings.Fields(t)
	for i, p := range parts {
		if p == "+" {
			sum += operands1[i] + operands2[i] + operands3[i] + operands4[i]
		} else {
			sum += (operands1[i] * operands2[i] * operands3[i] * operands4[i])
		}
	}

	log.Println(sum)
}

func part2() {
}
