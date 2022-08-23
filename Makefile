proto:
	del /S *.pb.go
	protoc pkg/auth/pb/auth.proto --go_out=plugins=grpc:.
	protoc pkg/course/pb/course.proto --go_out=plugins=grpc:.
	protoc pkg/timetable/pb/timetable.proto --go_out=plugins=grpc:.

start:
	go run cmd/main.go
