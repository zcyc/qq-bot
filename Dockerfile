FROM golang:alpine AS builder

ENV GOPROXY="https://goproxy.cn"

WORKDIR /qq-bot

COPY . .

RUN go build -o ./build/qq_bot cmd/qq_bot/main.go

FROM alpine

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
        apk update && \
        apk add tzdata ca-certificates && \
        cp -r -f /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
        rm -rf /var/cache/apk/*

ADD conf /build/conf
WORKDIR /build
COPY --from=builder /qq-bot/build/qq_bot /build/qq_bot

CMD ["./qq_bot"]