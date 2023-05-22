package gostruct

import "fmt"

// 类型在后边
type Student struct {
	name string
	// age  int
}

func (stu *Student) hello(person string) string {
	return fmt.Sprintf("hello %s, I am %s", person, stu.name)
}

// 测试一下go的类
func StructTest() {
	//取地址，这里取地址之后是自动加*
	stu := &Student{
		name: "Tom",
	}
	//世界
	msg := stu.hello("jack")
	fmt.Println(msg)

}
