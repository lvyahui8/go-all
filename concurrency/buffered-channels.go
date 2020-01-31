package main

import "fmt"

/**
缓冲 channel
channel 可以是 _带缓冲的_。为 make 提供第二个参数作为缓冲长度来初始化一个缓冲 channel：

ch := make(chan int, 100)
向缓冲 channel 发送数据的时候，只有在缓冲区满的时候才会阻塞。当缓冲区清空的时候接受阻塞。

修改例子使得缓冲区被填满，然后看看会发生什么。
*/
func main() {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	// channel可以带缓冲， 超过缓冲长度依然会阻塞
	// fatal error: all goroutines are asleep - deadlock!
	// c <- 3
	fmt.Println(<-c)
	fmt.Println(<-c)
}
