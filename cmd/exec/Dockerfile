FROM golang:1 as builder


WORKDIR /go/src/github.com/hmuendel/kubevaulter

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build \
--ldflags "-X main.COMMIT=$(git rev-parse HEAD) \
-X main.VERSION=$(cat ./cmd/rec/version)" 	\
./cmd/exec

RUN git rev-parse HEAD
RUN ls -l

FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN update-ca-certificates
RUN apk add --no-cache curl
WORKDIR /root/
COPY --from=builder /go/src/github.com/hmuendel/kubevaulter/exec  .
CMD ["./exec"]




