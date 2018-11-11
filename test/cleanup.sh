#!/usr/bin/env bash


echo deleting helm release
helm delete --purge kubevaulter-test
kubectl delete cm kubevaulter-test
rm init.log