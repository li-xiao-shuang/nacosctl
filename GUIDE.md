# 命令指南

使用 **nacosctl -h**  查看命令的提示

## login 命令

输入 nacosctl login -h 查看提示

```shell
Set username and password for nacos authentication

Usage:
  nacosctl login [username] [password] [flags]

Examples:
nacosctl login [resource]

Flags:
  -h, --help   help for login

Global Flags:
  -a, --address string   nacos server ip+port (default "127.0.0.1:8848")
```

执行示例：

```shell
☁  ~  nacosctl login nacos nacos
login successful
```

## config 命令

### 创建配置

输入 nacosctl create config -h 查看提示

```shell
Create a configuration

Usage:
  nacosctl create config [dataId] [content] [flags]

Examples:
nacosctl create config [dataId] [content]

Flags:
  -g, --group string         config group (default "DEFAULT_GROUP")
  -h, --help                 help for config
  -n, --namespaceId string   namespace id
  -t, --type string          config type (default "text")

Global Flags:
  -a, --address string   nacos server ip+port (default "127.0.0.1:8848")
```

执行示例：

```shell
☁  ~  nacosctl create config test_config  测试配置内容
done
```

### 查询配置

输入 nacosctl get config -h 查看提示

```shell
Get configuration information

Usage:
  nacosctl get config [dataId] [flags]

Examples:
nacosctl get config [dataId]

Flags:
  -g, --group string         config group (default "DEFAULT_GROUP")
  -h, --help                 help for config
  -n, --namespaceId string   namespace id
  -v, --version string       To get the previous version of the configuration, you need to specify the configuration id

Global Flags:
  -a, --address string   nacos server ip+port (default "127.0.0.1:8848")
```

执行示例：

```shell
☁  ~  nacosctl get config user.info
ID:      	433167120500219904
DataId:  	user.info
命名空间:
组名:    	DEFAULT_GROUP
MD5:     	c5b6b2f95e2e8b538b794c53ff2b614e
配置类型:	yaml
配置内容:	hh: 新增
```

### 查询配置的上一个版本

执行 nacosctl get config user.info -v=[配置id]

执行示例：

```shell
☁  ~  nacosctl get config user.info -v=433167120500219904
ID:      	1607
DataId:  	user.info
命名空间:
组名:    	DEFAULT_GROUP
MD5:     	7b1964fd7389703812be7ab29023a4c0
配置类型:
配置内容:	name: 哈哈1就斯蒂芬
         	age: 212
```

### 查询配置列表

输入 nacosctl get configs -h 查看提示

```shell
Get a list of configuration information in paginated form

Usage:
  nacosctl get configs [pageNo] [pageSize] [flags]

Examples:
nacosctl get configs [pageNo] [pageSize]

Flags:
  -g, --group string         config group (default "DEFAULT_GROUP")
  -h, --help                 help for configs
  -n, --namespaceId string   namespace id

Global Flags:
  -a, --address string   nacos server ip+port (default "127.0.0.1:8848")
```

执行示例：

```shell
ID:               	DataId:         	命名空间:	组名:        	配置类型:	配置内容:
428019419571793920	cipher-aes-test5	         	DEFAULT_GROUP	         	3eb10eec917021984dd88f2e9d8
433167120500219904	user.info       	         	DEFAULT_GROUP	         	hh: 新增
                  	                	         	             	         	hh: sdfd
433171487634710528	test-1          	         	DEFAULT_GROUP	         	啦啦啦
433172056990507008	example         	         	DEFAULT_GROUP	         	useLocalCache=true
                  	                	         	             	         	lxs=lxs
434908075100168192	test-dataId     	         	DEFAULT_GROUP	         	添加一个文本内容
435447497561755648	test-dev        	         	DEFAULT_GROUP	         	添加一个文本内容
435448380848623616	test-pro        	         	DEFAULT_GROUP	         	添加一个文本内容
457789458298060800	test123124      	         	DEFAULT_GROUP	         	aefef
461396094492889088	propert         	         	DEFAULT_GROUP	         	#
                  	                	         	             	         	# Copyright 1999-2018 Alibaba Group Holding Ltd.
                  	                	         	             	         	#
                  	                	         	             	         	# Licensed under the Apache License, Version 2....
```

### 更新配置

执行 nacosctl update config -h 查看提示

