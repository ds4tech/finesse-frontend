# covantis-sre


# Simple Web-server application written in Go lang. The purpose is to demonstrate variety of ways for CICD.

1. [Introduction](#intro)
2. [Build](#build) <br>
   2.1. [Exec](#build.exe) <br>
   2.2. [Docker](#build.docker)
3. [Deploy](#deploy) <br>
 3.1. [Kubernetes](#deploy.k8s) <br>
4. [Usage](#usage)
5. [Continous Integration](#ci)


## Introduction <a name="intro"></a>

Simple Go Calculator project with some math function:<a name="intro"></a>
- Echo

## BUILD <a name="build"></a>

### Executable <a name="build.exe"></a>
```
go build -o main cmd/calculator/main.go
./main

http://localhost:8888/
```

### Docker container <a name="build.docker"></a>
```
docker build -t go-calc -f build/package/Dockerfile .
docker run -d -p 80:8888 go-calc

http://localhost/
```

## DEPLOY <a name="deploy"></a>

### Kubernetes <a name="deploy.k8s"></a>
```
kubectl create -f deployments/kubernetes/k8s-replicaSet.yml

http://192.168.99.100:30000/
```


## USEAGE <a name="usage"></a>

1. Echo
- (x int64, y int64)
- http://localhost:8888/sum
- JSON: {"num1":"2","num2":"4"}
- curl:
```
curl --location --request POST "http://localhost:8888/sum" \
--header "Content-Type: application/json" \
--data "{\"num1\":\"2\",
\"num2\":\"4\"}"
```
## Continous Integration <a name="ci"></a>
Pipeline script written in Groovy is placed in [build/ci directory](https://github.com/ds4tech/pipeline-calculator-ws/blob/master/build/ci/pipeline.yaml). It is dedicated for Jenkins Pipeline JOB. <br>
To run Jenkins on AWS, run terraform scripts in deployments/aws/jenkins-ec2_instance
