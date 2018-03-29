# keep-up-graphql

[![Build Status](https://travis-ci.org/abbeyhrt/keep-up-graphql.svg?branch=master)](https://travis-ci.org/abbeyhrt/keep-up-graphql)

## Up & Running

`keep-up-graphql` requires the following tools and technologies for local development:

* Golang, can use Homebrew for installation
* [Docker](https://docs.docker.com/docker-for-mac/install/) for local development

### Local Database setup

Before running the server, you’ll need to make sure to start up the local database defined in `docker-compose.yml`. You can do this by running:

```
bash
# Run the database in detached mode
docker-compose up -d
```

You can verify that postgres is up by running `docker ps` and seeing the postgres container in the displayed list.

### Starting the server

You can start the server locally by running:

```
go
# Source your local `.env` file
source .env
go run cmd/pubapid/main.go
```

### Local `.env` file

We require a couple ENV variables in order to properly run the server. You can start off by copying `.env.template` to `.env`. For secrets, reach out to a team member in order to get the exact values. An important note, make sure not to commit the `.env` file!
