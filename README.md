## About Service
This is a service that provides API for management Laboratory. It is built with [Go](https://go.dev/). Backend Services communicate with Database using [GORM](https://gorm.io/) ORM. The Database used is [PostgreSQL](https://www.postgresql.org/).

## Repository Structure

The repository is structured as follows:

- `main` is main branch of the repository and contains the latest stable version of the code.


## Installation

### Prerequisites
- Go 1.18
- Docker
- Docker Compose


### Build Command
```bash
go build -tags netgo -ldflags '-s -w' -o main
./main
```


### Running the application
```bash
docker-compose up
```
