FROM vault


COPY ca.crt /usr/share/ca-certificates/extra
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY certs/kubevaulter-test-vault-bundle.pem /vault/tls/server.crt
COPY certs/kubevaulter-test-vault.key /vault/tls/server.key

