package gosection

import "fmt"

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
		fmt.Print(arr[i])
	}

}
