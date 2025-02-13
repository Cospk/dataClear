# 数据清洗组件

## 项目简介

这是一个基于 Go 语言实现的数据清洗组件，用于从不同来源（如文件或网络）接收数据，并通过一系列处理器（Processor）对数据进行处理，最后将处理后的数据存储到后端（如文件或数据库）。

## 主要功能

数据接收：支持从文件中读取数据。

数据处理：通过多个处理器（如 Filter、Fill、Aggregator）对数据进行处理。

Filter：过滤掉不符合条件的数据。

Fill：补充数据的某些字段。

Aggregator：对数据进行聚合操作（如求和、计数、最大值、最小值等）。

数据存储：将处理后的数据存储到文件中。

热加载配置：支持配置文件的热加载，修改配置文件后无需重启程序。

流式处理：支持逐行读取大文件，避免内存溢出。

## 设计思路

> 核心是数据清洗组件，接受数据然后将数据进过一系列处理器处理然后将数据存储到后端。

项目特性要求：

- 模块化：将数据清洗组件拆分为多个模块，每个模块负责特定的功能。
- 可扩展性：支持新增处理器，无需修改现有代码。
- 热加载：支持配置文件的热加载，修改配置文件后无需重启程序。
- 流式处理：支持逐行读取大文件，避免内存溢出。



## 快速开始
1. 安装依赖
   确保已安装 Go 1.16 或更高版本。然后克隆项目并安装依赖：

```shell
git clone https://gitee.com/your-repo/data-cleaner.git
cd data-cleaner
go mod tidy
```
2. 配置文件
   在项目根目录下创建 config.yaml 文件，内容如下：

```yaml
input_file_path: "testdata/input.json"
output_file_path: "output.json"
filter:
  field: "bizid"
  condition: ">25"
fill:
  field: "os_type"
  value: "linux"
aggregator:
  field: "__value__"
  group_by_field: "zone"
  aggregations:
    sum: "sum"
    count: "count"
    min: "min"
    max: "max"
```
3. 测试数据
   在 testdata/input.json 文件中准备测试数据，例如：

```test
{"zone": "gz", "bizid": "20", "env": "prod", "__value__": 1024}
{"zone": "sz", "bizid": "26", "env": "dev", "__value__": 2048}
{"zone": "gz", "bizid": "26", "env": "prod", "__value__": 10}
{"zone": "gz", "bizid": "26", "env": "prod", "__value__": 20}
{"zone": "gz", "bizid": "26", "env": "prod", "__value__": 30}
{"zone": "sz", "bizid": "26", "env": "dev", "__value__": 15}
{"zone": "sz", "bizid": "26", "env": "dev", "__value__": 25}
```

4. 运行项目
   在项目根目录下运行以下命令来构建和运行项目：

```shell
go build -o data-cleaner ./cmd/data-cleaner
./data-cleaner
```

程序会读取 testdata/input.json 文件中的数据，处理后输出到 output.json 文件中。

---

## 项目结构
```shell
data-cleaner/
├── internal/
│   ├── config/              # 配置文件解析
│   │   ├── config.go
│   │   └── config.yaml
│   ├── processor/           # 处理器实现
│   │   ├── filter.go
│   │   ├── fill.go
│   │   ├── aggregator.go
│   │   └── processor.go     # Processor 接口定义
│   ├── pipeline/            # 数据管道
│   │   └── pipeline.go
│   ├── receiver/            # 数据接收器
│   │   └── file_receiver.go
│   ├── backend/             # 数据存储后端
│   │   └── file_backend.go
│   └── utils/               # 工具函数
│       └── utils.go
├── testdata/                # 测试数据
│   └── input.json
├── go.mod                   # Go 模块文件
├── go.sum                   # Go 模块依赖校验
├── main.go          # 主程序入口
└── README.md                # 项目说明文档
```
---

## 配置说明

**配置文件格式**

配置文件为 YAML 格式，支持以下字段：

+ input_file_path：输入文件路径。

+ output_file_path：输出文件路径。

+ filter：过滤器配置。
  + field：过滤字段。
  + condition：过滤条件。
  
+ fill：填充配置。
  + field：填充字段。
  + value：填充值。
  
+ aggregator：聚合器配置。
  + field：聚合字段。
  + group_by_field：分组字段。
  + aggregations：聚合函数（如 sum、count、min、max）。

   

**热加载**

程序支持热加载配置文件。修改 config.yaml 文件后，程序会自动重新加载配置，无需重启。

---

## 处理器说明
1. Filter 处理器
   用于过滤掉不符合条件的数据。例如，只保留 bizid 大于 25 的数据。

2. Fill 处理器
   用于补充数据的某些字段。例如，补充操作系统类型和版本。

3. Aggregator 处理器
   用于对数据进行聚合操作。例如，对 __value__ 字段进行求和、计数、最大值、最小值等操作。