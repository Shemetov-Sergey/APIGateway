FROM golang:alpine as builder

WORKDIR /APIGateway

ENV HTTP_PROXY $HTTP_PROXY
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./cmd/

RUN #go build -o main ./cmd/

FROM alpine

WORKDIR /APIGateway

COPY --from=builder /APIGateway/main /APIGateway/main

RUN chmod +x /APIGateway/main

CMD ["./main"]