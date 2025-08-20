// Package main 提供Android平台自动化的示例
package main

import (
    "fmt"
    "github.com/zhangsan-ai/go-aibote/pkg/androidbot"
    "github.com/zhangsan-ai/go-aibote/pkg/common"
)

// script 定义了一个Android自动化脚本
// 它接收一个common.Bot类型的参数
// 返回error类型
// 在实际使用中，这个函数会包含具体的Android自动化操作
func script(bot common.Bot) error {
    // 将common.Bot类型转换为androidbot.AndroidBot类型
    androidBot, ok := bot.(androidbot.AndroidBot)
    if !ok {
        return fmt.Errorf("failed to cast Bot to AndroidBot")
    }

    // 获取最近任务列表
    tasks, err := androidBot.RecentTasks()
    if err != nil {
        return fmt.Errorf("failed to get recent tasks: %v", err)
    }

    // 打印最近任务信息
    fmt.Println("Recent tasks:")
    for i, task := range tasks {
        fmt.Printf("Task %d: ID=%s, Name=%s, PackageName=%s\n", i+1, task.ID, task.Name, task.PackageName)
    }

    // 点击指定坐标（示例）
    // err = androidBot.Tap(500, 500)
    // if err != nil {
    //     return fmt.Errorf("failed to tap: %v", err)
    // }

    // 滑动屏幕（示例）
    // err = androidBot.Swipe(500, 1000, 500, 500)
    // if err != nil {
    //     return fmt.Errorf("failed to swipe: %v", err)
    // }

    // 获取已安装的应用包列表
    // packages, err := androidBot.GetInstalledPackages()
    // if err != nil {
    //     return fmt.Errorf("failed to get installed packages: %v", err)
    // }
    // fmt.Println("Installed packages count:", len(packages))

    // 启动应用（示例）
    // err = androidBot.StartApp("com.example.app")
    // if err != nil {
    //     return fmt.Errorf("failed to start app: %v", err)
    // }

    // 停止应用（示例）
    // err = androidBot.StopApp("com.example.app")
    // if err != nil {
    //     return fmt.Errorf("failed to stop app: %v", err)
    // }

    // 截取屏幕截图（示例）
    // err = androidBot.TakeScreenshot("./screenshot.png")
    // if err != nil {
    //     return fmt.Errorf("failed to take screenshot: %v", err)
    // }

    // 通过XPath查找元素（示例）
    // x, y, err := androidBot.FindElementByXPath("//button[@text='OK']")
    // if err != nil {
    //     return fmt.Errorf("failed to find element by XPath: %v", err)
    // }
    // if x >= 0 && y >= 0 {
    //     fmt.Printf("Element found at (%d, %d)\n", x, y)
    // }

    // 发送按键事件（示例）
    // err = androidBot.SendKeyEvent(4) // 返回键
    // if err != nil {
    //     return fmt.Errorf("failed to send key event: %v", err)
    // }

    // 输入文本（示例）
    // err = androidBot.InputText("Hello, Android!")
    // if err != nil {
    //     return fmt.Errorf("failed to input text: %v", err)
    // }

    return nil
}

func main() {
    // 创建AndroidBot实例，使用默认配置
    bot, err := androidbot.NewAndroidBot()
    if err != nil {
        fmt.Printf("Failed to create AndroidBot: %v\n", err)
        return
    }

    // 启动TCP服务器（可选）
    // 可以通过网络连接来控制这个AndroidBot实例
    // err = bot.StartServer()
    // if err != nil {
    //     fmt.Printf("Failed to start server: %v\n", err)
    //     return
    // }
    // defer bot.StopServer()

    // 执行Android自动化脚本
    err = bot.ExecuteScript(script)
    if err != nil {
        fmt.Printf("Script execution failed: %v\n", err)
        return
    }

    fmt.Println("Android automation script executed successfully")
}