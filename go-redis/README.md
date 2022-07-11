# Go-Redis项目

### 主要学习参考了HDT3213的[godis](https://github.com/HDT3213/godis)

* config
    - [config.go](config/config.go) 配置文件解析
* interface 一些模块间的接口定义
    * resp
        - [conn.go](interface/resp/conn.go)
        - [reply.go](interface/resp/reply.go)
    * tcp
        - [handler.go](interface/tcp/handler.go)
* lib 各种工具，比如logger、同步和通配符
    * log
        - [files.go](lib/logger/files.go)
        - [logger.go](lib/logger/logger.go) 日志
    * sync
        * atomic
            - [bool.go](lib/sync/atomic/bool.go) 对bool进行封装
        * wait
            - [wait.go](lib/sync/wait/wait.go) 对WaitGroup封装
* resp
    * parser
        - [parser.go](resp/parser/parser.go)
    * reply
        - [consts.go](resp/reply/consts.go)
        - [error.go](resp/reply/error.go)
        - [reply.go](resp/reply/reply.go)
* tcp tcp服务器实现
    - [echo.go](tcp/echo.go) 一个简单的 Echo 服务器，它会接受客户端连接并将客户端发送的内容原样传回客户端
    - [server.go](tcp/server.go) 可优雅关闭的服务器
* [main.go](main.go) 主程序
* [redis.conf](redis.conf) redis配置文件