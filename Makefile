run: 
			@go mod tidy
			@go run main.go

build: 
			@go build ./...
			@echo "built trangonel!"