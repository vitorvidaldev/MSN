# Message Sending Network - MSN

Message sending application, that allows multiple users to communicate with each other. Developed with:
- Golang 1.16.7
- Docker
- Docker-compose
- PostgreSQL
- Redis
- Elasticsearch

## Features

Below you will find a list of features available in this API.

### User Features
- Create new user
- Login new user
- Delete user

### Conversation Features
- Create new conversation with another user
- Create new conversation with multiple users
- Delete coversation

### Message Features
- Send message to conversation
- Delete message
- Search message history

### Analytics
- Get usage data
- Get history report (daily, weekly, monthly)

## How do I execute the application?

In the root of this repository, execute the following commands:

```
docker-compose up -d
```

```
go run main.go
```

You can access `http://localhost:8080/health/` in the browser of your choice to verify if the application is up.

## Author

Vitor Vidal - More about me [here](https://github.com/vitorvidaldev).

## License

This project uses the MIT License. See `License` file in this repository.
