

migration-up:
	migrate -source file://migrations -database "mysql://anhnv:123456@tcp(ec2-54-179-88-203.ap-southeast-1.compute.amazonaws.com:3306)/healthnet?charset=utf8mb4&parseTime=True&loc=Local" up

migration-down:
	migrate -source file://migrations -database "mysql://anhnv:123456@tcp(ec2-54-179-88-203.ap-southeast-1.compute.amazonaws.com:3306)/healthnet?charset=utf8mb4&parseTime=True&loc=Local" down

api:
	go run main.go api