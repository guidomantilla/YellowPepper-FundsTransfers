build: validate compile

validate: format vet lint test

format:
	go fmt ./...

vet:
	go vet ./...

lint:
	golangci-lint run ./...

test:
	go test -covermode count -coverprofile coverage.out ./...

coverage-local: test
	go tool cover -html=coverage.out -o coverage.html
	open coverage.html

compile:
	go build -a -o Main.app .

serve:
	go run . serve

prepare:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go mod download

