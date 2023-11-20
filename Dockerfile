FROM golang:1.21

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o go_redis .

CMD ["./go_redis"]
