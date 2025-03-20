package gogoroutine

import (
	"fmt"
	"sync"
	"time"
)

// 1.使用sync
var wg sync.WaitGroup

// 2.使用信道
// 这里的10表示信道ch的缓冲区大小，即可以存储10个字符串类型的数据。
// 有缓冲信道的特点：
// 1. 当缓冲区未满时，发送操作不会阻塞
// 2. 当缓冲区不为空时，接收操作不会阻塞
// 3. 适合用于生产者和消费者模式，可以缓解生产者和消费者的速度差异

var ch = make(chan string, 10)

// ===============test1：单个执行内部写Done=========
func download1(url string) {
	fmt.Println("start to download", url)
	time.Sleep(time.Second) //模拟耗时操作
	//单个内部是done
	wg.Done()
}


// 最终等待所有的结果
func GoRoutineTest1() {
	for i := 0; i < 3; i++ {
		wg.Add(1)
		//这里的go相当于启动一个协程
		go download1("a.com/" + string(rune(i+'0')))
	}
	//等待所有结果
	wg.Wait()
	fmt.Println("Done~")
}






// ======== 这个就是简单的信道使用了,go emit之后依次等待返回=========
func download2(url string) {
	fmt.Println("start to download", url)
	time.Sleep(time.Second)
	//只需要在那边存值就可以了
	ch <- url // 将 url 发送给信道
}

func GoRoutineTest2() {
	//启动携程
	for i := 0; i < 3; i++ {
		//rune是int32的别名
		go download2("a.com/" + string(rune(i+'0')))
	}

	//输出结果
	for i := 0; i < 3; i++ {
		msg := <-ch
		fmt.Println("finish", msg)
		fmt.Println("Done!")
	}

}
