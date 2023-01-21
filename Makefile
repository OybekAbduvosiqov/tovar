run:
	go run cmd/main.go

swag-gen:
	swag init -g api/api.go -o api/docs

migration-up:
	migrate -path ./migrations/postgres/ -database 'postgres://oybek:oybek@localhost:5432/tovar?sslmode=disable' up

migration-down:
	migrate -path ./migrations/postgres/ -database 'postgres://oybek:oybek@localhost:5432/tovar?sslmode=disable' down
