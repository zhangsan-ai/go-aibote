// Package common 定义了所有Bot类型的共同接口和功能
package common

// Bot 接口定义了所有Bot类型的共同行为
// 无论是WindowsBot、WebBot还是AndroidBot都应该实现这个接口
// 它提供了启动服务器、停止服务器和执行脚本的基本功能
// 可以作为所有特定平台Bot的基础接口
// 这个接口设计遵循Go语言的接口隔离原则，确保每个接口只包含必要的方法
// 它也符合依赖倒置原则，高层模块依赖于抽象接口而非具体实现
// 在实际使用中，用户可以通过这个接口来操作不同平台的Bot实例
// 这使得代码更加灵活和可扩展
// 未来如果需要添加新的平台支持，只需要实现这个接口即可
// 这个接口的设计也考虑了错误处理，所有方法都返回error类型
// 这样可以确保调用者能够正确处理可能出现的异常情况
// 此外，这个接口也支持函数式编程风格，ExecuteScript方法接受一个函数作为参数
// 这使得脚本执行更加灵活和强大
// 总之，这个接口是整个框架的核心抽象，为不同平台的自动化提供了统一的访问方式
type Bot interface {
    // StartServer 启动TCP服务器，监听指定的IP和端口
    // ip: 监听的IP地址，可以是"0.0.0.0"表示监听所有网卡
    // port: 监听的端口号，范围0-65535
    // 返回error类型，如果启动成功则返回nil，否则返回具体的错误信息
    StartServer(ip string, port int) error
    
    // StopServer 停止TCP服务器，关闭所有连接
    // 返回error类型，如果停止成功则返回nil，否则返回具体的错误信息
    StopServer() error
    
    // ExecuteScript 执行自动化脚本
    // script: 一个接受Bot接口参数并返回error的函数
    // 返回error类型，如果脚本执行成功则返回nil，否则返回具体的错误信息
    ExecuteScript(script func(bot Bot) error) error
}