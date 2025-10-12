package main

import (
	"math"
)

type scoreResult struct {
	Total       float64 `json:"total"`
	SuccessRate float64 `json:"successRate"`
	ErrorRate   float64 `json:"errorRate"`
	Latency     float64 `json:"latency"`
	Qps         float64 `json:"qps"`
}

// 权重常量：用于不同评分项的权重
const (
	SuccessRateScoreWeight = 35
	ErrorRateScoreWeight   = 10
	LatencyScoreWeight     = 50
	QpsScoreWeight         = 5
)

// 分数计算的常量阈值
const (
	LatencyRangeMax      = 1000 // 超过 *ms 以上的平均延迟得 0 分
	LatencyRangeMin      = 0.1  // 小于 *ms 的平均延迟得 0 分
	LatencyFullMarkPoint = 50   // 小于 *ms 的平均延迟满分
	MaxQps               = 100  // * QPS 为满分
)

// 定义错误值
var ErrNoRequests = scoreResult{}

// ScoreBenchmarkResult 计算 DNS 服务器的评分
func ScoreBenchmarkResult(r jsonResult) scoreResult {
	// 检查成功响应数是否为 0
	if r.TotalSuccessResponses == 0 {
		return ErrNoRequests
	}

	// 计算成功率：成功响应次数占总请求次数的比例
	successRate := float64(r.TotalSuccessResponses) / float64(r.TotalRequests)
	// 计算成功率评分：线性映射
	successRateScore := successRate * 100

	// 计算错误率：错误响应和 IO 错误占总请求次数的比例
	errorRate := float64(r.TotalErrorResponses+r.TotalIOErrors) / float64(r.TotalRequests)
	// 错误率评分计算：线性映射
	errorRateScore := 100 * (1 - errorRate)
	// 确保最终分数在0-100之间
	errorRateScore = math.Max(0, math.Min(100, errorRateScore))

	// 计算延迟评分：综合平均延迟和标准差，考虑延迟的稳定性
	var latencyScore float64
	// 综合平均值和中位数
	meanMS := float64((r.LatencyStats.MeanMs + r.LatencyStats.P50Ms) / 2)

	if meanMS < LatencyRangeMin || meanMS > LatencyRangeMax {
		// 无效的平均延迟，得分为0
		latencyScore = 0
	} else {
		// 如果平均延迟在满分阈值和 0.1ms 之间，线性计算分数
		baseScore := 100 - (meanMS-LatencyFullMarkPoint)*100/(LatencyRangeMax-LatencyFullMarkPoint)
		// 考虑标准差，引入较轻的惩罚因子，使得延迟波动大的情况得分稍低
		stabilityFactor := 1 / (1 + 0.5*math.Pow(float64(r.LatencyStats.StdMs)/meanMS, 2))
		latencyScore = baseScore * (0.8 + 0.2*stabilityFactor)
	}
	// 确保最终分数在0-100之间
	latencyScore = math.Max(0, math.Min(100, latencyScore))

	// 如果 p95 延迟也非常高，进一步降低分数（处理极端延迟的情况）
	if r.LatencyStats.P95Ms > LatencyRangeMax {
		latencyScore *= 0.85 // 延迟不稳定，进一步扣分
	}

	// QPS 评分：使用对数函数映射，考虑最大 QPS
	qpsScore := 100 * math.Log(1+r.QueriesPerSecond) / math.Log(1+MaxQps)
	// 确保分数不超过 100
	qpsScore = math.Min(100, qpsScore)

	// 综合总分：根据各项评分的权重计算总分
	totalScore := (successRateScore*SuccessRateScoreWeight +
		errorRateScore*ErrorRateScoreWeight +
		latencyScore*LatencyScoreWeight +
		qpsScore*QpsScoreWeight) / 100

	// 返回评分结果
	return scoreResult{
		Total:       Round(totalScore, 2),
		SuccessRate: Round(successRateScore, 2),
		ErrorRate:   Round(errorRateScore, 2),
		Latency:     Round(latencyScore, 2),
		Qps:         Round(qpsScore, 2),
	}
}
