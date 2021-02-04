# Secrets for apps deployed into K8S

This repo is set of examples and prototypes for secrets management

## Repo status

---
    Active development
---

## Requirements

- [minikube](https://minikube.sigs.k8s.io/docs/start/)
- [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/)
- [helm](https://helm.sh/docs/intro/install/)
- [vault](https://vaultproject.io)
- [docker](https://docs.docker.com/get-docker/)

## How to start

- Run `make start-minikube`
- If you wish to have access to K8S dashboard run `make start-dashboard`
- Next you need to start registry
  - `cd registry`
  - `make start`
- Then build application
  - `cd ../web-app`
  - `make build-n-push arch=ubuntu`
  - `make build-n-push arch=scratch`
- And deployments
  - `cd ../basic-setup`
  - `make create-namespace`
  - `make apply-vault`
  - `make create-secret`
  - `make apply-bank-vaults`
  - `make apply-web-app`
