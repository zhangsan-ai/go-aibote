# Go-Aibote库整理概述

## 1. 项目简介

Go-Aibote是一个支持Windows、Web和Android平台的RPA(机器人流程自动化)框架的Go语言实现，利用Go语言的并发特性和类型安全优势，提供高效、可靠的自动化解决方案。

## 2. 项目结构

Go-Aibote采用清晰的分层架构，遵循Go项目的标准布局。核心代码位于`pkg`目录中，提供公开的API接口；示例代码位于`cmd`目录中，展示库的实际使用方法。

```
go-aibote/
├── cmd/                    # 可执行文件和示例代码
│   └── windows-bot/        # Windows自动化示例
├── pkg/                    # 公开的包和接口
│   ├── common/             # 通用接口定义
│   ├── windowsbot/         # Windows平台自动化
│   ├── webbot/             # Web平台自动化
│   └── androidbot/         # Android平台自动化
├── go.mod                  # Go模块定义
└── README.md               # 项目文档
```

## 3. 核心接口设计

### 3.1 通用接口

#### common.Bot接口

`Bot`接口是所有自动化机器人的基础抽象，定义了启动服务器、停止服务器和执行脚本的核心功能。

```go
// Bot 接口定义了所有Bot类型的共同行为
type Bot interface {
    // 启动TCP服务器，监听指定的IP和端口
    StartServer(ip string, port int) error
    
    // 停止TCP服务器，关闭所有连接
    StopServer() error
    
    // 执行自动化脚本
    ExecuteScript(script func(bot Bot) error) error
}
```

#### common.Client接口

`Client`接口定义了与底层驱动程序通信的方法，是Bot与驱动程序交互的桥梁。

```go
// Client 接口定义了与驱动程序通信的方法
type Client interface {
    // 连接到指定的IP和端口的驱动程序
    Connect(ip string, port int) error
    
    // 向驱动程序发送命令并接收响应
    SendCommand(cmd string, params ...interface{}) (string, error)
    
    // 关闭与驱动程序的连接
    Close() error
}
```

### 3.2 平台特定接口

#### WindowsBot接口

`WindowsBot`接口继承自`common.Bot`，提供Windows平台特定的自动化功能，如窗口操作、元素定位等。

```go
// WindowsBot 接口定义了Windows平台自动化的方法
type WindowsBot interface {
    common.Bot
    
    // 查找所有可见的Windows窗口
    FindWindows() ([]Window, error)
    
    // 获取元素名称
    GetElementName(hwnd string, xpath string) (string, error)
    
    // 获取元素文本
    GetElementValue(hwnd string, xpath string) (string, error)
    
    // 获取元素矩形
    GetElementRect(hwnd string, xpath string) (Rect, error)
    
    // 关闭本地驱动程序
    CloseDriverLocal() error
    
    // 关闭驱动程序
    CloseDriver() error
    
    // 其他Windows自动化方法...
}
```

#### WebBot接口

`WebBot`接口继承自`common.Bot`，提供Web平台特定的自动化功能，如页面导航、元素操作等。

```go
// WebBot 接口定义了Web平台自动化的方法
type WebBot interface {
    common.Bot
    
    // 导航到指定的URL
    Goto(url string) error
    
    // 查找单个网页元素
    FindElement(selector string) (WebElement, error)
    
    // 查找多个网页元素
    FindElements(selector string) ([]WebElement, error)
    
    // 获取当前网页的标题
    GetTitle() (string, error)
    
    // 获取当前网页的URL
    GetURL() (string, error)
    
    // 刷新当前网页
    Refresh() error
    
    // 其他Web自动化方法...
}
```

#### AndroidBot接口

`AndroidBot`接口继承自`common.Bot`，提供Android平台特定的自动化功能，如任务管理、坐标操作等。

```go
// AndroidBot 接口定义了Android平台自动化的方法
type AndroidBot interface {
    common.Bot
    
    // 显示手机最近任务列表
    RecentTasks() ([]Task, error)
    
    // 在指定坐标点点击
    Tap(x, y int) error
    
    // 滑动屏幕
    Swipe(x1, y1, x2, y2 int) error
    
    // 获取已安装的应用包列表
    GetInstalledPackages() ([]string, error)
    
    // 启动指定的应用
    StartApp(packageName string) error
    
    // 停止指定的应用
    StopApp(packageName string) error
    
    // 其他Android自动化方法...
}
```

