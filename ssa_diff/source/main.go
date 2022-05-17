package main

import "fmt"

type StInsertTest struct {
	Name string
	Age  int
}

func StInsertFunc() {
	st := StInsertTest{Name: "fg", Age: 12}
	fmt.Println(st)
}

func main() {
	StInsertFunc()
}
