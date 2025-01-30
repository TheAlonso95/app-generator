# Variables
APP_NAME=generator
BUILD_DIR=bin
CMD_DIR=cmd
MAIN_FILE=$(CMD_DIR)/main.go

# Commands
.PHONY: all build run test lint clean

all: build

build: $(MAIN_FILE)
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_FILE)

run: build
	@echo "Running $(APP_NAME)..."
	@$(BUILD_DIR)/$(APP_NAME)

test:
	@echo "Running tests..."
	@go test ./...

lint:
	@echo "Linting code..."
	@golangci-lint run

clean:
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)