```shell
Update configuration information

Usage:
  nacosctl update config [dataId] [content] [flags]

Examples:
nacosctl update config [dataId] [content]

Flags:
  -g, --group string         config group (default "DEFAULT_GROUP")
  -h, --help                 help for config
  -n, --namespaceId string   namespace id
  -t, --type string          config type (default "text")

Global Flags:
  -a, --address string   nacos server ip+port (default "127.0.0.1:8848")
```

执行示例：

```shell
☁  ~  nacosctl update config cipher-aes-test5 更新的配置内容
done
```

### 删除配置

执行 nacosctl delete config -h 查看提示

```shell
Delete a configuration

Usage:
  nacosctl delete config [dataId] [flags]

Examples:
nacosctl delete config [dataId]

Flags:
  -g, --group string         config group (default "DEFAULT_GROUP")
  -h, --help                 help for config
  -n, --namespaceId string   namespace id

Global Flags:
  -a, --address string   nacos server ip+port (default "127.0.0.1:8848")
```

执行示例：

```shell
☁  ~  nacosctl delete config test123124
done
```

## namespace 命令

### 创建命名空间

执行 nacosctl create namespace -h 查看提示

```shell
Create a namespace, If namespaceId is not specified, it will be automatically generated

Usage:
  nacosctl create namespace [namespaceName] [namespaceDesc] [flags]

Examples:
nacosctl create namespace [namespaceName] [namespaceDesc] [flags]

Flags:
  -h, --help                 help for namespace
  -n, --namespaceId string   namespace id

Global Flags:
  -a, --address string   nacos server ip+port (default "127.0.0.1:8848")
```

执行示例：

```shell
☁  ~  nacosctl create namespace prod 线上
done
```

### 查询命名空间

执行 nacosctl get namespace -h 查看提示

```shell
Get the specified namespace information

Usage:
  nacosctl get namespace [namespaceId] [flags]

Examples:
nacosctl get namespace [namespaceId]

Flags:
  -h, --help   help for namespace

Global Flags:
  -a, --address string   nacos server ip+port (default "127.0.0.1:8848")
```

执行示例：

```shell
☁  ~  nacosctl get namespace 31ef4a30-475c-47cc-bf6e-b20b11049b59
命名空间ID                          	命名空间名称	配置额度	使用数量
31ef4a30-475c-47cc-bf6e-b20b11049b59	prod        	200     	2
```

### 查询命名空间列表

执行 nacosctl get namespaces -h 查看提示

```shell
Returns all namespace information

Usage:
  nacosctl get namespaces [flags]

Examples:
nacosctl get namespaces

Flags:
  -h, --help   help for namespaces

Global Flags:
  -a, --address string   nacos server ip+port (default "127.0.0.1:8848")
```

执行示例：

```shell
☁  ~  nacosctl get namespaces
命名空间ID                          	 命名空间名称	      配置额度	使用数量
                                    	public      	200     	11
0afdddf0-89f7-42d5-97fc-0f5cf20f1e39	test        	200     	0
225b8e29-bc5c-4354-86ff-9a3939e5c4e2	线上         	200     	1
2e65b1de-8a16-40f1-a174-f294987daf0d	prod        	200     	0
31ef4a30-475c-47cc-bf6e-b20b11049b59	prod        	200     	2
5783e427-138e-46c9-b39d-16c4d5223b8f	test1       	200     	0
```

### 更新命名空间

执行 nacosctl update namespace -h 查看命令提示

```shell
Update the name and description of the specified namespace

Usage:
  nacosctl update namespace [namespaceName] [namespaceDesc] [namespaceId]  [flags]

Examples:
nacosctl update namespace [namespaceName] [namespaceDesc] [namespaceId]

Flags:
  -h, --help   help for namespace

Global Flags:
  -a, --address string   nacos server ip+port (default "127.0.0.1:8848")
```

执行示例：

```shell
☁  ~  nacosctl update namespace prod-1 aaaaaa 2e65b1de-8a16-40f1-a174-f294987daf0d
done
```

### 删除命名空间

执行 nacosctl delete namespace -h 查看命令提示

```shell
delete the specified namespace

Usage:
  nacosctl delete namespace [namespaceId] [flags]

Examples:
nacosctl delete namespace [namespaceId]

Flags:
  -h, --help   help for namespace

Global Flags:
  -a, --address string   nacos server ip+port (default "127.0.0.1:8848")
```

执行示例：

```shell
☁  ~  nacosctl delete namespace 7b89541e-a46b-4bc1-93d7-1a0e4123a220
done
```