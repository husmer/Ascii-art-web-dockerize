# syntax=docker/dockerfile:1
# A microservice in Go packaged into a container image.
FROM golang:1.21.1

WORKDIR /koodJohvi/ascii-art-web-dockerize

COPY go.mod ./

RUN go mod download

COPY . ./

RUN go build -o main

EXPOSE 8080

# Run
CMD ["./main"]