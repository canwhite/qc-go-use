package goasyncawait

/**

//这些可以用make创建，但是我更喜欢字面量
intArr := [5]int{1, 2, 3, 4, 5}
slide := []int{}
map := map[string]int{"a": 1, "b": 2}
chan := chan string{}

*/

import (
	"fmt"
	"time" //引入的竟然是time
) 

/*
chan类型是Go语言中的通道类型，用于在不同goroutine之间进行通信和同步。

主要特点：
1. 类型声明：chan T 表示传递T类型的通道
2. 通道方向：
   - chan T：双向通道（可发送和接收）
   - chan<- T：只写通道
   - <-chan T：只读通道
3. 通道操作：
   - ch <- v：发送值v到通道ch
   - v := <-ch：从通道ch接收值
   - close(ch)：关闭通道
4. 通道特性：
   - 默认是无缓冲的，发送和接收会阻塞直到另一端准备好
   - 可以创建带缓冲的通道：make(chan T, bufferSize)
   - 关闭后不能再发送，但可以继续接收直到通道为空
   - 从关闭的通道接收会返回零值
5. 使用场景：
   - goroutine之间的数据传递
   - 实现生产者消费者模式
   - 替代锁实现并发控制
   - 构建并发管道

注意事项：
- 避免向已关闭的通道发送数据，会导致panic
- 无缓冲通道需要发送和接收端都准备好
- 合理使用select语句处理多个通道操作
*/

/*
从某种程度上来说，通道确实可以看作是Go语言中实现发布订阅模式的一种机制，但两者之间还是有一些区别：

相似之处：
1. 消息传递：通道和发布订阅都实现了消息的传递
2. 解耦：生产者和消费者/发布者和订阅者之间是解耦的
3. 异步：都可以实现异步通信

区别：
1. 通道是点对点的，发布订阅是一对多的
2. 通道是强类型的，发布订阅通常不限制消息类型
3. 通道有明确的发送接收关系，发布订阅更松散
4. 通道可以阻塞等待，发布订阅通常是事件驱动

通道更适用于：
- 需要严格同步的场景
- 需要保证消息顺序的场景
- 需要控制并发数量的场景

发布订阅更适用于：
- 广播式消息传递
- 事件驱动架构
- 需要灵活扩展订阅者的场景
*/


func doSomething(done chan bool) {
	// pretend this is some long operation
	time.Sleep(1 * time.Second)

	// 这里左指是发送值到通道
	done <- true // signal that we are done


}

func AsyncTest() {
	done := make(chan bool)
	
	//go可以理解为emit
	go doSomething(done)

	//所以<-done 表示等待done channel返回一个值,从而实现等待goroutine结束的效果。
	res := <-done // wait for goroutine to finish
	fmt.Println("done!", res)
}
