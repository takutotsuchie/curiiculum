FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download


COPY . .
RUN go get -u github.com/go-sql-driver/mysql
RUN go mod tidy

RUN go build -o app .

CMD ["./app"]