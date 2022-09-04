deps:
	go mod tidy && go mod download && go mod verify

build:
	go build -o ./bin/email-project cmd/server/main.go

test:
	go test ./...

run:
	./bin/email-project

docker-build:
	docker build --tag email-project .

docker-run:
	docker run --publish 8080:8080 email-project

docker-push:
	docker tag email-project registry.digitalocean.com/email-project/main:latest
	&& docker push registry.digitalocean.com/email-project/main:latest

run-postgres:
	helm install postgresql -f ./internal/postgres/config.yml bitnami/postgresql

kube-deploy:
	kubectl apply -f ./kubernetes/service.yml && \
	kubectl apply -f ./kubernetes/deployment.yml && \
	kubectl apply -f ./kubernetes/persistent-volume.yml && \
	make run-postgres

download-registry-secrets:
	kubectl create secret generic do-registry \
      --from-file=.dockerconfigjson=docker-config.json \
      --type=kubernetes.io/dockerconfigjson

cleanup:
	kubectl delete deployment email-project && \
	kubectl delete svc email-project-web-service && \
	helm uninstall postgresql && \
	kubectl delete -f ./kubernetes/persistent-volume.yml
