# Переменные
APP_NAME = go-refresh
BUILD_DIR = bin
MAIN_FILE = ./cmd/go-refresh/main.go

# Значения по умолчанию
IN ?= samples/sample1.in.txt
OUT ?= samples/sample1.out.txt
LANG_FLAG ?= --lang

# Цель по умолчанию (срабатывает при вводе `make`)
all: clean tidy build

# Сборка проекта
build:
	@echo "Собираем проект..."
	go build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_FILE) $(IN) $(OUT) $(LANG_FLAG)

# Запуск проекта без компиляции бинарника
run:
	go run $(MAIN_FILE) $(IN) $(OUT) $(LANG_FLAG)

# Запуск тестов
test:
	@echo "Запуск тестов..."
	go test -v ./...

# Очистка кэша и удаление бинарника
clean:
	@echo "Очистка директории сборки..."
	rm -rf $(BUILD_DIR)
	go clean

# Загрузка и обновление зависимостей
tidy:
	go mod tidy
