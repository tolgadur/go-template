# syntax=docker/dockerfile:1

FROM golang:1.18-alpine

WORKDIR /root

COPY go.mod go.sum ./
RUN go mod download && go mod verify
RUN apk --no-cache add curl

COPY . .

RUN go build -o /email-project cmd/server/main.go

EXPOSE 8080

CMD [ "/email-project" ]
