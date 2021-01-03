#!/bin/bash

cd cache

# Cleanup
rm etcd-v3.4.14-linux-amd64.tar.gz
rm -rf etcd-v3.4.14-linux-amd64
rm etcd.log
rm -rf default.etcd

curl -sL https://github.com/etcd-io/etcd/releases/download/v3.4.14/etcd-v3.4.14-linux-amd64.tar.gz | tar xz

./etcd-v3.4.14-linux-amd64/etcd > etcd.log 2>&1 &

echo "===> etcd PID:" $!

sleep 10

curl -L http://127.0.0.1:2379/version

cd ..
