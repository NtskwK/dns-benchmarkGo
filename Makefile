all: run

build-frontend:
	@echo building frontend...
	cd web && pnpm install &&pnpm build

build: build-frontend
	@echo building dnsbenchmark...
	go mod tidy
	GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -tags release -o release/dns-benchmarkGo-darwin-amd64 .
	GOOS=darwin GOARCH=arm64 go build -ldflags "-s -w" -tags release -o release/dns-benchmarkGo-darwin-arm64 .
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -tags release -o release/dns-benchmarkGo-linux-amd64 .
	GOOS=linux GOARCH=arm64 go build -ldflags "-s -w" -tags release -o release/dns-benchmarkGo-linux-arm64 .
	GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -tags release -o release/dns-benchmarkGo-windows-amd64.exe .
	GOOS=windows GOARCH=arm64 go build -ldflags "-s -w" -tags release -o release/dns-benchmarkGo-windows-arm64.exe .

run:
	go run .

update-geodata:
	@echo updating geolocation data...
	curl https://cdn.jsdelivr.net/gh/Loyalsoldier/geoip@release/Country.mmdb -o ./res/Country.mmdb

update-domains:
	@echo updating domain data...
	curl https://cdn.jsdelivr.net/gh/Tantalor93/dnspyre@master/data/10000-domains -o ./res/domains.txt
	cd res && python sort.py

update-dnspyre:
	@echo updating dnspyre submodule...
	git submodule update --remote dnspyre

update: update-geodata update-domains update-dnspyre

.PHONY: all build run clean update-geodata update-domains update-dnspyre update
