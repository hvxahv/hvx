windows := go build -o ./.release
linux := SET CGO_ENABLED=0 SET GOOS=linux SET GOARCH=arm64 go build -o ./.release

build: windows linux

clear:
	rm build/*

windows:
	$(windows) ./hvx
	$(windows) ./cmd/account

linux:
	$(linux) ./cmd/hvx
	$(linux) ./cmd/account
