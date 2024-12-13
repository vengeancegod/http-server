FROM golang:latest as builder

LABEL maintainer = "amoCRM"

WORKDIR /app

# Копируем go.mod и go.sum
COPY go.mod go.sum ./

# Устанавливаем зависимости
RUN go mod tidy

# Копируем весь код
COPY . ./

# Собираем основной бинарник
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o app ./cmd/app/main.go

# Собираем бинарник для воркеров
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o workers ./cmd/workers/workers.go

# Этап финального образа
FROM alpine:latest

WORKDIR /root/

# Копируем бинарники из стадии сборки
COPY --from=builder /app/app /root/app
COPY --from=builder /app/workers /root/workers

# Проверяем, что файлы скопированы
RUN ls -l /root/

# Делаем бинарники исполнимыми
RUN chmod +x /root/app /root/workers

# Открываем порты
EXPOSE 8080 8081 

# Запускаем основной бинарник по умолчанию
CMD ["/root/app"]