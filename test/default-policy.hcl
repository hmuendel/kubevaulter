path "secret/write/*" {
  capabilities = ["create", "update", "read"]
}

path "secret/mixed/*" {
  capabilities = ["read"]
}

path "secret/mixed/write" {
  capabilities = ["create", "update"]
}

path "secret/reading/foo" {
  capabilities = ["create"]
}
