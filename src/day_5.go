package main

import (
	"common"
	"fmt"
	"log"
	"strconv"
)

type Move struct {
	Count int
	From  int
	To    int
}

type Stack []byte

func (s *Stack) Push(v byte) {
	*s = append(*s, v)
}

func (s *Stack) Pop() byte {
	ret := (*s)[len(*s)-1]
	*s = (*s)[0 : len(*s)-1]
	return ret
}

func (s *Stack) Top() byte {
	return (*s)[len(*s)-1]
}

func main() {
	lines := common.GetLines(5)
	emptyi := 0
	for i, line := range lines {
		// Find empty line
		if line == "" {
			emptyi = i
			break
		}
	}
	ind := emptyi - 2
	line := lines[emptyi-1]
	stacksize, _ := strconv.Atoi(string(line[len(line)-2]))
	stack := make([]Stack, stacksize)
	stack2 := make([]Stack, stacksize)
	for i := ind; i >= 0; i-- {
		cur := lines[i]
		for j := 0; j < stacksize; j++ {
			bt := cur[4*(j+1)-3]
			if bt != ' ' {
				stack[j].Push(bt)
				stack2[j].Push(bt)
			}
		}
	}
	for i := emptyi + 1; i < len(lines); i++ {
		cur := lines[i]
		var move Move
		n, _ := fmt.Sscanf(cur, "move %d from %d to %d", &move.Count, &move.From, &move.To)
		move.To -= 1
		move.From -= 1
		if n != 3 {
			log.Fatalf("Failed to parse move: %s", cur)
		}
		for j := 0; j < move.Count; j++ {
			stack[move.To].Push(stack[move.From].Pop())
		}
		var temp Stack
		for j := 0; j < move.Count; j++ {
			temp.Push(stack2[move.From].Pop())
		}
		for k := 0; k < move.Count; k++ {
			stack2[move.To].Push(temp[len(temp)-1-k])
		}
	}
	var res, res2 string
	for i := 0; i < stacksize; i++ {
		res += string(stack[i].Top())
		res2 += string(stack2[i].Top())
	}

	fmt.Printf("%s\n%s\n", res, res2)
}
