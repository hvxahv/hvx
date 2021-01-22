protocGoOut := protoc --go_out=. --go_opt=paths=source_relative
protocGoGRPCOut := --go-grpc_out=. --go-grpc_opt=paths=source_relative

gen proto:
	$(protocGoOut) $(protocGoGRPCOut) ./api/util/v1/*.proto
	$(protocGoOut) $(protocGoGRPCOut) ./api/kernel/v1/*.proto
