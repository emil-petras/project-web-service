# project-web-service

## list of related repositories
https://github.com/emil-petras/project-db-service
https://github.com/emil-petras/project-idempotency-service
https://github.com/emil-petras/project-proto

this service depends on project-db-service and project-idempotency-service

## run using docker compose
docker compose up -d

## run using docker file
docker build --rm -t project-web-service . 
docker run -p 8999:8999 -d project-web-service

## usage

-To check whether the service is working:
GET localhost:8999/ping

-To get an auth token:
POST localhost:8999/login
{
    "username": "mark",
	"password": "pass123"
}

-Once you get a token, you can use it to deposit and withdraw.
POST localhost:8999/deposit
{
    "token" : "69f64a44-caaa-462f-a649-533d98be8932",
	"amount": 15,
	"timestamp": 11854
}