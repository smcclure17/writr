FROM golang:1.20.3-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY app.go ./app.go
COPY pkg/ ./pkg/

RUN CGO_ENABLED=0 GOOS=linux go build app.go

EXPOSE 8081
EXPOSE 6379

CMD ["./app"]