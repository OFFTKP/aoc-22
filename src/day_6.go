package main

import (
	"fmt"
)

func main() {
	lines := GetLines(6)
	str := lines[0]
	start, end := 0, len(str)-1
	for i := start; i < end-4; i++ {
		cur := str[i : i+4]
		mp := make(map[byte]bool)
		found := true
		for j := 0; j < 4; j++ {
			ch := cur[j]
			if mp[ch] {
				found = false
				break
			}
			mp[ch] = true
		}
		if found {
			fmt.Printf("%d\n", i+4)
			break
		}
	}
	for i := start; i < end-14; i++ {
		cur := str[i : i+14]
		mp := make(map[byte]bool)
		found := true
		for j := 0; j < 14; j++ {
			ch := cur[j]
			if mp[ch] {
				found = false
				break
			}
			mp[ch] = true
		}
		if found {
			fmt.Printf("%d\n", i+14)
			break
		}
	}
}
