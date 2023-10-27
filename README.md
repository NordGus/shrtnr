# Shrtnr

A URL Shortener service to practice Message Driven Software Architecture and distributed systems in Go with the ROM Stack

## Design Constrains

The service should only hold a maximum of N short urls. When a new url is added and the service can't store more, the oldest is drop to make space.

> Let's implement some queues

## Setup

Download the javascript dependencies (You can skip this step if you are running the `devcontainer` environment):

```shell
yarn install
```

Now download the go dependencies (You can skip this step if you are running the `devcontainer` environment):

```shell
go mod download
```

Create and prepare the database by running the [Migrator](#migrator) service (check [docs](#migrator) for additional configuration):

```shell
go run cmd/migrator/main.go
```

Start the [Management](#management) service to serve the management Web UI on http://localhost:3000 (check [docs](#management) for additional configuration):

```shell
go run cmd/management/main.go
```

Start the [Redirector](#redirector) service to listen for redirection calls on http://localhost:4269 (check [docs](#redirector) for additional configuration):

```shell
go run cmd/redirector/main.go
```

## Services

`Shrtnr` is compose of 3 different services.

- Migrator: to manage database creation and modification.
- Management: The web UI to manage the URLs that are registered in the system.
- Redirector: The redirection service that takes the requests and redirects them to the proper target.

### Migrator

Service that manages the database setup and migrations. You can use it by running:

```shell
go run cmd/migrator/main.go
```

It supports the flag `--db-file-path` to define the location of the SQLite database data file, and it's name. By default, it uses `./data/shrtnr.db`.

You can set a custom path by passing the flag like this:

```shell
go run cmd/migrator/main.go --db-file-path=/path/to/your/database/file
```

*Important:* You need to run this service if you haven't run any `Shrtnr` services before

### Management

Services that serves the management Web UI to manage all URLs in the system. To start it by running:

```shell
go run cmd/management/main.go
```



### Redirector



## Usage

### Disclaimer
I do not recommend to open any of the services to the internet. I didn't implement User Auth on purpose. I designed this system as an exercise to develop something simple with the ROM Stack and *maybe* use it as part of my Home Lab network.



---
Built with the ROM Stack
