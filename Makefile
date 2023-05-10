protogen:
	rm -rf ./pb/*.go && \
	protoc ./proto/*.proto --go_out=. --go-grpc_out=.
