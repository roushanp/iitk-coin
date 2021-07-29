FROM golang:latest

LABEL maintainer="Roushan <roushanprakash123@gmail.com>"

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN go build

CMD ["./iitk-coin"]