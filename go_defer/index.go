package godefer

import "fmt"

// 你可以理解为defer主要做一些保底操作，这和await还是区别很大的
// 使用在一些函数前
func DeferTest() {

	defer fmt.Println(9)
	fmt.Println(1)
	fmt.Println(2)

}
