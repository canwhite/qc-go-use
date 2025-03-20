package gocallback

/*
*

 1. 定义自定义类型

 2. 为现有类型定义别名(type alias)

    //这里是自定义类型
    type User struct {
        ID int
        Name string
    }

    type Sayer interface {
        Say()
    }

    type Handler func(int)

    // 这里是别名的意思
    type MyInt = int
*/
type Handler func(int) // 定义Handler类型为func(int)

func call(h Handler) { // 接收Handler类型的参数h
	h(10) // 调用h函数
}

func CallBackTest() {
	// 这是一个匿名函数
	// 匿名函数的作用域理解：
	// 1. 匿名函数可以访问其定义所在作用域的变量
	// 2. 在Go中，匿名函数形成一个闭包，可以捕获并保留外部变量的引用
	// 3. 即使外部函数执行完毕，被捕获的变量仍然存在
	// 4. 匿名函数内部定义的变量只在函数内部可见
	{
		outerVar := 5
        //这玩意和js贼像
		func() {
			innerVar := 10
			println("访问外部变量:", outerVar) // 可以访问外部变量
			println("访问内部变量:", innerVar) // 可以访问内部变量
		}()
		// println(innerVar) // 这里会报错，无法访问内部变量
	}
	f := func(x int) {
		println(x)
	}
	call(f) // 传递f函数作为参数
}
