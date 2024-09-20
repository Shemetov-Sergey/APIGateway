proto:
	protoc --go_out=. ./pkg/**/pb/*.proto

proto_grpc:
	protoc --go_out=. --go-grpc_out=. pkg/**/pb/*.proto

server:
	go run cmd/main.go