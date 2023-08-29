# Makefile

# Variáveis
APP_NAME := go_ddd_example
MAIN_FILE := cmd/main.go
GO_FLAGS := -race
SQLC_CMD = sqlc
MIGRATE_CMD = migrate
DB_URL = "postgresql://postgres:123123@localhost:5342/database?sslmode=disable"
MIGRATIONS_DIR = internal/infrastructure/migrations

# Comandos
.PHONY: all build run clean
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


.PHONY: generate
generate:
	$(SQLC_CMD) generate

.PHONY: create-migration
create-migration:
	$(MIGRATE_CMD) create -ext sql -dir $(MIGRATIONS_DIR) -seq $(name)

.PHONY: apply-migrations
apply-migrations:
	$(MIGRATE_CMD) -database $(DB_URL) -path $(MIGRATIONS_DIR) up

install-migrate:
	go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest
