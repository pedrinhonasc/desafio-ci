FROM golang:alpine as builder

WORKDIR /go/src

COPY ./main.go .

RUN GOOS=linux go build -o /go/bin/soma main.go

FROM scratch

COPY --from=builder /go/bin/soma /go/bin/soma

ENTRYPOINT [ "/go/bin/soma" ]