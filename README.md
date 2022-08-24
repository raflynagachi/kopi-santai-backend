## Project Description
RESTful API Food Ordering System for Kopi Santai built with Golang and PostgreSQL.

### ER Diagram  
https://dbdiagram.io/embed/6305fffcf1a9b01b0fd380d8

### Coverage
![Coverage](https://heroku.example.com/cover.png "coverage")
  
  
## Setup Project
1. Run in postgres [tables.sql](https://heroku.example.com/tables.sql) to create database and tables
2. Run in postgres [populate.sql](https://heroku.example.com/populate.sql) to seed database
3. Run go project ```ENV=dev go run .```

## How to run test
Run go project in testing environment  
```ENV=testing go test ./... -cover```

## How to import postman
Exported json postman automatically set bearer token every login or register with user account.
Make sure to run login/register first to get the token.
1. Download [exported postman](https://heroku.example.com/final-project-backend.postman_collection.json)
2. Open postman -> File -> Import (Ctrl+O)
3. Drag and drop downloaded file to the box
4. Check your collection with name _example.json_
5. Let's play the postman
