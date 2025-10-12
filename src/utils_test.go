package main

import (
	"os"
	"testing"
)

func TestFormatListFile(t *testing.T) {
	// 创建临时测试文件
	tempFile, err := os.CreateTemp("", "test_list_*.txt")
	if err != nil {
		t.Fatalf("创建临时文件失败: %v", err)
	}
	defer os.Remove(tempFile.Name())

	// 写入测试数据
	testData := []byte("line1\n  line2  \n# comment\n\nline3")
	if _, err := tempFile.Write(testData); err != nil {
		t.Fatalf("写入测试数据失败: %v", err)
	}
	tempFile.Close()

	// 测试 FormatListFile 函数
	result, err := FormatListFile(tempFile.Name())
	if err != nil {
		t.Fatalf("FormatListFile 失败: %v", err)
	}

	expected := []string{"line1", "line2", "line3"}
	if len(result) != len(expected) {
		t.Fatalf("预期结果长度 %d, 实际结果长度 %d", len(expected), len(result))
	}

	for i, v := range expected {
		if result[i] != v {
			t.Errorf("第 %d 行不匹配, 预期 %q, 实际 %q", i+1, v, result[i])
		}
	}
}

func TestFormatListData(t *testing.T) {
	testData := []byte("line1\n  line2  \n# comment\n\nline3")
	result, err := FormatListData(&testData)
	if err != nil {
		t.Fatalf("FormatListData 失败: %v", err)
	}

	expected := []string{"line1", "line2", "line3"}
	if len(result) != len(expected) {
		t.Fatalf("预期结果长度 %d, 实际结果长度 %d", len(expected), len(result))
	}

	for i, v := range expected {
		if result[i] != v {
			t.Errorf("第 %d 行不匹配, 预期 %q, 实际 %q", i+1, v, result[i])
		}
	}
}
