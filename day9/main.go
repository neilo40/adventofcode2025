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

type Coord struct {
	X, Y int
}

func CoordFromString(s string) *Coord {
	parts := strings.Split(s, ",")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	return &Coord{X: x, Y: y}
}

func part1() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)

	coords := make([]*Coord, 0, 496)
	for scanner.Scan() {
		t := scanner.Text()
		coords = append(coords, CoordFromString(t))
	}

	max := 0

	for i := 0; i < len(coords); i++ {
		for j := i + 1; j < len(coords); j++ {
			size := helper.AbsInt(coords[i].X-coords[j].X+1) * helper.AbsInt(coords[i].Y-coords[j].Y+1)
			if size > max {
				max = size
			}
		}
	}

	log.Println(max)
}

func part2() {
}
