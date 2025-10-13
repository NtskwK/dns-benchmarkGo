package utils

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"strings"
)

// FormatListFile 格式化列表文件
func FormatListFile(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("读取文件失败: %w", err)
	}
	return FormatListData(&data)
}

// FormatListData 格式化列表字节
func FormatListData(data *[]byte) ([]string, error) {
	lines := make([]string, 0, 100) // 预分配容量，减少内存分配
	scanner := bufio.NewScanner(bytes.NewReader(*data))
	scanner.Buffer(make([]byte, 4096), 1048576) // 设置更大的缓冲区

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" && !strings.HasPrefix(line, "#") {
			lines = append(lines, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("扫描数据失败: %w", err)
	}

	return lines, nil
}

// Round 四舍五入
func Round(x float64, precision int) float64 {
	scale := math.Pow10(precision)
	return math.Round(x*scale) / scale
}

type LatencyStats struct {
	MinMs  int64 `json:"minMs"`
	MeanMs int64 `json:"meanMs"`
	StdMs  int64 `json:"stdMs"`
	MaxMs  int64 `json:"maxMs"`
	P99Ms  int64 `json:"p99Ms"`
	P95Ms  int64 `json:"p95Ms"`
	P90Ms  int64 `json:"p90Ms"`
	P75Ms  int64 `json:"p75Ms"`
	P50Ms  int64 `json:"p50Ms"`
}

type JsonResult struct {
	// 用到了的 dnspyre 输出 JSON 格式的字段结构体定义
	TotalRequests            int64            `json:"totalRequests"`
	TotalSuccessResponses    int64            `json:"totalSuccessResponses"`
	TotalNegativeResponses   int64            `json:"totalNegativeResponses"`
	TotalErrorResponses      int64            `json:"totalErrorResponses"`
	TotalIOErrors            int64            `json:"totalIOErrors"`
	TotalIDmismatch          int64            `json:"totalIDmismatch"` // dnspyre v3.4.1
	TotalTruncatedResponses  int64            `json:"totalTruncatedResponses"`
	ResponseRcodes           map[string]int64 `json:"responseRcodes,omitempty"`
	QuestionTypes            map[string]int64 `json:"questionTypes"`
	QueriesPerSecond         float64          `json:"queriesPerSecond"`
	BenchmarkDurationSeconds float64          `json:"benchmarkDurationSeconds"`
	LatencyStats             LatencyStats     `json:"latencyStats"`

	// add:地理信息
	IPAddress string      `json:"ip"`
	Geocode   string      `json:"geocode"`
	Score     ScoreResult `json:"score"`
}

// 自定义 BenchmarkResult 类型，用于 JSON 序列化
type BenchmarkResult map[string]JsonResult

func (b *BenchmarkResult) String() (string, error) {
	jsonData, err := json.Marshal(b)
	if err != nil {
		return "", err
	}
	// return template.JSEscapeString(string(jsonData)), nil
	return string(jsonData), nil
}
