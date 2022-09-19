deps:
	go mod tidy && go mod download && go mod verify

build:
	go build -o ./bin/email-project cmd/server/main.go && \
	make docker-build

test:
	go test ./...

run:
	./bin/email-project

download-registry-secrets:
	kubectl create secret generic do-registry \
      --from-file=.dockerconfigjson=docker-config.json \
      --type=kubernetes.io/dockerconfigjson

docker-build:
	docker build --tag email-project .

docker-run:
	docker run --publish 8080:8080 email-project

docker-push:
	docker tag email-project registry.digitalocean.com/email-project/main:latest && \
 	docker push registry.digitalocean.com/email-project/main:latest

kube-run-postgres:
	kubectl apply -f ./kubernetes/persistent-volume.yaml && \
	helm install postgresql -f ./db/values.yaml bitnami/postgresql --set volumePermissions.enabled=true

kube-delete-postgres:
	helm delete postgresql && \
	kubectl delete -f ./kubernetes/persistent-volume.yaml

kube-run-app:
	make download-registry-secrets && \
	kubectl apply -f ./kubernetes/service.yaml && \
	kubectl apply -f ./kubernetes/deployment.yaml

kube-cleanup:
	kubectl delete deployment email-project && \
	kubectl delete svc email-project-web-service && \
	make kube-delete-postgres

minikube-start:
	minikube start && \
	minikube tunnel

minikube-stop:
	minikube stop && \
	minikube delete

create-local-db:
	psql  -c 'DROP DATABASE IF EXISTS app_db;' && \
	psql  -c 'CREATE DATABASE app_db;'