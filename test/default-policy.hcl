path "secret/write/*" {
  capabilities = ["create", "update"]
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
