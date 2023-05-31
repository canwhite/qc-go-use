package goasyncawait

/**
import (
	"ioutil"
	"net/http"
)

func myAsynchronousFunction(res chan *Response) {
	response, _ := http.Get("http://example.com/")
	defer response.Body.Close()
	text, _ = ioutil.ReadAll(response.Body)
	res <- text
}

//你可以理解它为有点异步单线程的意思
func myAwaitingFunction() {
	//传入一个channel对象，用于拿值
	responseChannel := make(chan *Response)
	//同步执行完
	go myAsynchronousFunction(responseChannel)
	//获取值
	resp := <-responseChannel



}
**/
