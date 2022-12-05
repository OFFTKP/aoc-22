package main

import (
	"common"
	"fmt"
	"log"
	"strings"
)

func getPrio(c int) int {
	if c >= 'a' && c <= 'z' {
		return c - 'a' + 1
	} else if c >= 'A' && c <= 'Z' {
		return 26 + c - 'A' + 1
	} else {
		panic(0)
	}
}

func main() {
	lines := common.GetLines(3)
	sum := 0
	sum2 := 0
	for _, line := range lines {
		mid := len(line) / 2
		left := line[0:mid]
		right := line[mid:]
		found := false
		for i := 0; i < mid; i++ {
			if strings.ContainsRune(left, rune(right[i])) {
				sum += getPrio(int(right[i]))
				found = true
				break
			}
		}
		if !found {
			log.Fatalf("\n%s\n%s", left, right)
		}
	}
	for i := 0; i < len(lines); i += 3 {
		line1 := lines[i]
		line2 := lines[i+1]
		line3 := lines[i+2]
		max := common.MaxOf(len(line1), len(line2), len(line3))
		m := make(map[int]int, max)
		for _, c := range line1 {
			if m[int(c)] == 0 {
				m[int(c)] = 1
			}
		}
		for _, c := range line2 {
			if m[int(c)] == 1 {
				m[int(c)] = 2
			}
		}
		for _, c := range line3 {
			if m[int(c)] == 2 {
				m[int(c)] = 3
			}
		}
		for ch, val := range m {
			if val == 3 {
				sum2 += getPrio(ch)
				break
			}
		}
	}
	fmt.Printf("%d\n%d\n", sum, sum2)
}
