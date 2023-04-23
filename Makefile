build:
	mkdir -p out
	go build -o ./out/todo

run-http-server: build
	./out/todo http

run-grpc-server: build
	./out/todo grpc


