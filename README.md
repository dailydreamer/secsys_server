# Requirement 

[Golang](https://golang.org/) 1.7+

[Govendor](https://github.com/kardianos/govendor) For package managment.

# TODO

- [x] add cors and jwt middleware
- [x] complete api design
- [ ] implemnet controllers

# Usage 

## Manual Test Server

Start server for test.

```sh
go run index.go
```

`POST /signup`

```sh
curl -X POST -H "Content-Type:application/json" -d '{"phone": "1", "password": "1"}' http://localhost:3000/v1/signup -i
```

`POST /login`

```sh
curl -X POST -H "Content-Type:application/json" -d '{"phone": "1", "password": "1"}' http://localhost:3000/v1/login -i
```

`GET /users/:userID` without token

```sh
curl -X GET http://localhost:3000/v1/users/6380f347-cca1-4d20-b13d-59f632a0d28b -i
```

`GET /users/:userID` with token

```sh
curl -X GET -H "Authorization:BEARER eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjYzODBmMzQ3LWNjYTEtNGQyMC1iMTNkLTU5ZjYzMmEwZDI4YiIsImV4cCI6MTQ5NDk1MzgwNCwiaXNzIjoic2Vjc3lzIn0.EAtH8gLCRea6feM_EeKzdYD9XkJ4dk64WLRkTDQeTXU" http://localhost:3000/v1/users/6380f347-cca1-4d20-b13d-59f632a0d28b -i
```