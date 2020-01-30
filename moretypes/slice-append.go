package main

import "fmt"

func printSlice2(s string, x []int) {
	fmt.Printf("%s len = %d cap = %d %v\n", s, len(x), cap(x), x)
}

/**
向 slice 添加元素
向 slice 添加元素是一种常见的操作，因此 Go 提供了一个内建函数 `append`。 内建函数的文档对 append 有详细介绍。

func append(s []T, vs ...T) []T
append 的第一个参数 s 是一个类型为 T 的数组，其余类型为 T 的值将会添加到 slice。

append 的结果是一个包含原 slice 所有元素加上新添加的元素的 slice。

如果 s 的底层数组太小，而不能容纳所有值时，会分配一个更大的数组。 返回的 slice 会指向这个新分配的数组。

（了解更多关于 slice 的内容，参阅文章slice：使用和内幕。）
*/
func main() {
	var a []int
	printSlice2("a", a)

	a = append(a, 0)
	printSlice2("a", a)

	a = append(a, 1)
	printSlice2("a", a)

	a = append(a, 2, 3, 4)
	printSlice2("a", a)
}
