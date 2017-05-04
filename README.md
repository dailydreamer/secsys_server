# Requirement 

[Golang](https://golang.org/) 1.7+

[Govendor](https://github.com/kardianos/govendor) For package managment.

# TODO

- [x] add cors and jwt middleware
- [x] complete api design
- [x] golang json null omitempty survey
- [x] implement middleware
- [x] implement controllers
- [ ] insert log and message auto creating into controllers
- [ ] verification
- [ ] file upload

# Usage 

## Manual Test Server

Init database.

```sh
psql postgres://bzjgtwxllnppiq:31a68d8748d99ded0a1d42f95a6c28eef56f0cf5fd892250f1905f99bde7ce95@ec2-54-235-90-107.compute-1.amazonaws.com:5432/d5imq2f3o1cs7a -f ./scripts/initdb.sql
```

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
curl -X GET http://localhost:3000/v1/users/9cf0eb50-49d9-43b7-858f-9d97bd082230 -i
```

`GET /users/:userID` with token

```sh
curl -X GET -H "Authorization:BEARER eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjljZjBlYjUwLTQ5ZDktNDNiNy04NThmLTlkOTdiZDA4MjIzMCIsImlzQWRtaW4iOnRydWUsImV4cCI6MTQ5NTIxMjAyNywiaXNzIjoic2Vjc3lzIn0.4usS8PZUvA7AZNIX0ErpzLAds29rLPtWevkNTWKvDUw" http://localhost:3000/v1/users/9cf0eb50-49d9-43b7-858f-9d97bd082230 -i
```