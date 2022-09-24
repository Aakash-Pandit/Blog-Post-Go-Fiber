FROM golang:1.18-alpine AS builder
RUN apk update
RUN apk add git
RUN apk --no-cache add ca-certificates
WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/main .

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/bin/main .

CMD ["./main"]
