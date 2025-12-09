package main

import (
	"bufio"
	"container/heap"
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

type Point struct {
	X, Y, Z int
}

func (p *Point) DistanceTo(p2 *Point) int {
	// we don't care about the actual distance, so skip the sqrt
	return helper.AbsInt(helper.SqInt(p2.X-p.X) + helper.SqInt(p2.Y-p.Y) + helper.SqInt(p2.Z-p.Z))
}

func PointFromString(s string) *Point {
	parts := strings.Split(s, ",")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	z, _ := strconv.Atoi(parts[2])
	return &Point{X: x, Y: y, Z: z}
}

type Pair struct {
	A, B     *Point
	Distance int
}

type PairHeap []Pair

func (h PairHeap) Len() int           { return len(h) }
func (h PairHeap) Less(i, j int) bool { return h[i].Distance < h[j].Distance }
func (h PairHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *PairHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Pair))
}

func (h *PairHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func part1() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)

	// 1000 choose 2 is 500k distance calcs so no need to be fancy here

	pts := make([]Point, 0, 1000)
	for scanner.Scan() {
		t := scanner.Text()
		pts = append(pts, *PointFromString(t))
	}

	ph := &PairHeap{}
	heap.Init(ph)

	for i := 0; i < len(pts); i++ {
		for j := i + 1; j < len(pts); j++ {
			p1 := pts[i]
			p2 := pts[j]
			heap.Push(ph, Pair{A: &p1, B: &p2, Distance: p1.DistanceTo(&p2)})
		}
	}

	for range 1000 {
		p := heap.Pop(ph)
		pAsPair := p.(Pair)

	}
}

func part2() {
}
