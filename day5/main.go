package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/neilo40/adventofcode2025/helper"
)

func main() {
	helper.DownloadInput()
	part1()
	part2()
}

type Rule struct {
	start int
	end   int
}

func RuleFromString(rs string) Rule {
	parts := strings.Split(rs, "-")
	start, _ := strconv.Atoi(parts[0])
	end, _ := strconv.Atoi(parts[1])
	return Rule{start: start, end: end}
}

func (r *Rule) IsInRange(num int) bool {
	return num >= r.start && num <= r.end
}

func (r *Rule) InRangeFromHere(num int) (int, int) {
	if num > r.end {
		return 0, num
	}
	if num > r.start {
		return (r.end - num) + 1, r.end + 1
	}
	return (r.end - r.start) + 1, r.end + 1
}

func part1() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)

	rules := make([]Rule, 0, 190)
	for scanner.Scan() {
		t := scanner.Text()
		if t == "" {
			break
		}
		rules = append(rules, RuleFromString(t))
	}

	freshCount := 0

	for scanner.Scan() {
		t := scanner.Text()
		num, _ := strconv.Atoi(t)
		for _, r := range rules {
			if r.IsInRange(num) {
				freshCount++
				break
			}
		}
	}

	log.Println(freshCount)
}

func part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)

	rules := make([]Rule, 0, 190)
	for scanner.Scan() {
		t := scanner.Text()
		if t == "" {
			break
		}
		rules = append(rules, RuleFromString(t))
	}

	sort.Slice(rules, func(i, j int) bool {
		return rules[i].start < rules[j].start
	})

	freshCount := 0
	curNum := 0
	fresh := 0
	for _, r := range rules {
		fresh, curNum = r.InRangeFromHere(curNum)
		freshCount += fresh
	}

	log.Println(freshCount)
}
