# keep-up

[![CircleCI](https://circleci.com/gh/abbeyhrt/keep-up.svg?style=shield)](https://circleci.com/gh/abbeyhrt/keep-up)

<!-- To run doctoc, just do `npx doctoc README.md` in this directory! -->

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

## Table of Contents

- [Up & Running](#up--running)
  - [Working on a service](#working-on-a-service)
  - [Starting the GraphQL service](#starting-the-graphql-service)
  - [Local `.env` file](#local-env-file)
- [Troubleshooting](#troubleshooting)
  - [What if I receive a `502 Bad Gateway` response when using the proxy?](#what-if-i-receive-a-502-bad-gateway-response-when-using-the-proxy)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Up & Running

`keep-up` requires the following tools and technologies for local development:

- Node.js recommended to use `nvm` to install `v10`
- Yarn, alternative to `npm`. Can use Homebrew for installation with the
  `--without-node` flag
- Golang, can use Homebrew for installation
- [Docker](https://docs.docker.com/docker-for-mac/install/) for local development

After installing the software above, you should be ready to go. By default,
you'll most likely want the `postgres` and `proxy` services running. `postgres`
provides the underlying database for local development, and `proxy` allows you
to access services at a centralized URL, namely `localhost:3000`.

You can run these two services through `docker-compose` by running:

```bash
docker-compose up -d --build postgres proxy
```

You can verify they are stable by running `docker ps` and seeing the
postgres and proxy containers in the displayed list.

Most likely, you'll also need to run migrations as well to make sure that your
local copy of the database matches the latest version for keep-up. You can find
more details about how we do migrations [here](./docs/migrations.md).

For convenience, here is the default `migrate` command:

```bash
migrate \
  -path ./graphql/internal/database/migrations/ \
  -database postgres://postgres:postgres@0.0.0.0:5432/keep_up_dev?sslmode=disable up
```

### Working on a service

If you're working on an individual service, like the UI, then you can run the
alternative service using `docker-compose`. For example, if you are working on
the UI server then you might want to just run the default `graphql` service.

You can run this by doing:

```bash
docker-compose up -d --build graphql
```

While everything else is running, you can run specific services by visiting the
respective folders and running the corresponding start commands. For the UI
Service, this would be accomplished by running:

```bash
cd ui
PROTOCOL=http PORT=3000 yarn dev
```

_Note: if you ever need to restart everything, you can run:_

```bash
docker-compose down -v --remove-orphans
```

### Starting the GraphQL service

You can start the server locally by running:

```bash
# Source your local `.env` file
source .env
go run ./graphql/cmd/pubapid/main.go
```

If you want to just run the server while working on the UI, you can use
`docker-compose` by doing:

```bash
docker-compose up -d
```

### Local `.env` file

We require a couple ENV variables in order to properly run the server. You can
start off by copying `.env.template` to `.env`. For secrets, reach out to a team
member in order to get the exact values. An important note, make sure not to
commit the `.env` file!

## Troubleshooting

#### What if I receive a `502 Bad Gateway` response when using the proxy?

This most likely means that one of the upstream proxies are not currently running. Make sure that the UI and GraphQL services are running!

#### What if I receive a `400 Bad Request` response when using the proxy?

This most likely means that one of the upstream proxies are not currently running. Make sure that the UI service is running!
