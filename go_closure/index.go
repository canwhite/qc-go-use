package goclosure

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
