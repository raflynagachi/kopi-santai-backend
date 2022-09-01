## Project Description
RESTful API Food Ordering System for Kopi Santai built with Golang and PostgreSQL.

### ER Diagram  
https://dbdiagram.io/embed/6310b9300911f91ba515ec28

### Coverage
![Coverage](https://heroku.example.com/assets/docs/coverage.png "coverage")
  
  
## Setup Project
1. Run in postgres [tables.sql](https://heroku.example.com/tables.sql) to create database and tables
2. Run in postgres [populate.sql](https://heroku.example.com/populate.sql) to seed database
3. Run go project ```ENV=dev go run .```

## How to run test
Run go project in testing environment  
```ENV=testing go test ./... -cover```  
  
## OpenAPI Documentation
[OpenAPI documentation](https://heroku.example.com/docs)  
  
## How to import postman
Exported json postman automatically set bearer token every login or register with user account.
Make sure to run login/register first to get the token.
1. Download [exported postman]("https://heroku.example.com/KopiSantai-OpenAPI3.0.postman_collection.json")
2. Open postman -> File -> Import (Ctrl+O)
3. Drag and drop the file to the box in postman
4. Check your API documentation with name Kopi Santai - OpenAPI 3.0
5. Let's play the postman
