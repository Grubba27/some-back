# soundproof-back

- Candidate name: Gabriel Grubba, [linkedin](https://www.linkedin.com/in/gabriel-grubba/?locale=en_US)

## Description

Simple API with a few routes: 

- auth: login and register

- health: check if the API is up

- user: get user info, update user info, and get users by public adress(metamask)

to test run the folling commands, having goland installed and .env file with the following variables:

```bash
ENV="stag"
DATABASE_URL="pgurl"
JWT_SECRET="secret"
```
then run the following command to start the server:

```bash

go run main.go

```