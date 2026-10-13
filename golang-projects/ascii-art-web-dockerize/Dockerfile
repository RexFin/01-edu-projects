FROM golang:1.21-alpine AS builder

WORKDIR /build
COPY . .

RUN go build -o server ./cmd/app/main.go

FROM debian:bookworm-slim

LABEL maintainer="#nkishiba #assrakhimova #agabduli"
LABEL description="Web application for generating ASCII art"
LABEL version="1.0.0"

WORKDIR /app

COPY --from=builder /build/server .
COPY --from=builder /build/static ./static
COPY --from=builder /build/templates ./templates
COPY --from=builder /build/banners ./banners

EXPOSE 8080

CMD ["./server"]