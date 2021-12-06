validate: format vet lint test

format:
	go fmt ./...

vet:
	go vet ./...

lint:
	golangci-lint run ./...

test:
	go test ./...

build:
	go build -a -o Main.app .

serve:
	go run . serve

install:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.43.0
	go install github.com/kisielk/errcheck@v1.6.0
	go mod download

