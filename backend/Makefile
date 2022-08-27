build:
	go mod download && go mod verify && go build -o ./bin/email-project cmd/server/main.go

docker-build:
	docker build --tag email-project .

docker-run:
	docker run --publish 8080:8080 email-project

docker-push:
	docker tag email-project registry.digitalocean.com/email-project/main:latest && docker push registry.digitalocean.com/email-project/main:latest