# Requirement 

[Golang](https://golang.org/) 1.7+

[Govendor](https://github.com/kardianos/govendor) For package managment.

# TODO

- [x] add cors and jwt middleware
- [ ] implement as api docs

# Usage 

## Manual Test Server

Start server for test.

```sh
go run index.go
```

`POST /auth/signup`

```sh
curl -X POST -H "Content-Type:application/json" -d '{"phone": "1", "password": "1"}' http://localhost:3000/v1/auth/signup -i
```

`POST /auth/login`

```sh
curl -X POST -H "Content-Type:application/json" -d '{"phone": "1", "password": "1"}' http://localhost:3000/v1/auth/login -i
```

`GET /user` without token

```sh
curl -X GET -H  http://localhost:3000/v1/user -i
```

`GET /user` with token

```sh
curl -X GET -H "Authorization:BEARER eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjYzODBmMzQ3LWNjYTEtNGQyMC1iMTNkLTU5ZjYzMmEwZDI4YiIsImV4cCI6MTQ5NDk1MzgwNCwiaXNzIjoic2Vjc3lzIn0.EAtH8gLCRea6feM_EeKzdYD9XkJ4dk64WLRkTDQeTXU" http://localhost:3000/v1/user -i
```