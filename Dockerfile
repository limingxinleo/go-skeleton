FROM golang:1.13 as builder

LABEL maintainer="limx <h@hyperf.io>"
ENV GOPROXY https://goproxy.cn,direct
ENV GO111MODULE=on

WORKDIR /go/cache

ADD go.mod .
ADD go.sum .
RUN go mod download

WORKDIR /go/builder

ADD . .

RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -installsuffix cgo -o app main.go

FROM scratch

ENV GIN_MODE=release

COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /go/builder/app /

EXPOSE 9501

ENTRYPOINT ["/app"]