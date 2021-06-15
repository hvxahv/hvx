windows := go build -o ./build
linux := GOOS=linux GOARCH=amd64 go build -o ./build

build: windows linux

clear:
	rm build/*

windows:
	$(windows) ./app/accounts

linux:
	$(linux) ./app/accounts
