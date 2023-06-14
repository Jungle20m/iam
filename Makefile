

migration-up:
	migrate -source file://migrations -database "mysql://anhnv:123456@tcp(ec2-54-179-88-203.ap-southeast-1.compute.amazonaws.com:3306)/healthnet?charset=utf8mb4&parseTime=True&loc=Local" up

migration-down:
	migrate -source file://migrations -database "mysql://anhnv:123456@tcp(ec2-54-179-88-203.ap-southeast-1.compute.amazonaws.com:3306)/healthnet?charset=utf8mb4&parseTime=True&loc=Local" down


grpc-gen:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative .\internal\modules\auth\grpc-transport\protoc\auth.proto

api:
	go run main.go api