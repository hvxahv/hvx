gbow := go build -o ./build/windows

build windows:
	rm ./build/windows/*.exe
	$(gbow) ./app/accounts
	$(gbow) ./app/gateway
	$(gbow) ./app/status
	$(gbow) ./app/bot