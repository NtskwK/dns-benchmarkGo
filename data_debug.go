//go:build !release

package main

import (
	"os"

	"github.com/oschwald/geoip2-golang"
)

// geoDataPath 是GeoLite2国家数据库的路径
// 这个常量定义了GeoLite2国家数据库文件的相对路径
// 该数据库用于IP地理位置查询，可以确定IP地址所属的国家
const geoDataPath = "./res/Country.mmdb"

// GetGeoData 打开并返回GeoIP2数据库读取器
// 这个函数尝试打开GeoLite2国家数据库文件，并返回一个可以用于查询的读取器
// 返回值:
//   - *geoip2.Reader: 成功时返回GeoIP2数据库的读取器
//   - error: 如果打开文件失败，返回相应的错误
func GetGeoData() (*geoip2.Reader, error) {
	return geoip2.Open(geoDataPath)
}

// sampleServersDataPath 是示例DNS服务器列表的路径
// 这个常量定义了包含DNS服务器列表的文本文件的相对路径
// 该文件包含了一系列DNS服务器的地址，用于DNS性能测试
const sampleServersDataPath = "./res/providers.txt"

// GetSampleServersData 读取并返回示例DNS服务器列表的内容
// 这个函数读取包含DNS服务器列表的文本文件，并返回其内容
// 返回值:
//   - []byte: 成功时返回文件的完整内容
//   - error: 如果读取文件失败，返回相应的错误
func GetSampleServersData() ([]byte, error) {
	return os.ReadFile(sampleServersDataPath)
}

// domainsDataPath 是域名数据的路径
// 这个常量定义了包含域名数据的文本文件的相对路径
// 该文件包含了一系列域名，用于DNS性能测试
const domainsDataPath = "./res/domains.txt"

// GetDomainsData 读取并返回域名数据的内容
// 这个函数读取包含域名数据的文本文件，并返回其内容
func GetDomainsData() ([]byte, error) {
	return os.ReadFile(domainsDataPath)
}

// templateHTMLPath 是HTML模板文件的路径
// 这个常量定义了包含HTML模板文件的相对路径
// 该模板用于生成测试结果的HTML报告
const templateHTMLPath = "./res/template.html"

// GetTemplateHTML 读取并返回HTML模板文件的内容
// 这个函数读取包含HTML模板文件的内容，并返回其内容
// 返回值:
//   - []byte: 成功时返回文件的完整内容
//   - error: 如果读取文件失败，返回相应的错误
func GetTemplateHTML() ([]byte, error) {
	return os.ReadFile(templateHTMLPath)
}
