windows := go build -o ./build
linux := SET CGO_ENABLED=0 SET GOOS=linux SET GOARCH=arm64 go build -o ./build

build: windows linux

clear:
	rm build/*

windows:
	$(windows) ./app/accounts

linux:
	$(linux) ./app/accounts
