package gosingleflight

// Pre :注意，基础测试文件后边加_test会被识别
/**
go的基准测试
1. 命名以Benchmark开头的函数,该函数会被go test自动执行。
2. 函数接收testing.B接口类型的参数。
3. 在函数中重复执行需要测试的代码b.N次。
4. 使用go test -bench=Benchmark名称 来运行基准测试。
eg1:
BenchmarkFib-8               500000          2344 ns/op         112 B/op           2 allocs/op
这说明BenchmarkFib-8这个测试执行了500000遍,每次执行Fib(30)平均耗时2344纳秒,分配了112字节的内存和2个对象。

eg2:
go test -run='^$' -bench=. -count=1 -benchtime=1000x -benchmem > slow.txt
- go test : 运行测试的基本命令
- -run='^$' : 匹配不执行任何测试用例,只运行基准测试，因为是一个空的正则表达式，不会匹配
- -bench=. : 运行所有基准测试
- -count=1 : 每个基准测试至少执行1次
- -benchtime=1000x : 每个基准测试运行至少1000纳秒(1微秒)
- -benchmem : 打印内存分配统计信息
- > slow.txt : 把输出重定向到slow.txt文件


--基准测试用于验证优化结果，测试用例验证业务是否成功


*/

import (
	"sync"
	"testing"
	"time"
)

//实际上没有必要一开始就引入模块，用的时候再引也是一样的道理

type user struct{}

// 写一个异步操作，这部分是独立的
func getUserById(id int) user {
	//模拟数据库查询耗时
	time.Sleep(time.Millisecond)
	/**
	- user{}:返回一个临时实例,包含类型的零值字段。
	- &user{}:返回类型的地址,代表一块内存,要通过地址来访问实例和修改实例字段。
	如果我们获取后想要长久保存，就需要获取它的地址
	*/
	return user{}
}

//然后就是具体的并发操作

func BenchmarkBufferWithPool(b *testing.B) {
	var wg sync.WaitGroup
	//然后并发操作
	for i := 0; i < b.N; i++ {
		wg.Add(1) //这个可以理解为并发技术器，wg.Add(1)的作用是增加等待组的计数器,表示有一个并发操作加入
		//运行一个自执行函数
		go func() {
			defer wg.Done() //所有工作都进行完了之后运行此操作
			_ = getUserById(1024)
		}()
	}

	//等待所有操作完成
	wg.Wait()

}
