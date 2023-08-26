
# Donna Notes
Donna Notes is a simple CRUD application built in Golang and GRPC, with the aim of studying these technologies and potentially serving as a notes API for another project.

![Donna Notes](https://i.imgur.com/Ezz9xGY.png)


## Running locally

Clone the project

```bash
  git clone https://github.com/devigor/donna-notes-service
```

Clone the project

```bash
  cd donna-notes-service
```

Start the server

```bash
  go run cmd/main.go
```

## Environment Variables

To run this project, you will need to add the following environment variables to your .env file.

`DATABASE_URL`


## Stack

**Back-end:** Golang, GRPC

**Database:** PostgreSQL

**Infra:** Docker


## What have I learned?

In this project, the challenge was to put into practice what I've studied about Go and GRPC by creating a simple API. I also took the opportunity to refresh my knowledge of Docker and Docker Compose to run the entire application within Docker. Additionally, I wanted to use a database I had never worked with before, PostgreSQL.

The key takeaways and learnings from this project include:

- Pointers in Go
- Services and messages in .proto files
- Communication between two services within a container
- Understanding how GRPC facilitates communication.


