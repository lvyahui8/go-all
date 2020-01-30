package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	fields := strings.Fields(s)
	m := make(map[string]int)
	for _, item := range fields {
		_, exist := m[item]
		if !exist {
			m[item] = 1
		} else {
			m[item]++
		}
	}
	return m
}

func main() {
	fmt.Println(WordCount("How are you ? What your name ? How about you ?"))
}
