
## Description

Simple API with a few routes: 

- auth: login and register

- health: check if the API is up

- user: get user info, update user info, and get users by public adress(metamask)

## Running locally

to test run the folling commands, having goland installed and .env file with the following variables:

```bash
PROD="false"
DATABASE_URL="some"
JWT_SECRET="secret"
RAILWAY="false"
PORT="8080"

```

for generating the swagger documentation run the following command:

```bash
swag init
```

then run the following command to start the server:

```bash

go run main.go

```
## Deployed

You can test the API on the following link:
[swagger](https://soundproof-back-production.up.railway.app/swagger/index.html)

```txt
soundproof-back-production.up.railway.app/swagger/index.html
```

## Deploying

```bash
railway up
```
was the command used to deploy the API but you can just use the following command to deploy anywhere:

```bash
go build -o out
```

then run the out file

```bash
./out
```
