package gointerface

import "fmt"

//定义方法，继承和实现方法

// 定义接口，后边使用值接收者实现
type Point interface {
	X() int
	Y() int
}

// 实现
type MyPoint struct {
	x int
	y int
}

// 这里使用值接收者而不是指针接收者的原因：
// 1. MyPoint 结构体较小，使用值传递不会带来明显的性能开销
// 2. 方法实现中不需要修改结构体的字段值
// 3. 使用值接收者可以保证方法的线程安全性，因为每次调用都会操作原始值的副本
// 4. 这种访问方式叫做 "值接收者方法" (Value Receiver Method)

// 使用值接收者的场景：
// - 当结构体较小，且不需要修改结构体字段时
// - 当需要保证方法操作的不可变性时
// - 当方法逻辑不依赖于结构体的指针特性时

func (p MyPoint) X() int {
	return p.x
}

func (p MyPoint) Y() int {
	return p.y
}

func InterfaceTest() {
	//类型是Point，向上兼容了

	var p Point = MyPoint{x: 10, y: 100}
	fmt.Println(p.X())

	/**
	    var p1 Point = MyPoint{x: 10, y: 20}
	    var p2 Point = &MyPoint{x: 10, y: 20}

	    // p1的类型是MyPoint
	    fmt.Printf("%T\n", p1)

	    // p2的类型是*MyPoint
	    fmt.Printf("%T\n", p2)

		//总结：可见第一种是直接的结构体，第二种是指针类型的
	*/

}
