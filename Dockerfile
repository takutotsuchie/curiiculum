FROM golang:1.20-alpine

WORKDIR /app

RUN go mod download
# sqlboilerのダウンロード
RUN go install github.com/volatiletech/sqlboiler/v4@latest
RUN go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest


COPY . .
RUN go mod tidy

# RUN go build -o app .

# CMD ["./app"]