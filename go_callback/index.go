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
	// 定义一个匿名函数并赋值给f
	f := func(x int) {
		println(x)
	}
	call(f) // 传递f函数作为参数
}
