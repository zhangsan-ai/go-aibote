# Go-Aibote

Go-Aibote是一个支持Windows、Web和Android平台的RPA(机器人流程自动化)框架的Go语言实现，利用Go语言的并发特性和跨平台优势，提供高效、可靠的自动化解决方案。

## 项目特性

1. **多平台支持**：支持Windows、Web和Android三种平台的自动化操作
2. **高性能**：利用Go语言的并发特性，提供高效的自动化解决方案
3. **TCP通信**：采用TCP协议与底层驱动程序通信，确保高效可靠的数据传输
4. **丰富的API**：提供丰富的自动化API，包括元素操作、窗口管理、键鼠控制等
5. **类型安全**：利用Go语言的类型系统，提供更好的代码安全性和开发体验

## 项目结构

```
go-aibote/
├── cmd/                    # 可执行文件目录
│   ├── windows-bot/        # Windows自动化示例
│   ├── web-bot/            # Web自动化示例
│   └── android-bot/        # Android自动化示例
├── internal/               # 内部实现代码
│   ├── common/             # 共享的内部组件
│   ├── windows/            # Windows平台内部实现
│   ├── web/                # Web平台内部实现
│   └── android/            # Android平台内部实现
├── pkg/                    # 公开的包和接口
│   ├── common/             # 通用接口定义
│   ├── windowsbot/         # WindowsBot接口和实现
│   ├── webbot/             # WebBot接口和实现
│   └── androidbot/         # AndroidBot接口和实现
├── go.mod                  # Go模块定义
└── go.sum                  # 依赖版本锁定
```

## 安装

```bash
# 克隆仓库
git clone https://github.com/yourusername/go-aibote.git

# 进入项目目录
cd go-aibote

# 安装依赖
go mod download
```

## 快速开始

### Windows平台自动化示例

```go
package main

import (
    "fmt"
    "log"
    "github.com/yourusername/go-aibote/pkg/windowsbot"
)

func main() {
    // 创建WindowsBot实例
    bot, err := windowsbot.NewWindowsBot(
        windowsbot.WithLogLevel(windowsbot.LogLevelDebug),
        windowsbot.WithLogStorage(true),
    )
    if err != nil {
        log.Fatalf("Failed to create WindowsBot: %v", err)
    }
    
    // 定义自动化脚本
    script := func(b windowsbot.WindowsBot) error {
        // 查询所有窗口
        windows, err := b.FindWindows()
        if err != nil {
            return err
        }
        
        // 打印窗口信息
        fmt.Println("Found windows:")
        for _, window := range windows {
            fmt.Printf("  Hwnd: %s, Title: %s\n", window.Hwnd, window.Title)
        }
        
        return nil
    }
    
    // 启动服务并执行脚本
    err = bot.StartServer("0.0.0.0", 9999)
    if err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
    defer bot.StopServer()
    
    err = bot.ExecuteScript(script)
    if err != nil {
        log.Fatalf("Script execution failed: %v", err)
    }
}
```

### Web平台自动化示例

```go
package main

import (
    "fmt"
    "log"
    "github.com/yourusername/go-aibote/pkg/webbot"
)

func main() {
    // 创建WebBot实例
    bot, err := webbot.NewWebBot(
        webbot.WithBrowser(webbot.BrowserChrome),
        webbot.WithUserDataDir("./UserData"),
        webbot.WithImplicitWait(10.0),
    )
    if err != nil {
        log.Fatalf("Failed to create WebBot: %v", err)
    }
    
    // 定义自动化脚本
    script := func(b webbot.WebBot) error {
        // 导航到百度
        err := b.Goto("https://www.baidu.com")
        if err != nil {
            return err
        }
        
        // 获取页面标题
        title, err := b.GetTitle()
        if err != nil {
            return err
        }
        fmt.Printf("Page title: %s\n", title)
        
        return nil
    }
    
    // 启动服务并执行脚本
    err = bot.StartServer("0.0.0.0", 9999)
    if err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
    defer bot.StopServer()
    
    err = bot.ExecuteScript(script)
    if err != nil {
        log.Fatalf("Script execution failed: %v", err)
    }
}
```

### Android平台自动化示例

```go
package main

import (
    "fmt"
    "log"
    "github.com/yourusername/go-aibote/pkg/androidbot"
)

func main() {
    // 创建AndroidBot实例
    bot, err := androidbot.NewAndroidBot()
    if err != nil {
        log.Fatalf("Failed to create AndroidBot: %v", err)
    }
    
    // 定义自动化脚本
    script := func(b androidbot.AndroidBot) error {
        // 获取最近任务列表
        tasks, err := b.RecentTasks()
        if err != nil {
            return err
        }
        
        // 打印任务信息
        fmt.Println("Recent tasks:")
        for _, task := range tasks {
            fmt.Printf("  Name: %s, Package: %s\n", task.Name, task.PackageName)
        }
        
        return nil
    }
    
    // 启动服务并执行脚本
    err = bot.StartServer("0.0.0.0", 8888)
    if err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
    defer bot.StopServer()
    
    err = bot.ExecuteScript(script)
    if err != nil {
        log.Fatalf("Script execution failed: %v", err)
    }
}
```

## 与PyAibote的区别

1. **语言差异**：使用Go语言替代Python，提供更好的性能和并发支持
2. **接口设计**：采用Go语言的接口设计模式，更加符合Go的编程风格
3. **错误处理**：使用Go风格的错误处理机制，而非Python的异常机制
4. **并发模型**：利用Go的goroutine和channel实现高效的并发处理
5. **类型系统**：利用Go的静态类型系统，提供更好的代码安全性

## 注意事项

1. **驱动程序**：Go-Aibote需要配合原PyAibote的驱动程序使用(WindowsDriver.exe、WebDriver.exe等)
2. **端口配置**：确保Go-Aibote服务监听的端口与驱动程序配置的端口一致
3. **调试模式**：在开发环境中可以启用调试模式，生产环境中建议关闭
4. **资源管理**：使用defer语句确保服务器和客户端资源正确关闭

## 贡献指南

欢迎贡献代码、报告问题或提出建议。请遵循以下步骤：

1. Fork本仓库
2. 创建特性分支
3. 提交更改
4. 推送到远程分支
5. 创建Pull Request

## 许可证

本项目采用MIT许可证，详情请参见LICENSE文件。