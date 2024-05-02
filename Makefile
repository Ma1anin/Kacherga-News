include .env

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

BINARY_NAME=newsapp
BINARY_UNIX=$(BINARY_NAME)_unix

all: help

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

run:
	build
	./$(BINARY_NAME)

test: 
	$(GOTEST) -v ./test/test.go

clean:
	$(GOCLEAN)
	rm -f @(BINARY_NAME)
	rm -f @(BINARY_UNIX)

.PHONY: help
help:
    @echo "  >  make build - build project"
    @echo "  >  make run - build and run project"
    @echo "  >  make clean - clean binary file"
