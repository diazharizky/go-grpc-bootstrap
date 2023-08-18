.PHONY: protogen run lint

clean:
	rm -rf ./pb/*.go

protogen:
	make clean && \
	protoc ./proto/*.proto --go_out=. --go-grpc_out=.

run:
	go run main.go

lint:
	golangci-lint run --timeout 3m

tidy:
	go mod tidy
