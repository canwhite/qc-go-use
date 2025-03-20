package go_string

import (
	"fmt"
	"strings" //可以看出strings是一个静态类
	"unicode/utf8"
)

//DES: 字符串的特殊性，是只读的字节切片，所以不能更改，
//strings静态类中有很多可用方法
func StringExamples() {
	// 1. 字符串声明与初始化
	str1 := "Hello, 世界"
	fmt.Println("字符串示例1:", str1)

	// 2. 获取字符串长度
	// 注意：len() 返回的是字节数，对于中文字符需要使用 utf8.RuneCountInString
	fmt.Println("字符串字节长度:", len(str1))
	fmt.Println("字符串字符长度:", utf8.RuneCountInString(str1))

	// 3. 字符串拼接
	// 字符串是不可变的，一旦定义就不能直接修改
	str2 := "Go语言"
	str3 := str1 + " " + str2
	fmt.Println("字符串拼接:", str3)

	// 4. 字符串比较
	if str1 == "Hello, 世界" {
		fmt.Println("字符串相等")
	} else {
		fmt.Println("字符串不相等")
	}

	// 5. 字符串包含判断
	if strings.Contains(str1, "世界") {
		fmt.Println("字符串包含'世界'")
	}

	// 6. 字符串分割
	words := strings.Split("Go is a great language", " ")
	fmt.Println("字符串分割:", words)

	// 7. 字符串替换
	// strings.Replace函数中的最后一个参数1表示替换的次数
	// -1表示替换所有匹配项
	// 0表示不替换
	// n(n>0)表示最多替换n次
	newStr := strings.Replace(str1, "Hello", "你好", 1)
	fmt.Println("字符串替换:", newStr)

	// 8. 字符串大小写转换
	upper := strings.ToUpper("hello")
	lower := strings.ToLower("WORLD")
	fmt.Println("大小写转换:", upper, lower)

	// 9. 字符串修剪
	trimmed := strings.Trim("   Golang   ", " ")
	fmt.Println("字符串修剪:", trimmed)

	// 10. 字符串遍历
	for i, r := range str1 {
		fmt.Printf("字符索引:%d, Unicode值:%U, 字符:%c\n", i, r, r)
	}
}



