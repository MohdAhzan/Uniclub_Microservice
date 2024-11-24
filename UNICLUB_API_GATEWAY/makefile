run:
	go run cmd/main.go
proto:
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/usersvc/user-svc.proto
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/inventorysvc/inv-svc.proto

