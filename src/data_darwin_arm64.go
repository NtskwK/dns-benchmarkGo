package main

import _ "embed"

//go:embed res/dnspyre-darwin-arm64
var dnspyreBinData []byte

func GetDnspyreBin() ([]byte, string) {
	return dnspyreBinData, "dnspyre-darwin-arm64"
}
