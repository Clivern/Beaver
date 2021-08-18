<p align="center">
    <img src="/assets/img/gopher.png" width="200" />
    <h3 align="center">Beaver</h3>
    <p align="center">A Real Time Messaging Server.</p>
    <p align="center">
        <a href="https://github.com/Clivern/Beaver/actions/workflows/build.yml">
            <img src="https://github.com/Clivern/Beaver/actions/workflows/build.yml/badge.svg">
        </a>
        <a href="https://github.com/Clivern/Beaver/actions">
            <img src="https://github.com/Clivern/Beaver/workflows/Release/badge.svg">
        </a>
        <a href="https://github.com/Clivern/Beaver/releases">
            <img src="https://img.shields.io/badge/Version-2.0.0-red.svg">
        </a>
        <a href="https://goreportcard.com/report/github.com/Clivern/Beaver">
            <img src="https://goreportcard.com/badge/github.com/Clivern/Beaver?v=2.0.0">
        </a>
        <a href="https://godoc.org/github.com/clivern/beaver">
            <img src="https://godoc.org/github.com/clivern/beaver?status.svg">
        </a>
        <a href="https://github.com/Clivern/Beaver/blob/master/LICENSE">
            <img src="https://img.shields.io/badge/LICENSE-MIT-orange.svg">
        </a>
    </p>
</p>
<br/>
<p align="center">
    <img src="/assets/charts/chart.png?v=2.0.0" width="60%" />
</p>

Beaver is a real-time messaging server. With beaver you can easily build scalable in-app notifications, realtime graphs, multiplayer games, chat applications, geotracking and more in web applications and mobile apps.


## Documentation

#### Run Beaver on Ubuntu

