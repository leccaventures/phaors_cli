APP_NAME := pharoscli
CMD_DIR := ./cmd/pharoscli
BIN_DIR := ./bin
BINARY := $(BIN_DIR)/$(APP_NAME)
DIST_DIR := ./dist
FIXTURE_SCRIPT := ./scripts/export-fixtures.sh

.PHONY: build run test fmt clean help install release fixtures

build:
	mkdir -p $(BIN_DIR)
	go build -o $(BINARY) $(CMD_DIR)

install:
	go install $(CMD_DIR)

release:
	mkdir -p $(DIST_DIR)
	GOOS=darwin GOARCH=amd64 go build -o $(DIST_DIR)/$(APP_NAME)-darwin-amd64 $(CMD_DIR)
	GOOS=darwin GOARCH=arm64 go build -o $(DIST_DIR)/$(APP_NAME)-darwin-arm64 $(CMD_DIR)
	GOOS=linux GOARCH=amd64 go build -o $(DIST_DIR)/$(APP_NAME)-linux-amd64 $(CMD_DIR)
	GOOS=windows GOARCH=amd64 go build -o $(DIST_DIR)/$(APP_NAME)-windows-amd64.exe $(CMD_DIR)

run:
	go run $(CMD_DIR) $(ARGS)

test:
	go test ./...

fixtures: build
	bash $(FIXTURE_SCRIPT) $(BINARY)

fmt:
	gofmt -w ./cmd ./internal

clean:
	rm -rf $(BIN_DIR) $(DIST_DIR)

help:
	@printf "Targets:\n"
	@printf "  make build    Build %s into %s\n" "$(APP_NAME)" "$(BINARY)"
	@printf "  make install  Install %s with go install\n" "$(APP_NAME)"
	@printf "  make release  Cross-build binaries into %s\n" "$(DIST_DIR)"
	@printf "  make run      Run the CLI (pass ARGS='...')\n"
	@printf "  make test     Run Go tests\n"
	@printf "  make fixtures Export JSON fixtures into ./test\n"
	@printf "  make fmt      Format Go sources\n"
	@printf "  make clean    Remove build artifacts\n"
