FROM golang:1.16-alpine as build

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /go/release
ADD . .
RUN go mod download
RUN go build -o app ./main.go

FROM alpine as prod

RUN apk add tzdata
COPY --from=build /go/release/app /
COPY --from=build /go/release/config/config-docker.yaml ./config/config.yaml

# 启动服务
CMD ["/app"]
