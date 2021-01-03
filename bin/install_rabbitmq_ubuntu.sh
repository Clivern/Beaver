#!/bin/bash

apt-get update
echo "deb http://www.rabbitmq.com/debian/ testing main" >> /etc/apt/sources.list
wget -O- https://www.rabbitmq.com/rabbitmq-release-signing-key.asc | sudo apt-key add -

apt-get update
sudo apt-get install -y rabbitmq-server

sudo systemctl enable rabbitmq-server
sudo systemctl start rabbitmq-server

sudo rabbitmqctl add_user admin password
sudo rabbitmqctl set_user_tags admin administrator
sudo rabbitmqctl set_permissions -p / admin ".*" ".*" ".*"

sudo rabbitmq-plugins enable rabbitmq_management

# Login into http://127.0.0.1:15672/
# admin/password
