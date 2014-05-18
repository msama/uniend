init:
	go get github.com/crowdmob/goamz/aws
	go get github.com/crowdmob/goamz/dynamodb
	go get github.com/go-martini/martini

build:
	go build -o bin/frontend ./frontend
	go build -o bin/admin ./admin

test:
    go test -race -timeout=1m -short ./frontend/...
    go test -race -timeout=1m -short ./admin/...

run-admin:
	./bin/admin

run-frontend:
	./bin/frontend