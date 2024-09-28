FROM golang:alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /entrypoint /app/cmd/main.go

FROM alpine:latest AS runtime

LABEL maintainer="SeyedHossein HosseiniMotlagh <hossein_225@yahoo.com>"

WORKDIR /app

COPY --from=builder /entrypoint .

EXPOSE 50053

ENTRYPOINT ["/app/entrypoint"]