package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

type Operation int

type ItemsV []int

func (items *ItemsV) PopFront() int {
	item := (*items)[0]
	*items = (*items)[1:]
	return item
}

const (
	Increment Operation = iota
	Multiplication
	Power
)

type Monkey struct {
	Items              ItemsV
	IncrementOperation Operation
	IncrementNum       int
	TestDivisor        int
	TrueMonkey         int
	FalseMonkey        int
	TimesInspected     int
}

var superModulo int = 1

var monkeys []Monkey

func parseMonkeys(lines []string) {
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		var monkey Monkey
		i += 1
		line = lines[i]
		ssplit := strings.Split(line, ": ")[1]
		csplit := strings.Split(ssplit, ", ")
		isplit := make([]int, len(csplit))
		for j := 0; j < len(csplit); j++ {
			k, _ := strconv.Atoi(csplit[j])
			isplit[j] = int(k)
		}
		monkey.Items = isplit
		i += 1
		line = lines[i]
		var op byte
		var num string
		fmt.Sscanf(line, "  Operation: new = old %c %s", &op, &num)
		if num == "old" {
			monkey.IncrementOperation = Power
		} else {
			numi, err := strconv.Atoi(num)
			if err != nil {
				log.Fatalln(err)
			}
			switch op {
			case '+':
				monkey.IncrementOperation = Increment
			case '*':
				monkey.IncrementOperation = Multiplication
			}
			monkey.IncrementNum = int(numi)
		}
		i += 1
		line = lines[i]
		var div int
		fmt.Sscanf(line, "  Test: divisible by %d", &div)
		monkey.TestDivisor = int(div)
		superModulo *= int(div)
		i += 1
		line = lines[i]
		var true_monkey, false_monkey int
		fmt.Sscanf(line, "    If true: throw to monkey %d", &true_monkey)
		i += 1
		line = lines[i]
		fmt.Sscanf(line, "    If false: throw to monkey %d", &false_monkey)
		i += 1
		monkey.TrueMonkey = true_monkey
		monkey.FalseMonkey = false_monkey
		monkeys = append(monkeys, monkey)
	}
}

func printMonkeys() {
	for i := 0; i < len(monkeys); i++ {
		monkey := monkeys[i]
		fmt.Printf("Monkey %d:\n", i)
		fmt.Printf("  Starting items: %v\n", monkey.Items)
		fmt.Printf("  Operation: new = old ")
		switch monkey.IncrementOperation {
		case Increment:
			fmt.Printf("+ %d\n", monkey.IncrementNum)
		case Multiplication:
			fmt.Printf("* %d\n", monkey.IncrementNum)
		case Power:
			fmt.Printf("* old\n")
		}
		fmt.Printf("  Test: divisible by %d\n", monkey.TestDivisor)
		fmt.Printf("    If true: throw to monkey %d\n", monkey.TrueMonkey)
		fmt.Printf("    If false: throw to monkey %d\n", monkey.FalseMonkey)
		fmt.Println()
	}
}

func execute(lines []string, part2 bool) {
	times := 20
	if part2 {
		times = 10000
	}
	for k := 0; k < times; k++ {
		for i := 0; i < len(monkeys); i++ {
			var monkey *Monkey = &monkeys[i]
			for {
				if len(monkey.Items) == 0 {
					break
				}
				item := monkey.Items.PopFront()
				monkey.TimesInspected += 1
				switch monkey.IncrementOperation {
				case Increment:
					item += monkey.IncrementNum
				case Multiplication:
					item *= monkey.IncrementNum
				case Power:
					item *= item
				}
				if !part2 {
					item /= 3
				} else {
					item %= superModulo
				}
				if item%monkey.TestDivisor == 0 {
					monkeys[monkey.TrueMonkey].Items = append(monkeys[monkey.TrueMonkey].Items, item)
				} else {
					monkeys[monkey.FalseMonkey].Items = append(monkeys[monkey.FalseMonkey].Items, item)
				}
			}
		}
	}
}

func main() {
	lines := GetLines(11)
	parseMonkeys(lines)
	printMonkeys()
	execute(lines, true)
	var inspections []int
	for i := 0; i < len(monkeys); i++ {
		inspections = append(inspections, monkeys[i].TimesInspected)
	}
	fmt.Printf("Inspections: %v\n", inspections)
	sort.Slice(inspections, func(i, j int) bool {
		return inspections[i] > inspections[j]
	})
	fmt.Printf("Monkey business: %d\n", inspections[0]*inspections[1])
}
