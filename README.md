# Webserver

Status of latest build from branches:<br/> 
main:
[![CircleCI](https://dl.circleci.com/status-badge/img/gh/ds4tech/webserver-frontend/tree/main.svg?style=svg)](https://dl.circleci.com/status-badge/redirect/gh/ds4tech/webserver-frontend/tree/main)
dev:
[![CircleCI](https://dl.circleci.com/status-badge/img/gh/ds4tech/webserver-frontend/tree/dev.svg?style=svg)](https://dl.circleci.com/status-badge/redirect/gh/ds4tech/webserver-frontend/tree/dev)


## PreReq
```
export CALCULATOR_URL="http://localhost:8888"
```

### Simple Web-server application written in Go lang.

1. [Introduction](#intro)
2. [Build](#build) <br>
   2.1. [Exec](#build.exe) <br>
   2.2. [Docker](#build.docker)
3. [Deploy](#deploy) <br>
 3.1. [Kubernetes](#deploy.k8s) <br>
4. [Usage](#usage)
5. [Continous Integration](#ci)


## Introduction <a name="intro"></a>

Simple Webserver Go project:<a name="intro"></a>
API:
- /-/health - returns server version 
- echo - /api/echo?text=foo --> returns a JSON object with the key "text

Main page shows form which allows to input values which are sent to calculator webservice.

## BUILD <a name="build"></a>

### Executable <a name="build.exe"></a>
```
go build -o webserver cmd/main.go 
./webserver
```

http://localhost:9090/

### Docker container <a name="build.docker"></a>
```
docker build . -t ds4tech/webserver:0.0.1
docker run -it --rm -p 9090:9090 --name hello-webserver ds4tech/webserver:0.0.1
```

http://localhost:9090/

## DEPLOY <a name="deploy"></a>

### Kubernetes <a name="deploy.k8s"></a>
```
kubectl apply -f deployment/kubernetes/manifest.yaml
kubectl port-forward svc/hello-webserver 9090
```

http://localhost:9090/

### Helm <a name="deploy.k8s"></a>
```
helm install hello-webserver ./deployment/helm/charts/hello-webserver
kubectl port-forward svc/hello-webserver 9090
```

http://localhost:9090/

## USEAGE <a name="usage"></a>

1. Echo
```
curl -X GET "http://localhost:9090/api/echo?text=testingJson"
```

## Continous Integration <a name="ci"></a>
Pipeline script written in yaml file for Circle CI is placed in [build/ci directory](https://github.com/ds4tech/covantis-sre/blob/main/.circleci/config.yml).  <br>
