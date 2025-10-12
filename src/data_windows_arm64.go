package main

import _ "embed"

//go:embed res/dnspyre-windows-arm64.exe
var dnspyreBinData []byte

func GetDnspyreBin() ([]byte, string) {
	return dnspyreBinData, "dnspyre-windows-arm64.exe"
}
