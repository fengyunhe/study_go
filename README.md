# Help

## 创建和引入外部模块

此处假设是需要在本项目引入 `go.firegod.cn/grettings` 模块

1. 创建一个目录 `creetings`
2. 进入目录 `creetings`
3. 在 `grettins.go`中编写方法，方法名称第一个字母是大写
4. 在引入模块的项目项目下执行 `go mod init go.firegod.cn/grettings`
5. 因为模块是本地路径，所以需要做一个替换，执行 `go mod edit -replace go.firegod.cn/greetings=../greetings`
6. 执行`go mod tidy`
7. 引入成功
