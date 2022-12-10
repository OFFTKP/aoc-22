package main

import "fmt"

var cycles, x int = 0, 1
var sum int = 0
var screen [40][6]byte

func draw() {
	screenx := (cycles - 1) % 40
	screeny := (cycles - 1) / 40
	spritex := x % 40
	if screenx >= spritex-1 && screenx <= spritex+1 {
		screen[screenx][screeny] = '#'
	} else {
		screen[screenx][screeny] = '.'
	}
}

func main() {
	lines := GetLines(10)
	for _, line := range lines {
		val := 0
		instr := ""
		fmt.Sscanf(line, "%s %d", &instr, &val)
		if instr == "noop" {
			cycles++
		} else if instr == "addx" {
			cycles++
		}
		if cycles == 20 || cycles == 60 || cycles == 100 ||
			cycles == 140 || cycles == 180 || cycles == 220 {
			sum += x * cycles
			fmt.Printf("cycle %d: x: %d, sum: %d\n", cycles, x, sum)
		}
		draw()
		if instr == "addx" {
			cycles++
			if cycles == 20 || cycles == 60 || cycles == 100 ||
				cycles == 140 || cycles == 180 || cycles == 220 {
				sum += x * cycles
				fmt.Printf("cycle %d: x: %d, sum: %d\n", cycles, x, sum)
			}
			draw()
			x += val
		}
	}
	fmt.Println(sum)
	for y := 0; y < 6; y++ {
		for x := 0; x < 40; x++ {
			fmt.Printf("%c", screen[x][y])
		}
		fmt.Println()
	}
}
