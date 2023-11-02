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

Start `vite` dev server:

```shell
yarn dev
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

- **Migrator**: to manage database creation and modification.
- **Management**: The web UI to manage the URLs that are registered in the system.
- **Redirector**: The redirection service that takes the requests and redirects them to the proper target.

### Migrator

Manages the database setup and migrations. You can start it by running:

```shell
go run cmd/migrator/main.go
```

Flags:

- `--db-file-path` defines the location of the SQLite database data file and its name. Default value: `./data/shrtnr.db`

    You can set a custom path by passing the flag like this:

    ```shell
    go run cmd/migrator/main.go --db-file-path=/path/to/your/database/file
    ```

*Important:* You need to run this service if you haven't run any `Shrtnr` services before.

### Management

Serves the Web UI to manage all URLs in the system. You can start it by running:

```shell
go run cmd/management/main.go
```

Flags:

- `--env` defines the environment where the application is running. Default value: `development`

    You can change it like this:

    ```shell
    go run cmd/management/main.go --env=environment
    ```
    
    *Important:* If you set it to `production`, you need to bundle the client code before so the build can embed the bundle files in the executable.

- `--port` defines the port where the Web UI server will listen for requests. Default value: `3000`

  You can change it like this:

    ```shell
    go run cmd/management/main.go --port=420
    ```

- `--db-file-path` defines the SQLite database's data file location and name. Default value: `./data/shrtnr.db`

  You can change it like this:

    ```shell
    go run cmd/management/main.go --db-file-path=/path/to/your/database/file
    ```

- `--capacity` defines the limit of URLs the service can store. Default value: `2500`

  You can change it like this:

    ```shell
    go run cmd/management/main.go --capacity=69
    ```

- `--search-term-limit` defines the limit of term results the search cache returns when called. Default value: `10`

  You can change it like this:

    ```shell
    go run cmd/management/main.go --search-term-limit=42
    ```

- `--search-concurrency` defines the limit of concurrent processes when checking each trie cache for search terms. Default value: `30`

  You can change it like this:
  
    ```shell
    go run cmd/management/main.go --search-concurrency=1977
    ```

- `--redirect-service-url` defines the redirector service's URL. Default value: `http://localhost:4269`

    You can change it like this:

    ```shell
    go run cmd/management/main.go --redirect-service-url=https://do.main
    ```

### Redirector

Serves the Web Service to handle redirects of URLs stored in the system. You can start it by running:

```shell
go run cmd/redirector/main.go
```

Flags:

- `--env` defines the environment where the application is running. Default value: `development`

  You can change it like this:

    ```shell
    go run cmd/redirector/main.go --env=environment
    ```

  *Important:* If you set it to `production`, you need to bundle the client code before so the build can embed the bundle files in the executable.

- `--port` defines the port where the web server will listen for requests. Default value: `4269`

  You can change it like this:

    ```shell
    go run cmd/redirector/main.go --port=420
    ```

- `--db-file-path` defines the SQLite database's data file location and name. Default value: `./data/shrtnr.db`

  You can change it like this:

    ```shell
    go run cmd/redirector/main.go --db-file-path=/path/to/your/database/file
    ```

## Documentation

For further Documentation:

- [Management Service](/docs/management)
- [Redirector Service](/docs/redirector)

## Post-morten 

I tried to apply what I understand for Domain-Driven Design with a touch of Hexagonal Architecture and Event-Driven Architecture. While also using HTMX as my frontend library and model to handle Web UI reactivity and rendering complementing it with WebComponents for the stuff I couldn't handle with my knowledge of HTMX. 

### Where it went well

- It helped me to separate business logic from the implementation of my adapters. Where I was able to simply connect the domain logic into the HTTP endpoints.


- The implementation of an Event-Driven SoftWare Architecture gave me a way to synchronize the caches in the different subsystems.


- HTMX let me build a reactive Web UI the way I wanted without using an JSON API and another frontend library like React. 
  - My client-side code is simpler.
  - I was easier to wrap my head around compare to Turbo.


- Using WebComponents for the final stretch of interactivity was awesome.
  - I never taught that this web platform native technology could help me create some behaviors that in the past I considered I would need a frontend library like React. 


- Tailwind CSS helped me a lot when it came to write build my UI and make it presentable.
    - For a CSS dummy like myself it was cool to easily write cool CSS styles and behaviors so beautiful.

### Where I failed or stalled

- What I understand as Hexagonal Architecture clearly is flawed and I need to investigate more in the field.
  - The way I implemented the `repository` is not universal but at the same time the way the application accepts the structures I can implement an `port` for different types of SQL databases.


- I implemented incorrectly the idea of Domain-Driven Design. 
  - I clearly mixed core logic with a `module` centric design that made it impossible for me to write any type of test suite to prevent me to shot myself in the foot or take 2-3 hours debugging stuff I broke or simply write super buggy code.


- The `module` centric design I used on the application is clearly how to implement the `Singleton Pattern` in Go.
  - Yeah, basically I build an entire system with `Singletons` pretty bad idea.
  - There's some parts of the system that probably can use this pattern.

- My understanding of HTMX and WebComponents is clearly in complete. This was clear to me when implementing the Notification toast.
  - The way I implemented the "timer line" to indicate the user how long until the Toast will clear automatically, is super hacky and inconsistent.

### Was it fun?

Mistakes were made, but I enjoyed the process. I feel proud on how this App turned out.

### What's next?

- Update the ROM Stack to the new standard created by working with this App.


- Read a lot and learn about Domain-Driven Design, Hexagonal Architecture and Software Architecture in General.


- Plan the next application to continue growing the ROM Stack.

---
> ### Disclaimer
> I do not recommend to open any of the services to the internet. I didn't implement User Auth on purpose. I designed this system as an exercise to develop something simple with the ROM Stack and *maybe* use it as part of my Home Lab network. - [@NordGus](https://github.com/NordGus)

*Built with the ROM Stack*
