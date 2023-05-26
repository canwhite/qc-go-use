package gomap

import "fmt"

func RunTest() {
	//定义 ：map[KeyType]ValueType
	//这种声明方式主要适合section、map、channel
	ages := make(map[string]int)

	ages["zack"] = 31
	ages["zhangsan"] = 18
	ages["lisi"] = 23
	fmt.Println(ages)

	//删除
	delete(ages, "lisi")
	fmt.Println(ages)

	//长度
	fmt.Println(len(ages))

	//for循环
	for key, value := range ages {
		fmt.Println(key)
		fmt.Println(value)
	}

}
