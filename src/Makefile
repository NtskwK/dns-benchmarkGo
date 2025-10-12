all: run

build-frontend:
	@echo building frontend...
	cd web
	pnpm install
	pnpm build
	cd ..

build: build-frontend
	@echo building dnsbenchmark...
	go mod tidy
	GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -tags release -o dnspy-darwin-amd64 .
	GOOS=darwin GOARCH=arm64 go build -ldflags "-s -w" -tags release -o dnspy-darwin-arm64 .
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -tags release -o dnspy-linux-amd64 .
	GOOS=linux GOARCH=arm64 go build -ldflags "-s -w" -tags release -o dnspy-linux-arm64 .
	GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -tags release -o dnspy-windows-amd64.exe .
	GOOS=windows GOARCH=arm64 go build -ldflags "-s -w" -tags release -o dnspy-windows-arm64.exe .

run:
	go run .

update-geodata:
	@echo updating geolocation data...
	curl https://cdn.jsdelivr.net/gh/Loyalsoldier/geoip@release/Country.mmdb -o ./res/Country.mmdb

update-domains:
	@echo updating domain data...
	curl https://cdn.jsdelivr.net/gh/Tantalor93/dnspyre@master/data/10000-domains -o ./res/domains.txt

configuration:
	@echo updating dnspyre executable...
	@echo please do not modify the dnspyre version number in use, otherwise dnspy may not function properly
	go get github.com/tantalor93/dnspyre/v3@latest
	GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o ./res/dnspyre-darwin-amd64 github.com/tantalor93/dnspyre/v3
	GOOS=darwin GOARCH=arm64 go build -ldflags "-s -w" -o ./res/dnspyre-darwin-arm64 github.com/tantalor93/dnspyre/v3
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ./res/dnspyre-linux-amd64 github.com/tantalor93/dnspyre/v3
	GOOS=linux GOARCH=arm64 go build -ldflags "-s -w" -o ./res/dnspyre-linux-arm64 github.com/tantalor93/dnspyre/v3
	GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o ./res/dnspyre-windows-amd64.exe github.com/tantalor93/dnspyre/v3
	GOOS=windows GOARCH=arm64 go build -ldflags "-s -w" -o ./res/dnspyre-windows-arm64.exe github.com/tantalor93/dnspyre/v3


update: update-geodata update-domains

.PHONY: all build run clean update-geodata update-caller update-domains
