package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	grid := [][]rune{}
	emptyRow := strings.Repeat(".", 140)
	grid = append(grid, []rune(emptyRow))

	for scanner.Scan() {
		t := scanner.Text()
		row := "." + t + "."
		grid = append(grid, []rune(row))
	}

	grid = append(grid, []rune(emptyRow))

	count := 0

	for i := 1; i < 139; i++ {
		for j := 1; j < 139; j++ {
			if grid[i][j] != '@' {
				continue
			}
			neighbours := 0
			for ii := i - 1; ii <= i+1; ii++ {
				for jj := j - 1; jj <= j+1; jj++ {
					if grid[ii][jj] == '@' {
						neighbours++
					}
				}
			}
			if neighbours < 5 {
				count++
			}
		}
	}

	fmt.Println(count)
}

func part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)

	grid := [][]rune{}
	emptyRow := strings.Repeat(".", 140)
	grid = append(grid, []rune(emptyRow))

	for scanner.Scan() {
		t := scanner.Text()
		row := "." + t + "."
		grid = append(grid, []rune(row))
	}

	grid = append(grid, []rune(emptyRow))

	totalCount := 0

	for {
		thisCount := 0
		for i := 1; i < 139; i++ {
			for j := 1; j < 139; j++ {
				if grid[i][j] != '@' {
					continue
				}
				neighbours := 0
				for ii := i - 1; ii <= i+1; ii++ {
					for jj := j - 1; jj <= j+1; jj++ {
						if grid[ii][jj] == '@' || grid[ii][jj] == 'x' {
							neighbours++
						}
					}
				}
				if neighbours < 5 {
					thisCount++
					grid[i][j] = 'x'
				}
			}
		}
		if thisCount == 0 {
			break
		}
		totalCount += thisCount
		for i := 1; i < 139; i++ {
			for j := 1; j < 139; j++ {
				if grid[i][j] == 'x' {
					grid[i][j] = '.'
				}
			}
		}
	}

	fmt.Println(totalCount)

}
