#

## 项目介绍

|      包          |      释义      |
| :-------------: | :---: |
|     gin         | web   |
|     gorm        | 数据库 |
|     zap         | 日志   |
|     viper       | 配置  |
|     jwt         | 认证  |
|     casbin      | 权限  |

## 项目结构

```shell
├── app.go         # 入口程序
├── conf           # 配置文件
│   ├── conf.toml  # web的基本配置
│   └── model.conf # casbin的配置文件
├── constant       # 常量
│   └── enums.go
├── global         # 整个项目都需要的东西
│   └── global.go
├── go.mod
├── go.sum
├── initialize     # 项目初始化相关
│   ├── config.go  # 初始化配置
│   ├── db.go      # 初始化db,其中内置了数据表的迁移和数据库初始化部分
│   ├── log.go     # 初始化log
│   ├── redis.go   # 初始化redis
│   └── router.go  # 初始化router
├── middleware     # 中间件
│   ├── auth.go    # 认证中间件
│   └── timer.go   # 耗时计算中间件
├── models         # 数据库
│   ├── Base.go
│   ├── Group.go
│   └── User.go
├── readme.md
├── router         # router内容
│   ├── router.go
│   ├── v1
│   │   ├── deployment.go
│   │   └── namespace.go
│   └── v2
├── schema
│   ├── conf.go
│   ├── jwt.go
│   └── user_info.go
├── utils         # 工具包
│   ├── casbin.go
│   ├── error.go  # 统一错误处理
│   ├── jwt.go
│   ├── log.go    # log的配置
│   └── response.go # 返回封装
└── zap.log       # 日志
```
