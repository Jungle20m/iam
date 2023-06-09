

migration-up:
	migrate -source file://migrations -database "mysql://anhnv:123456@tcp(localhost:3308)/healthnet?charset=utf8mb4&parseTime=True&loc=Local" up

migration-down:
	migrate -source file://migrations -database "mysql://anhnv:123456@tcp(localhost:3308)/healthnet?charset=utf8mb4&parseTime=True&loc=Local" down

api:
	go run main.go