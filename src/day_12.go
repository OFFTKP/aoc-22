package main

import (
	"fmt"
	"sort"
)

type Position struct {
	X int
	Y int
}

type Hill struct {
	Height  byte
	Pos     Position
	Visited bool
	Parent  *Hill
}

const (
	Up int = iota
	Right
	Down
	Left
)

var start Position
var end Position
var grid [][]Hill
var path [][]bool

func (h *Hill) getNeighbor(dir int) (bool, *Hill) {
	if h.Pos.X == 0 && dir == Left {
		return false, nil
	} else if h.Pos.X == len(grid[0])-1 && dir == Right {
		return false, nil
	} else if h.Pos.Y == 0 && dir == Up {
		return false, nil
	} else if h.Pos.Y == len(grid)-1 && dir == Down {
		return false, nil
	}
	switch dir {
	case Up:
		temp := &grid[h.Pos.Y-1][h.Pos.X]
		if h.Height+1 >= temp.Height {
			return true, temp
		}
	case Right:
		temp := &grid[h.Pos.Y][h.Pos.X+1]
		if h.Height+1 >= temp.Height {
			return true, temp
		}
	case Down:
		temp := &grid[h.Pos.Y+1][h.Pos.X]
		if h.Height+1 >= temp.Height {
			return true, temp
		}
	case Left:
		temp := &grid[h.Pos.Y][h.Pos.X-1]
		if h.Height+1 >= temp.Height {
			return true, temp
		}
	}
	return false, nil
}

func (h *Hill) getNeighbors() []*Hill {
	var neighbors []*Hill
	for i := 0; i < 4; i++ {
		b, neighbor := h.getNeighbor(i)
		if b {
			neighbors = append(neighbors, neighbor)
		}
	}
	sort.Slice(neighbors, func(i, j int) bool {
		return neighbors[i].Height < neighbors[j].Height
	})
	return neighbors
}

type Queue []*Hill

func (q *Queue) Len() int { return len(*q) }

func (q *Queue) Push(h *Hill) {
	*q = append(*q, h)
}

func (q *Queue) Pop() *Hill {
	h := (*q)[0]
	*q = (*q)[1:]
	return h
}

func search(strt Position) bool {
	var q Queue
	grid[strt.Y][strt.X].Visited = true
	q.Push(&grid[strt.Y][strt.X])
	for q.Len() > 0 {
		h := q.Pop()
		if h.Pos == end {
			return true
		}
		for _, neighbor := range h.getNeighbors() {
			if !neighbor.Visited {
				neighbor.Visited = true
				neighbor.Parent = h
				q.Push(neighbor)
			}
		}
	}
	return false
}

func search_impl(strt Position) int {
	// Reset visited status
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			grid[i][j].Visited = false
		}
	}
	b := search(strt)
	if b {
		count := 0
		h := &grid[end.Y][end.X]
		for {
			if h.Pos.X == strt.X && h.Pos.Y == strt.Y {
				break
			}
			path[h.Pos.Y][h.Pos.X] = true
			h = h.Parent
			count++
			if count > 1000 {
				break
			}
		}
		return count
	}
	return 0
}

func init() {
	lines := GetLines(12)
	grid = make([][]Hill, len(lines))
	path = make([][]bool, len(lines))
	for i, line := range lines {
		grid[i] = make([]Hill, len(line))
		path[i] = make([]bool, len(line))
		for j := 0; j < len(line); j++ {
			if line[j] == 'S' {
				start = Position{j, i}
				grid[i][j].Height = 'a'
			} else if line[j] == 'E' {
				end = Position{j, i}
				grid[i][j].Height = 'z'
			} else {
				grid[i][j].Height = line[j]
			}
			grid[i][j].Pos = Position{j, i}
		}
	}
}

func main() {
	// Part 1
	fmt.Printf("Part 1: %d\n", search_impl(start))
	// Part 2
	var positions []Position
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j].Height == 'a' {
				positions = append(positions, Position{j, i})
			}
		}
	}
	var results []int
	for _, j := range positions {
		n := search_impl(j)
		if n > 0 {
			results = append(results, n)
		}
	}
	sort.Ints(results)
	fmt.Printf("Part 2: %d\n", results[0])
}
