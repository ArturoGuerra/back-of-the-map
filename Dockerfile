FROM golang:alpine AS builder 

WORKDIR /build
COPY . .
RUN apk add --update make
RUN make build

FROM alpine:latest
WORKDIR /app
COPY --from=builder /build/rolewatcher /app

CMD ["./rolewatcher"]
