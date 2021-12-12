## 🗒 GoLang To-Do List Sample App w/ Redis

[![Build Status](https://app.travis-ci.com/bvisonl/godo-list.svg?branch=master)](https://app.travis-ci.com/bvisonl/godo-list)

[![Coverage Status](https://coveralls.io/repos/github/bvisonl/godo-list/badge.svg)](https://coveralls.io/github/bvisonl/godo-list)

### 📝 Description

A simple to-do list app written in GoLang with Redis as a backend showing a few GoLang features.

### 📝 Features

- [✅] [Redis](https://redis.io/) as a backend
- [✅] [Gorilla](https://godoc.org/github.com/gorilla/mux) as a router
- [✅] [GoLang](https://golang.org/) as a language
- [✅] [Git](https://git-scm.com/) as a version control system
- [✅] [Travis CI](https://travis-ci.org/) as a continuous integration system
- Using the following features to handle concurrency & application lifecycle:
  - Using [goroutine](https://golang.org/pkg/runtime/#Goexit)
  - Using [sync.WaitGroup](https://golang.org/pkg/sync/#WaitGroup)
  - Using [sync.Mutex](https://golang.org/pkg/sync/#Mutex)
  - Using [channel](https://golang.org/pkg/sync/#Channel)
- Using the following features to handle HTTP requests:
  - Using [http.Handler](https://golang.org/pkg/net/http/#Handler)
  - Using [http.HandlerFunc](https://golang.org/pkg/net/http/#HandlerFunc)
  - Using [http.Server](https://golang.org/pkg/net/http/#Server)
  - Using [http.ListenAndServe](https://golang.org/pkg/net/http/#ListenAndServe)
- Using the [testing](https://golang.org/pkg/testing/) suite to test the app and provide coverage reports
- Using the [gopkg.in/yaml.v2](https://gopkg.in/yaml.v2) package to parse the YAML config file
- Using the [github.com/go-redis/redis] as a Redis client
- Using docker to deploy the app

### 📝 Dependencies

- [gopkg.in/yaml.v2](https://gopkg.in/yaml.v2)
- [github.com/go-redis/redis]()
- [github.com/gorilla]()
- [docker](https://www.docker.com/)
- [golang]()

### 📝 Installation

To install the app, run the following command:

### 📝 Usage

To run the app, run the following command:

```bash
go run main.go
```

### 📝 Testing Coverage

To see the coverage report, run the following command:

```bash
go tool cover -html=coverage.out
```
