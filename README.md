## Project Description
RESTful API Food Ordering System for Kopi Santai built with Golang and PostgreSQL.

### ER Diagram  
https://dbdiagram.io/embed/631486110911f91ba52f63df  
![ERD](https://git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/-/raw/master/assets/docs/erd.png "ERD")

### Coverage
![Coverage](https://git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/-/raw/master/assets/docs/coverage.png "coverage")
  
  
## Setup Project on local
1. Run in postgres [tables.sql](https://git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/-/raw/master/tables.sql) to create database and tables
2. Run in postgres [populate.sql](https://git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/-/raw/master/populate.sql) to seed database
3. Run go project ```DATABASE_URL= ENV=dev go run .```

## How to run test on local
Run go project in testing environment
```DATABASE_URL= ENV=testing go test ./... -cover```  
  
## OpenAPI Documentation
[OpenAPI documentation](https://kopi-santai.herokuapp.com/docs)  
  
## How to import postman
Exported json postman automatically set bearer token every login or register with user account.
Make sure to run login/register first to get the token.
1. Download [exported postman]("https://git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/-/blob/master/Kopi%20Santai%20-%20OpenAPI%203.0.postman_collection.json")
2. Open postman -> File -> Import (Ctrl+O)
3. Drag and drop the file to the box in postman
4. Check your API documentation with name Kopi Santai - OpenAPI 3.0
5. Let's play the postman
