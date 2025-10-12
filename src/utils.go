package main

import (
	"bufio"
	"bytes"
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
