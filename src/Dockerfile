FROM golang:1-alpine

WORKDIR /src/app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o main ./cmd

CMD ["./main"]