FROM golang:1.20-alpine

WORKDIR /app

RUN go mod download
# sqlboilerのダウンロード
RUN go get -u -t github.com/volatiletech/sqlboiler/v4
RUN go get -u -t github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql


COPY . .
RUN go mod tidy

# RUN go build -o app .

# CMD ["./app"]