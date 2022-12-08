package main

import (
	"fmt"
)

type Trees struct {
	arr [][]byte
}

func (t *Trees) getScore(x, y int) int {
	val := t.arr[y][x]
	r, l, u, d := 0, 0, 0, 0

	for i := x + 1; i < len(t.arr[0]); i++ {
		r++
		if t.arr[y][i] >= val {
			break
		}
	}
	for i := x - 1; i >= 0; i-- {
		l++
		if t.arr[y][i] >= val {
			break
		}
	}
	for j := y + 1; j < len(t.arr); j++ {
		d++
		if t.arr[j][x] >= val {
			break
		}
	}
	for j := y - 1; j >= 0; j-- {
		u++
		if t.arr[j][x] >= val {
			break
		}
	}
	return r * l * u * d
}

func main() {
	lines := GetLines(8)
	col_count := len(lines[0])
	row_count := len(lines)
	var arr [][]byte = make([][]byte, row_count)
	for i := 0; i < row_count; i++ {
		arr[i] = make([]byte, col_count)
		for j := 0; j < col_count; j++ {
			arr[i][j] = lines[i][j] - '0'
		}
	}
	sum := 0
	collisions := make(map[int]bool)
	for i := 1; i < row_count-1; i++ {
		cur := arr[i][0]
		// L -> R
		for j := 1; j < col_count-1; j++ {
			val := arr[i][j]
			if val > cur {
				cur = val
				if !collisions[j+i*col_count] {
					collisions[j+i*col_count] = true
					sum++
				}
			}
		}
		// R -> L
		cur = arr[i][col_count-1]
		for j := col_count - 2; j > 0; j-- {
			if arr[i][j] > cur {
				cur = arr[i][j]
				if !collisions[j+i*col_count] {
					collisions[j+i*col_count] = true
					sum++
				}
			}
		}
	}
	for j := 1; j < col_count-1; j++ {
		cur := arr[0][j]
		// U -> D
		for i := 1; i < row_count-1; i++ {
			if arr[i][j] > cur {
				cur = arr[i][j]
				if !collisions[i*col_count+j] {
					collisions[i*col_count+j] = true
					sum++
				}
			}
		}
		// D -> U
		cur = arr[row_count-1][j]
		for i := row_count - 2; i > 0; i-- {
			if arr[i][j] > cur {
				cur = arr[i][j]
				if !collisions[i*col_count+j] {
					collisions[i*col_count+j] = true
					sum++
				}
			}
		}
	}
	sum += row_count*2 + col_count*2 - 4

	fmt.Printf("%d\n", sum)
	// Part 2
	trees := Trees{arr: arr}
	max_score := 0
	for i := 1; i < row_count-1; i++ {
		for j := 1; j < col_count-1; j++ {
			score := trees.getScore(j, i)
			if score > max_score {
				max_score = score
			}
		}
	}
	fmt.Printf("%d\n", max_score)
}
