# go-grpc-demo

gRPC server along with getaway Mux for http proxy in Go

## Install

```
go mod download
```

## Run server

```
cd cmd/server
./run
```

## Auth test

```bash
#!/bin/bash

# request token with `admin` role
AUTH_DATA_1=$(curl http://localhost:8080/auth?role=admin -H "Accept: application/json")
TOKEN_1=$(echo "$AUTH_DATA_1" | gawk '{ match($0, /:"(.+)"/, arr); if(arr[1] != "") print arr[1] }')
curl http://localhost:8080/users -H "Accept: application/json" -H "Authorization: Bearer $TOKEN_1"

# request token with `user` role
AUTH_DATA_2=$(curl http://localhost:8080/auth?role=user -H "Accept: application/json")
TOKEN_2=$(echo "$AUTH_DATA_2" | gawk '{ match($0, /:"(.+)"/, arr); if(arr[1] != "") print arr[1] }')
curl http://localhost:8080/users -H "Accept: application/json" -H "Authorization: Bearer $TOKEN_2"
```
