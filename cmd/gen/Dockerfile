FROM golang:1.9 as builder


WORKDIR /go/src/github.com/hmuendel/kubevaulter

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build \
--ldflags "-X main.COMMIT=$(git rev-parse HEAD) \
-X main.VERSION=$(cat ./cmd/gen/version)" \
./cmd/gen

RUN git rev-parse HEAD


FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN update-ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/hmuendel/kubevaulter/gen .
CMD ["./gen"]




