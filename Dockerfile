FROM golang:1.20.3-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY websockets.go ./websockets.go
COPY pkg/ ./pkg/

RUN CGO_ENABLED=0 GOOS=linux go build websockets.go

EXPOSE 8081
EXPOSE 6379

CMD ["./websockets"]