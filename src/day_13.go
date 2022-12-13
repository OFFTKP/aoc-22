package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
)

type Obj struct {
	objects []Obj
	value   int
	is_key  bool
}

type State int

const (
	True State = iota
	False
	Continue
)

func parse_impl(line string, object *Obj) int {
	var intString string
	object.objects = make([]Obj, 0)
	object.value = -1
	for i := 0; i < len(line); i++ {
		c := line[i]
		switch c {
		case '[':
			object.objects = append(object.objects, Obj{})
			i += parse_impl(line[i+1:], &object.objects[len(object.objects)-1])
		case ']':
			if len(intString) > 0 {
				integer, err := strconv.Atoi(intString)
				if err != nil {
					log.Fatalln(err)
				}
				object.objects = append(object.objects, Obj{value: integer})
				intString = ""
			}
			return i + 1
		case ',':
			if len(intString) > 0 {
				integer, err := strconv.Atoi(intString)
				if err != nil {
					log.Fatalln(err)
				}
				object.objects = append(object.objects, Obj{value: integer})
				intString = ""
			}
		default:
			intString += string(c)
		}
	}
	return 0
}

func print_impl(object Obj) {
	if object.value != -1 {
		fmt.Printf("%d", object.value)
		return
	}
	fmt.Print("[")
	for i, obj := range object.objects {
		print_impl(obj)
		if i < len(object.objects)-1 {
			fmt.Printf(",")
		}
	}
	fmt.Printf("]")
}

func compare_impl(object_left Obj, object_right Obj) State {
	if object_left.value != -1 && object_right.value != -1 {
		// Integer comparison
		if object_left.value < object_right.value {
			// Right order
			return True
		} else if object_left.value > object_right.value {
			// Wrong order
			return False
		} else {
			// Continue to next input
			return Continue
		}
	} else {
		if object_left.value != -1 {
			object_left.objects = append(object_left.objects, Obj{value: object_left.value})
			object_left.value = -1
		}
		if object_right.value != -1 {
			object_right.objects = append(object_right.objects, Obj{value: object_right.value})
			object_right.value = -1
		}
		// List comparison
		i := 0
		for ; i < len(object_left.objects); i++ {
			if i >= len(object_right.objects) {
				return False
			}
			state := compare_impl(object_left.objects[i], object_right.objects[i])
			if state != Continue {
				return state
			}
		}
		if i < len(object_right.objects) {
			// Left run out of items, right order
			return True
		} else if len(object_left.objects) > len(object_right.objects) {
			// Left has more items, wrong order
			return False
		} else {
			// Continue to next input
			return Continue
		}
	}
}

func main() {
	lines := GetLines(13)
	// Part 1
	objects := make([]Obj, 0)
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		var object Obj
		parse_impl(line[1:], &object)
		objects = append(objects, object)
	}
	sum := 0
	index := 1
	for i := 0; i < len(objects); i += 2 {
		// print_impl(object)
		if compare_impl(objects[i], objects[i+1]) == True {
			sum += index
		}
		index++
	}
	fmt.Println(sum)
	// Part 2
	object1 := Obj{
		objects: []Obj{
			{value: 2},
		},
		value:  -1,
		is_key: true,
	}
	object2 := Obj{
		objects: []Obj{
			{value: 6},
		},
		value:  -1,
		is_key: true,
	}
	objects = append(objects, object1)
	objects = append(objects, object2)
	sort.Slice(objects, func(i, j int) bool {
		return compare_impl(objects[i], objects[j]) == True
	})
	key := 1
	for i := 0; i < len(objects); i++ {
		if objects[i].is_key {
			key *= (i + 1)
			fmt.Println("Multiplier: ", i+1)
		}
	}
	fmt.Println(key)
}
