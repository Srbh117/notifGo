run:
	@go run .

build:
	@go build -buildvcs=false -o ./bin/


testp:
	@go run cmd/testpublish/main.go