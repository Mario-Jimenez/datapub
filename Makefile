build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o datapub-srv main.go

run:
	./datapub-srv

docker-build:
	docker build . -t docker.pkg.github.com/mario-jimenez/datapub/datapub-srv:0.0.1

docker-push:
	docker push docker.pkg.github.com/mario-jimenez/datapub/datapub-srv:0.0.1
