package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	lines := GetLines(1)
	calories := make([]int, len(lines))
	index := 0
	for _, str := range lines {
		if str == "" {
			index++
			continue
		}
		n, _ := strconv.Atoi(str)
		calories[index] += n
	}
	sort.Slice(calories, func(i, j int) bool {
		return calories[i] > calories[j]
	})
	fmt.Println(calories[0])
	fmt.Println(calories[0] + calories[1] + calories[2])
}
