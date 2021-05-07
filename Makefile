run:
	go run main.go

test:
	go test ./...

proto:
	protoc --go_out=paths=source_relative:. $(file)