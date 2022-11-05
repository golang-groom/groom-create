build:
	go build -o groom-create cmd/groom-create/main.go
install: build 
	go install ./cmd/gurm-create
test: build
	go test ./... -cover
	go vet ./...
format: 
	go fmt ./...
default: build 
.PHONY: build test format
