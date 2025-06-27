package goslide

import (
	"fmt"
)

func SectionTest() {

	// 使用make创建切片
	// make([]类型, 长度, 容量)
	// 长度：当前切片包含的元素个数, TODO: 长度设置好了会有默认值
	// 容量：底层数组的总大小
	
	slice1 := make([]int, 3) // 长度为3，容量默认为3的整型切片
	slice2 := make([]string, 2, 5) // 长度为2，容量为5的字符串切片

	fmt.Println("make创建的slice1:", slice1)
	fmt.Println("make创建的slice2:", slice2)

	// make创建和字面量创建的区别：
	// 1. make创建可以指定长度和容量，字面量创建长度和容量固定
	// 2. make创建的切片元素会初始化为零值，字面量创建可以指定初始值
	// 3. make创建更灵活，适合需要预分配内存的场景
	// 4. 字面量创建更直观，适合已知初始值的场景



	//创建空的切片
	arr := []int{1, 2, 3, 4, 5}

	//2之前+3到结尾
	// ... 是Go语言中的可变参数展开符，用于将切片中的元素展开成独立的参数
	// 在这个例子中，arr[3:]... 表示将arr[3:]切片中的所有元素作为独立的参数传递给append函数
	// 是的，Go语言中的切片操作是左闭右开的，即包含起始索引，但不包含结束索引
	// 例如 arr[1:3] 会包含索引1和2的元素，但不包含索引3的元素
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
	//1）先取值
	top := stack[len(stack)-1]
	//2）再压缩
	stack = stack[:len(stack)-1]
	fmt.Println(top, stack)

	//--for range是一种很好用的形式--
	// 在Go语言中，range返回两个值：索引和元素值
	// 使用下划线_表示忽略索引值，只取元素值
	// 这样可以避免声明一个不使用的变量

	for _, num := range stack {
		fmt.Println(num)
	}


	//元素包含也是这样的流程
	// 查询元素
	// 方式1：遍历查找
	target := 2
	found := false
	for _, num := range arr {
		if num == target {
			found = true
			break
		}
	}
	if found {
		fmt.Printf("元素 %d 存在于切片中\n", target)
	} else {
		fmt.Printf("元素 %d 不存在于切片中\n", target)
	}

	// 方式2：使用二分查找（前提是切片已排序）
	sortedArr := []int{1, 2, 3, 4, 5}
	left, right := 0, len(sortedArr)-1
	for left <= right {
		mid := (left + right) / 2
		if sortedArr[mid] == target {
			fmt.Printf("元素 %d 存在于切片中，索引为 %d\n", target, mid)
			break
		} else if sortedArr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}


}
