# e-commerce backend
Backend for e-commerce app

# Development Setup

## Tools

- Go version 1.3
- Go Modules for dependency management.
- migrate for database migration. https://github.com/golang-migrate/migrate

## Migration

Make sure to use the latest database schema by running:

```shell
migrate \
  -source file://migrations \
  -database "mysql://[username]:[password]@tcp([host]:[port])/[database]" up
```

To modify database schema, make new migration files by using command below:

```shell
migrate create -ext sql -dir migrations [migration_name]
```

There is already script to run migrations in Makefile using test or development configuration

```shell
make migrate-test-up 
```

```shell
make migrate-dev-up 
```

You can change the database configs for migrations

## Run Program

This program needs environment variables. See example below.

```shell
PORT=8800

READER_HOST=localhost
READER_PORT=3306
READER_USER=root
READER_PASSWORD=password

WRITER_HOST=localhost
WRITER_PORT=3306
WRITER_USER=root
WRITER_PASSWORD=password

MYSQL_DATABASE_NAME=db_name

TIMEOUT_ON_SECONDS=120
OPERATION_ON_EACH_CONTEXT=500
```

You can store environment variables to a file, such as ".env". But, if you want to differentiate environment variables used by docker, you can give another name, such as ".env-no-docker". Run the program with commands below.

```shell
./cmds/env .env go run main.go
```

# Test Program
To ensure code quality, one must test the code by its units. One can use docker to initiate the test dependency such as MySQL. To do test, one also need env configuration. See example below that is suited for the docker-compose.test configuration. Name the below file to `env-test`

```shell
MYSQL_DATABASE_NAME=test
READER_HOST=localhost
READER_PORT=3316
READER_USER=test
READER_PASSWORD=test
WRITER_HOST=localhost
WRITER_PORT=3316
WRITER_USER=test
WRITER_PASSWORD=test
MYSQL_MIGRATOR_USERNAME=test
MYSQL_MIGRATOR_PASSWORD=test
MYSQL_HOST=localhost
MYSQL_PORT=3316
```

Firstly migrate the db using
```shell
make migrate-test-up
```

Run the test using
```shell
make test
```

# Deployment
One can deploy this application using Docker. Make sure docker is installed on your system. There is already an example of env file for production called env-prod.
Create .env file from env-prod:
```shell
cp env-prod .env
```
Then deploy the database depency docker container using:

``` shell
./cmds/docker-up-prod.sh
```

Make sure migrate is installed on the system to do database migration. And then do this:
``` shell
make migrate-prod-up
```

Finally deploy the Go application using:
``` shell
./cmds/deploy.sh
```