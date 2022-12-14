package main

import (
	"fmt"
	"strconv"
	"strings"
)

type State int

const (
	Empty State = iota
	Wall
	Sand
)

var space [][]State
var lowestY int
var restCount int

func main() {
	lines := GetLines(14)
	space = make([][]State, 300)
	for i := 0; i < len(space); i++ {
		space[i] = make([]State, 1000)
	}
	for _, line := range lines {
		split := strings.Split(line, " -> ")
		var start_x, start_y int
		fmt.Sscanf(split[0], "%d,%d", &start_x, &start_y)
		for j := 1; j < len(split); j++ {
			if start_y > lowestY {
				lowestY = start_y
			}
			var x, y int
			fmt.Sscanf(split[j], "%d,%d", &x, &y)
			signx := sign(x - start_x)
			signy := sign(y - start_y)
			if signx != 0 && signy == 0 {
				for k := start_x; k != x; k += signx {
					space[start_y][k] = Wall
				}
				space[y][x] = Wall
			} else if signy != 0 && signx == 0 {
				for k := start_y; k != y; k += signy {
					space[k][start_x] = Wall
				}
				space[y][x] = Wall
			} else if signx == 0 && signy == 0 {
				space[start_y][start_x] = Wall
			} else {
				for k := start_x; k != x; k += signx {
					for l := start_y; l != y; l += signy {
						space[l][k] = Wall
					}
				}
				space[y][x] = Wall
			}
			start_x = x
			start_y = y
			if start_y > lowestY {
				lowestY = start_y
			}
		}
	}
	lowestY += 2
	fmt.Println("Floor:" + strconv.Itoa(lowestY))
	for {
		// Change to false for part 1
		if fallSand(true) {
			break
		}
	}
	// printSpace()
	fmt.Println(restCount)
}

func fallSand(part2 bool) bool {
	for {
		b := fallSingleSand(part2)
		if b {
			return true
		}
	}
}

func fallSingleSand(part2 bool) bool {
	sand_x := 500
	sand_y := 0
	for {
		if space[sand_y+1][sand_x] == Empty {
			sand_y++
			if !part2 {
				if sand_y > lowestY {
					return true
				}
			}
		} else if space[sand_y+1][sand_x-1] == Empty {
			sand_y++
			sand_x--
		} else if space[sand_y+1][sand_x+1] == Empty {
			sand_y++
			sand_x++
		} else {
			space[sand_y][sand_x] = Sand
			restCount++
			if part2 {
				if sand_y == 0 && sand_x == 500 {
					return true
				}
			}
			return false
		}
		if part2 {
			if sand_y == lowestY {
				space[sand_y][sand_x] = Sand
				return false
			}
		}
	}
}

func printSpace() {
	for y := 0; y < lowestY; y++ {
		for x := 200; x < 800; x++ {
			switch space[y][x] {
			case Empty:
				fmt.Print(" ")
			case Wall:
				fmt.Print("#")
			case Sand:
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
