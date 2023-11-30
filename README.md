# Go CRUD Users Service
The User Management Service is a highly efficient and scalable microservice written in Go, designed to handle the complete set of CRUD (Create, Read, Update, Delete) operations for user data management. It serves as a foundational component for systems requiring user management functionality, such as authentication, profile management, and access control

## Key Features
- *Create Users*: Enables the creation of new user accounts. It accepts user details (like name, email, password, etc.), performs validation, and stores them securely in the database. Passwords are hashed for security purposes.
- *Read Users*: Supports querying user details. This can be done either by retrieving a specific user by their unique identifier. The service ensures sensitive data like passwords are never exposed in the retrieval process.
- *Update Users*: Allows modifications to existing user. TODO: Users or administrators can update details such as email, name, or passwords. The service ensures proper authentication and authorization before allowing updates.
- *Delete Users*: Facilitates the secure deletion of user accounts. This includes the option hard deletion (permanent removal from the database).
- *Error Handling*: Comprehensive error handling provides clear, actionable feedback for both client applications and end-users.
- *Database Integration*: Compatible with mysql database.

## Technologies
- Golang 1.20
- Gorm
- swaggo
- MySQL
- Gorilla Mux
- testify

## Requirement
- Docker
- docker-compose
- Golang `>=` 1.20
- golangci-lint

## Setup
You need to clone the repository <br />
<b>Repo : git@github.com:brcodingdev/go-crud-users.git </b>

## Run

```bash
# run the service
$ make run
```

## Swagger
After run the application

Open your browser:
http://localhost:8010/swagger/index.html#/

### Test

```bash
# run tests
$ make test
```

### Lint

```bash
# run tests
$ make lint
```

## TODO
- More Unit Tests
- JWT
- More Validations
- List users features
- Logs for observability