# qc-go-use
go and modules


## package

```
-相当于初始化package.json,go的方法是
go mod init qc-go-use
包的寻址是从go.mod文件开始的，然后内部拼接就可以
比如我们是tgo/xxx
-------tgo外部使用-------
- 如果其他项目要使用我们的tgo模块，可以按照以下步骤操作：

1. 首先确保我们的模块已经推送到远程仓库（如GitHub）
2. 在其他项目的go.mod文件中添加依赖：
   ```go
   require github.com/yourusername/qc-go-use v1.0.0
   ```
3. 然后执行命令下载依赖：
   ```bash
   go mod tidy    //tidy整洁的
   ``` 
4. 在代码中导入并使用我们的模块：
   ```go
   import "github.com/yourusername/qc-go-use/tgo/xxx"
   ```

PS: 3-4 可以压缩成一步，go get xxx，会自动更新依赖




-加载包

go get [-u] 包地址
-u的意思是是否将包更新到最新
PS：或者先import写上，报错后点击地址自己加载   



-瘦身
go mod tidy
```

## use
[gin](https://github.com/canwhite/qc-gin-use)    
[algorithm](https://github.com/canwhite/qc-go-algorithm)

