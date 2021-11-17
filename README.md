


目录:
- [nacos-cli]()
- [构建安装]()
  - [源码安装]()
  - [二进制安装]()
- [命令指南]()
  - [config命令]()
    - [add]()
    - [get]()
    - [del]()
    - [list]()
    - [version]()

# nacos-cli

**Nacos** 命令行工具，基于 **Nacos OpenApi** 封装的命令行工具。提供一些对配置、服务注册与发现、命名空间的简单操作 。借助 **go** 语言的跨平台交叉编译，将编译好的二进制文件直接放到指定系统下就可以直接运行， 无需环境部署。

# 构建安装

## 源码安装

从 **github** 将源码 **clone** 到本地, 编译源码适用于想自定义命令时候，可以拿到源码二次开发后重新构建。这种方式就需要使用 **go** 环境了。

```bash
git clone https://github.com/is-gypsophila/nacos-cli.git
```

使用  **go build -o nacos-cli** 编译项目，**-o** 执行指定输出文件为 **nacos-cli** 。

到这里是不是有点熟悉，是不是想起 **reids-cli** ？没错就是模仿 **redis-cli** 起的名字，希望能够像 **reids-cli** 一样好用。

将编译好的 **nacos-cli** 可执行二进制文件放到 **/usr/local/bin** 的目录下，就可以使用了。

## 二进制安装

直接下载编译好的二进制文件 **nacos-cli** ，然后将它放到 **/usr/local/bin** 目录下。原理同上。

> 由于执行命令需要用到 Nacos 的服务 address、port 两个参数，这里有两种指定方式。第一种是通过命令的 Flag 指定，第二种是在 $HOME 的目录下创建 config.yaml , 在 yaml 文件中指定 address、port。例如：
>

```yaml
server:
  address: 127.0.0.1
  port: 8081
```

# 命令指南

## config 命令

可以通过 **nacos-cli config -h** 查看命令怎么使用

```bash
Usage:
  nacos-cli config [command]

Available Commands:
  add         Add a configuration
  del         Delete a configuration
  get         Get configuration information
  list        Configuration information list
  version     Query the configuration information of the previous version

Flags:
  -h, --help   help for config
```

### add

添加一个配置

```bash
Usage:
  nacos-cli config add [namespaceId] [dataId] [group] [content] [type] [flags]

Flags:
  -b, --address string   nacos server ip address (default "127.0.0.1")
  -h, --help             help for add
  -p, --port string      nacos server port (default "8848")
```

```bash
例子 ：
nacos-cli config add '' test-dataId '' 添加一个文本内容 '' -b 127.0.0.1 -p 8848

输出：
success
```

- [namespaceId] 不指定默认为 public
- [group] 不指定默认为 DEFAULT_GROUP
- [type] 不指定默认为 text

### get

获取指定配置信息

```bash
Usage:
  nacos-cli config get [namespaceId] [dataId] [group] [flags]

Flags:
  -b, --address string   nacos server ip address (default "127.0.0.1")
  -h, --help             help for get
  -p, --port string      nacos server port (default "8848");
```

```bash
例子：
nacos-cli config get '' test-dataId ''

输出：
{
    "id": "434908075100168192",
    "dataId": "test-dataId",
    "group": "DEFAULT_GROUP",
    "content": "添加一个文本内容",
    "md5": "fd331d460d327ad1087fb865e4d9bfb2",
    "tenant": "",
    "type": "text"
}
```

### del

删除配置

```bash
Usage:
  nacos-cli config del [namespaceId] [dataId] [group] [flags]

Flags:
  -b, --address string   nacos server ip address (default "127.0.0.1")
  -h, --help             help for del
  -p, --port string      nacos server port (default "8848")
```

```bash
例子：
nacos-cli config del '' test-1 ''

输出：
success
```

### list

查询配置列表，可分页获取

```bash
Usage:
  nacos-cli config list [pageNo] [pageSize] [flags]

Flags:
  -b, --address string   nacos server ip address (default "127.0.0.1")
  -h, --help             help for list
  -p, --port string      nacos server port (default "8848")
```

```bash
例子：
nacos-cli config list 1 10

输出：
+--------------------+------------------+---------------+------------------------------------------------------------------+
| id                 | dataId           | group         | content                                                          |
+--------------------+------------------+---------------+------------------------------------------------------------------+
| 428019419571793920 | cipher-aes-test5 | DEFAULT_GROUP | 3eb10eec917021984dd88f2e9d87c1b005c3cd9fa52c15fdb389b136950a9a18 |
| 434257417337851904 | test-2           | DEFAULT_GROUP | 修改内容                                                          |
| 433167120500219904 | user.info        | DEFAULT_GROUP | name: 哈哈1就斯蒂芬...                                             |
| 433171487634710528 | test-1           | DEFAULT_GROUP | 啦啦啦 啦啦啦                                                      |
| 433172056990507008 | example          | DEFAULT_GROUP | useLocalCache=true...                                            |
| 434307109421600768 | test-6666        | DEFAULT_GROUP | server:...                                                       |
| 434908075100168192 | test-dataId      | DEFAULT_GROUP | 添加一个文本内容                                                    |
+--------------------+------------------+---------------+------------------------------------------------------------------+
```

> 如果content有换行 则只展示第一行 后面内容用...标识，如果长度大于80就截断，后续内容也用...表示。
>

### version

查看配置的上一个版本

```bash
Usage:
  nacos-cli config version [id] [namespaceId] [dataId] [group] [flags]

Flags:
  -b, --address string   nacos server ip address (default "127.0.0.1")
  -h, --help             help for version
  -p, --port string      nacos server port (default "8848")
```

```bash
例子：
nacos-cli config version 434307109421600768 '' test-6666 ''

输出：
{
    "id": "701",
    "dataId": "test-6666",
    "group": "DEFAULT_GROUP",
    "content": "server:\n  address: 127.0.0.1\n  port: 8081\nserver:\n  address: 127.0.0.1\n  port: 8081\n  server:\n  address: 127.0.0.1\n  port: 8081\nserver:\n  address: 127.0.0.1\n  port: 8081",
    "md5": "3e282efbb24f97c55791cfc5f1ddfce2",
    "tenant": "",
    "type": ""
}
```