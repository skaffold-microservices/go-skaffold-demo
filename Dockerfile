FROM golang:1.10.1-alpine3.7 as builder
COPY server.go .
RUN go build -o /server server.go

FROM alpine:3.7  
COPY --from=builder /server .
CMD ["./server"]
