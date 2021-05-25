gbow := go build -o ./build/windows
gbol := $$Env:GOOS = "linux"; $$Env:GOARCH = "amd64" ; go build -o ./build/linux

hvx_out := go build -o ./build/windows

build-windows:
	rm ./build/windows/*.exe
	$(gbow) ./app/accounts
	$(gbow) ./app/gateway
	$(gbow) ./app/status
	$(gbow) ./app/bot

build-linux:
	rm ./build/linux/*
	$(gbol)	./app/accounts
	$(gbol) ./app/gateway
	$(gbol) ./app/status
	$(gbol) ./app/bot

hvx-windows:
	rm ./build/windows/*.exe
	$(hvx_out) ./app/accounts