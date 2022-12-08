package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Node struct {
	Name        string
	Size        int
	Children    []*Node
	Previous    *Node
	IsDirectory bool
}

type Directory struct {
	Name string
	Size int
}

var directories []Directory

func main() {
	lines := GetLines(7)
	var start Node
	start.Name = "/"
	start.Size = 0
	var current *Node = &start
	for _, line := range lines {
		words := strings.Split(line, " ")
		if len(words) != 3 {
			var temp *Node = new(Node)
			if words[0] == "dir" {
				temp.Name = words[1]
				temp.Size = 0
				temp.Previous = current
				temp.IsDirectory = true
			} else {
				temp.Size, _ = strconv.Atoi(words[0])
				temp.Name = words[1]
				temp.Previous = current
				temp.IsDirectory = false
			}
			temp.Children = make([]*Node, 0)
			current.Children = append(current.Children, temp)
			continue
		} else {
			if words[1] == "cd" {
				arg := words[2]
				if arg == "/" {
					current = &start
					continue
				} else if arg == ".." {
					if current.Previous != nil {
						current = current.Previous
					} else {
						log.Fatalln("Already at root")
					}
				} else {
					found := false
					for _, child := range current.Children {
						if child.Name == arg {
							current = child
							found = true
							break
						}
					}
					if !found {
						log.Fatalf("No such directory: %s\nSize of children: %d\n", arg, len(current.Children))
					}
				}
			}
		}
	}
	rootsum := 0
	for _, child := range start.Children {
		rootsum += recurse(child)
	}
	ceil := 30000000 - (70000000 - rootsum)
	min := 30000000
	sum := 0
	for _, dir := range directories {
		if dir.Size <= 100000 {
			sum += dir.Size
		}
		if dir.Size >= ceil && dir.Size < min {
			min = dir.Size
		}
	}
	fmt.Printf("%d\n%d\n", sum, min)
}

func recurse(node *Node) int {
	var sum int = node.Size
	if node.IsDirectory {
		for _, child := range node.Children {
			sum += recurse(child)
		}
		node.Size = sum
		directories = append(directories, Directory{node.Name, sum})
	} else {
		return node.Size
	}
	return sum
}
