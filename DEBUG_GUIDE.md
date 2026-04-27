# Go 程序调试指南 (MacBook Pro M3)

## ✅ 环境配置完成

- **Go 版本**: 1.24.7 darwin/arm64
- **调试器**: Delve 1.25.1
- **IDE**: VS Code

## 📝 VS Code 调试配置

已在 `.vscode/launch.json` 中配置了 3 个调试配置：

1. **Debug Current File** - 调试当前文件
2. **Debug Tests** - 调试单元测试
3. **Attach to Process** - 附加到运行中的进程

## 🚀 如何使用调试

### 方法 1: VS Code 图形化调试 (推荐)

1. 在代码中点击要设置断点的行号左边 → 会出现红点
2. 按 `F5` 或点击 Run → Run and Debug
3. 选择 "Debug Current File" 配置
4. 代码会在断点处暂停，你可以：
   - ⏸️ 查看变量值
   - ⏭️ 单步执行 (F10)
   - ⬇️ 进入函数 (F11)
   - ⬆️ 跳出函数 (Shift+F11)
   - ▶️ 继续执行 (F5)

### 方法 2: 命令行调试 (Delve CLI)

```bash
# 进入你的程序目录
cd /Users/zhubo/go_project/demo/hello_demo1

# 启动调试会话
dlv debug

# 常用命令:
# (dlv) break main.main     - 设置断点
# (dlv) continue            - 继续执行
# (dlv) next                - 单步执行
# (dlv) step                - 进入函数
# (dlv) print a             - 打印变量值
# (dlv) locals              - 显示本地变量
# (dlv) quit                - 退出调试
```

### 方法 3: 调试单元测试

1. 在 VS Code 中打开 `*_test.go` 文件
2. 按 `F5` 并选择 "Debug Tests" 配置

## 🔧 常见问题排查

### 问题 1: 无法设置断点？
- 确保文件已保存
- 确保调试配置中的 `mode` 是 `"debug"`

### 问题 2: 调试器卡住或无响应？
- 尝试停止调试器 (Shift+F5)
- 重新启动 VS Code

### 问题 3: 变量无法展开或显示 `<error>`？
- 这是 Apple Silicon 上有时出现的问题
- 使用 `print` 命令代替，或使用 `locals` 查看所有本地变量

## 📚 调试技巧

1. **条件断点**: 右键点击断点 → Edit Breakpoint → 输入条件 (e.g., `x > 10`)
2. **日志断点**: 右键点击断点 → Edit Breakpoint → 勾选 "Log Message"
3. **Watch 表达式**: 在 Debug 面板点击 "+" → 输入表达式 (e.g., `len(slice)`)

## 🎯 快速开始

```go
// 创建一个简单的程序来练习调试
package main

import "fmt"

func add(a, b int) int {
    return a + b  // <- 在这里设置断点
}

func main() {
    x := 5
    y := 3
    result := add(x, y)  // <- 也可以在这里设置断点
    fmt.Println(result)
}
```

1. 复制上面的代码到文件
2. 在 `return a + b` 行点击行号左边设置断点
3. 按 F5 启动调试
4. 观察变量值的变化

## ✨ 下一步

现在你可以:
- 打开任何 Go 文件
- 设置断点 (点击行号左边)
- 按 F5 开始调试
- 享受完整的 Go 调试体验!

有问题? 查看 [Delve 文档](https://github.com/go-delve/delve)
