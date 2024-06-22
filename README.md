# API Go

API in Go with JWT and MySql

Available endpoints:
- /api/v1/users
    - /register (POST): create a new User
    - /login (GET): get a JWT user token

- /api/v1/tasks
    - / (POST): create a new task (must be logged)
    - /{id} (GET): get a task by taskID (must be logged)


To run this project (setup the db with `docker compose up`):
```
make run
``` 

To test this project:
```
make test
``` 

## Build by 🛠️

* [Go](https://go.dev/)
