package gogoroutine

import (
	"fmt"
	"sync"
	"time"
)

// 1.使用sync
var wg sync.WaitGroup

// 2.使用信道
var ch = make(chan string, 10)

// --test1
// 单个执行之后写Done
func download1(url string) {
	fmt.Println("start to download", url)
	time.Sleep(time.Second) //模拟耗时操作
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

// --test2
func download2(url string) {
	fmt.Println("start to download", url)
	time.Sleep(time.Second)
	//只需要在那边存值就可以了
	ch <- url // 将 url 发送给信道
}

func GoRoutineTest2() {
	//启动携程
	for i := 0; i < 3; i++ {
		go download2("a.com/" + string(rune(i+'0')))
	}

	//输出结果
	for i := 0; i < 3; i++ {
		msg := <-ch
		fmt.Println("finish", msg)
		fmt.Println("Done!")
	}

}
