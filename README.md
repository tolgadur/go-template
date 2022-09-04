# email-project
This is a fun project with no purpose yet.

## how to run
You can build and run this either locally with `make build` and `make run` or with docker with `make docker-build` and 
`make docker-run`. The standard port that is exposed is 8080. The server will be available at `localhost:8080`.

If you want to run this with kubernetes you have to push the docker image to a registry with `make docker-push` and 
then run `make kube-deploy` to deploy the application to your kubernetes cluster. If you want to test this locally, you can use
`minikube start`. As this application relies on a loadbalancer service, make sure to use `minikube tunnel` in this case.

## example requests
```bash
❯ curl localhost:8080/Tolga
{"message":"Hello, Tolga"}
```

## postgres
To start the postgres server, run `make run-postgres`. This will start a persistent volume and 
start the helm chart with all necessary rights. You can then connect to the postgres server with: 

```bash
❯ export POSTGRES_PASSWORD=$(kubectl get secret --namespace default postgresql-dev -o jsonpath="{.data.password}" | base64 --decode)
❯ kubectl run postgresql-dev-client --rm --tty -i --restart='Never' --namespace default --image docker.io/bitnami/postgresql:14.1.0-debian-10-r80 --env="PGPASSWORD=$POSTGRES_PASSWORD" \
--command -- psql --host postgresql -U app1 -d app_db -p 5432
❯ \conninfo # command to check if you are connected to the correct database
```
