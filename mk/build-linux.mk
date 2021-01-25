gbol := $$Env:GOOS = "linux"; $$Env:GOARCH = "amd64" ; go build -o ./build/linux

build-linux:
	rm ./build/linux/*
	$(gbol)	./app/accounts
	$(gbol) ./app/gateway
	$(gbol) ./app/status
	$(gbol) ./app/bot