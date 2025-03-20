package goclosure


// add 函数是一个闭包函数的示例
// 返回值定义 func(int) int 表示返回一个函数，这个函数：
// - 接收一个int类型参数
// - 返回一个int类型结果
// 这种定义方式允许我们返回一个闭包函数



func add(x int) func(int) int {
	
	sum := x
	return func(y int) int { // 返回一个闭包
		sum += y // 读取外层sum变量
		return sum
	}
}

func ClosureTest() {
	f := add(10)    // f是一个闭包
	result := f(20) // 通过f调用闭包,sum变量值增加为30
	println(result) // 30
}
