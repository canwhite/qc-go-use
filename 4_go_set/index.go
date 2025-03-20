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
}
