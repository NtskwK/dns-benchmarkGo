package main

import (
	"testing"
)

func TestCheckGeo(t *testing.T) {
	// 初始化GeoIP数据库
	geoDB, err := InitGeoDB()
	if err != nil {
		t.Fatalf("初始化GeoIP数据库失败: %v", err)
	}
	defer geoDB.Close()

	serversOK := []string{
		"1.1.1.1:53",
		"114.114.114.114",

		"119.29.29.29",
		"2402:4e00::",
		"https://dns.google/dns-query",
		"tls://dns.cloudflare.com",
		"quic://dns.google:853",
		"1.1.1.1:5353",
		// 特殊情况
		"https://freedns.controld.com/p3",
		"https://dns.bebasid.com/unfiltered",
		"2620:119:53::53",
		"https://doh.cleanbrowsing.org/doh/family-filter/",
		// 无理取闹
		"https://1:1:1:1:1:1",
	}
	serversErr := []string{
		"192.168.1.1",
		"https://dns.goooooogle/dns-query",
		"",
	}

	// 测试正确的服务器地址
	for _, server := range serversOK {
		t.Run(server, func(t *testing.T) {
			ip, geoCode, err := CheckGeo(geoDB, server, true)
			if err != nil {
				t.Errorf("CheckGeo(%s) 失败: %v", server, err)
			} else {
				t.Logf("CheckGeo(%s) 成功: IP=%s, GeoCode=%s", server, ip, geoCode)
			}
		})
	}

	// 测试错误的服务器地址
	for _, server := range serversErr {
		t.Run(server, func(t *testing.T) {
			ip, geoCode, err := CheckGeo(geoDB, server, true)
			if err == nil {
				t.Errorf("CheckGeo(%s) 应该失败，但成功了: IP=%s, GeoCode=%s", server, ip, geoCode)
			} else {
				t.Logf("CheckGeo(%s) 预期失败: %v", server, err)
			}
		})
	}
}
