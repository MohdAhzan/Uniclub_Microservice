

run:
	go run cmd/main.go

hello:
	echo "hellloooo"
build:
	go build -o ./cmd/uniclub_project ./cmd
buildrun:
	go build -o ./cmd/uniclub_project ./cmd && ./cmd/uniclub_project
proto:
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/user-svc.proto

