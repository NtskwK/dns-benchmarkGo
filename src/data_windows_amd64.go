package main

import _ "embed"

//go:embed res/dnspyre-windows-amd64.exe
var dnspyreBinData []byte

func GetDnspyreBin() ([]byte, string) {
	return dnspyreBinData, "dnspyre-windows-amd64.exe"
}
