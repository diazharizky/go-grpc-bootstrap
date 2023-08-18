.PHONY: protogen run lint

protogen:
	rm -rf ./pb/*.go && \
	protoc ./proto/*.proto --go_out=. --go-grpc_out=.

run:
	go run main.go

lint:
	golangci-lint run --timeout 3m

tidy:
	go mod tidy
