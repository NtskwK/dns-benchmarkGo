package main

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/oschwald/geoip2-golang"
	log "github.com/sirupsen/logrus"
	"github.com/tantalor93/dnspyre/v3/pkg/dnsbench"
)

// 具体工作实现
func runDnspyre(geoDB *geoip2.Reader, preferIPv4 bool, noAAAA bool, server, domainsPath string, duration, concurrency int, probability float64) jsonResult {

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

	// 配置 dnspyre benchmark
	queryTypes := []string{"A"}
	if !noAAAA {
		queryTypes = append(queryTypes, "AAAA")
	}

	// 使用 @ 前缀引用域名文件
	queries := []string{"@" + domainsPath}

	bench := &dnsbench.Benchmark{
		Server:         server,
		Types:          queryTypes,
		Duration:       time.Duration(duration) * time.Second,
		Concurrency:    uint32(concurrency),
		Probability:    probability,
		Queries:        queries,
		Count:          0, // 使用 Duration 模式
		Silent:         true,
		JSON:           true,
		Recurse:        true,
		Edns0:          1232,
		RequestTimeout: 5 * time.Second,
		ConnectTimeout: 1 * time.Second,
		ReadTimeout:    3 * time.Second,
		WriteTimeout:   1 * time.Second,
	}

	log.WithFields(log.Fields{
		"目标": server,
	}).Infof("\x1b[32m%s 开始测试\x1b[0m", server)

	// 运行 benchmark
	ctx := context.Background()
	benchStart := time.Now()
	stats, err := bench.Run(ctx)
	benchDuration := time.Since(benchStart)

	if err != nil {
		log.WithFields(log.Fields{
			"目标": server,
			"错误": err,
		}).Errorf("\x1b[31m%s 测试失败\x1b[0m", server)
		return jsonResult{}
	}

	// 转换结果为 JSON 格式
	result := convertDnspyreResult(stats, benchDuration)

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

// convertDnspyreResult 将 dnspyre 的 ResultStats 转换为 jsonResult
func convertDnspyreResult(stats []*dnsbench.ResultStats, benchDuration time.Duration) jsonResult {
	// 合并所有 stats
	var totalCounters dnsbench.Counters
	qtypeTotals := make(map[string]int64)

	// 合并计数器
	for _, stat := range stats {
		if stat.Counters != nil {
			totalCounters.Total += stat.Counters.Total
			totalCounters.IOError += stat.Counters.IOError
			totalCounters.Success += stat.Counters.Success
			totalCounters.Negative += stat.Counters.Negative
			totalCounters.Error += stat.Counters.Error
			totalCounters.IDmismatch += stat.Counters.IDmismatch
			totalCounters.Truncated += stat.Counters.Truncated
		}

		// 合并 query types
		for k, v := range stat.Qtypes {
			qtypeTotals[k] += v
		}
	}

	// 合并 histograms
	var mergedHist *dnsbench.ResultStats
	if len(stats) > 0 {
		mergedHist = stats[0]
		for i := 1; i < len(stats); i++ {
			if stats[i].Hist != nil && mergedHist.Hist != nil {
				mergedHist.Hist.Merge(stats[i].Hist)
			}
		}
	}

	// 构建结果
	result := jsonResult{
		TotalRequests:           totalCounters.Total,
		TotalSuccessResponses:   totalCounters.Success,
		TotalNegativeResponses:  totalCounters.Negative,
		TotalErrorResponses:     totalCounters.Error,
		TotalIOErrors:           totalCounters.IOError,
		TotalIDmismatch:         totalCounters.IDmismatch,
		TotalTruncatedResponses: totalCounters.Truncated,
		QuestionTypes:           qtypeTotals,
		QueriesPerSecond:        math.Round(float64(totalCounters.Total)/benchDuration.Seconds()*100) / 100,
		BenchmarkDurationSeconds: roundDuration(benchDuration).Seconds(),
	}

	// 添加延迟统计
	if mergedHist != nil && mergedHist.Hist != nil {
		hist := mergedHist.Hist
		result.LatencyStats = latencyStats{
			MinMs:  roundDuration(time.Duration(hist.Min())).Milliseconds(),
			MeanMs: roundDuration(time.Duration(hist.Mean())).Milliseconds(),
			StdMs:  roundDuration(time.Duration(hist.StdDev())).Milliseconds(),
			MaxMs:  roundDuration(time.Duration(hist.Max())).Milliseconds(),
			P99Ms:  roundDuration(time.Duration(hist.ValueAtQuantile(99))).Milliseconds(),
			P95Ms:  roundDuration(time.Duration(hist.ValueAtQuantile(95))).Milliseconds(),
			P90Ms:  roundDuration(time.Duration(hist.ValueAtQuantile(90))).Milliseconds(),
			P75Ms:  roundDuration(time.Duration(hist.ValueAtQuantile(75))).Milliseconds(),
			P50Ms:  roundDuration(time.Duration(hist.ValueAtQuantile(50))).Milliseconds(),
		}
	}

	return result
}

// roundDuration 将 duration 四舍五入
func roundDuration(duration time.Duration) time.Duration {
	switch {
	case duration > time.Second:
		return duration.Round(10 * time.Millisecond)
	case duration > time.Millisecond:
		return duration.Round(time.Millisecond / 10)
	case duration > time.Microsecond:
		return duration.Round(time.Microsecond / 10)
	default:
		return duration
	}
}
