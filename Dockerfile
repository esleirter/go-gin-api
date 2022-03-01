FROM golang:1.17.7-alpine as builder
WORKDIR /app
COPY . .
RUN     apk add git --no-cache && \
        go env -w CGO_ENABLED=0 GOOS=linux GOARCH=amd64 && go mod tidy && \
        go build -a -installsuffix cgo \
        -ldflags '-w -extldflags "-static"' -o /go/bin/main main.go

FROM    alpine:3.15.0
WORKDIR /app
COPY    --from=builder /go/bin/main /app/
EXPOSE  3000
CMD     ["./main"]