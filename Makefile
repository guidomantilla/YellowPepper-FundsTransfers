serve:
	go run . serve

test:
	go test -v -race ./...
	gofmt -d .
	go vet ./...

build:
	go build -a -o Main.app .

install:
	go mod download

