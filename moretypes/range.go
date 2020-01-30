package main

import "fmt"

var pow = []int{1, 2, 3, 4, 5, 6, 7, 8, 8}

/**
range
for 循环的 range 格式可以对 slice 或者 map 进行迭代循环。
*/
func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
