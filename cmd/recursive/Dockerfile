FROM golang:1.9 as builder


WORKDIR /go/src/github.com/hmuendel/kubevaulter

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build \
--ldflags "-X main.COMMIT=$(git rev-parse HEAD) \
-X main.VERSION=$(cat ./cmd/recursive/version)" \
./cmd/recursive

RUN git rev-parse HEAD


FROM alpine:latest
RUN mkdir -p /usr/local/share/ca-certificates
RUN echo -----BEGIN CERTIFICATE-----\
           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\
           D18g0Z/p0JwFGg==\
           -----END CERTIFICATE----- > /usr/local/share/ca-certificates/vault.crt
RUN apk --no-cache add ca-certificates
RUN update-ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/hmuendel/kubevaulter/recursive .
CMD ["./recursive"]




