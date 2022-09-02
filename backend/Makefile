deps:
	go mod tidy && go mod download && go mod verify

build:
	go build -o ./bin/email-project cmd/server/main.go

docker-build:
	docker build --tag email-project .

docker-run:
	docker run --publish 8080:8080 email-project

docker-push:
	docker tag email-project registry.digitalocean.com/email-project/main:latest && docker push registry.digitalocean.com/email-project/main:latest

kube-deploy:
	kubectl apply -f deployment.yml

delete-kube-deploy:
	kubectl delete deployment email-project

download-registry-secrets:
	kubectl create secret generic do-registry \
      --from-file=.dockerconfigjson=docker-config.json \
      --type=kubernetes.io/dockerconfigjson
