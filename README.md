# LogLinter

Custom Go linter and `golangci-lint` plugin for static analysis of log messages (`log`, `log/slog`). 
Разработано в рамках тестового задания на позицию Intern Go Developer.

## Описание

`loglinter` анализирует аргументы функций логирования и проверяет их на соответствие строгим корпоративным стандартам. Линтер умеет "распаковывать" конкатенацию строк (например, `"user: " + password`) и проверять имена переменных, переданных в логгер.

### Реализованные правила (Rules Engine):
1. **CheckLowercase**: Лог-сообщение должно начинаться со строчной буквы.
2. **CheckEnglish**: Лог-сообщение должно быть только на английском языке (используется `unicode.RangeTable`).
3. **CheckSymbols**: Запрет на использование спецсимволов (`!`, `?`, `...`) и эмодзи (через `unicode.IsSymbol`).
4. **CheckSensitive**: Предотвращение утечек чувствительных данных (password, token, api_key, secret), в том числе при конкатенации и передаче через переменные.

## Архитектура

Проект построен с упором на расширяемость и чистый код:
- **AST Parsing (`analyzer.go`)**: Использует рекурсивный обход бинарных выражений (`*ast.BinaryExpr`) через Type Switch, что позволяет анализировать сложные конструкции вроде `log.Info("data: " + api_key)`.
- **Rules Engine (`internal/rules`)**: Бизнес-логика проверок вынесена в отдельный пакет. Добавление нового правила сводится к написанию одной функции с сигнатурой `func(string) string` и добавлению её в массив `activeRules`.

## Как запустить (Standalone CLI)

Проект использует стандартный фреймворк `golang.org/x/tools/go/analysis`.

```bash
# Сборка CLI-утилиты
go build -o loglinter ./cmd/loglinter/main.go

# Запуск проверки для конкретного файла или директории
./loglinter ./...

## Интеграция с golangci-lint (Module Plugin System)

Проект поддерживает современную интеграцию через [Module Plugin System](https://golangci-lint.run/plugins/module-plugins/), появившуюся в `golangci-lint` v1.57+. Это позволяет компилировать линтер напрямую в бинарник `golangci-lint`, избегая проблем с CGO.

**Конфигурация:**
В корне проекта находится файл `.custom-gcl.yml`, который описывает правила импорта нашего плагина. Также обновлен `.golangci.yml` (тип плагина изменен на `module`).

**Инструкция по сборке кастомного линтера:**

1. Запустите сборку `golangci-lint` с интегрированным плагином (рекомендуется использовать ту же версию, что и в `.custom-gcl.yml`, например `v1.61.0`):
   ```bash
   go run [github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0](https://github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0) custom
