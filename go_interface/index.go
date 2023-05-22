package gointerface

import "fmt"

//定义方法，继承和实现方法

// 定义接口
type Point interface {
	X() int
	Y() int
}

// 实现
type MyPoint struct {
	x int
	y int
}

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
