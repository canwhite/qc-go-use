package gotest

func Test() {
	/**
		//ps : package main

		//如果一个Go文件希望被编译成一个完整的、独立运行的可执行程序,
		//那么需要使用package main包裹,并且提供一个main函数作为程序入口



		// calc.go
		package main

		func add(num1 int, num2 int) int {
			return num1 + num2
		}


		//然后具体的测试方法
		// calc_test.go
		package main

		import "testing"

		func TestAdd(t *testing.T) {
			if ans := add(1, 2); ans != 3 {
				t.Error("add(1, 2) should be equal to 3")
			}
		}

		//一个文件夹可以作为 package，同一个 package 内部变量、类型、方法等定义可以相互看到。
		//这里我们直接调用了add，因为他们平级，都是main包里的




		//运行测试用例的时候
		$ go test -v   //这里-v查看详情信息
		=== RUN   TestAdd
		--- PASS: TestAdd (0.00s)
		PASS
		ok      example 0.040s








	**/

}
