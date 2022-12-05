package main

import (
	"common"
	"fmt"
)

const (
	_ int = iota
	Rock
	Paper
	Scissors
)

func win(my int) int {
	return my%3 + 1
}

func lose(my int) int {
	return (my+1)%3 + 1
}

func fight(my int, enemy int) int {
	if my == enemy {
		return my + 3
	}
	if my == win(enemy) {
		return my + 6
	}
	return my
}

func fight2(my byte, enemy int) int {
	if my == 'X' {
		return lose(enemy)
	} else if my == 'Y' {
		return 3 + enemy
	} else if my == 'Z' {
		return 6 + win(enemy)
	}
	panic(0)
}

func main() {
	lines := common.GetLines(2)
	var sum int = 0
	var sum2 int = 0
	for _, line := range lines {
		var my int = int(line[2]) - 'X' + 1
		var enemy int = int(line[0]) - 'A' + 1
		result := fight(my, enemy)
		result2 := fight2(line[2], enemy)
		sum += result
		sum2 += result2
	}
	fmt.Println(sum)
	fmt.Println(sum2)
}
