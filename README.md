# go-web-template
Go语言web开发模板

```
go-web-template
├─api
│  └─http
├─cmd   放main.go和配置文件, 作为启动入口
├─conf  放配置文件对应的golang struct, 使用的是toml
├─dao   data access object, 数据库访问方法, redis, memcache访问方法, 还有一些RPC调用也放在这里面
├─model 放结构体, 比如Http参数转换用的struct, DB存储对应的struct, 各层之间传递用的struct
├─library   项目内部包
│  ├─database
│  │  └─orm 
│  └─log 日志库
├─server    主要是提供协议转换, 聚合. 逻辑还是再service层做
│  └─http   提供http服务
└─service   对于后端服务来说, 该目录提供服务的实现, 对于http服务, 该目录提供http服务的实现.
```