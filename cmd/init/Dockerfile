FROM golang:1.9 as builder


WORKDIR /go/src/github.com/hmuendel/kubevaulter

COPY . .

RUN COMMIT=$(git rev-parse HEAD) VERSION=$(cat ./cmd/init/version) CGO_ENABLED=0 GOOS=linux go build \
--ldflags "-X github.com/hmuendel/kubevaulter/cmd/init/main.COMMIT=$COMMIT -X github.com/hmuendel/kubevaulter/cmd/init/main.VERSION=$VERSION" \
./cmd/init

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/hmuendel/kubevaulter/init .
CMD ["./init"]




