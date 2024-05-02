BINARY_NAME="newsapp"

build:
	go build -o cmd/web main.go

run:
	go run cmd/web/main.go

win:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ${BINARY_NAME}

help:
    @echo "make - clean, format, build"
    @echo "make build - just build"
    @echo "make run - just run"
    @echo "make clean - clean binary file"
    @echo "make gotool - run fmt and vet tools"
    @echo "make mac - build for mac"
    @echo "make win - build for windows"