FROM golang:1.24.1 AS build

WORKDIR /app

COPY . .  

RUN go build -o server .

FROM golang:1.24.1

COPY --from=build /app /app  

WORKDIR /app

EXPOSE 8080

CMD ["./server"]