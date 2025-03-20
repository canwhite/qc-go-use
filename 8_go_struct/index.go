package gostruct

import "fmt"


// =========1）指针接收者==========
// 类型在后边
type Student struct {
	name string
	// age  int
}

// 前边加 stu *Student 的意义： 
// 1. 表示这是一个指针接收者方法
// 2. 使用指针接收者时，方法内部可以修改结构体的字段
// 3. 可以避免方法调用时的值拷贝，提高性能
// 4. 当结构体较大时，使用指针接收者更高效
// 5. 指针接收者方法只能在可寻址的结构体实例上调用
// 6. 与值接收者相比，指针接收者方法可以修改原始结构体

func (stu *Student) hello(person string) string {
	return fmt.Sprintf("hello %s, I am %s", person, stu.name)
}

// 测试一下go的类
func StructTest() {
	//取地址，这里取地址之后是自动加*
	//所以创建实例的时候就是一个取地址符，再初始化
	stu := &Student{
		name: "Tom",
	}
	//世界
	msg := stu.hello("jack")
	fmt.Println(msg)

}



// =========2）嵌入结构体继承==========
// Go语言没有传统意义上的继承，但可以通过嵌入结构体实现类似继承的效果

// 定义一个父结构体
type Person struct {
	name string
	age  int
}

// 定义一个子结构体，嵌入Person
type Teacher struct {
	Person // 嵌入Person结构体，相当于继承
	subject string
}

// Teacher可以直接访问Person的字段和方法
func (t *Teacher) teach() {
	fmt.Printf("Teacher %s is teaching %s\n", t.name, t.subject)
}



