package main

import "fmt"

var i, j int = 1, 2

func main() {
	// 省略类型， 变量从初始值中获得类型
	var c, python, java = true, false, "no!"
	fmt.Print(i, j, c, python, java)
}
