# Используем официальный образ Go
FROM golang:1.23-alpine

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы go.mod и go.sum
COPY go.mod go.sum ./

# Скачиваем зависимости
RUN go mod download

# Копируем исходный код в контейнер
COPY . .

# Собираем приложение
RUN go build -o laratest_go .

# Указываем команду для запуска приложения
CMD ["./laratest_go"]