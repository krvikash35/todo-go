build:
	mkdir -p out
	go build -o ./out/todo

gen-grpc-go:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative server/grpc/todo.proto

run-http-server: build
	./out/todo http

run-grpc-server: gen-grpc-go build
	./out/todo grpc