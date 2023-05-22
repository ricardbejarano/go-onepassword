all: lint tidy test build

lint:
	goimports -w .
	gofmt -w -s .

tidy:
	go mod tidy

test:
	go vet ./...
	go test ./...

build:
	CGO_ENABLED=0 go build -o bin/ ./...

clean:
	rm -r bin
