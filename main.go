package main

import (
	//gos "tgo/go_slide"
	// gom "tgo/go_map"
	// god "tgo/go_defer"
	govar "tgo/0_go_var"
	goset "tgo/4_go_set"
)

func main() {
	govar.VarTest()       // Uncomment to use the import
	// gos.SectionTest() //数组和切片
	//gostruct.StructTest() //class
	//gointerface.InterfaceTest()
	// gor.Test1()
	// goroutine.GoRoutineTest1()
	// goroutine.GoRoutineTest2()
	// gom.RunTest()
	// god.DeferTest()
	goset.RunTest()
}
