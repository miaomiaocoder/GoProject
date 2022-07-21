# Go-Redis项目

### 主要学习参考了HDT3213的[godis](https://github.com/HDT3213/godis)
go-redis 是一个用 Go 语言实现的 Redis 服务器。
### 项目目录
* config
    - [config.go](config/config.go) 配置文件解析
* aof AOF持久化实现
    - [aof.go](aof/aof.go)
* cluster 集群
    - [client_pool.go](cluster/client_pool.go)
    - [cluster_database.go](cluster/cluster_database.go)
    - [com.go](cluster/com.go) 节点间通信
    - [del.go](cluster/del.go) delete命令原子性实现
    - [flushdb.go](cluster/flushdb.go)
    - [ping.go](cluster/ping.go)
    - [rename.go](cluster/rename.go)
    - [router.go](cluster/router.go)
    - [select.go](cluster/select.go)
* config 配置文件解析
    - [config.go](config/config.go)
* database 存储引擎核心 
    - [database.go](database/database.go)
    - [db.go](database/db.go)
    - [echo_database.go](database/echo_database.go)
    - [keys.go](database/keys.go)
    - [ping.go](database/ping.go)
    - [string.go](database/string.go)
* datastruct redis的各类数据结构实现
    * dict
        - [dict.go](datastruct/dict/dict.go)
        - [syng_dict.go](datastruct/dict/sync_dict.go)
* interface 一些模块间的接口定义
    * database
        - [database.go](interface/database/database.go)
    * resp
        - [conn.go](interface/resp/conn.go)
        - [reply.go](interface/resp/reply.go)
    * tcp
        - [handler.go](interface/tcp/handler.go)
* lib 各种工具，比如logger、同步和通配符
    * consistenhash
        - [consistenthash.go](lib/consistenhash/consistenthash.go)
    * logger
        - [files.go](lib/logger/files.go)
        - [logger.go](lib/logger/logger.go) 日志
    * sync
        * atomic
            - [bool.go](lib/sync/atomic/bool.go) 对bool进行封装
        * wait
            - [wait.go](lib/sync/wait/wait.go) 对WaitGroup封装
    * utils
        - [utils.go](lib/utils/utils.go)
    * wildcard
        - [wildcard.go](lib/wildcard/wildcard.go)
* resp
    * client
        - [client.go](resp/client/client.go)
    * connection
        - [conn.go](resp/connection/conn.go)
    * handler
        - [handler.go](resp/handler/handler.go)
    * parser
        - [parser.go](resp/parser/parser.go) resp协议解析器实现
    * reply
        - [consts.go](resp/reply/consts.go)
        - [error.go](resp/reply/error.go)
        - [reply.go](resp/reply/reply.go)
* tcp tcp服务器实现
    - [echo.go](tcp/echo.go) 一个简单的 Echo 服务器，它会接受客户端连接并将客户端发送的内容原样传回客户端
    - [server.go](tcp/server.go) 可优雅关闭的服务器
* [main.go](main.go) 主程序
* [redis.conf](redis.conf) redis配置文件
* [appendonly.aof](appendonly.aof)