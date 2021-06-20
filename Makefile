run:
	go run server.go

test:
	go test ./...

proto:
	protoc --go_out=plugins=grpc:. $(file)