// Package main 提供Web平台自动化的示例
package main

import (
    "fmt"
    "github.com/zhangsan-ai/go-aibote/pkg/common"
    "github.com/zhangsan-ai/go-aibote/pkg/webbot"
)

// script 定义了一个Web自动化脚本
// 它接收一个common.Bot类型的参数
// 返回error类型
// 在实际使用中，这个函数会包含具体的Web自动化操作
func script(bot common.Bot) error {
    // 将common.Bot类型转换为webbot.WebBot类型
    webBot, ok := bot.(webbot.WebBot)
    if !ok {
        return fmt.Errorf("failed to cast Bot to WebBot")
    }

    // 打开指定URL
    err := webBot.Goto("https://www.example.com")
    if err != nil {
        return fmt.Errorf("failed to navigate to URL: %v", err)
    }

    // 查找元素示例
    // element, err := webBot.FindElement("css selector", "input[type='text']")
    // if err != nil {
    //     return fmt.Errorf("failed to find element: %v", err)
    // }

    // 点击元素示例
    // err = webBot.ClickElement(element)
    // if err != nil {
    //     return fmt.Errorf("failed to click element: %v", err)
    // }

    // 输入文本示例
    // err = webBot.InputText(element, "Hello, World!")
    // if err != nil {
    //     return fmt.Errorf("failed to input text: %v", err)
    // }

    // 获取元素文本示例
    // text, err := webBot.GetElementText(element)
    // if err != nil {
    //     return fmt.Errorf("failed to get element text: %v", err)
    // }
    // fmt.Println("Element text:", text)

    // 关闭浏览器驱动
    // err = webBot.CloseDriver()
    // if err != nil {
    //     return fmt.Errorf("failed to close driver: %v", err)
    // }

    return nil
}

func main() {
    // 创建WebBot实例，使用默认配置
    bot, err := webbot.NewWebBot()
    if err != nil {
        fmt.Printf("Failed to create WebBot: %v\n", err)
        return
    }

    // 启动TCP服务器（可选）
    // 可以通过网络连接来控制这个WebBot实例
    // err = bot.StartServer()
    // if err != nil {
    //     fmt.Printf("Failed to start server: %v\n", err)
    //     return
    // }
    // defer bot.StopServer()

    // 执行Web自动化脚本
    err = bot.ExecuteScript(script)
    if err != nil {
        fmt.Printf("Script execution failed: %v\n", err)
        return
    }

    fmt.Println("Web automation script executed successfully")
}