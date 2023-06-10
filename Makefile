BINARY_NAME=server

build:
	go build -o bin/${BINARY_NAME} main.go

run:
	go run main.go

clean:
	go clean
	rm bin/${BINARY_NAME}

test:
	go test -v ./...

setup: build
	go install github.com/swaggo/swag/cmd/swag@latest

api-docs:
	swag init