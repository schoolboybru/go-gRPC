run:
	go run main.go

test:
	go test ./...

proto:
	protoc --go_out=plugins=grpc:. $(file)