GO ?= go

CORE_PACKAGES := ./internal/... ./pkg/...
ENTRY_PACKAGES := ./examples ./homework1 ./homework2 ./homework3 ./homework4 ./homework5 ./homework6

.PHONY: all build test examples deps lint fmt coverage clean help

all: build test ## Сборка и тесты

build: ## Сборка каркаса и программ запуска
	$(GO) build $(CORE_PACKAGES) ./final_tests/...
	$(GO) build $(ENTRY_PACKAGES)

test: ## Запуск тестов
	$(GO) test -v ./...

examples: ## Запуск примера работы с Z3
	$(GO) run ./examples/basic_z3_example.go

deps: ## Загрузка зависимостей Go
	$(GO) mod download

lint: ## Проверка каркаса и программ запуска через go vet
	$(GO) vet $(CORE_PACKAGES) $(ENTRY_PACKAGES)

fmt: ## Форматирование Go-кода
	$(GO) fmt ./...

coverage: ## Отчёт о покрытии тестами
	$(GO) test -coverprofile=coverage.out ./...
	$(GO) tool cover -html=coverage.out -o coverage.html

clean: ## Очистка результатов сборки и покрытия
	$(GO) clean ./...
	$(RM) coverage.out coverage.html

help: ## Список команд
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z0-9_-]+:.*?## / {printf "  %-12s %s\n", $$1, $$2}' $(MAKEFILE_LIST)
