package main

import (
	"fmt"
	"time"
)

/**
当 select 中的其他条件分支都没有准备好的时候，`default` 分支会被执行。

为了非阻塞的发送或者接收，可使用 default 分支：

select {
case i := <-c:
    // 使用 i
default:
    // 从 c 读取会阻塞
}
*/
func main() {
	tick := time.Tick(3000 * time.Microsecond)
	boom := time.After(9000 * time.Microsecond)

	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("   .")
			time.Sleep(50 * time.Microsecond)
		}
	}
}
