FROM golang:1.21-alpine as BUILD

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/bin/linux_amd64/main


FROM alpine:latest
# 复制资源
COPY --from=BUILD /app/config /config
COPY --from=BUILD /app/mds /mds
COPY --from=BUILD /app/bin/linux_amd64/main app/main

# Env设置 - App
ENV APP_USE_ENV="true"
ENV APP_PORT=8080
ENV APP_DEBUG="false"
ENV APP_FRONTEND_URL="http://www.example.com"
ENV APP_BACKEND_URL="http://api.example.com"

# 暴露端口
EXPOSE 8080

ENTRYPOINT ["app/main"]