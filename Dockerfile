FROM golang:1.24-alpine AS builder

RUN apk add --no-cache git upx

WORKDIR /src

RUN git clone https://github.com/mattn/xpost.git .

RUN go mod download
RUN go mod tidy

RUN CGO_ENABLED=0 GOARCH=amd64 go build -ldflags "-s -w" -o xpost .

RUN upx xpost
FROM alpine:3.20

RUN apk add --no-cache jq curl

COPY --from=builder /src/xpost /usr/local/bin/xpost
RUN chmod +x /usr/local/bin/xpost

CMD ["/bin/sh"]
