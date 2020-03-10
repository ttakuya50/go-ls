build:
	go build -o go-ls -v *.go

test:
	go test -race -v ./...