# qc-go-use

go and modules

## package

```

-初始化
go mod init qc-go-use
包的寻址是从go.mod文件开始的，类比js项目的package.json ,然后内部拼接就可以

-包管理
Go 语言使用模块（module）作为包管理的基本单位。每个模块都有一个 go.mod 文件，用于定义模块路径和依赖关系。主要特点如下：

1. 模块路径：每个模块都有一个唯一的导入路径，通常与代码仓库地址对应
2. 版本管理：使用语义化版本号（v1.2.3）来管理依赖版本
3. 依赖缓存：下载的依赖会存储在 $GOPATH/pkg/mod 目录下
4. 最小版本选择：自动选择满足所有依赖要求的最小版本

示例：
假设我们有一个项目需要导入 github.com/gin-gonic/gin 包：

1. 初始化模块：
   go mod init myproject

2. 在代码中导入：
   import "github.com/gin-gonic/gin"

3. 自动下载依赖：
   go mod tidy

4. go.mod 文件内容示例：
   module myproject

   go 1.20

   require (
       github.com/gin-gonic/gin v1.9.0
   )



如果我们要自己实现一个包，可以按照以下步骤：

1. 创建包目录结构：
   mypackage/
   ├── go.mod
   ├── mypackage.go

2. 初始化包模块：
   cd mypackage
   go mod init github.com/username/mypackage

3. 编写包代码（mypackage.go）：
   package mypackage

   func Hello() string {
       return "Hello from mypackage!"
   }

4. 发布包到GitHub（可选）：
   git init
   git add .
   git commit -m "Initial commit"
   git remote add origin https://github.com/username/mypackage.git
   git push -u origin main

5. 在其他项目中使用：
   import "github.com/username/mypackage"

   func main() {
       fmt.Println(mypackage.Hello())
   }

6. 更新版本：
   git tag v1.0.0
   git push origin v1.0.0

这样，你就可以创建、发布和使用自己的Go包了。





-加载包
go get [-u] 包地址
-u的意思是是否将包更新到最新
PS：或者先import写上，报错后点击地址自己加载


-run
- 运行单个文件
go run main.go  # 直接运行，默认会先编译再运行

- 运行整个项目
go run .  # 直接运行，默认会先编译再运行

- 编译项目
go build  # 生成可执行文件

- 编译并运行
go build && ./qc-go-use  # 先编译再运行

注意：使用go run命令时，默认会先编译代码到临时目录再运行，但不会生成可执行文件



-瘦身
go mod tidy
```

## use

[gin](https://github.com/canwhite/qc-gin-use)  
[algorithm](https://github.com/canwhite/qc-go-algorithm)