Download [the latest beaver binary](https://github.com/Clivern/Beaver/releases). Make it executable from everywhere.

```zsh
$ export BEAVER_LATEST_VERSION=$(curl --silent "https://api.github.com/repos/Clivern/Beaver/releases/latest" | jq '.tag_name' | sed -E 's/.*"([^"]+)".*/\1/' | tr -d v)

$ curl -sL https://github.com/Clivern/Beaver/releases/download/v{$BEAVER_LATEST_VERSION}/beaver_{$BEAVER_LATEST_VERSION}_Linux_x86_64.tar.gz | tar xz
```

Then install `redis` cluster or a single node. Update the following config file with redis configs.


Create the configs file `config.yml` from `config.dist.yml`. Something like the following:

```yaml
# App configs
app:
    # Env mode (dev or prod)
    mode: ${BEAVER_APP_MODE:-prod}
    # HTTP port
    port: ${BEAVER_API_PORT:-8080}
    # Hostname
    hostname: ${BEAVER_API_HOSTNAME:-127.0.0.1}
    # TLS configs
    tls:
        status: ${BEAVER_API_TLS_STATUS:-off}
        pemPath: ${BEAVER_API_TLS_PEMPATH:-cert/server.pem}
        keyPath: ${BEAVER_API_TLS_KEYPATH:-cert/server.key}

    # API Configs
    api:
        key: ${BEAVER_API_KEY:-6c68b836-6f8e-465e-b59f-89c1db53afca}

    # Beaver Secret
    secret: ${BEAVER_SECRET:-sWUhHRcs4Aqa0MEnYwbuQln3EW8CZ0oD}

    # Runtime, Requests/Response and Beaver Metrics
    metrics:
        prometheus:
            # Route for the metrics endpoint
            endpoint: ${BEAVER_METRICS_PROM_ENDPOINT:-/metrics}

    # Application Database
    database:
        # Database driver
        driver: ${BEAVER_DB_DRIVER:-redis}

        # Redis Configs
        redis:
            # Redis address
            address: ${BEAVER_DB_REDIS_ADDR:-localhost:6379}
            # Redis password
            password: ${BEAVER_DB_REDIS_PASSWORD:- }
            # Redis database
            db: ${BEAVER_DB_REDIS_DB:-0}

    # Log configs
    log:
        # Log level, it can be debug, info, warn, error, panic, fatal
        level: ${BEAVER_LOG_LEVEL:-info}
        # Output can be stdout or abs path to log file /var/logs/beaver.log
        output: ${BEAVER_LOG_OUTPUT:-stdout}
        # Format can be json
        format: ${BEAVER_LOG_FORMAT:-json}
```

The run the `beaver` with `systemd`

```zsh
$ beaver api -c /path/to/config.yml
```

### API Endpoints

Create a Config `app_name`:

```bash
$ curl -X POST \
    -H "Content-Type: application/json" \
    -H "X-API-Key: 6c68b836-6f8e-465e-b59f-89c1db53afca" \
    -d '{"key":"app_name","value":"Beaver"}' \
    "http://localhost:8080/api/config"
```

Get a Config `app_name`:

```bash
$ curl -X GET \
    -H "Content-Type: application/json" \
    -H "X-API-Key: 6c68b836-6f8e-465e-b59f-89c1db53afca" \
    "http://localhost:8080/api/config/app_name"

{"key":"app_name","value":"Beaver"}
```

Update a Config `app_name`:

```bash
$ curl -X PUT \
    -H "Content-Type: application/json" \
    -H "X-API-Key: 6c68b836-6f8e-465e-b59f-89c1db53afca" \
    -d '{"value":"Beaver"}' \
    "http://localhost:8080/api/config/app_name"
```

Delete a Config `app_name`:

```bash
$ curl -X DELETE \
    -H "Content-Type: application/json" \
    -H "X-API-Key: 6c68b836-6f8e-465e-b59f-89c1db53afca" \
    "http://localhost:8080/api/config/app_name"
```

Create a Channel:

```bash
# Private Channel
$ curl -X POST \
    -H 'Content-Type: application/json' \
    -H 'X-API-Key: 6c68b836-6f8e-465e-b59f-89c1db53afca' \
    -d '{"name": "app_x_chatroom_1", "type": "private"}' \
    'http://localhost:8080/api/channel'

# Public Channel
$ curl -X POST \
    -H 'Content-Type: application/json' \
    -H 'X-API-Key: 6c68b836-6f8e-465e-b59f-89c1db53afca' \
    -d '{"name": "app_y_chatroom_1", "type": "public"}' \
    'http://localhost:8080/api/channel'

# Presence Channel
$ curl -X POST \
    -H 'Content-Type: application/json' \
    -H 'X-API-Key: 6c68b836-6f8e-465e-b59f-89c1db53afca' \
    -d '{"name": "app_z_chatroom_5", "type": "presence"}' \
    'http://localhost:8080/api/channel'
```

Get a Channel:

```bash
$ curl -X GET \
    -H 'Content-Type: application/json' \
    -H 'X-API-Key: 6c68b836-6f8e-465e-b59f-89c1db53afca' \
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
    -H 'X-API-Key: 6c68b836-6f8e-465e-b59f-89c1db53afca' \
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
    -H 'X-API-Key: 6c68b836-6f8e-465e-b59f-89c1db53afca' \
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
    -H 'X-API-Key: 6c68b836-6f8e-465e-b59f-89c1db53afca' \
    -d '{"type": "private"}' \
    'http://localhost:8080/api/channel/app_y_chatroom_1'
```

Delete a Channel `app_y_chatroom_1`:

```bash
$ curl -X DELETE \
    -H 'Content-Type: application/json' \
    -H 'X-API-Key: 6c68b836-6f8e-465e-b59f-89c1db53afca' \
    -d '' \
    'http://localhost:8080/api/channel/app_y_chatroom_1'
```

Create a Client and add to `app_x_chatroom_1` Channel:

```bash
$ curl -X POST \
    -H 'Content-Type: application/json' \
    -H 'X-API-Key: 6c68b836-6f8e-465e-b59f-89c1db53afca' \
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
    -H 'X-API-Key: 6c68b836-6f8e-465e-b59f-89c1db53afca' \
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
    -H 'X-API-Key: 6c68b836-6f8e-465e-b59f-89c1db53afca' \
    -d '{"channels": ["app_z_chatroom_5"]}' \
    'http://localhost:8080/api/client/69775af3-5f68-4725-8162-09cab63e8427/subscribe'
```

Unsubscribe a Client `69775af3-5f68-4725-8162-09cab63e8427` from a Channel `app_z_chatroom_5`:

```bash
$ curl -X PUT \
    -H 'Content-Type: application/json' \
    -H 'X-API-Key: 6c68b836-6f8e-465e-b59f-89c1db53afca' \
    -d '{"channels": ["app_z_chatroom_5"]}' \
    'http://localhost:8080/api/client/69775af3-5f68-4725-8162-09cab63e8427/unsubscribe'
```

Delete a Client:

```bash
$ curl -X DELETE \
    -H 'Content-Type: application/json' \
    -H 'X-API-Key: 6c68b836-6f8e-465e-b59f-89c1db53afca' \
    -d '' \
    'http://localhost:8080/api/client/69775af3-5f68-4725-8162-09cab63e8427'
```

Publish to a Channel `app_x_chatroom_1`:

```bash
$ curl -X POST \
    -H 'Content-Type: application/json' \
    -H 'X-API-Key: 6c68b836-6f8e-465e-b59f-89c1db53afca' \
    -d '{"channel": "app_x_chatroom_1", "data": "{\"message\": \"Hello World\"}"}' \
    'http://localhost:8080/api/publish'
```

Broadcast to Channels `["app_x_chatroom_1"]`:

```bash
$ curl -X POST \
    -H 'Content-Type: application/json' \
    -H 'X-API-Key: 6c68b836-6f8e-465e-b59f-89c1db53afca' \
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


## Client:

- [Go Client](https://github.com/domgolonka/beavergo) Thanks [@domgolonka](https://github.com/domgolonka)


## Versioning

For transparency into our release cycle and in striving to maintain backward compatibility, Beaver is maintained under the [Semantic Versioning guidelines](https://semver.org/) and release process is predictable and business-friendly.

See the [Releases section of our GitHub project](https://github.com/clivern/beaver/releases) for changelogs for each release version of Beaver. It contains summaries of the most noteworthy changes made in each release.


## Bug tracker

If you have any suggestions, bug reports, or annoyances please report them to our issue tracker at https://github.com/clivern/beaver/issues


## Security Issues

If you discover a security vulnerability within Beaver, please send an email to [hello@clivern.com](mailto:hello@clivern.com)


## Contributing

We are an open source, community-driven project so please feel free to join us. see the [contributing guidelines](CONTRIBUTING.md) for more details.


## License

© 2018, Clivern. Released under [MIT License](https://opensource.org/licenses/mit-license.php).

**Beaver** is authored and maintained by [@Clivern](http://github.com/clivern).
