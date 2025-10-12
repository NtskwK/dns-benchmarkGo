package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/oschwald/geoip2-golang"
	log "github.com/sirupsen/logrus"
)

// 具体工作实现
func runDnspyre(geoDB *geoip2.Reader, preferIPv4 bool, noAAAA bool, binPath, server, domainsPath string, duration, concurrency int, probability float64) jsonResult {

	log.WithFields(log.Fields{
		"目标": server,
		"时间": fmt.Sprintf("%ds", duration),
		"并发": concurrency,
		"概率": fmt.Sprintf("%.2f", probability),
	}).Infof("\x1b[32m%s 开始测试\x1b[0m", server)
	// 先获取服务器地理信息
	ip, geoCode, err := CheckGeo(geoDB, server, preferIPv4)
	if err != nil {
		log.WithFields(log.Fields{
			"目标": server,
			"错误": err,
		}).Errorf("\x1b[31m%s 解析失败\x1b[0m", server)
		return jsonResult{}
	} else {
		log.WithFields(log.Fields{
			"目标": server,
			"IP": ip,
			"代码": geoCode,
		}).Infof("\x1b[32m%s 成功解析\x1b[0m", server)
	}

	// 运行 dnspyre
	args := []string{
		"--json",
		"--no-distribution",
		"-t", "A",
		"--duration", fmt.Sprintf("%ds", duration),
		"-c", fmt.Sprintf("%d", concurrency),
		"@" + domainsPath,
		"--probability", fmt.Sprintf("%.2f", probability),
		"--server", server,
	}
	if !noAAAA {
		args = append(args, "-t", "AAAA")
	}

	cmd := exec.Command(binPath, args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	log.WithFields(log.Fields{
		"目标": server,
	}).Infof("\x1b[32m%s 开始测试\x1b[0m", server)
	err = cmd.Run()

	if err != nil {
		log.WithFields(log.Fields{
			"目标":     server,
			"错误":     err,
			"stderr": stderr.String(),
		}).Errorf("\x1b[31m%s 测试失败\x1b[0m", server)
		return jsonResult{}
	}

	ret := stdout.Bytes()
	retLen := len(ret)
	// 检查 dnspyre 输出格式是否正确
	if retLen == 0 || ret[0] != '{' || !bytes.HasSuffix(ret, []byte("}\n")) {
		log.WithFields(log.Fields{
			"目标": server,
			"错误": "dnspyre 输出格式错误",
		}).Errorf("\x1b[31m%s 测试失败\x1b[0m", server)
		return jsonResult{}
	}

	// 转换为 JSON 格式
	var result jsonResult
	err = json.Unmarshal(ret, &result)
	if err != nil {
		log.WithFields(log.Fields{
			"目标": server,
			"错误": err,
		}).Errorf("\x1b[31m%s 测试失败\x1b[0m", server)
		return jsonResult{}
	}

	// 添加地理信息
	result.Geocode = geoCode
	result.IPAddress = ip

	// 打分
	result.Score = ScoreBenchmarkResult(result)

	if result.Score.Total == 0 {
		log.WithFields(log.Fields{
			"目标": server,
			"错误": "无法连接服务器",
		}).Errorf("\x1b[31m%s 测试失败\x1b[0m", server)
	} else {
		log.WithFields(log.Fields{
			"目标":    server,
			"总分":    fmt.Sprintf("%.2f", result.Score.Total),
			"成功率得分": fmt.Sprintf("%.2f", result.Score.SuccessRate),
			"错误率得分": fmt.Sprintf("%.2f", result.Score.ErrorRate),
			"延迟得分":  fmt.Sprintf("%.2f", result.Score.Latency),
			"QPS得分": fmt.Sprintf("%.2f", result.Score.Qps),
		}).Infof("\x1b[32m%s 测试完成\x1b[0m", server)
	}
	return result
}
