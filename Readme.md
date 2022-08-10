# Go Starter
## Getting Started !!


### Install
Step 1
```shell
- Setting enviroment for postgres db in file dal/psql.go (username, port, etc ...)
- Create new db in your postgres
- Run file seed/sample.sql to initialize database
```
Step 2
```shell
go build
make gen-query-gorm
```
### Run
```shell
go run main.go
```
