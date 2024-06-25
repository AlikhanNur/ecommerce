
# E-commerce API Development Using Gin Framework and MongoDB

This project aims to develop a robust E-commerce platform leveraging the Gin framework for API development and MongoDB for data storage. The platform will exclusively operate through APIs, facilitating seamless interactions between clients and the server-side application logic. The project will be extensively tested using Postman to ensure API functionality, reliability, and performance.


## Run project

run project with bellow commands

```bash
  docker-compose up -d
  go run main.go
```

## API Reference

#### Sign up function api call POST

```http
  http://localhost:8000/users/signup
```

```bash
  {
  "first_name": "Alikhan",
  "last_name": "Nuraddinov",
  "email": "alikhan@gmail.com",
  "password": "alikhanddinov",
  "phone": "87073331912"
}
```
response

```bash
  {
 {
    "user": {
        "_id": "667a70cd3239f066d110c89f",
        "first_name": "Alikhan",
        "last_name": "Nuraddinov",
        "password": "$2a$14$39F.C7NrVPMGyxh3ezcAXurLEQTNON/KWkjC.K0cQLpjodZi767bS",
        "email": "alikhan@gmail.com",
        "phone": "87073331912",
        "token": "",
        "Refresh_Token": "",
        "created_at": "2024-06-25T12:25:01+05:00",
        "updtaed_at": "2024-06-25T12:25:01+05:00",
        "user_id": "667a70cd3239f066d110c89f",
        "usercart": [],
        "address": [],
        "orders": []
    }
}
}
```

#### Login function api call POST request

```http
 http://localhost:8000/users/login
```



