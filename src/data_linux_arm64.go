package main

import _ "embed"

//go:embed res/dnspyre-linux-arm64
var dnspyreBinData []byte

func GetDnspyreBin() ([]byte, string) {
	return dnspyreBinData, "dnspyre-linux-arm64"
}
