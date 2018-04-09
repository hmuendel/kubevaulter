#!/bin/bash -x

vault server -dev -dev-root-token-id=test &
pid=$!

vault login test



kill $pid