**Local Database Setup**

Things it needed:

* docker-compose.yml

  The docker-compose.yml file tells docker some basic information about the docker image we are creating

```
  version: '3'
  services:
    postgres:
      image: postgres:9.6.1
      ports:
      - "${DB_PORT:-5432}:5432"

      volumes:
      - "local-db-volume:/var/lib.postgresql/data"
      environment:
        POSTGRES_USER: postgres
        POSTGRES_PASSWORD: postgres
        POSTGRES_DB: keep_up_dev
  volumes:
    local-db-volume:
      driver: local
```

The volumes part of the docker file directs docker to make a local folder that maops over to the /var/lib.postgresql/data folder in the container.

* database.go

  For the local databse, I created the database.go file. The key part of this is the New function that creates a new instance of a database with the connection string. `sql.Open(connStr)` opens the db connection while db.Ping pings the database. This New function is then used in the main.go file.

* edited main.go

  I edited main.go by adding

  ```
  		if cfg.Env == "development" {
  			_, err := database.New(cfg.Postgres.String())
  			if err != nil {
  				log.Fatal(err)
  			}
  		}
  ```

  which creates a new database if the current environment is in development, not production. That is becasue you don't want this database starting up upon deployment since it's not production ready.

I can confirm that this database works in development by running `docker-compose up -d` in my terminal and then running `go cmd/pubapid/main.go`. If there server starts with no error, then the databse was pinged successfully!

**Notes:**

Make sure the names and passwords to everything matches and don't forget to run `docker-compose up -d`.
To check that you have, run `docker ps` which says wheter or not a connection is open.
