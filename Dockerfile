FROM golang:1.18 AS builder
WORKDIR /go/src 
ENV GO111MODULE=on GOPROXY=https://goproxy.cn CGO_ENABLED=0   
ADD . .
RUN go mod tidy && go build -tags netgo -o /bin/pulsar-client cmd/pulsar-client/main.go

FROM alpine:3.10
COPY --from=builder /bin/pulsar-client /bin/pulsar-client
ENTRYPOINT ["/bin/pulsar-client"]