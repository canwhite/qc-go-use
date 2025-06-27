package gostruct

import "fmt"


// =========1）指针接收者==========
// 类型在后边
// 在Go语言中，type关键字用于定义新的类型，而struct是定义结构体的关键字
// 使用type struct的组合有以下原因：
// 1. 明确类型定义：type Student struct 表示定义了一个名为Student的新类型
// 2. 提高代码可读性：通过type定义，可以清晰地知道这是一个新的类型
// 3. 支持类型方法：只有通过type定义的类型才能定义方法
// 4. 支持类型别名：type可以用于创建类型别名，如 type MyInt = int
// 5. 支持类型转换：通过type定义的类型可以进行类型转换
// 6. 支持接口实现：只有通过type定义的类型才能实现接口
// 7. 支持类型检查：编译器可以通过type定义进行更严格的类型检查
// 因此，type struct的组合是Go语言中定义结构体类型的标准方式

type Student struct {
	name string
	// age  int
}

//* 有两个作用
//1. 地址声明
//2. 解引用

// 在Go语言中，当使用指针接收者定义方法时，编译器会自动处理指针的解引用操作
// 这意味着在方法内部可以直接访问结构体字段，而不需要显式地使用*操作符
// 例如，在hello方法中，stu.name 实际上等同于 (*stu).name
// 这种语法糖使得代码更简洁易读，同时保持了指针传递的效率优势

// 指针接收者的主要优点：
// 1. 避免值拷贝，提高性能（特别是对于大型结构体）
// 2. 允许方法修改接收者的状态
// 3. 支持在nil指针上调用方法（需要额外处理）

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



