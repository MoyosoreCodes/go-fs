build:
	go build -v -o bin/fs main.go

run: build
	./bin/fs

test: 
	go test ./... -v

test-coverage:
	go test -coverprofile=coverage.out && go tool cover -html=coverage.out