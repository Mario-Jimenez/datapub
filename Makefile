build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o datapub-srv main.go

run:
	./datapub-srv
