GO = go
GOBUILD = $(GO) build
GORUN = $(GO) run
GOTEST=$(GO) test
BINARY_NAME = app
APP_PATH = ./cmd/server/main.go

LOG_PATH = ./logs/

.DEFAULT_GOAL := all

all: build run


clear: 
	rm -f logs/*.log *.exe $(BINARY_NAME)

build:
	$(GOBUILD) -o $(BINARY_NAME) $(APP_PATH)

run:
	./$(BINARY_NAME) --log-lvl="debug" --log-file="${LOG_PATH}"

