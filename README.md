# keep-up

[![Build Status](https://travis-ci.org/abbeyhrt/keep-up.svg?branch=master)](https://travis-ci.org/abbeyhrt/keep-up)

## Table of Contents

<!-- To run doctoc, just do `npx doctoc README.md` in this directory! -->

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->


- [Up & Running](#up--running)
  - [Local Database setup](#local-database-setup)
  - [Starting the server](#starting-the-server)
  - [Local `.env` file](#local-env-file)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Up & Running

`keep-up` requires the following tools and technologies for local development:

* Golang, can use Homebrew for installation
* [Docker](https://docs.docker.com/docker-for-mac/install/) for local development

### Local Database setup

Before running the server, you’ll need to make sure to start up the local database defined in `docker-compose.yml`. You can do this by running:

```
# Run the database in detached mode
docker-compose run -d postgres
```

You can verify that postgres is up by running `docker ps` and seeing the postgres container in the displayed list.

### Starting the server

You can start the server locally by running:

```
# Source your local `.env` file
source .env
go run ./graphql/cmd/pubapid/main.go
```

### Local `.env` file

We require a couple ENV variables in order to properly run the server. You can start off by copying `.env.template` to `.env`. For secrets, reach out to a team member in order to get the exact values. An important note, make sure not to commit the `.env` file!
