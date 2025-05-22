FROM golang:1.23.6-alpine3.21 AS builder

ENV CGO_ENABLED=1

WORKDIR /app
COPY . ./

RUN apk add --no-cache gcc musl-dev
RUN go mod download
RUN go build -o app

FROM alpine:3.21.3
WORKDIR /app
COPY --from=builder /app/app .
EXPOSE 8080
CMD ["./app"]
