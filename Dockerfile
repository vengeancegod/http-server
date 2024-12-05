FROM golang:latest as builder

LABEL maintainer = "amoCRM"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . ./

RUN GOOS=linux GOARCH=amd64 go build -o main ./cmd/app

FROM alpine:latest

# Устанавливаем mysql-client и необходимые зависимости
RUN apk --no-cache add mysql-client libc6-compat

WORKDIR /root/

# Копируем бинарник из стадии builder
COPY --from=builder /app/main /root/

# Проверяем, что файл скопирован
RUN ls -l /root/

# Делаем бинарник исполнимым
RUN chmod +x /root/main

EXPOSE 8080

# Запускаем бинарник
CMD ["/root/main"]
