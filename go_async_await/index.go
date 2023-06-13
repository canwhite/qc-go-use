package goasyncawait

import (
	"fmt"
	"time"
)

func doSomething(done chan bool) {
	// pretend this is some long operation
	time.Sleep(1 * time.Second)
	done <- true // signal that we are done
}

func AsyncTest() {
	done := make(chan bool)
	go doSomething(done)

	//所以<-done 表示等待done channel返回一个值,从而实现等待goroutine结束的效果。
	res := <-done // wait for goroutine to finish
	fmt.Println("done!", res)
}
