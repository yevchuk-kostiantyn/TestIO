## TestIO

TestIO is a prototype of the testing knowledge system.

## Current Status

TestIO is in the development stage.

## Tests

To run tests:
1. `cd $GOPATH/src/github.com/yevchuk-kostiantyn/TestIO`
2. `go test ./...`

## Setup

The program requires:
1. Golang Programming Language (version 1.9.1 and higher)
2. Docker (version 1.12.6 and higher, build 78d1802)
3. Minikube (version: v0.24.1)
4. Kubectl

Grab the latest version of the program by running `git clone https://github.com/yevchuk-kostiantyn/TestIO.git`. The repo should be in `$GOPATH/src/github.com/yevchuk-kostiantyn` directory!

To run the TestIO program: 
1. `cd $GOPATH/src/github.com/yevchuk-kostiantyn/TestIO`
2. `docker build -t testio .`
3. `docker run --publish 1997:1997 --name testiorun --rm testio`
4. open `http://localhost:1997/` in a web browser.

To run the TestIO program in the Kubernetes cluster:
1. `kubectl run kubernetes-testio --image=docker.io/kostiantynyevchuk/testio:latest --port=1997`
2. Check if deployment was successful by running `kubectl get deployments`
3. In the extra terminal window run `kubectl proxy`
4. `export POD_NAME=$(kubectl get pods -o go-template --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')`
5. `curl http://localhostpi/v1/proxy/namespaces/default/pods/$POD_NAME/`


## Contributors
1. Kostiantyn Yevchuk: kostyantyn.yewchuk@gmail.com