# nacosctl

**Nacos** 命令行工具，基于 **Nacos OpenApi** 封装的命令行工具。提供一些对配置、服务注册与发现、命名空间等命令操作 。借助 **go** 语言的跨平台交叉编译，将编译好的二进制文件直接放到指定系统下就可以直接运行， 无需环境部署。

![](https://tva1.sinaimg.cn/large/008i3skNly1gz63zy4pmtj310o0in75z.jpg)

# 构建安装

## 源码安装

从 **github** 将源码 **clone** 到本地, 编译源码适用于想自定义命令时候，可以拿到源码二次开发后重新构建。

```bash
git clone https://github.com/li-xiao-shuang/nacosctl.git
```

使用  **go build nacosctl.go** 编译项目 。

将编译好的 **nacosctl** 可执行二进制文件放到 **/usr/local/bin** 的目录下，就可以使用了。

## 二进制安装

直接下载编译好的二进制文件 **[nacosctl](https://github.com/li-xiao-shuang/nacosctl/releases/download/1.1.1/nacosctl.tar.gz)** ，然后将它放到 **/usr/local/bin** 目录下。

## Homebrew 安装

执行命令：
```
brew tap li-xiao-shuang/nacosctl
brew install nacosctl
```

# 指南
[命令指南](GUIDE.md)
