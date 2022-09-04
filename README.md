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
