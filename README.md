# Crud_With_Gorm
//Stages of Project
/*
1. CRUD --> Completed
2. Validation --> Completed
3. Service layer --> Completed
4. Repository layer --> completed
5. Swgger --> Parsial completed
6. JWT Auth. --> Done
7. Docker
8. Testing
9. Redis caching
10. gRPC 
*/

/*// Swagger idded for to test our endpoints

1. swag init
//Delete below two things from docs/docs.go line number 205
2. LeftDelim:        "{{",
	RightDelim:       "}}",
3. go run main.go */

## Setup
1. Copy `env.example` to `.env`
2. Update DB credentials and JWT secret
3. Run `go run main.go`

/*
login details
POST-->http://localhost:8000/login
user 1:
{
  "email": "admin@example.com",
  "password": "admin123"
}
User 2: 
{
    "email": "test@example.com",
    "password": "test123"
}

*connect docker container
run in cmd
docker exec -it gin-gorm-rest-docker-db-1 psql -U postgres -d postgres

*/