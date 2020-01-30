package main

import "fmt"

func fibonacci() func() int {
	x := 1
	y := -1
	return func() int {
		z := x + y
		y = x
		x = z
		return z
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
