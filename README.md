# Ecommerce

Dropshipping, you can sell whatever you want

To run the app:
```bash
docker-compose up -d
go run main.go
````
If you have MongoDB installed you can run just application

runs the app at: [http://localhost:8000](http://localhost:8000)

You need sign in first

## Functionality

- JWT authentication
- Docker connection to MongoDB
- Http-endpoints with gin-gonic/gin
- Detailed Structs
- Searching, Adding, Removing, Buying, Updating products in MongoDB

## Postman Endpoints

- Registration: POST [http://localhost:8000/users/signup](http://localhost:8000/users/signup)
```json
{
  "first_name": "Tema",
  "last_name": "Yerlanuly",
  "email": "temirlan.eraly@gmail.com",
  "password": "123123",
  "phone": "+77743"
}
````
- Login:  POST [http://localhost:8000/users/login](http://localhost:8000/users/login)
```json
{
  "email": "temirlan.eraly@gmail.com",
  "password": "123123"
}
```
- GET [http://localhost:8000/users/productview](http://localhost:8000/users/productview)
- POST [http://localhost:8000/admin/addproduct](http://localhost:8000/admin/addproduct)
- GET [http://localhost:8000/users/search?name=](http://localhost:8000/users/search?name=)
- PUT [http://localhost:8000/edithomeaddress?id](http://localhost:8000/edithomeaddress?id=)

List of all endpoints will be displayed by gin package when you will run app