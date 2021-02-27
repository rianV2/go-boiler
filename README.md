# Go Boiler

Golang boiler plate api service.

### Requirements

* Golang >=1.16
* MySQL >=5.7
* `make` command (optional)
* Docker (optional)

### Documentations

* [API](docs/api.md)

### Feature
* [Go Standar Project Layout](https://github.com/golang-standards/project-layout)
* [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

# Setting up 

1. Make sure MySql is running
2. Create development config file by copying `.env.example` to `.env` and edit the values to match you environments
3. Create database
4. Run this command in project root to run the database migration: `$ make migrate`
5. To run the REST API run: `$ make run-api`
6. To run the test: `$ make test-all`