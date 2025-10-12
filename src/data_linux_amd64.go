package main

import _ "embed"

//go:embed res/dnspyre-linux-amd64
var dnspyreBinData []byte

func GetDnspyreBin() ([]byte, string) {
	return dnspyreBinData, "dnspyre-linux-amd64"
}
