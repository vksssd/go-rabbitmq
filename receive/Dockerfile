FROM golang:1.16 AS builder

WORKDIR /app


COPY . . 
RUN go mod tidy

RUN go build -o receive .
RUN chmod +x receive


RUN CGO_ENABLED=0 GOOS=linux go build -o receive .

FROM alpine:3.12
RUN apk --no-cache add ca-certificates libc6-compat

WORKDIR  /app

COPY --from=builder /app/receive .

EXPOSE 8080

RUN chmod +x receive

CMD ["go","run","receive.go"]
