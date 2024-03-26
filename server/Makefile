GO = go
GOBUILD = $(GO) build
GORUN = $(GO) run
GOTEST=$(GO) test
BINARY_NAME = app
APP_PATH = ./cmd/server/main.go

LOG_PATH = ./logs/

.DEFAULT_GOAL := all

all: build run clean clear 
cl: build_client run_client clear 

clear: 
	rm -f logs/*.log *.exe $(BINARY_NAME) ./client

build:
	$(GOBUILD) -o $(BINARY_NAME) $(APP_PATH)

run:
	./$(BINARY_NAME) --log-lvl="debug" --log-file="${LOG_PATH}"

build_client:
	$(GOBUILD) -o client ./cmd/client/client.go 

run_client:
	./client