<p align="center">
    <img alt="Beaver Logo" src="https://raw.githubusercontent.com/Clivern/Beaver/master/assets/img/logo.png" height="80" />
    <h3 align="center">Beaver</h3>
    <p align="center">A Real Time Messaging Server.</p>
</p>

Beaver is a real-time messaging server. With beaver you can easily build scalable in-app notifications, realtime graphs, multiplayer games, chat applications, geotracking and more in web applications and mobile apps.

<p align="center">
    <img alt="Beaver Single Node" src="https://raw.githubusercontent.com/Clivern/Beaver/feature/enhancements/assets/charts/mixed.png" />
</p>

## Documentation

### Config & Run The Application

Beaver uses [Go Modules](https://github.com/golang/go/wiki/Modules) to manage dependencies. First Create a dist config file.

```bash
$ cp config.yml config.dist.yml
```

Then add your `app.*`, `log.*`, `redis_*`, `api.*` ...etc.

```yml
app:
    mode: dev
    port: 8080
    domain: example.com
    secret: sWUhHRcs4Aqa0MEnYwbuQln3EW8CZ0oD

log:
    level: info
    path: var/logs

redis:
    addr: localhost:6379
    password:
    db: 0

api:
    token: sWUhHRcs4Aqa0MEnYwbuQln3EW8CZ0oD
```

And then run the application.

```bash
$ go build beaver.go
$ ./beaver

// OR

$ go run beaver.go

// To Provide a custom config file
$ ./beaver -config=/custom/path/config.dist.yml
$ go run beaver.go -config=/custom/path/config.dist.yml
```

Or [download a pre-built Beaver binary](https://github.com/Clivern/Beaver/releases) for your operating system.

```bash
$ curl -sL https://github.com/Clivern/Beaver/releases/download/x.x.x/beaver_x.x.x_OS.tar.gz | tar xz
$ ./beaver -config=config.dist.yml
```

Also running beaver with docker still an option.

```bash
$ mkdir -p $HOME/srv/beaver
$ mkdir -p $HOME/srv/beaver/configs
$ mkdir -p $HOME/srv/beaver/logs

$ cd $HOME/srv/beaver

$ curl -OL https://raw.githubusercontent.com/Clivern/Beaver/master/Dockerfile
$ curl -OL https://raw.githubusercontent.com/Clivern/Beaver/master/docker-compose.yml
$ curl -OL https://raw.githubusercontent.com/Clivern/Beaver/master/config.yml

$ cp config.yml ./configs/config.dist.yml
$ rm config.yml
# Update log.path to be the absolute path to config file on host machine ($HOME/srv/beaver/logs)
$ sed -i "s|var/logs|${HOME}/srv/beaver/logs|g" ./configs/config.dist.yml

# Build and run containers
$ cd $HOME/srv/beaver/
$ docker-compose build
$ docker-compose up -d
```


### API Endpoints

Create a Config `app_name`:

```bash
$ curl -X POST \
    -H "Content-Type: application/json" \
    -H "X-AUTH-TOKEN: sWUhHRcs4Aqa0MEnYwbuQln3EW8CZ0oD" \
    -d '{"key":"app_name","value":"Beaver"}' \
    "http://localhost:8080/api/config"
```

Get a Config `app_name`:

```bash
$ curl -X GET \
    -H "Content-Type: application/json" \
    -H "X-AUTH-TOKEN: sWUhHRcs4Aqa0MEnYwbuQln3EW8CZ0oD" \
    "http://localhost:8080/api/config/app_name"

{"key":"app_name","value":"Beaver"}
```

Update a Config `app_name`:

```bash
$ curl -X PUT \
    -H "Content-Type: application/json" \
    -H "X-AUTH-TOKEN: sWUhHRcs4Aqa0MEnYwbuQln3EW8CZ0oD" \
    -d '{"value":"Beaver"}' \
    "http://localhost:8080/api/config/app_name"
```

Delete a Config `app_name`:

```bash
$ curl -X DELETE \
    -H "Content-Type: application/json" \
    -H "X-AUTH-TOKEN: sWUhHRcs4Aqa0MEnYwbuQln3EW8CZ0oD" \
    "http://localhost:8080/api/config/app_name"
```

Create a Channel:

```bash
# Private Channel
$ curl -X POST \
    -H 'Content-Type: application/json' \
    -H 'X-AUTH-TOKEN: sWUhHRcs4Aqa0MEnYwbuQln3EW8CZ0oD' \
    -d '{"name": "app_x_chatroom_1", "type": "private"}' \
    'http://localhost:8080/api/channel'

# Public Channel
$ curl -X POST \
    -H 'Content-Type: application/json' \
    -H 'X-AUTH-TOKEN: sWUhHRcs4Aqa0MEnYwbuQln3EW8CZ0oD' \
    -d '{"name": "app_y_chatroom_1", "type": "public"}' \
    'http://localhost:8080/api/channel'

# Presence Channel
$ curl -X POST \
    -H 'Content-Type: application/json' \
    -H 'X-AUTH-TOKEN: sWUhHRcs4Aqa0MEnYwbuQln3EW8CZ0oD' \
    -d '{"name": "app_z_chatroom_5", "type": "presence"}' \
    'http://localhost:8080/api/channel'
```

Get a Channel:

```bash
$ curl -X GET \
    -H 'Content-Type: application/json' \
    -H 'X-AUTH-TOKEN: sWUhHRcs4Aqa0MEnYwbuQln3EW8CZ0oD' \
    -d '' \
    'http://localhost:8080/api/channel/app_x_chatroom_1'
{
    "created_at":1545573214,
    "listeners_count":0,
    "name":"app_x_chatroom_1",
    "subscribers_count":0,
    "type":"private",
    "updated_at":1545573214
}

$ curl -X GET \
    -H 'Content-Type: application/json' \
    -H 'X-AUTH-TOKEN: sWUhHRcs4Aqa0MEnYwbuQln3EW8CZ0oD' \
    -d '' \
    'http://localhost:8080/api/channel/app_y_chatroom_1'
{
    "created_at":1545573219,
    "listeners_count":0,
    "name":"app_y_chatroom_1",
    "subscribers_count":0,
    "type":"public",
    "updated_at":1545573219
}

$ curl -X GET \
    -H 'Content-Type: application/json' \
    -H 'X-AUTH-TOKEN: sWUhHRcs4Aqa0MEnYwbuQln3EW8CZ0oD' \
    -d '' \
    'http://localhost:8080/api/channel/app_z_chatroom_5'
{
    "created_at": 1545573225,
    "listeners": null,
    "listeners_count": 0,
    "name": "app_z_chatroom_5",
    "subscribers": null,
    "subscribers_count": 0,
    "type": "presence",
    "updated_at": 1545573225
}
```

Update a Channel `app_y_chatroom_1`:

```bash
$ curl -X PUT \
    -H 'Content-Type: application/json' \
    -H 'X-AUTH-TOKEN: sWUhHRcs4Aqa0MEnYwbuQln3EW8CZ0oD' \
    -d '{"type": "private"}' \
    'http://localhost:8080/api/channel/app_y_chatroom_1'
```

Delete a Channel `app_y_chatroom_1`:

```bash
$ curl -X DELETE \
    -H 'Content-Type: application/json' \
    -H 'X-AUTH-TOKEN: sWUhHRcs4Aqa0MEnYwbuQln3EW8CZ0oD' \
    -d '' \
    'http://localhost:8080/api/channel/app_y_chatroom_1'
```

Create a Client and add to `app_x_chatroom_1` Channel:

```bash
$ curl -X POST \
    -H 'Content-Type: application/json' \
    -H 'X-AUTH-TOKEN: sWUhHRcs4Aqa0MEnYwbuQln3EW8CZ0oD' \
    -d '{"channels": ["app_x_chatroom_1"]}' \
    'http://localhost:8080/api/client'
{
    "channels": [
        "app_x_chatroom_1"
    ],
    "created_at": 1545575142,
    "id": "69775af3-5f68-4725-8162-09cab63e8427",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjoiNjk3NzVhZjMtNWY2OC00NzI1LTgxNjItMDljYWI2M2U4NDI3QDE1NDU1NzUxNDIiLCJ0aW1lc3RhbXAiOjE1NDU1NzUxNDJ9.EqL-nWwu5p7hJXWrKdZN3Ds2cxWVjNYmeP1mbl562nU",
    "updated_at": 1545575142
}
```

Get a Client `69775af3-5f68-4725-8162-09cab63e8427`:

```bash
$ curl -X GET \
    -H 'Content-Type: application/json' \
    -H 'X-AUTH-TOKEN: sWUhHRcs4Aqa0MEnYwbuQln3EW8CZ0oD' \
    -d '' \
    'http://localhost:8080/api/client/69775af3-5f68-4725-8162-09cab63e8427'
{
    "channels": [
        "app_x_chatroom_1"
    ],
    "created_at": 1545575142,
    "id": "69775af3-5f68-4725-8162-09cab63e8427",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjoiNjk3NzVhZjMtNWY2OC00NzI1LTgxNjItMDljYWI2M2U4NDI3QDE1NDU1NzUxNDIiLCJ0aW1lc3RhbXAiOjE1NDU1NzUxNDJ9.EqL-nWwu5p7hJXWrKdZN3Ds2cxWVjNYmeP1mbl562nU",
    "updated_at": 1545575142
}
```

Subscribe a Client `69775af3-5f68-4725-8162-09cab63e8427` to a Channel `app_z_chatroom_5`:

```bash
$ curl -X PUT \
    -H 'Content-Type: application/json' \
    -H 'X-AUTH-TOKEN: sWUhHRcs4Aqa0MEnYwbuQln3EW8CZ0oD' \
    -d '{"channels": ["app_z_chatroom_5"]}' \
    'http://localhost:8080/api/client/69775af3-5f68-4725-8162-09cab63e8427/subscribe'
```

Unsubscribe a Client `69775af3-5f68-4725-8162-09cab63e8427` from a Channel `app_z_chatroom_5`:

```bash
$ curl -X PUT \
    -H 'Content-Type: application/json' \
    -H 'X-AUTH-TOKEN: sWUhHRcs4Aqa0MEnYwbuQln3EW8CZ0oD' \
    -d '{"channels": ["app_z_chatroom_5"]}' \
    'http://localhost:8080/api/client/69775af3-5f68-4725-8162-09cab63e8427/unsubscribe'
```

Delete a Client:

```bash
$ curl -X DELETE \
    -H 'Content-Type: application/json' \
    -H 'X-AUTH-TOKEN: sWUhHRcs4Aqa0MEnYwbuQln3EW8CZ0oD' \
    -d '' \
    'http://localhost:8080/api/client/69775af3-5f68-4725-8162-09cab63e8427'
```

Publish to a Channel `app_x_chatroom_1`:

```bash
$ curl -X POST \
    -H 'Content-Type: application/json' \
    -H 'X-AUTH-TOKEN: sWUhHRcs4Aqa0MEnYwbuQln3EW8CZ0oD' \
    -d '{"channel": "app_x_chatroom_1", "data": "{\"message\": \"Hello World\"}"}' \
    'http://localhost:8080/api/publish'
```

Broadcast to Channels `["app_x_chatroom_1"]`:

```bash
$ curl -X POST \
    -H 'Content-Type: application/json' \
    -H 'X-AUTH-TOKEN: sWUhHRcs4Aqa0MEnYwbuQln3EW8CZ0oD' \
    -d '{"channels": ["app_x_chatroom_1"], "data": "{\"message\": \"Hello World\"}"}' \
    'http://localhost:8080/api/broadcast'
```

Sample Frontend Client

```js
function Socket(url){
    ws = new WebSocket(url);
    ws.onmessage = function(e) { console.log(e); };
    ws.onclose = function(){
        // Try to reconnect in 5 seconds
        setTimeout(function(){Socket(url)}, 5000);
    };
}

Socket("ws://localhost:8080/ws/$ID/$TOKEN");
```


## Badges

[![Build Status](https://travis-ci.org/Clivern/Beaver.svg?branch=master)](https://travis-ci.org/Clivern/Beaver)
[![GitHub license](https://img.shields.io/github/license/Clivern/Beaver.svg)](https://github.com/Clivern/Beaver/blob/master/LICENSE)
[![Version](https://img.shields.io/badge/Version-1.1.1-red.svg)](https://github.com/Clivern/Beaver/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/Clivern/Beaver)](https://goreportcard.com/report/github.com/Clivern/Beaver)

## For contributors

To run beaver locally for development or even testing, please follow the following:

```bash
# Use src/github.com/clivern/beaver
$ mkdir -p $GOPATH/src/github.com/clivern/beaver
$ git clone https://github.com/Clivern/Beaver.git $GOPATH/src/github.com/clivern/beaver
$ cd $GOPATH/src/github.com/clivern/beaver

# Create a feature branch
$ git branch feature/x
$ git checkout feature/x

$ export GO111MODULE=on
$ cp config.yml config.dist.yml
$ cp config.yml config.test.yml

# Add redis to config.test.yml and config.dist.yml

# to run beaver
$ go run beaver.go
$ go build beaver.go

# To run test cases
$ make ci

# To create a testing redis & rabbitMQ container
$ docker run -d --name redis -p 6379:6379 redis
$ docker pull rabbitmq
$ docker run -d --hostname my-rabbit --name some-rabbit -p 4369:4369 -p 5671:5671 -p 5672:5672 -p 15672:15672 rabbitmq
$ docker exec some-rabbit rabbitmq-plugins enable rabbitmq_management
# Login at http://localhost:15672/ (or the IP of your docker host)
# using guest/guest
```

Then Create a PR with the master branch.

## Changelog

* Version 1.1.1:
```
Fix Dockerfile.
```

* Version 1.1.0:
```
Switch to go 1.11 modules.
Use goreleaser to deliver pre-built binaries.
Upgrade dependencies.
```

* Version 1.0.0:
```
Initial Release.
```


## Acknowledgements

© 2018, Clivern. Released under [MIT License](https://opensource.org/licenses/mit-license.php).

**Beaver** is authored and maintained by [@Clivern](http://github.com/clivern).
