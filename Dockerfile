FROM golang:latest

LABEL maintainer="Roushan <roushanprakash123@gmail.com>"

LABEL version="1.0"

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN go build

CMD ["./iitk-coin"]