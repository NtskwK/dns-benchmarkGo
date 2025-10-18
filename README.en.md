# dnspy - A Testing Tool for Locally Testing the Accessibility and Performance of DNS Servers Worldwide

[English](./README.en.md) | [中文](./README.md)

Frustrated with domestic DNS being hijacked by operators, needing reliable services to support normal internet access.

There aren't many existing tools, and dnsjumper has issues like only supporting Windows, fewer data sources, and fewer evaluation dimensions.

So I made a tool to test the servers that can be used normally on the local network and their performance, written in Golang supporting Windows, macOS, Linux.

It also comes with a visualization analysis website that lets you see at a glance which DNS servers you can use 😊, friendly reminder: click on the bar chart in the data analysis panel to copy the server address.

Usage: Follow the instructions below to download the testing tool, get the test result json file, open the analysis panel website, upload the data for analysis. The website does not store data.

## Data Analysis Dashboard Preview

![Data Analysis Dashboard Preview](https://github.com/user-attachments/assets/c743f7ba-4d77-4d16-8515-02c0dc99ddfa)

[Data Analysis Dashboard with Sample Data](https://bench.dash.2020818.xyz)

## Usage

![dnspy](https://github.com/user-attachments/assets/a499d2fc-ffcd-4b71-a0dd-d6e5839792dd)

Download the `dns-benchmarkGo-*` file according to your system architecture from the [releases](https://github.com/NtskwK/dns-benchmarkGo/releases) page in this repository. For example, for macOS with M series processors, download the `dnspy-darwin-arm64` file.

Then **you must disable all proxy software's Tun mode and virtual network card mode, otherwise it will affect the test results.**
Then **you must disable all proxy software's Tun mode and virtual network card mode, otherwise it will affect the test results.**
Then **you must disable all proxy software's Tun mode and virtual network card mode, otherwise it will affect the test results.**

Rename the file to `dnspy` (Windows is `dnspy.exe`), then open a terminal, navigate to the directory where this file is located. Execute the command to start testing

```bash
unset http_proxy https_proxy all_proxy HTTP_PROXY HTTPS_PROXY ALL_PROXY
./dnspy
```

Follow the prompts to start the test.

By default, multi-threading mode is used to speed up testing. However, the default parameter of 10 threads requires at least 1 MB/s upstream and downstream network and at least 4-core processor.
If the network or processor is not good, it will lead to inaccurate test results, you must reduce the number of threads through the `-w` parameter.

After the test is completed, it will be output to a JSON file in the current directory with a name like `dnspy_result_2024-11-07-17-32-13.json`.

Enter `Y` or `y` or just press Enter as prompted by the program, and the data analysis dashboard website will automatically open. Click the `Read Analysis` button in the top right corner of the website, select your JSON file, and you can see the visualized test results.

### Running in Source Code Mode

#### 1. Clone repository and initialize submodules

  ```bash
  git clone --recurse-submodules https://github.com/NtskwK/dns-benchmarkGo.git
  cd dns-benchmark/src
  ```

  If you have already cloned the repository, use the following command to initialize submodules:

  ```bash
  git submodule update --init --recursive
  ```

#### 2. Run

  ```bash
  go run .
  ```

## Available Parameters

```batch
~> dnspy-windows-amd64.exe -h  

Usage examples:

dnspy

Start testing directly using built-in worldwide domains

dnspy -s 114.114.114.114

Test a single server

dnspy dnspy_benchmark_2024-10-22-08-18.json

Visualize analysis of test results

Parameter description:
  -c, --concurrency int   Number of concurrent tests per test
                           (default 10)
  -d, --domains string    File path to store domain data for batch testing
                          Must be a file path relative to the current program working directory
                          Internal file format is one per line
                          If not modified, use the built-in 10000 popular domains
                           (default "@sampleDomains@")
  -t, --duration int      Duration of each test, in seconds
                           (default 10)
  -f, --file string       File path to store server data for batch testing
                          Must be a file path relative to the current program working directory
                          Internal file format is one per line

  -g, --geo string        Standalone function: Use GeoIP database for IP or domain geolocation query

      --json              Output logs in json format

  -l, --level string      Log level
                          Options: debug,info,warn,error,fatal,panic
                           (default "info")
      --no-aaaa           Do not resolve AAAA records for each test

      --old-html          Deprecated, not recommended
                          It is recommended to use as <Example 1> the program directly parses and outputs the data json file and follows the prompts to view the visualized data analysis
                          If you need to view the visualized data analysis next time, you can use the program to open the json file as <Example 3>
                          This parameter uses the old method to output a single HTML file to the same directory as the data json
                          Can be double-clicked to open and view

  -o, --output string     Output result file path
                          Must be a file path relative to the current program working directory
                          If not specified, output to dnspy_result_<current time>.json under the current working directory

      --prefer-ipv4       Prefer IPv4 addresses when converting DNS server domain names to IP addresses
                           (default true)
  -s, --server strings    Manually specify servers to test, supports multiple

  -w, --worker int        How many DNS servers to test at the same time
                           (default 20)
```

## Compilation

Required environment for compilation:

- You need `Go` environment and `curl` command on your computer, preferably `make` command, otherwise you may need to manually execute the contents of the makefile

- Ability to access Github to download resource files

- If the following problem occurs on Windows, please use `gitbash` to execute the command content.

> 'GOOS' is not recognized as an internal or external command,
operable program or batch file.

### Compilation Process

#### 1. Clone repository and initialize submodules

  ```bash
  # Clone this repository (including submodules)
  git clone --recurse-submodules https://github.com/NtskwK/dns-benchmarkGo.git
  cd dns-benchmark/src
  ```

  If you have already cloned the repository, use the following command to initialize submodules:

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

#### 3. Compile

  ```bash
  make build
  ```
