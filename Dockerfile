ARG GOVERSION=alpine
FROM golang:$GOVERSION
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn,direct
WORKDIR $GOPATH/src/camelgem/mail
COPY . $GOPATH/src/camelgem/mail
RUN go build .
ENTRYPOINT ["./mail"]
CMD ["--from=example@qq.com", "--to=example@qq.com", "--provider=QQ", "--host=SMTP", "--subject=xxxxx", "--body=xxxxx."]



