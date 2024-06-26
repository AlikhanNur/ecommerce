
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
    "Response": "Successfully Signed Up!!"
}
```

#### Login function api call POST request

```http
 http://localhost:8000/users/login
```

```bash
{
  "email": "dii@gmail.com",
  "password": "alikhanddinov"
}
```
reponse:

```bash
{
    "_id": "667a8c88102171563b3d9fe9",
    "first_name": "Alikhan",
    "last_name": "Nuraddinov",
    "password": "$2a$14$gJj2Vx4e8.pXwjGLqsq5JOtXiX1I.e/HCayKe33q8wPvhsMbDYrUG",
    "email": "dii@gmail.com",
    "phone": "88",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6ImRpaUBnbWFpbC5jb20iLCJGaXJzdE5hbWUiOiJBbGlraGFuIiwiTGFzdE5hbWUiOiJOdXJhZGRpbm92IiwiVWlkIjoiNjY3YThjODgxMDIxNzE1NjNiM2Q5ZmU5IiwiZXhwIjoxNzE5MzkzODAwfQ.05aCgBXpxfhWeD-Gh795Mrkf5NioG0ndm2dkKFcOO64",
    "Refresh_Token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IiIsIkZpcnN0TmFtZSI6IiIsIkxhc3ROYW1lIjoiIiwiVWlkIjoiIiwiZXhwIjoxNzE5OTEyMjAwfQ._IsNBPSiBEMMegyxh3CovbaW6Arpzp5xVLPx4jAjreA",
    "created_at": "2024-06-25T09:23:20Z",
    "updtaed_at": "2024-06-25T09:23:20Z",
    "user_id": "667a8c88102171563b3d9fe9",
    "usercart": [],
    "address": [],
    "orders": []
}
```

#### Admin add products POST request

```http
 http://localhost:8000/admin/addproduct
```

```bash
{
  "product_name": "Alienware x15",
  "price": 2500,
  "rating": 10,
  "image": "alienware.jpg"
}
```

response: 

```bash
 {
  "Successfully added our Product Admin!!"
}
```

#### View all the products in db GET request

```http
  http://localhost:8000/users/productview
```

response: 

```bash
[
    {
        "Product_ID": "667a910d102171563b3d9fea",
        "product_name": "Alienware x15",
        "price": 2500,
        "rating": 10,
        "image": "alienware.jpg"
    },
    {
        "Product_ID": "667a9232102171563b3d9feb",
        "product_name": "Iphone 15 pro max",
        "price": 990,
        "rating": 9,
        "image": "iphone.jpg"
    },
    {
        "Product_ID": "667a924f102171563b3d9fec",
        "product_name": "Iphone 15 pro",
        "price": 790,
        "rating": 9,
        "image": "iphonePro.jpg"
    },
    {
        "Product_ID": "667a928e102171563b3d9fed",
        "product_name": "mac m3 air 13.6",
        "price": 999,
        "rating": 10,
        "image": "macM3.jpg"
    }
]
```


#### Search Product by regex function GET request

```http
  http://localhost:8000/users/search?name=Ip
```
response: 

```bash
 [
    {
        "Product_ID": "667a9232102171563b3d9feb",
        "product_name": "Iphone 15 pro max",
        "price": 990,
        "rating": 9,
        "image": "iphone.jpg"
    },
    {
        "Product_ID": "667a924f102171563b3d9fec",
        "product_name": "Iphone 15 pro",
        "price": 790,
        "rating": 9,
        "image": "iphonePro.jpg"
    }
]
```

### Adding the Products to the Cart (GET REQUEST)

```http
  http://localhost:8000/addtocart?id=xxxproduct_idxxx&userID=xxxxxxuser_idxxxxxx
```

response

```bash
  "successfully added to cart"
```

### Removing item from the Cart GET REQUEST

```http
 http://localhost:8000/removeitem?id=xxxxxxx&userID=xxxxxxxxxxxx
```

response 

```bash
  "successfully removed item"
```

### Listing the item in the users cart (GET REQUEST) and total price

```http
 http://localhost:8000/listcart?id=xxxxxxuser_idxxxxxxxxxx
```
response
```bash
  1980[
    {
        "Product_ID": "667a9232102171563b3d9feb",
        "product_name": "Iphone 15 pro max",
        "price": 990,
        "rating": 9,
        "image": "iphone.jpg"
    },
    {
        "Product_ID": "667a9232102171563b3d9feb",
        "product_name": "Iphone 15 pro max",
        "price": 990,
        "rating": 9,
        "image": "iphone.jpg"
    }
]
```


### Addding the Address (POST REQUEST)

```http
 http://localhost:8000/addadress?id=user_id**\*\***\***\*\***
```
response
```bash
  address added
```

### Editing the Home Address(PUT REQUEST)
```http
 http://localhost:8000/edithomeaddress?id=xxxxxxxxxxuser_idxxxxxxxxxxxxxxx
```

### Editing the Work Address(PUT REQUEST)
```http
 http://localhost:8000/editworkaddress?id=xxxxxxxxxxuser_idxxxxxxxxxxxxxxx
```

### Delete Addresses(GET REQUEST)
```http
 http://localhost:8000/deleteaddresses?id=xxxxxxxxxuser_idxxxxxxxxxxxxx
```

### Cart Checkout Function and placing the order(GET REQUEST)
```http
 http://localhost:8000cartcheckout?id=xxuser_idxxx
```

### Instantly Buying the Products(GET REQUEST)
```http
http://localhost:8000?userid=xxuser_idxxx&pid=xxxxproduct_idxxxx
```



