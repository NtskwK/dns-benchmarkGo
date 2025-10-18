# dnspy 

一款用于本地测试全世界的 DNS 服务器的可访问性和性能的测试工具

[English](./README.en.md) | [中文](./README.md)

使用 Golang 编写支持 Windows、macOS、Linux。

并且附带可视化分析网站让你一目了然的知道可以用哪些 DNS 服务器😊。（调用系统设置的默认浏览器）

**Web dashboard is not support on GUI-less Linux**

**必须关闭所有代理软件的 Tun 模式、虚拟网卡模式，否则会影响测试结果。**

> 若出现多个服务器的测试延迟为小于5ms，请排查网络服务是否被当地运营商劫持的可能。

## 数据分析面板预览

![数据分析面板预览](https://github.com/user-attachments/assets/c743f7ba-4d77-4d16-8515-02c0dc99ddfa)

[数据分析面板，内含示例数据](https://bench.dash.2020818.xyz)

## 使用方式

![dnspy](https://github.com/user-attachments/assets/a499d2fc-ffcd-4b71-a0dd-d6e5839792dd)

在本仓库的 [releases](https://github.com/NtskwK/dns-benchmarkGo/releases) 页面中按你的系统架构进行下载后点击运行即可。

程序默认使用多线程模式，以加快测试速度。但是默认参数 10 个线程需要至少上下行 1 MB/s 网络和至少 4 核心处理器。
如果网络或处理器不好，会导致测试结果不准确，请通过`-w` 参数降低线程数。

测试完成后会输出到当前目录下形如 `dnspy_result_2024-11-07-17-32-13.json` 的 JSON 文件中。

按程序提示输入 `Y` 或 `y` 或直接回车，会自动打开数据分析面板网站，点击网站右上角的 `读取分析` 按钮，选择你刚才的 JSON 文件，就可以看到可视化测试结果了。

### 源码模式下运行

#### 1. 克隆本仓库并初始化子模块

  ```bash
  git clone --recurse-submodules https://github.com/NtskwK/dns-benchmarkGo.git
  cd dns-benchmarkGo
  ```

  如果已经克隆了仓库，可以使用以下命令初始化子模块：

  ```bash
  git submodule update --init --recursive
  ```

#### 2. 运行

  ```bash
  go run .
  ```

## 可用参数

```batch
$dns-benchmarkGo --help

参数说明:
  -c, --concurrency int   每个测试并发数
                           (default 10)
  -d, --domains string    要批量测试的域名数据存储的文件路径
                          必须是相对当前程序工作路径的文件路径
                          文件内部格式是一行一条
                          不修改则使用内置的10000个热门域名
                           (default "@sampleDomains@")
  -t, --duration int      每个测试持续时间,单位秒
                           (default 10)
  -f, --file string       要批量测试的服务器数据存储的文件路径
                          必须是相对当前程序工作路径的文件路径
                          文件内部格式是一行一条

  -g, --geo string        独立功能: 使用 GeoIP 数据库进行 IP 或域名归属地查询

      --json              以json格式输出日志

  -l, --level string      日志级别
                          可选 debug,info,warn,error,fatal,panic
                           (default "info")
      --no-aaaa           每个测试不解析 AAAA 记录

  -o, --output string     输出结果的文件路径
                          必须是相对当前程序工作路径的文件路径
                          不指定则输出到当前工作路径下的 dnspy_result_<当前时间>.json

      --prefer-ipv4       在DNS服务器的域名转换为IP地址过程中优先使用IPv4地址
                           (default true)
  -s, --server strings    手动指定要测试的服务器,支持多个

  -w, --worker int        同一时间测试多少个 DNS 服务器
                           (default 20)
```

## 编译

编译所需环境：

- 你的电脑上需要有 `Go` 环境、`curl` 命令，最好有`make`命令，不然你可能需要手动执行`makefile`中的内容

- 能够访问 Github 下载资源文件

- 若在Windows出现以下问题，请改用`gitbash`执行命令内容。

> 'GOOS' is not recognized as an internal or external command,
operable program or batch file.

### 编译过程

#### 1. 克隆本仓库并初始化子模块

  ```bash
  # 克隆本仓库（包含子模块）
  git clone --recurse-submodules https://github.com/NtskwK/dns-benchmarkGo.git
  cd dns-benchmarkGo
  ```

  如果已经克隆了仓库，可以使用以下命令初始化子模块：

  ```bash
  git submodule update --init --recursive
  cd src
  ```

#### 2. 更新数据文件（可选）

  ```bash
  # 更新所有
  make update 
  # 更新geodata
  make update-geodata 
  # 更新测试用域名
  make update-domains
  # 更新 dnspyre 子模块
  make update-dnspyre
  ```

#### 3. 进行编译

  ```bash
  make build
  ```

## Thanks

- [dns-benchmark](https://github.com/xxnuo/dns-benchmark)
- [dnspyre](https://github.com/Tantalor93/dnspyre)
