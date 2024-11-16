# syntax=docker/dockerfile:1

FROM golang:1.23.2

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY main.go ./
COPY pkg ./pkg

RUN CGO_ENABLED=0 GOOS=linux go build -o /euler-rebalancing-api

EXPOSE 8080

# Run
CMD ["/euler-rebalancing-api"]
