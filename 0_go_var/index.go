package govar

import "fmt" //输出工具

func VarTest() {

	/*
		var a int     // 如果没有赋值，默认为0
		var a int = 1 // 声明时赋值
		var a = 1     // 声明时赋值
	*/

	//字面量形式,字符串是只读的字节切片，所以不能更改
	msg := "hello world"
	fmt.Println(msg)


	a := 10
	a = 20 
	fmt.Println(a)

	/**
	字符串类型：string
	字节类型：byte (等价于uint8)
	布尔值类型：boolean，(true 或 false)
	整型类型： int(取决于操作系统), int8, int16, int32, int64, uint8, uint16, …
	浮点数类型：float32, float64
	空值：nil
	*/

}
