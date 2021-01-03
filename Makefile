SHELL=powershell.exe
# VARIABLE
protocGoOut := protoc --go_out=. --go_opt=paths=source_relative
protocGoGRPCOut := --go-grpc_out=. --go-grpc_opt=paths=source_relative
gbow := go build -o ./build/windows

gen:
	$(protocGoOut) $(protocGoGRPCOut) ./api/util/v1/*.proto
	$(protocGoOut) $(protocGoGRPCOut) ./api/kernel/v1/*.proto

build win:
	rm ./build/windows/*.exe
	$(gbow) ./app/accounts
	$(gbow) ./app/gateway
	$(gbow) ./app/status
	$(gbow) ./app/bot

clean:
	rm ./build/windows/*.exe