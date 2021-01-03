#!/bin/bash

cd cache

# Cleanup
rm etcd-v3.4.14-darwin-amd64.zip
rm -rf etcd-v3.4.14-darwin-amd64
rm etcd.log
rm -rf default.etcd

curl -OL https://github.com/etcd-io/etcd/releases/download/v3.4.14/etcd-v3.4.14-darwin-amd64.zip
unzip etcd-v3.4.14-darwin-amd64.zip

./etcd-v3.4.14-darwin-amd64/etcd > etcd.log 2>&1 &

echo "===> etcd PID:" $!

sleep 10

curl -L http://127.0.0.1:2379/version

cd ..
