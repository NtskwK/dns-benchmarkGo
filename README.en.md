# dns-benchmarkGo

A testing tool for evaluating the accessibility and performance of DNS servers worldwide on your local machine.

[English](./README.en.md) | [中文](./README.md)

Written in Golang with support for Windows, macOS, and Linux.

Includes a visual analysis website that gives you a clear overview of which DNS servers you can use 😊. (Opens in your system's default browser)

**Web dashboard is not supported on GUI-less Linux**

**You must disable TUN mode and virtual network card mode in all proxy software, otherwise it will affect the test results.**

> If multiple servers show test latency less than 5ms, please check if network services are being hijacked by your local ISP.

## Data Analysis Dashboard Preview

![Data Analysis Dashboard Preview](https://github.com/user-attachments/assets/c743f7ba-4d77-4d16-8515-02c0dc99ddfa)

[Data Analysis Dashboard with Sample Data](https://bench.dash.2020818.xyz)

## Usage

![dnspy](https://github.com/user-attachments/assets/a499d2fc-ffcd-4b71-a0dd-d6e5839792dd)

Download the appropriate version for your system architecture from the [releases](https://github.com/NtskwK/dns-benchmarkGo/releases) page and run it.

The program uses multi-threading mode by default to accelerate testing. However, the default parameter of 10 threads requires at least 1 MB/s upload/download speed and at least a 4-core processor.
If your network or processor performance is insufficient, it may lead to inaccurate test results. Please use the `-w` parameter to reduce the number of threads.

After testing is complete, results will be output to a JSON file in the current directory with a name like `dnspy_result_2024-11-07-17-32-13.json`.

Following the program prompts, enter `Y` or `y` or press Enter directly to automatically open the data analysis dashboard website. Click the `Load Analysis` button in the top-right corner of the website and select your JSON file to view the visualized test results.

### Running from Source Code

#### 1. Clone this repository and initialize submodules

```bash
git clone --recurse-submodules https://github.com/NtskwK/dns-benchmarkGo.git
cd dns-benchmarkGo
```

If you've already cloned the repository, you can initialize submodules with:

```bash
git submodule update --init --recursive
```

#### 2. Run

```bash
go run .
```

## Available Parameters

```batch
$dns-benchmarkGo --help

Parameter Description:
  -c, --concurrency int   Concurrency for each test
                           (default 10)
  -d, --domains string    File path for storing domain data to batch test
                          Must be a file path relative to the current program working directory
                          File format: one entry per line
                          If not modified, uses built-in 10,000 popular domains
                           (default "@sampleDomains@")
  -t, --duration int      Duration for each test in seconds
                           (default 10)
  -f, --file string       File path for storing server data to batch test
                          Must be a file path relative to the current program working directory
                          File format: one entry per line

  -g, --geo string        Independent feature: Use GeoIP database for IP or domain geolocation query

      --json              Output logs in JSON format

  -l, --level string      Log level
                          Options: debug,info,warn,error,fatal,panic
                           (default "info")
      --no-aaaa           Do not resolve AAAA records for each test

  -o, --output string     Output file path for results
                          Must be a file path relative to the current program working directory
                          If not specified, outputs to dnspy_result_<current_time>.json in current working directory

      --prefer-ipv4       Prefer IPv4 addresses when converting DNS server domain names to IP addresses
                           (default true)
  -s, --server strings    Manually specify servers to test, supports multiple

  -w, --worker int        How many DNS servers to test simultaneously
                           (default 20)
```

## Building

Build Requirements:

- Your computer needs a `Go` environment, `curl` command, preferably `make` command, otherwise you may need to manually execute the contents of the `makefile`

- Ability to access Github to download resource files

- If you encounter the following issue on Windows, please use `gitbash` to execute the commands.

> 'GOOS' is not recognized as an internal or external command,
operable program or batch file.

### Build Process

#### 1. Clone this repository and initialize submodules

```bash
# Clone this repository (including submodules)
git clone --recurse-submodules https://github.com/NtskwK/dns-benchmarkGo.git
cd dns-benchmarkGo
```

If you've already cloned the repository, you can initialize submodules with:

```bash
git submodule update --init --recursive
cd src
```

#### 2. Update data files (optional)

```bash
# Update all
make update 
# Update geodata
make update-geodata 
# Update test domains
make update-domains
# Update dnspyre submodule
make update-dnspyre
```

#### 3. Build

```bash
make build
```

## Thanks

- [dns-benchmark](https://github.com/xxnuo/dns-benchmark)
- [dnspyre](https://github.com/Tantalor93/dnspyre)