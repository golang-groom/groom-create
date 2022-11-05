PROJNAME={{ .ProjName }}
build:
	go build  -o ${PROJNAME} cmd/${PROJNAME}/main.go
run:
	go run cmd/${PROJNAME}/main.go
test:
	go test -v ./...
format:
	go fmt ./...
.PHONY: build test format
