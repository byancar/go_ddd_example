# Makefile

# Variáveis
APP_NAME := go_ddd_example
MAIN_FILE := cmd/main.go
GO_FLAGS := -race

# Comandos
all: clean build run

build:
	@echo "Building $(APP_NAME)..."
	@go build $(GO_FLAGS) -o $(APP_NAME) $(MAIN_FILE)

run:
	@echo "Running $(APP_NAME)..."
	@./$(APP_NAME)

clean:
	@echo "Cleaning..."
	@rm -f $(APP_NAME)

.PHONY: all build run clean