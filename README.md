# GoLang Server Boilerplate

## Features:
- Connected to Postgres
- Gorm integration & Migrations
- API serving

## Prereqs:
- Golang
- Postgres

## Setup:
1. ```go get github.com/ZaneWithSpoon/goServerBoiler```
2. ``cd $GOPATH/src/github.com/ZaneWithSpoon/goServerBoiler``
3. ```go get ./...```
4. ```touch password.txt```

Just update the password.txt and main.go to connect to the right postgres DB and ```go run main.go```

http://localhost:3001/test