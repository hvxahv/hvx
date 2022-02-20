SHELL=bash

protocOut := protoc --go_out=. --go_opt=paths=source_relative
protocGRPCOut := --go-grpc_out=. --go-grpc_opt=paths=source_relative

NAMES := account notify saved message activity channel article

proto:
	$(foreach n,$(NAMES), $(protocOut) $(protocGRPCOut) ./api/$(n)/v1alpha1/*.proto && @echo $(n) success! ${\n} &)


