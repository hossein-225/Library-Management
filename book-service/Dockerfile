FROM golang:alpine AS builder

RUN apk add git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# RUN CGO_ENABLED=0 GOOS=linux go build -o /book-service ./cmd/main.go
RUN go build -o /book-service ./cmd/main.go

FROM alpine:latest AS runtime

LABEL maintainer="SeyedHossein HosseiniMotlagh <hossein_225@yahoo.com>"

WORKDIR /app

COPY --from=builder /book-service .

EXPOSE 50051

ENTRYPOINT ["./entrypoint"]