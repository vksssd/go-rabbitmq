FROM golang:1.16 AS builder

WORKDIR /app

COPY . .
RUN go mod tidy 
RUN go build -o send main.go
RUN chmod +x send


FROM alpine:3.12

RUN apk --no-cache add ca-certificates libc6-compat

WORKDIR /app
COPY --from=builder /app/send .

EXPOSE 8080

RUN chmod +x send


CMD ["./send"]

