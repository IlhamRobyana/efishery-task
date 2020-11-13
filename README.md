# efishery-task
Built with Golang, NodeJS, PostgreSQL, Redis, Echo, Express, and other third-party libraries.
## Installation
Make sure Golang and NodeJS is already installed.
### Auth API
Install the dependendices with ```dep ensure```
### Fetching API
Install the dependencies with ```npm install```
## Peresquite
Before running the application, set the dot env file first using .env.example as the layout.

Of course PostgreSQL and Redis needed to be setup first.

And then migrate the data structure to PostgreSQL

```cd auth-api/migration```

```go run main.go```

## Running
Go to the root folder of the project and enter the following:
### Auth API
```go run auth.go```
### Fetching API
```npm start```
