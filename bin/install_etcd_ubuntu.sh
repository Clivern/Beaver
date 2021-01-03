#!/bin/bash

ETCD_VER=v3.4.14

# choose either URL
GOOGLE_URL=https://storage.googleapis.com/etcd
GITHUB_URL=https://github.com/etcd-io/etcd/releases/download
DOWNLOAD_URL=${GOOGLE_URL}

rm -f /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz
rm -rf /tmp/etcd-download-test && mkdir -p /tmp/etcd-download-test

curl -L ${DOWNLOAD_URL}/${ETCD_VER}/etcd-${ETCD_VER}-linux-amd64.tar.gz -o /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz
tar xzvf /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz -C /tmp/etcd-download-test --strip-components=1
rm -f /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz

/tmp/etcd-download-test/etcd --version
/tmp/etcd-download-test/etcdctl version

sudo cp /tmp/etcd-download-test/etcd /usr/local/bin/
sudo cp /tmp/etcd-download-test/etcdctl /usr/local/bin/

sudo mkdir -p /var/lib/etcd/
sudo mkdir /etc/etcd

sudo groupadd --system etcd
sudo useradd -s /sbin/nologin --system -g etcd etcd

sudo chown -R etcd:etcd /var/lib/etcd/

echo "[Unit]
Description=etcd key-value store
Documentation=https://github.com/etcd-io/etcd
After=network.target

[Service]
User=etcd
Type=notify
Environment=ETCD_DATA_DIR=/var/lib/etcd
Environment=ETCD_NAME=%m
ExecStart=/usr/local/bin/etcd --listen-client-urls http://0.0.0.0:2379 --advertise-client-urls http://0.0.0.0:2379
Restart=always
RestartSec=10s
LimitNOFILE=40000

[Install]
WantedBy=multi-user.target" > /etc/systemd/system/etcd.service

sudo systemctl daemon-reload
sudo systemctl start etcd.service

# Then enable authentication on etcd server
# etcdctl user add root
# etcdctl auth enable
