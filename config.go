package main

import (
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"
)

// Config 结构体用于存储所有的配置选项
type Config struct {
	LogJSON         bool     // 日志格式
	LogLevel        string   // 日志级别
	PreferIPv4      bool     // 在DNS服务器的域名转换为IP地址过程中优先使用IPv4地址
	ServersDataPath string   // 要测试的服务器数据路径,必须是相对当前程序工作路径的文件路径,文件内部格式是一行一条
	DomainsDataPath string   // 要批量测试的域名数据存储的文件路径,必须是相对当前程序工作路径的文件路径,文件内部格式是一行一条
	Duration        int      // 每个测试持续时间,单位秒
	Concurrency     int      // 每个测试并发数
	NoAAAARecord    bool     // 每个测试不解析 AAAA 记录
	Servers         []string // 手动指定要测试的服务器,支持多个
	Workers         int      // 同一时间测试多少个 DNS 服务器
	OutputPath      string   // 输出结果的文件路径,必须是相对当前程序工作路径的文件路径
	// 功能参数
	InputResultJsonPath string // 输入结果 json 文件路径,必须是相对当前程序工作路径的文件路径
	FnGeo               string // 使用 GeoIP 数据库进行 IP 归属地查询
	// Web 服务相关参数
	StaticDirPath   string // 本地静态文件目录路径,用于服务如图片等静态资源
	StaticURLPath   string // 静态文件服务的URL路径前缀,例如 /static/
}

func InitFlags() (Config, error) {
	cfg := Config{}
	flag.BoolVar(&cfg.LogJSON, "json", false, "\x1b[32m以json格式输出日志\x1b[0m\n")
	flag.StringVarP(&cfg.LogLevel, "level", "l", "info", "\x1b[32m日志级别\n可选 debug,info,warn,error,fatal,panic\x1b[0m\n")
	flag.BoolVar(&cfg.PreferIPv4, "prefer-ipv4", true, "\x1b[32m在DNS服务器的域名转换为IP地址过程中优先使用IPv4地址\x1b[0m\n")
	flag.StringVarP(&cfg.ServersDataPath, "file", "f", "", "\x1b[32m要批量测试的服务器数据存储的文件路径\n必须是相对当前程序工作路径的文件路径\n文件内部格式是一行一条\x1b[0m\n")
	flag.StringSliceVarP(&cfg.Servers, "server", "s", []string{}, "\x1b[32m手动指定要测试的服务器,支持多个\x1b[0m\n")
	flag.StringVarP(&cfg.DomainsDataPath, "domains", "d", "@sampleDomains@", "\x1b[32m要批量测试的域名数据存储的文件路径\n必须是相对当前程序工作路径的文件路径\n文件内部格式是一行一条\n不修改则使用内置的10000个热门域名\x1b[0m\n")
	flag.IntVarP(&cfg.Duration, "duration", "t", 10, "\x1b[32m每个测试持续时间,单位秒\x1b[0m\n")
	flag.IntVarP(&cfg.Concurrency, "concurrency", "c", 10, "\x1b[32m每个测试并发数\x1b[0m\n")
	flag.IntVarP(&cfg.Workers, "worker", "w", 20, "\x1b[32m同一时间测试多少个 DNS 服务器\x1b[0m\n")
	flag.BoolVar(&cfg.NoAAAARecord, "no-aaaa", false, "\x1b[32m每个测试不解析 AAAA 记录\x1b[0m\n")
	flag.StringVarP(&cfg.OutputPath, "output", "o", "", "\x1b[32m输出结果的文件路径\n必须是相对当前程序工作路径的文件路径\n不指定则输出到当前工作路径下的 dnspy_result_<当前时间>.json\x1b[0m\n")
	flag.StringVarP(&cfg.FnGeo, "geo", "g", "", "\x1b[32m独立功能: 使用 GeoIP 数据库进行 IP 或域名归属地查询\x1b[0m\n")
	// Web 服务相关参数
	flag.StringVar(&cfg.StaticDirPath, "static-dir", "", "\x1b[32m本地静态文件目录路径\n用于服务如图片等静态资源,必须是相对当前程序工作路径的文件路径\x1b[0m\n")
	flag.StringVar(&cfg.StaticURLPath, "static-url", "/static/", "\x1b[32m静态文件服务的URL路径前缀\n例如 /static/, 访问时会拼接此路径\x1b[0m\n")

	// 使用说明
	flag.Usage = func() {
		fmt.Print(
			"dns-benchmarkGo\n\n" +
				"Homepage: https://github.com/NtskwK/dns-benchmarkGo\n\n" +
				"使用示例:\n\n" +
				"\x1b[33mdns-benchmarkGo\x1b[0m\n\n" +
				"\x1b[32m使用内置的世界所有域名直接启动测试\x1b[0m\n\n" +
				"\x1b[33mdns-benchmarkGo -s 114.114.114.114\x1b[0m\n\n" +
				"\x1b[32m测试单个服务器\x1b[0m\n\n" +
				"\x1b[33mdns-benchmarkGo dns-benchmarkGo_benchmark_2024-10-22-08-18.json\x1b[0m\n\n" +
				"\x1b[32m对测试结果进行可视化分析\x1b[0m\n\n" +
				"参数说明:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	otherFlags := flag.Args()
	exitTrigger := false

	if cfg.FnGeo != "" {
		exitTrigger = true
		geoDB, err := InitGeoDB()
		if err != nil {
			log.WithFields(log.Fields{
				"错误": err,
			}).Error("读取 GeoIP 数据库失败")
			return cfg, err
		}
		ip, country, err := CheckGeo(geoDB, cfg.FnGeo, cfg.PreferIPv4)
		if err != nil {
			log.WithFields(log.Fields{
				"错误": err,
			}).Error("查询失败")
			return cfg, err
		}
		log.WithFields(log.Fields{
			"IP":   ip,
			"Code": country,
		}).Infof("\x1b[32m查询结果:\x1b[0m")
	}

	for _, v := range otherFlags {
		if strings.HasSuffix(v, ".json") {
			cfg.InputResultJsonPath = v
		}
	}

	if exitTrigger {
		os.Exit(0)
	}

	if cfg.ServersDataPath == "" && len(cfg.Servers) == 0 {
		log.Error("你没有指定要测试的服务器数据存储的文件路径或手动输入要测试的服务器")
		log.Info("是否使用内置的世界 DNS 服务器数据开始测试(服务器很多,需要测试一段时间)? [y/N]")
		var input string
		fmt.Scanln(&input)
		if input == "Y" || input == "y" {
			cfg.ServersDataPath = "@sampleServers@"
		} else {
			// log.Error("没有有效数据,程序退出")
			return cfg, fmt.Errorf("没有有效数据")
		}
	}

	return cfg, nil
}