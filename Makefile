BINARY_NAME=filecloud

build:
	go build -o bin/${BINARY_NAME} cmd/app/main.go

run:build
	./bin/${BINARY_NAME}

clean:
	go clean
	rm bin/${BINARY_NAME}