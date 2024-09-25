FROM golang:alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /auth-service ./cmd/main.go
# RUN go build -o /auth-service ./cmd/main.go

FROM alpine:latest AS runtime

LABEL maintainer="SeyedHossein HosseiniMotlagh <hossein_225@yahoo.com>"

WORKDIR /app

COPY --from=builder /auth-service .

EXPOSE 50054

ENTRYPOINT ["./auth-service"]