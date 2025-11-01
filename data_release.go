//go:build release

package main

import (
	_ "embed"

	"github.com/oschwald/geoip2-golang"
)

//go:embed res/Country.mmdb
var GeoData []byte

// GetGeoData 函数返回一个可用于IP地理位置查询的GeoIP2数据库读取器
// 这个函数使用嵌入的GeoData字节切片创建一个geoip2.Reader实例
// 返回值:
//   - *geoip2.Reader: 成功时返回GeoIP2数据库的读取器
//   - error: 如果创建读取器失败，返回相应的错误
func GetGeoData() (*geoip2.Reader, error) {
	return geoip2.FromBytes(GeoData)
}

//go:embed res/providers.txt
var SampleServersData []byte

// GetSampleServersData 函数返回嵌入的DNS服务器列表数据
// 这个函数直接返回SampleServersData字节切片，无需从文件系统读取
// 返回值:
//   - []byte: 返回包含DNS服务器列表的字节切片
//   - error: 始终返回nil，因为数据已经嵌入，不会发生读取错误
func GetSampleServersData() ([]byte, error) {
	return SampleServersData, nil
}

//go:embed res/domains.txt
var DomainsData []byte

// GetDomainsData 函数返回嵌入的域名数据
// 这个函数直接返回DomainsData字节切片，无需从文件系统读取
// 返回值:
//   - []byte: 返回包含域名数据的字节切片
//   - error: 始终返回nil，因为数据已经嵌入，不会发生读取错误
func GetDomainsData() ([]byte, error) {
	return DomainsData, nil
}
