SHELL := /bin/sh

GO ?= go
BIN_DIR := bin
EXEEXT :=

ifeq ($(OS),Windows_NT)
EXEEXT := .exe
endif

CLI_BIN := $(BIN_DIR)/mux-cli$(EXEEXT)
SSH_BIN := $(BIN_DIR)/mux-ssh$(EXEEXT)

.PHONY: test build tidy docker-build docker-run

test:
	$(GO) test ./...

build: $(CLI_BIN) $(SSH_BIN)

tidy:
	$(GO) mod tidy

docker-build:
	docker build -t mux-ssh .

docker-run:
	docker run --rm -p 22:22 -p 80:80 mux-ssh

$(BIN_DIR):
	mkdir -p $(BIN_DIR)

$(CLI_BIN): | $(BIN_DIR)
	$(GO) build -o $(CLI_BIN) ./cmd/cli

$(SSH_BIN): | $(BIN_DIR)
	$(GO) build -o $(SSH_BIN) ./cmd/ssh
