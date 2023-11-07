FROM golang:1.21.1

WORKDIR /app

COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o tweedledo-backend cmd/main.go

EXPOSE 8080

CMD ["./tweedledo-backend"]