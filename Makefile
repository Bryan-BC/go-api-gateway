proto:
	C:\protoc-21.5-win64\bin\protoc pkg/auth/pb/auth.proto --go_out=plugins=grpc:.

start:
	go run cmd/main.go