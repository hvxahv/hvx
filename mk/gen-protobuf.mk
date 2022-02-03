protocGoOut := protoc --go_out=. --go_opt=paths=source_relative
protocGoGRPCOut := --go-grpc_out=. --go-grpc_opt=paths=source_relative

NAMES := account device notify saved message

gen proto:
	$(foreach n,$(NAMES), $(protocGoOut) $(protocGoGRPCOut) ./api/$(n)/v1alpha1/*.proto && @echo $(n) success! ${\n} &)


