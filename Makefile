build:
	GOFLAGS=-mod=mod go build -o bin/snoop main.go
run:
	GOFLAGS=-mod=mod go run main.go
