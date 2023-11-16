# Simple REST Backend for library project

## Setup

Prerequisites:
* go programming language installed (1.21.4 at the time of creation) https://go.dev/dl/
* postgres docker image https://hub.docker.com/_/postgres

Navigate to the directory you want to create project in and run following command

`git clone https://github.com/dam900/BD_library_backend.git`

Then start your postgres docker image and edit the connection string in 
`<project_dir>/storage/storage.go`
```
func NewPostgresStore() (*PostgresStorage, error) {
    connStr := "user=postgres dbname=postgres password=hello_world sslmode=disable" // <- this
    db, err := sql.Open("postgres", connStr) 
    ...
```

Then you can finally run

`go run server.go`

This should start an Echo server on port `:1323` 

## Overview




