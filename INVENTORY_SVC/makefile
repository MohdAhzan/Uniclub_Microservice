

run:
	go run cmd/main.go

build:
	go build -o ./cmd/uniclub_project ./cmd
buildrun:
	go build -o ./cmd/uniclub_project ./cmd && ./cmd/uniclub_project
proto:
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/inv-svc.proto

