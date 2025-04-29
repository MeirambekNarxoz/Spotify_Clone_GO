FROM golang:1.23-alpine

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o main ./cmd/server

EXPOSE 8081

CMD ["./main"]
