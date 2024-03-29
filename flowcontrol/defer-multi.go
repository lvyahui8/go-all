package main

import "fmt"

/**
defer 栈
延迟的函数调用被压入一个栈中。当函数返回时， 会按照后进先出的顺序调用被延迟的函数调用。

阅读博文了解更多关于 defer 语句的信息。
*/
func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Printf("%d ", i)
	}

	fmt.Println("done")
}
