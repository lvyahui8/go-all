package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IPAddr [4]byte

/**
练习：Stringers
让 IPAddr 类型实现 fmt.Stringer 以便用点分格式输出地址。

例如，`IPAddr{1,`2,`3,`4}` 应当输出 `"1.2.3.4"`。
*/
func (ip IPAddr) String() string {
	items := make([]string, 4)
	for i, item := range ip {
		items[i] = strconv.Itoa(int(item))
	}
	return strings.Join(items, ".")
}

func main() {
	addrs := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDns": {8, 8, 8, 8},
	}
	for n, a := range addrs {
		fmt.Printf("%s\t:%v\n", n, a)
	}
}
