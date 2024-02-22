default:
	just --list

build:
	@go build -o monkey .

test:
	@go test -v ./...

fmt:
	@go fmt ./...

@run *arg:
    go run main.go $@