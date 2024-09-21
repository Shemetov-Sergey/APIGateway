FROM golang:alpine as builder

WORKDIR /APIGateway

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/

FROM alpine

WORKDIR /APIGateway

COPY --from=builder /APIGateway/main /APIGateway/main
COPY --from=builder /APIGateway/pkg/config/envs/*.env /APIGateway/

RUN chmod +x /APIGateway/main

CMD ["./main"]