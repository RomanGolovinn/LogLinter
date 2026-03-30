.PHONY: build-cli build-plugin test test-plugin

# Сборка независимой утилиты
build-cli:
	go build -o loglinter ./cmd/loglinter/main.go

# Запуск обычных тестов
test:
	go test -v ./...

# Сборка .so плагина
build-plugin:
	CGO_ENABLED=1 go build -buildmode=plugin -o plugin/loglinter.so plugin/loglinter.go


test-plugin: build-plugin
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run ./testdata/src/a/...