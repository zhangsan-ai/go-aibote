// Package main 提供WindowsBot的使用示例
package main

import (
    "fmt"
    "log"
    "github.com/zhangsan-ai/go-aibote/pkg/common"
    "github.com/zhangsan-ai/go-aibote/pkg/windowsbot"
)

func main() {
    // 创建WindowsBot实例
    // 使用函数选项模式配置WindowsBot
    // 设置日志级别为DEBUG，启用日志存储，启用调试模式
    bot, err := windowsbot.NewWindowsBot(
        windowsbot.WithLogLevel(windowsbot.LogLevelDebug),
        windowsbot.WithLogStorage(true),
        windowsbot.WithDebugMode(true),
    )
    
    if err != nil {
        log.Fatalf("Failed to create WindowsBot instance: %v", err)
    }
    
    // 定义自动化脚本函数
    // 注意：script函数必须接受common.Bot类型参数才能传递给ExecuteScript方法
    // 这是因为common.Bot接口定义的ExecuteScript方法要求参数类型为func(bot Bot) error
    script := func(bot common.Bot) error {
        // 将通用Bot接口转换为WindowsBot接口，以便使用Windows特有的方法
        // 这是因为ExecuteScript方法是通过common.Bot接口调用的，需要类型转换才能使用特定平台的功能
        windowsBot, ok := bot.(windowsbot.WindowsBot)
        if !ok {
            return fmt.Errorf("failed to convert Bot to WindowsBot")
        }
        
        // 查询所有窗口句柄
        windows, err := windowsBot.FindWindows()
        if err != nil {
            return fmt.Errorf("failed to find windows: %w", err)
        }
        
        // 打印找到的窗口信息
        fmt.Println("Found windows:")
        for _, window := range windows {
            fmt.Printf("  Hwnd: %s, Title: %s, ClassName: %s\n", 
                window.Hwnd, window.Title, window.ClassName)
        }
        
        // 在实际应用中，可以根据窗口标题或类名找到特定窗口
        // 然后对窗口中的元素进行操作
        
        // 示例：假设找到了一个窗口，尝试获取窗口中的元素信息
        // if len(windows) > 0 {
        //     window := windows[0]
        //     elementName, err := windowsBot.GetElementName(window.Hwnd, "//button[@id='submit']")
        //     if err != nil {
        //         return fmt.Errorf("failed to get element name: %w", err)
        //     }
        //     fmt.Printf("Element name: %s\n", elementName)
        // }
        
        // 关闭驱动程序(可以根据需要选择合适的方法)
        // err = windowsBot.CloseDriver()
        // if err != nil {
        //     return fmt.Errorf("failed to close driver: %w", err)
        // }
        
        return nil
    }
    
    // 启动TCP服务器，监听指定的IP和端口
    // "0.0.0.0"表示监听所有网卡
    // 端口9999是示例端口，需要与客户端驱动程序配置一致
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
    
    // 执行自动化脚本
    // 注意：这里传递的script函数必须与common.Bot接口中ExecuteScript方法定义的参数类型匹配
    err = bot.ExecuteScript(script)
    if err != nil {
        log.Fatalf("Failed to execute script: %v", err)
    }
    
    fmt.Println("Windows automation script executed successfully")
}