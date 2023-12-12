FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux

RUN apk update --no-cache

WORKDIR /build

COPY . .

RUN go mod download
RUN go build -ldflags="-s -w" -o /app/wash-payment cmd/main.go

FROM alpine

RUN apk update --no-cache

WORKDIR /app

COPY environment/firebase /app/firebase
COPY migrations migrations
COPY --from=builder /app/wash-payment wash-payment

EXPOSE 8080
CMD ["/app/wash-payment"]
