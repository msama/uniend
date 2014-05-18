init:
	go get github.com/crowdmob/goamz/aws
	go get github.com/crowdmob/goamz/dynamodb
	go get github.com/go-martini/martini
	go get github.com/bitly/go-simplejson

build:
	go build -o bin/frontend ./frontend
	go build -o bin/admin ./admin

test:
    go test -race -timeout=1m -short ./frontend/...
    go test -race -timeout=1m -short ./admin/...

run-admin:
	./bin/admin -config=./dev-properties.json

run-frontend:
	./bin/frontend -config=./dev-properties.json