# calculator-on-go

## Простой веб-сервис для вычисления арифметических выражений

### Описание
Этот проект реализует веб-сервис, который вычисляет арифметические выражения, переданные пользователем через HTTP-запрос. Веб-сервис поддерживает базовые арифметические операции, такие как сложение, вычитание, умножение и деление.

### Структура проекта
- **cmd/** — точка входа приложения.
- **pkg/calculation/errors.go** — специальные ситуативные случаи ошибок.
- **internal/** — внутренняя логика и модули приложения.
- **pkg/** — вспомогательные пакеты и утилиты.

### Запуск сервиса

#### Установите Go
Перед тем как начать, убедитесь, что у вас установлен Go. Для этого выполните команду:

```bash
go version
Если Go не установлен, скачайте и установите его с официального сайта: https://golang.org/dl/

Склонируйте проект с GitHub
bash
Копировать код
git clone https://github.com/MiklilkiM/calculator-on-go.git
cd calculator-on-go
Запуск сервера
Перейдите в папку проекта и запустите сервер:

bash
Копировать код
go run ./cmd/main.go
Сервис будет доступен по адресу: http://localhost:8080/api/v1/calculate.

Альтернативный запуск
Вы можете использовать скрипты для сборки и запуска:

Для Linux/MacOS:

bash
Копировать код
./build/build.sh
Для Windows:

bash
Копировать код
.\build\build.bat
Эндпоинты
POST /api/v1/calculate
Описание: Этот эндпоинт принимает JSON с математическим выражением и возвращает результат вычислений.

Пример запроса с использованием PowerShell:
powershell
Копировать код
Invoke-RestMethod -Uri "http://localhost:8080/api/v1/calculate" `
-Method POST `
-Headers @{"Content-Type"="application/json"} `
-Body '{"expression": "2+2*2"}'
Пример успешного ответа:
json
Копировать код
{
  "result": "6.000000"
}
Пример ошибки 422:
Если выражение содержит некорректный символ (например, $), сервер вернет ошибку 422:

Пример запроса:

powershell
Копировать код
Invoke-RestMethod -Uri "http://localhost:8080/api/v1/calculate" `
-Method POST `
-Headers @{"Content-Type"="application/json"} `
-Body '{"expression": "2+2$"}'
Пример ответа:

json
Копировать код
{
  "error": "Invalid character in expression"
}
Тестирование
Для запуска тестов выполните команду:

bash
Копировать код
go test ./...
Примеры cURL-запросов:
Успешный запрос:

bash
Копировать код
curl -X POST http://localhost:8080/api/v1/calculate/ \
-H "Content-Type: application/json" \
-d '{"expression": "1 + 2"}'
Запрос с некорректным выражением:

bash
Копировать код
curl -X POST http://localhost:8080/api/v1/calculate/ \
-H "Content-Type: application/json" \
-d '{"expression": "invalid"}'
Запрос с пустым телом:

bash
Копировать код
curl -X POST http://localhost:8080/api/v1/calculate/ \
-H "Content-Type: application/json" \
-d ''
Запрос с ошибкой синтаксиса:

bash
Копировать код
curl -X POST http://localhost:8080/api/v1/calculate/ \
-H "Content-Type: application/json" \
-d '{invalid}'