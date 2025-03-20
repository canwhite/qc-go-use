package gomap

import "fmt"

func RunTest() {
	//定义 ：map[KeyType]ValueType
	//这种声明方式主要适合Array和slide、map、channel
	ages := make(map[string]int)
	//但是上述这些又都可以用字面量创建

	ages["zack"] = 31
	ages["zhangsan"] = 18
	ages["lisi"] = 23
	fmt.Println(ages)

	//删除的时候比较有趣，类似于c++撒
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
