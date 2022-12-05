package main

import (
	"common"
	"fmt"
)

func main() {
	lines := common.GetLines(4)
	sum := 0
	sum2 := 0
	for _, line := range lines {
		var l_l, l_r, r_l, r_r int
		fmt.Sscanf(line, "%d-%d,%d-%d", &l_l, &l_r, &r_l, &r_r)
		if (l_l <= r_l && l_r >= r_r) || (r_l <= l_l && r_r >= l_r) {
			sum += 1
		}
		if (r_r >= l_l && r_r <= l_r) || (r_l >= l_l && r_l <= l_r) || (l_l >= r_l && l_l <= r_r) || (l_r >= r_l && l_l <= r_r) {
			sum2 += 1
		}
	}
	fmt.Printf("%d\n%d\n", sum, sum2)
}
