package govar

import "fmt"

func VarTest() {
	/*
		var a int     // 如果没有赋值，默认为0
		var a int = 1 // 声明时赋值
		var a = 1     // 声明时赋值
	*/

	//字面量形式
	msg := "hello world"
	fmt.Println(msg)

	// -- var和:=的区别
	// 1. := 不能在函数外使用,只能在函数内部声明变量。`var` 可以在函数内外都可以声明变量。
	// 2. 使用 := 声明的变量根据右值自动推导类型,而 var 声明的变量如果不指定类型,会默认为 interface{} 类型。
	// 3. := 声明的变量在同一个作用域内不允许重复声明,否则会导致编译错误。`var` 声明的变量在同一作用域内可以重复声明。

	//简单类型

	/**
	空值：nil
	整型类型： int(取决于操作系统), int8, int16, int32, int64, uint8, uint16, …
	浮点数类型：float32, float64
	字节类型：byte (等价于uint8)
	字符串类型：string
	布尔值类型：boolean，(true 或 false)
	*/

}
