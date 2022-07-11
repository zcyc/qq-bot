# QQ 频道机器人

## 功能

* 一言
* 涩图

## 使用方式

### 1.修改配置

```
vim conf/config.toml
```

### 2.启动机器人

```bash
go run cmd/qq_bot/main.go
```

## 程序流程

### 1.连接网关

```
wss://api.sgroup.qq.com/websocket/
```

### 2.鉴权连接

```json
{
  "op": 2,
  "d": {
    "token": "APP_ID.APP_TOKEN",
    "intents": 513,
    "shard": [
      0,
      1
    ],
    "properties": {
      "$os": "linux",
      "$browser": "my_library",
      "$device": "my_library"
    }
  }
}
```

### 3.根据第一步的 heartbeat_interval 定时发送心跳包

```json
{
  "op": 1,
  "d": 251
}
```

`d` 为客户端收到的最新的消息的 `s`，如果是第一次连接，传 `null`

### 4.根据不同消息类型处理消息

```json
{
  "author": {
    "avatar": "http://thirdqq.qlogo.cn/0",
    "bot": false,
    "id": "1234",
    "username": "abc"
  },
  "channel_id": "100010",
  "content": "ndnnd",
  "guild_id": "18700000000001",
  "id": "0812345677890abcdef",
  "member": {
    "joined_at": "2021-04-12T16:34:42+08:00",
    "roles": [
      "1"
    ]
  },
  "timestamp": "2021-05-20T15:14:58+08:00",
  "seq": 101
}
```

推送次数有限制，尽量使用 msgId 来回复消息。

## 常见错误

### 图片无法发送

发消息接口有鉴黄，后续找一下其他接口。

### 请求超时

修改第`36`行超时时间。

```
vim integration/qq/client.go
```

## TODO

* 程序流程图
* 配置文件映射到 struct
* 消息分片
* 优雅的消息处理器
* 优雅的消息处理器注册
* 日志打印
* 监控、报警
* 单元测试

## 参考

* https://github.com/tencent-connect/botgo

* https://github.com/HKail/taskbot
