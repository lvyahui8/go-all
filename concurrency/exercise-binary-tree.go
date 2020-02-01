package main

import (
	"fmt"
	"sort"
)

/**
https://go-tour-zh.appspot.com/concurrency/7
*/

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func ll(tree *Tree, c chan int) []int {
	n := bfs(tree, c)
	close(c)
	items := make([]int, n)
	i := 0
	for item := range c {
		items[i] = item
		i++
	}
	sort.Ints(items)
	return items
}

func bfs(tree *Tree, c chan int) int {
	if tree == nil {
		return 0
	}
	c <- tree.Value
	sum := 1
	sum += bfs(tree.Left, c)
	sum += bfs(tree.Right, c)
	return sum
}

func same(items1, items2 []int) bool {
	if len(items1) != len(items2) {
		return false
	} else {
		for i := 0; i < len(items1); i++ {
			if items1[i] != items2[i] {
				return false
			}
		}
		return true
	}
}

func main() {
	tree1 := Tree{
		Left: &Tree{
			Left: &Tree{
				Value: 1,
			},
			Value: 1,
			Right: &Tree{
				Value: 2,
			},
		},
		Value: 3,
		Right: &Tree{
			Left: &Tree{
				Value: 5,
			},
			Value: 8,
			Right: &Tree{
				Value: 13,
			},
		},
	}
	tree2 := Tree{
		Left: &Tree{
			Left: &Tree{
				Left: &Tree{
					Value: 1,
				},
				Value: 1,
				Right: &Tree{
					Value: 2,
				},
			},
			Value: 3,
			Right: &Tree{
				Value: 5,
			},
		},
		Value: 8,
		Right: &Tree{
			Value: 13,
		},
	}
	c1 := make(chan int, 10)
	c2 := make(chan int, 10)
	items1 := ll(&tree1, c1)
	items2 := ll(&tree2, c2)
	fmt.Println("same:", same(items1, items2))

}
