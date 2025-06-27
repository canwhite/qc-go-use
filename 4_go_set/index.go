// package set 表示当前文件属于set包
// 包是Go语言中组织代码的基本单位，具有以下特点：
// 1. 每个Go文件必须属于一个包
// 2. 包名通常与目录名相同，但并非强制要求
// 3. 同一个目录下的所有Go文件必须属于同一个包
// 4. 包名应当简短且有意义，通常使用小写字母
// 5. 包的作用：
//    - 组织代码结构
//    - 控制访问权限（首字母大写的标识符可被外部访问）
//    - 提供命名空间，避免命名冲突
// 6. 在本例中，package set表示这个文件中的代码属于set包，
//    可以被其他包通过 import "tgo/4_go_set" 的方式引用

package set

// Go语言中的import语句可以采用两种形式：
// 1. 合并式import（推荐）：
//    import (
//        "fmt"
//        "math"
//    )
//    优点：代码更整洁，便于管理多个导入包

// 2. 单独式import：
//    import "fmt"
//    import "math"
//    缺点：当导入包较多时，代码会显得冗长

// 虽然两种方式都可以使用，但合并式import是Go社区的推荐做法，
// 因为它：
// - 使代码更简洁
// - 便于管理多个导入包
// - 符合Go语言的代码风格规范
// - 更容易发现未使用的导入（go工具会提示）

// 因此，建议使用合并式import的写法。

// Go语言本身没有内置的Set数据结构，但可以通过以下方式实现类似功能：

// 1. 使用map实现Set
// map的key是唯一的，可以用来模拟Set
// 示例：
// set := make(map[string]bool)
// set["a"] = true  // 添加元素
// delete(set, "a") // 删除元素
// if set["a"] {    // 检查元素是否存在
//     fmt.Println("a存在")
// }

// 2. 使用sync.Map实现并发安全的Set
// 适用于并发场景
// 示例：
// var set sync.Map
// set.Store("a", true)  // 添加元素
// set.Delete("a")       // 删除元素
// if _, ok := set.Load("a"); ok {  // 检查元素是否存在
//     fmt.Println("a存在")
// }

// 3. 使用第三方库
// 如github.com/scylladb/go-set
// 提供了更丰富的Set操作，如并集、交集、差集等
// 示例见RunTest函数

// 选择建议：
// - 简单场景：使用map
// - 并发场景：使用sync.Map
// - 需要丰富操作：使用第三方库

import (
	"fmt"
	"github.com/scylladb/go-set"
)

func RunTest() {
	
	s := set.NewStringSet("a", "b")
	s.Add("c")
	fmt.Println(s)

	// 创建新的Set
	s2 := set.NewStringSet()
	
	// 添加元素
	s2.Add("x")
	s2.Add("y")
	fmt.Println("添加后:", s2)

	// 删除元素
	s2.Remove("x")
	fmt.Println("删除x后:", s2)

	// 检查元素是否存在
	if s2.Has("y") {
		fmt.Println("y存在")
	}

	// 获取Set大小
	fmt.Println("Set大小:", s2.Size())

	// 清空Set
	s2.Clear()
	fmt.Println("清空后:", s2)

	// 合并两个Set
	s3 := set.NewStringSet("m", "n")
	s4 := set.NewStringSet("n", "p")
	s3.Merge(s4)
	fmt.Println("合并后:", s3)
	// 遍历Set
	fmt.Println("遍历Set s3:")
	for item := range s3.List() {
		fmt.Println(item)
	}

	// 使用List遍历
	fmt.Println("使用List遍历 s3:")
	for item := range s3.List() {
		fmt.Println(item)
	}
}
