# Webserver finesse-frontend

Simple Web-server application written in Go lang with pipelines deploying the app on GoogleCloud.

1. [CircleCI](#circleCI)
2. [Github Actions](#gha) <br>
   2.1. [setup IAM](#iam) <br>
3. [Application run](#app)

## CircleCI  <a name="circleCI"></a>
Remeber that there are two pipelines:<br>
	- github actions <br>
	- circleCI <br>
which duplicates the data and can fake metrics shown in Grafana dashboard.

### Status of latest build from branches
main:
[![CircleCI](https://dl.circleci.com/status-badge/img/gh/ds4tech/finesse-frontend/tree/main.svg?style=svg)](https://dl.circleci.com/status-badge/redirect/gh/ds4tech/finesse-frontend/tree/main)
<br/>
dev:
[![CircleCI](https://dl.circleci.com/status-badge/img/gh/ds4tech/finesse-frontend/tree/dev.svg?style=svg)](https://dl.circleci.com/status-badge/redirect/gh/ds4tech/finesse-frontend/tree/dev)

### Continous Integration <a name="ci"></a>
Pipeline script written in yaml file for Circle CI is placed in [build/ci directory](https://github.com/ds4tech/finesse-frontend/blob/dev/.circleci/config.yml).  <br>


## Github Actions <a name="gha"></a>
### PreReq to deploy this on CloudRun using.
Certain resources must be created before the pipeline can be triggered. Otherwise it will fail, compleining on missing resources.
 
### Google Account
Based on:
https://cloud.google.com/iam/docs/workload-identity-federation-with-deployment-pipelines#gcloud
https://cloud.google.com/blog/products/identity-security/secure-your-use-of-third-party-tools-with-identity-federation

#### Create WIF <a name="iam"></a>
```
gcloud iam workload-identity-pools create github-actions-pool \
--location="global" \
--description="The pool to authenticate GitHub actions." \
--display-name="GitHub Actions Pool"
```
#### OICD
```
gcloud iam workload-identity-pools providers create-oidc github-actions-oidc --workload-identity-pool="github-actions-pool" \
--issuer-uri="https://token.actions.githubusercontent.com/" \
--attribute-mapping="google.subject=assertion.sub,attribute.repository=assertion.repository,attribute.repository_owner=assertion.repository_owner,attribute.branch=assertion.sub.extract('/heads/{branch}/')" \
--location=global \
--attribute-condition="assertion.repository_owner=='ds4tech'"
```
Get provide name:
```
gcloud iam workload-identity-pools providers describe github-actions-oidc --location="global" --project="finesse-406710" --workload-identity-pool="github-actions-pool"
```
#### Create Service Account
```
gcloud iam service-accounts create finesse-frontend-sa --display-name="Finesse Application Service Account" --description="manages the application resources"
```

#### Add policy binding
```
gcloud iam service-accounts add-iam-policy-binding finesse-frontend-sa@finesse-406710.iam.gserviceaccount.com --role="roles/CustomWorkloadIdentityUser" \
--member="principalSet://iam.googleapis.com/projects/645493513259/locations/global/workloadIdentityPools/github-actions-pool/attribute.repository_owner/ds4tech"
```
More about Service account impersonation: https://cloud.google.com/iam/docs/workload-identity-federation#impersonation

```
gcloud iam service-accounts add-iam-policy-binding finesse-frontend-sa@finesse-406710.iam.gserviceaccount.com --role="roles/iam.workloadIdentityUser" \
--member="principal://iam.googleapis.com/projects/645493513259/locations/global/workloadIdentityPools/github-actions-pool/subject/repo:ds4tech/finesse-frontend:ref:refs/heads/main"
```

#### Create Artifact Registry
```
 gcloud beta artifacts repositories create finesse-frontend --repository-format=docker --location=europe-central2 --description="Docker repository"
```

#### TROUBLESHOOTING
https://github.com/google-github-actions/auth/blob/main/docs/TROUBLESHOOTING.md



## Application run <a name="app"></a>

Set env var
```
export CALCULATOR_URL="http://localhost:8888"
```

Simple Webserver Go project:
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

http://localhost:8080/

### Docker container <a name="build.docker"></a>
```
docker build . -t ds4tech/finesse-frontend:0.0.1
docker run -it --rm -p 8080:8080 --name finesse-frontend ds4tech/finesse-frontend:0.0.1
```

http://localhost:8080/

## DEPLOY <a name="deploy"></a>

### Kubernetes <a name="deploy.k8s"></a>
```
kubectl apply -f deployment/kubernetes/manifest.yaml
kubectl port-forward svc/finesse-frontend 8080
```

http://localhost:8080/

### Helm <a name="deploy.k8s"></a>
```
helm install finesse-frontend ./deployment/helm/charts/finesse-frontend
kubectl port-forward svc/finesse-frontend 8080
```

http://localhost:8080/

## USEAGE <a name="usage"></a>

1. Echo
```
curl -X GET "http://localhost:8080/api/echo?text=testingJson"
```

