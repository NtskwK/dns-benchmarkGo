package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/oschwald/geoip2-golang"
)

// AloneFunc_GenSampleServersIPCode 独立函数完成单一功能，与程序无关
//
// AloneFunc_GenSampleServersIPCode(GeoDB, "./res/providers.txt", "./res/providers.dat")
//
// 转化纯文本的 DNS 服务器列表为 服务器,IP,代码 的数据。txt -> dat
//
// 但是 dat 不是很适合使用，因为只是根据我的位置和服务商生成的 DNS 解析结果
//
// 可能会随时间改变
//
// 输入：GeoIP数据库、原始文件路径、输出文件路径
func AloneFunc_GenSampleServersIPCode(geoDB *geoip2.Reader, rawFilePath, outFilePath string) {
	// 打开输入文件
	inputFile, err := os.Open(rawFilePath)
	if err != nil {
		log.Fatalf("无法打开输入文件: %v", err)
	}
	defer inputFile.Close()

	// 创建输出文件
	outputFile, err := os.Create(outFilePath)
	if err != nil {
		log.Fatalf("无法创建输出文件: %v", err)
	}
	defer outputFile.Close()

	// 创建一个扫描器来读取输入文件
	scanner := bufio.NewScanner(inputFile)
	// 创建一个写入器来写入输出文件
	writer := bufio.NewWriter(outputFile)

	// 逐行读取输入文件并写入输出文件
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") {
			continue
		}
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		ip, code, err := CheckGeo(geoDB, line, true)
		_, err = writer.WriteString(line + "," + ip + "," + code + ";")
		if err != nil {
			log.Fatalf("写入输出文件时发生错误: %v", err)
		}
	}

	// 检查扫描过程中是否有错误
	if err := scanner.Err(); err != nil {
		log.Fatalf("读取输入文件时发生错误: %v", err)
	}

	// 刷新写入器，确保所有数据都写入文件
	writer.Flush()
}
