FROM golang:1.21.0-alpine3.18 as build
RUN apk update
WORKDIR /app
COPY go.mod go.sum .
COPY . .
RUN go build -o donna-md-service ./cmd/main.go

FROM alpine:3.18 as final
COPY --from=build /app/donna-md-service .
EXPOSE 50051
CMD ["./donna-md-service"]
