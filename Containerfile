# Build the App
FROM golang:1.18.2-alpine AS build

WORKDIR /app

COPY . .

RUN go build -o /app/golist-api

# Build the Image
FROM alpine:latest

WORKDIR /app

COPY --from=build /app/golist-api /app/golist-api

COPY --from=build /app/html /app/html

EXPOSE 8080

USER 1001

CMD ["/app/golist-api"]