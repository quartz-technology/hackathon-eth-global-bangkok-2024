# syntax=docker/dockerfile:1

FROM golang:1.19

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /euler-rebalancing-api

EXPOSE 8080

# Run
CMD ["/euler-rebalancing-api"]
