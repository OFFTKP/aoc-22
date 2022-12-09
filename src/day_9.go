package main

import (
	"fmt"
)

type Position struct {
	x int
	y int
}

var head Position
var tail Position
var tails [9]Position
var touched map[Position]bool
var touched2 map[Position]bool

func init() {
	head = Position{0, 0}
	tail = Position{0, 0}
	for i := 0; i < 9; i++ {
		tails[i] = Position{0, 0}
	}
	touched = make(map[Position]bool)
	touched2 = make(map[Position]bool)
}

func main() {
	lines := GetLines(9)

	for _, line := range lines {
		var dir byte
		var dist int
		var diradd Position
		fmt.Sscanf(line, "%c %d", &dir, &dist)
		switch dir {
		case 'R':
			diradd = Position{1, 0}
		case 'L':
			diradd = Position{-1, 0}
		case 'U':
			diradd = Position{0, 1}
		case 'D':
			diradd = Position{0, -1}
		}
		for i := 0; i < dist; i++ {
			head.x += diradd.x
			head.y += diradd.y
			resolveKnot(&head, &tail, true, false)
			resolveKnot(&head, &tails[0], false, true)
			for j := 1; j <= 7; j++ {
				resolveKnot(&tails[j-1], &tails[j], false, true)
			}
			resolveKnot(&tails[7], &tails[8], true, true)
			// visualize()
			// fmt.Scanln()
		}
	}
	fmt.Printf("%d\n", len(touched))
	fmt.Printf("%d\n", len(touched2))
}

func resolveKnot(leader, follower *Position, isTail bool, map2 bool) {
	dirx := sign(leader.x - follower.x)
	diry := sign(leader.y - follower.y)
	distx := abs(leader.x - follower.x)
	disty := abs(leader.y - follower.y)
	if distx > 1 && disty > 1 {
		follower.x += dirx
		follower.y += diry
	} else {
		if distx > 1 {
			follower.x += dirx
			if follower.y != leader.y {
				follower.y += diry
			}
		}
		if disty > 1 {
			follower.y += diry
			if follower.x != leader.x {
				follower.x += dirx
			}
		}
	}
	if isTail {
		if map2 {
			touched2[*follower] = true
		} else {
			touched[*follower] = true
		}
	}
}
