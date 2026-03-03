FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
ENV CGO_ENABLED=0 GOOS=linux
RUN go build -o weatherapp main.go

FROM alpine:latest
RUN addgroup -S wgroup  && adduser -S wuser -G wgroup
WORKDIR /home/wuser
COPY --from=builder /app/weatherapp .
USER wuser
EXPOSE 8080
CMD ["./weatherapp"]

