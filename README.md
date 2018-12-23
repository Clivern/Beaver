<p align="center">
    <img alt="Beaver Logo" src="https://raw.githubusercontent.com/Clivern/Beaver/master/assets/img/logo.png" height="80" />
    <h3 align="center">Beaver</h3>
    <p align="center">A Real Time Messaging Server.</p>
</p>

## Documentation

### Config & Run The Application

Beaver uses [dep](https://github.com/golang/dep) to manage dependencies so you need to install it

```bash
# For latest dep version
$ curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# For latest stable version
$ curl https://raw.githubusercontent.com/golang/dep/v0.5.0/install.sh | sh

$ dep ensure
```

Then Create a dist config file.

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

Also running beaver with docker still an option. Just don't forget to update environment variables on `docker-compose.yml` file. Then run the following stuff

```bash
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

Broadcast to Channels:

```bash
```

Publish to a Channel:

```bash
```

Frontend Client:

```js
```

## Badges

[![Build Status](https://travis-ci.org/Clivern/Beaver.svg?branch=master)](https://travis-ci.org/Clivern/Beaver)
[![GitHub license](https://img.shields.io/github/license/Clivern/Beaver.svg)](https://github.com/Clivern/Beaver/blob/master/LICENSE)
[![Version](https://img.shields.io/badge/Version-Under%20Development-red.svg)](https://github.com/Clivern/Beaver/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/Clivern/Beaver)](https://goreportcard.com/report/github.com/Clivern/Beaver)


## Changelog

* Version 1.0.0:
```
Initial Release.
```


## Acknowledgements

© 2018, Clivern. Released under [MIT License](https://opensource.org/licenses/mit-license.php).

**Beaver** is authored and maintained by [@Clivern](http://github.com/clivern).
