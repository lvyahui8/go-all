package main

import "fmt"

type MyReader struct {
}

/**
练习：Reader
实现一个 Reader 类型，它不断生成 ASCII 字符 'A' 的流。
*/

func (r *MyReader) Read(b []byte) (int, error) {
	l := len(b)
	for i := 0; i < l; i++ {
		b[i] = 'A'
	}
	return l, nil
}

func main() {
	b := make([]byte, 100)
	r := MyReader{}
	_, err := r.Read(b)
	if err == nil {
		fmt.Println(string(b))
	}
}
