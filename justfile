default:
	just --list

build:
	@go build -o monkey .

test:
	@go test -v ./...