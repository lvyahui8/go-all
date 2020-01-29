package main

import (
	"fmt"
	"runtime"
)

/**
switch
一个结构体（`struct`）就是一个字段的集合。

除非以 fallthrough 语句结束，否则分支会自动终止。

(这个跟java和c++的没有break自动穿透相反)
*/
func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s.", os)
	}
}
