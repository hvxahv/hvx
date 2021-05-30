# build to windows
b2w := go build -o ./build/windows
# build output windows
b2l := $$Env:GOOS = "linux"; $$Env:GOARCH = "amd64" ; go build -o ./build/linux

# build to windows command
bw:
	rm ./build/windows/*.exe
	$(b2w) ./app/accounts
	$(b2w) ./app/gateway
	$(b2w) ./app/status
	$(b2w) ./app/bot
	$(b2w)./app/hvx

# build to Linux command
bl:
	rm ./build/linux/*
	$(b2l)	./app/accounts
	$(b2l) ./app/gateway
	$(b2l) ./app/status
	$(b2l) ./app/bot
	$(b2l)./app/hvx
