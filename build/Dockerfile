FROM golang:1.19-alpine

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . .

RUN go build -o user-api cmd/main.go

EXPOSE 8080

CMD ["./user-api"]
