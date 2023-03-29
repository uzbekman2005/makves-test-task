run:
	go run cmd/main.go
	
swag:
	swag init -g api/router.go -o api/docs