PROJNAME={{ .ProjName }}
build:
	go build  -o ${PROJNAME} cmd/${PROJNAME}/main.go
run:
	go run cmd/${PROJNAME}/main.go
test:
	go test -v ./...
format:
	go fmt ./...
install: build
	install -m 755 ${PROJNAME} ~/.local/bin/${PROJNAME}

.PHONY: build test format install
