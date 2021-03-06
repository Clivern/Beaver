To Run Cassandra, Install Docker and docker-compose

```bash
$ apt-get update
$ apt install docker.io
$ systemctl enable docker
$ apt install docker-compose
```

Create `docker-compose.yml` file

```yaml
version: '3'

services:
  n1:
    image: 'cassandra:3.11'
    ports:
      - "9042:9042"
```

Run cassandra container

```bash
$ docker-compose up -d n1
$ docker ps

$ docker-compose exec n1 nodetool help
$ docker-compose exec n1 nodetool status
$ docker-compose exec n1 nodetool ring
```

Node Module Example

```golang
db := driver.NewCassandra().WithHosts(
    strings.Split(viper.GetString("app.database.cassandra.hosts"), ","),
).WithTimeout(
    viper.GetInt("app.database.cassandra.timeout"),
).WithAuth(
    viper.GetString("app.database.cassandra.username"),
    viper.GetString("app.database.cassandra.password"),
)


err = db.CreateSession()

if err != nil {
    panic(fmt.Sprintf(
        "Error while connecting cassandra: %s",
        err.Error(),
    ))
}

defer db.Close()

nodeModule := module.NewNodeModule(db)

id := gocql.TimeUUID()

fmt.Println(nodeModule.Insert(context.Background(), module.NodeModel{
    ID:        id,
    Address:   "http://127.0.0.1",
    Status:    "up",
    Hostname:  "clivern",
    CreatedAt: time.Now().Unix(),
    UpdatedAt: time.Now().Unix(),
}))

fmt.Println(nodeModule.Exists(context.Background(), id))
value, _ := nodeModule.GetById(context.Background(), id)

fmt.Println(value)
value.Status = "down"

fmt.Println(nodeModule.UpdateById(context.Background(), value))

value, _ = nodeModule.GetById(context.Background(), id)

fmt.Println(value)

fmt.Println(nodeModule.Count(context.Background()))
fmt.Println(nodeModule.DeleteById(context.Background(), id))
```


Channel Module Example

```golang
db := driver.NewCassandra().WithHosts(
    strings.Split(viper.GetString("app.database.cassandra.hosts"), ","),
).WithTimeout(
    viper.GetInt("app.database.cassandra.timeout"),
).WithAuth(
    viper.GetString("app.database.cassandra.username"),
    viper.GetString("app.database.cassandra.password"),
)


err = db.CreateSession()

if err != nil {
    panic(fmt.Sprintf(
        "Error while connecting cassandra: %s",
        err.Error(),
    ))
}

defer db.Close()

channelModule := module.NewChannelModule(db)
```


Client Module Example

```golang
db := driver.NewCassandra().WithHosts(
    strings.Split(viper.GetString("app.database.cassandra.hosts"), ","),
).WithTimeout(
    viper.GetInt("app.database.cassandra.timeout"),
).WithAuth(
    viper.GetString("app.database.cassandra.username"),
    viper.GetString("app.database.cassandra.password"),
)


err = db.CreateSession()

if err != nil {
    panic(fmt.Sprintf(
        "Error while connecting cassandra: %s",
        err.Error(),
    ))
}

defer db.Close()

clientModule := module.NewClientModule(db)
```


Message Module Example

```golang
db := driver.NewCassandra().WithHosts(
    strings.Split(viper.GetString("app.database.cassandra.hosts"), ","),
).WithTimeout(
    viper.GetInt("app.database.cassandra.timeout"),
).WithAuth(
    viper.GetString("app.database.cassandra.username"),
    viper.GetString("app.database.cassandra.password"),
)


err = db.CreateSession()

if err != nil {
    panic(fmt.Sprintf(
        "Error while connecting cassandra: %s",
        err.Error(),
    ))
}

defer db.Close()

messageModule := module.NewMessageModule(db)
```