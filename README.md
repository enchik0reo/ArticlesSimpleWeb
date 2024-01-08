# Articles

It's my first very simple full-stack web Go application.

## Features

- RESTful routing
- Server using gorilla/mux router
- Data persistence using PostgreSQL
- Dynamic HTML using Go templates
- Save and view articles

## Development

Software requirements:

- Go
- Docker

To start the application:

```sh
$ git clone github.com/enchik0reo/ArticlesLittleWeb
$ cd ArticlesLittleWeb

# Run docker-compose file
$ docker-compose up -d

# Run server on port 4000
$ go run ./cmd/articles/main.go
```
Go to http://localhost:4000/ and try it.

To stop service use `SIGTERM` signal (you can do Ctrl+C).