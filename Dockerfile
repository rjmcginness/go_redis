FROM golang:1.21 as builder

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o go_redis .

FROM scratch
COPY --from=builder /app/go_redis .
CMD ["/go_redis"]