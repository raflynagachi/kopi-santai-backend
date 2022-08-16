## Project Description
RESTful API Food Ordering System for Kopi Santai built with Golang and PostgreSQL.

### ER Diagram
![ER Diagram](https://heroku.example.com/erd.png "ERD")
  
### Coverage
![Coverage](https://heroku.example.com/cover.png "coverage")
  
  
## Setup Project
1. Run [raw.sql](https://heroku.example.com/raw.sql) to create database, tables and do seeding
2. Run go project ```ENV=dev go run .```

## How to run test
Run go project in testing environment  
```ENV=testing go test ./... -cover```

## How to import postman
Exported json postman automatically set bearer token every login or register with user account.
Make sure to run login/register first to get the token.
1. Download [exported postman](https://heroku.example.com/postman_collection.json)
2. Open postman -> File -> Import (Ctrl+O)
3. Drag and drop downloaded file to the box
4. Check your collection with name _example.json_
5. Let's play the postman
