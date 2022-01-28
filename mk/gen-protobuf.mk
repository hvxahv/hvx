protocGoOut := protoc --go_out=. --go_opt=paths=source_relative
protocGoGRPCOut := --go-grpc_out=. --go-grpc_opt=paths=source_relative

gen proto:
	$(protocGoOut) $(protocGoGRPCOut) ./api/account/v1alpha1/*.proto && \
$(protocGoOut) $(protocGoGRPCOut) ./api/device/v1alpha1/*.proto \
#$(protocGoOut) $(protocGoGRPCOut) ./api/articles/v1alpha1/*.proto && \
#$(protocGoOut) $(protocGoGRPCOut) ./api/channel/v1alpha1/*.proto && \
#$(protocGoOut) $(protocGoGRPCOut) ./api/activity/v1alpha1/*.proto
