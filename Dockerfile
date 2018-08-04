FROM golang:1.10.1-alpine3.7 as builder
COPY server.go .
# RUN go build -o /server server.go
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /server .

FROM scratch
COPY --from=builder /server /
CMD ["/server"]
