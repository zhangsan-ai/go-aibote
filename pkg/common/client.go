// Package common 定义了所有Bot类型的共同接口和功能
package common

// Client 接口定义了与驱动程序通信的方法
// 它是Bot与底层驱动程序(如WindowsDriver.exe、WebDriver.exe等)通信的桥梁
// 通过这个接口，Bot可以向驱动程序发送命令并接收响应
// 这个接口设计考虑了通信的基本需求：连接、发送命令和关闭连接
// 它也支持灵活的命令参数传递，使用可变参数和interface{}类型
// 这样可以适应不同命令可能需要的不同类型和数量的参数
// 所有方法都返回error类型，确保调用者能够正确处理可能出现的异常情况
// 在实际实现中，这个接口通常会基于TCP协议与驱动程序通信
// 但具体的通信协议细节对接口使用者是透明的
// 这种设计使得通信逻辑与业务逻辑分离，提高了代码的可维护性和可测试性
// 总之，这个接口是整个框架与底层驱动程序交互的核心抽象
// 它使得上层Bot接口能够以统一的方式与不同平台的驱动程序通信
type Client interface {
    // Connect 连接到指定的IP和端口的驱动程序
    // ip: 驱动程序所在的IP地址
    // port: 驱动程序监听的端口号
    // 返回error类型，如果连接成功则返回nil，否则返回具体的错误信息
    Connect(ip string, port int) error
    
    // SendCommand 向驱动程序发送命令并接收响应
    // cmd: 命令名称，如"getElementName"、"findWindows"等
    // params: 命令参数，支持任意类型和数量的参数
    // 返回两个值：响应结果字符串和error类型
    // 如果发送命令成功并收到响应，则返回响应结果和nil
    // 否则返回空字符串和具体的错误信息
    SendCommand(cmd string, params ...interface{}) (string, error)
    
    // Close 关闭与驱动程序的连接
    // 释放相关资源，如网络连接等
    // 返回error类型，如果关闭成功则返回nil，否则返回具体的错误信息
    Close() error
}