protogen:
	rm -rf ./pb/*.go && \
	protoc ./proto/*.proto --go_out=. --go-grpc_out=.

run:
	go run main.go

mock:
	cd $(dir) && mockery --name="$(name)" && cd -

test:
	./coverage.sh $(coverage)

migrate-up:
	migrate -database postgres://gogrpcbootstrap:gogrpcbootstrap@localhost:5432/gogrpcbootstrap?sslmode=disable -path ./migrations -verbose up

migrate-down:
	migrate -database postgres://gogrpcbootstrap:gogrpcbootstrap@localhost:5432/gogrpcbootstrap?sslmode=disable -path ./migrations -verbose down
