# goboil-clean
Golang boilerplate with clean architecture
- PORT : 3030
- PATH : /

## Installation

``` bash
# clone the repo
$ git clone 

# go into app's directory
$ cd goboil-clean
```

## Build & Run

Local environment
``` bash
# build 
$ go build

# run 
$ ENV=DEV go run main.go
$ ENV=DEV ./filego
```

Docker environment
``` bash
# build 
$ docker build -t goboil-clean-api:latest .

# run
$ docker compose -f deployments/docker-compose.yml up -d
```

## Documentation

Install environment
``` bash
# get swagger package 
$ go install github.com/swaggo/swag/cmd/swag@latest

# move to swagger dir
$ cd $GOPATH/src/github.com/swaggo/swag

# install swagger cmd 
$ go install cmd/swag
```

Generate documentation
``` bash
# generate swagger doc
$ swag init --propertyStrategy snakecase
```
to see the results, run app and access {{url}}/swagger/index.html

## Description 
This project built in clean architecture that contains :
1. Factory   
2. Middleware 
3. Handler
4. Binder
5. Validation
6. Usecase
7. Repository
8. Model
9. Database
9. Migration
10. Seed

This project have some default endpoint :
- Auth 
  - Login
  - Register
- Sample
  - Get (+ pagination, sort & filter)
  - GetByID
  - Create (+ transaction scope)
  - Update (+ transaction scope)
  - Delete

# Author
Muhammad Cholis Malik