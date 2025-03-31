FROM golang:1.24.1 AS builder

WORKDIR /app

COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

COPY . .

RUN go build -o server .

FROM golang:1.24.1 AS runner

WORKDIR /app

COPY --from=builder /app/server /app/server  

EXPOSE 8080

CMD ["./server"]