## 4. 数据结构

### 4.1 Windows平台

```go
// Rect 表示一个矩形区域
type Rect struct {
    X1, Y1 float64 // 左上角坐标
    X2, Y2 float64 // 右下角坐标
}

// Window 表示一个Windows窗口
type Window struct {
    Hwnd      string // 窗口句柄
    Title     string // 窗口标题
    ClassName string // 窗口类名
}
```

### 4.2 Web平台

```go
// WebElement 表示一个网页元素
type WebElement struct {
    ID          string // 元素ID
    XPath       string // 元素的XPath路径
    CSSSelector string // 元素的CSS选择器
    TagName     string // 元素的标签名
}
```

### 4.3 Android平台

```go
// Task 表示Android设备上的一个任务
type Task struct {
    ID          string // 任务ID
    Name        string // 任务名称
    PackageName string // 应用包名
}
```

## 5. 依赖管理

Go-Aibote使用Go模块系统进行依赖管理，主要依赖包括：

```go
module github.com/yourusername/go-aibote

go 1.21

require (
    github.com/pkg/errors v0.9.1        // 增强的错误处理
    github.com/shirou/gopsutil v3.21.11+incompatible // 系统信息获取
    go.uber.org/zap v1.26.0             // 高性能日志库
)
```

## 6. 使用示例

### 6.1 Windows自动化示例

以下是使用WindowsBot进行自动化的基本示例：

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
        windowsbot.WithDebugMode(true),
    )
    
    if err != nil {
        log.Fatalf("Failed to create WindowsBot instance: %v", err)
    }
    
    // 定义自动化脚本
    script := func(bot windowsbot.WindowsBot) error {
        // 查询所有窗口
        windows, err := bot.FindWindows()
        if err != nil {
            return fmt.Errorf("failed to find windows: %w", err)
        }
        
        // 打印窗口信息
        fmt.Println("Found windows:")
        for _, window := range windows {
            fmt.Printf("  Hwnd: %s, Title: %s, ClassName: %s\n", 
                window.Hwnd, window.Title, window.ClassName)
        }
        
        return nil
    }
    
    // 启动服务器
    err = bot.StartServer("0.0.0.0", 9999)
    if err != nil {
        log.Fatalf("Failed to start WindowsBot server: %v", err)
    }
    
    // 确保在程序退出时停止服务器
    defer func() {
        if err := bot.StopServer(); err != nil {
            log.Printf("Failed to stop WindowsBot server: %v", err)
        }
    }()
    
    // 执行脚本
    err = bot.ExecuteScript(script)
    if err != nil {
        log.Fatalf("Failed to execute script: %v", err)
    }
}
```

## 7. 设计模式与最佳实践

Go-Aibote库采用了多种Go语言的设计模式和最佳实践：

1. **接口设计模式**：通过定义清晰的接口，实现了高内聚、低耦合的设计
2. **函数选项模式**：使用函数选项模式配置Bot实例，提供灵活的初始化方式
3. **错误处理**：使用`github.com/pkg/errors`库提供丰富的错误上下文信息
4. **资源管理**：使用`defer`确保资源正确释放，避免资源泄漏
5. **类型安全**：充分利用Go语言的静态类型系统，提供更好的编译时检查

## 8. 注意事项与使用建议

1. **TCP通信**：库通过TCP协议与底层驱动程序通信，请确保端口配置正确且防火墙允许通信
2. **错误处理**：所有方法都返回错误信息，请确保正确处理这些错误
3. **资源管理**：使用完Bot实例后，请调用`StopServer`方法释放资源
4. **并发安全**：库的并发安全性取决于具体实现，请参考各模块的文档
5. **版本兼容性**：确保使用的Go版本与项目要求一致（Go 1.21+）

## 9. 下一步开发计划

1. 完成各平台Bot接口的具体实现
2. 编写单元测试和集成测试
3. 完善文档和示例代码
4. 优化性能和稳定性
5. 添加更多高级功能，如OCR识别、图像匹配等

## 10. 总结

Go-Aibote是一个功能完整的RPA框架，通过清晰的接口设计和Go语言的优势，为Windows、Web和Android平台的自动化提供了统一的解决方案。利用Go语言的并发特性和类型安全优势，它是进行跨平台自动化开发的理想选择。