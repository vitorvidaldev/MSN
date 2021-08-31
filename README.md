# MSN

Message sending application, that allows multiple users to communicate with each other

## Executing the application on your local machine

In the root of this repository, execute the following commands:

`docker-compose up -d`

`go run main.go`

You can access `http://localhost:8080/health/` in the browser of your choice to verify if the application is up.

## Postman Collection

```json
{
  "info": {
    "_postman_id": "9cb28f1c-2d4c-432c-8266-0493c4518107",
    "name": "MSN",
    "schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json"
  },
  "item": [
    {
      "name": "Get all users",
      "request": {
        "method": "GET",
        "header": [
          {
            "key": "Content-Type",
            "value": "application/json",
            "type": "text"
          }
        ],
        "url": {
          "raw": "localhost:8080/user/",
          "host": ["localhost"],
          "port": "8080",
          "path": ["user", ""]
        }
      },
      "response": []
    },
    {
      "name": "Get user by id",
      "request": {
        "method": "GET",
        "header": [
          {
            "key": "Content-Type",
            "value": "application/json",
            "type": "text"
          }
        ],
        "url": {
          "raw": "localhost:8080/user/60d8a2ecdff838fdcb116964",
          "host": ["localhost"],
          "port": "8080",
          "path": ["user", "60d8a2ecdff838fdcb116964"]
        }
      },
      "response": []
    },
    {
      "name": "Create user",
      "request": {
        "method": "POST",
        "header": [
          {
            "key": "Content-Type",
            "value": "application/json",
            "type": "text"
          }
        ],
        "body": {
          "mode": "raw",
          "raw": "{\r\n    \"Name\": \"test\",\r\n    \"Email\": \"test@test.com\",\r\n    \"Password\": \"test\"\r\n}"
        },
        "url": {
          "raw": "localhost:8080/user/",
          "host": ["localhost"],
          "port": "8080",
          "path": ["user", ""]
        }
      },
      "response": []
    },
    {
      "name": "Update user",
      "request": {
        "method": "PUT",
        "header": [
          {
            "key": "Content-Type",
            "value": "application/json",
            "type": "text"
          }
        ],
        "body": {
          "mode": "raw",
          "raw": "{\r\n    \"Name\": \"tesetstestestt\",\r\n    \"Email\": \"test@test.com\",\r\n    \"Password\": \"test\"\r\n}"
        },
        "url": {
          "raw": "localhost:8080/user/60d8a2ecdff838fdcb116964",
          "host": ["localhost"],
          "port": "8080",
          "path": ["user", "60d8a2ecdff838fdcb116964"]
        }
      },
      "response": []
    },
    {
      "name": "Delete user",
      "request": {
        "method": "DELETE",
        "header": [
          {
            "key": "Content-Type",
            "value": "application/json",
            "type": "text"
          }
        ],
        "url": {
          "raw": "localhost:8080/user/60d8a2ecdff838fdcb116964",
          "host": ["localhost"],
          "port": "8080",
          "path": ["user", "60d8a2ecdff838fdcb116964"]
        }
      },
      "response": []
    }
  ]
}
```

## Author

Vitor Vidal - vitorvidal.dev@gmail.com

## License

This project uses the MIT License. See `License` file in this repository.
