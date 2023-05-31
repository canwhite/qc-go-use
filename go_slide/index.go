package goslide

import (
	"fmt"
)

func SectionTest() {
	//创建空的切片
	arr := []int{1, 2, 3, 4, 5}
	// arr = append(arr, 5)
	// //删除0，后边的补上
	// arr = append(arr[:0], arr[1:]...)
	// fmt.Println(arr)
	//2之前+3到结尾
	arr = append(arr[:2], arr[3:]...)

	//PS : 字符串的长度获取和数组的长度获取是一样的
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	//最常用的可能就是push和pop了
	//push功能

	stack := []int{}
	stack = append(stack, 1)
	stack = append(stack, 2)
	stack = append(stack, 3)
	fmt.Println(stack) //

	//pop功能
	//取值
	top := stack[len(stack)-1]
	//压缩
	stack = stack[:len(stack)-1]
	fmt.Println(top, stack)

	//常规for循环就不说了，这里主要展示下

	for _, num := range stack {
		fmt.Println(num)
	}

}
