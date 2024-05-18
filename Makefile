build:
	@go build -o cmd/main

run: build
	@./cmd/main